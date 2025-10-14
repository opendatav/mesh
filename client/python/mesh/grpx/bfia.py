#
# Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
# TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
#
#

from concurrent.futures import Future
from typing import Any

import grpc
import mesh.codec.tools as tools
import mesh.tool as tool
import requests
from mesh.cause import MeshException, MeshCode
from mesh.context import Mesh
from mesh.grpx.channels import GrpcChannels
from mesh.kinds import Reference
from mesh.mpc.invoker import Execution
from mesh.ptp.y_pb2 import PeekInbound, PopInbound, PushInbound, ReleaseInbound, TransportOutbound


class BfiaBinding:
    def __init__(self, urn: str, args: Any):
        if urn.startswith("peek.chan.mesh"):
            self.path = "/org.ppc.ptp.PrivateTransferTransport/peek"
            self.h1path = "/v1/interconn/chan/peek"
            self.inbound = PeekInbound()
            self.inbound.topic = args.topic
            self.outbound = TransportOutbound
        elif urn.startswith("pop.chan.mesh"):
            self.path = "/org.ppc.ptp.PrivateTransferTransport/pop"
            self.h1path = "/v1/interconn/chan/pop"
            self.inbound = PopInbound()
            self.inbound.topic = args.topic
            self.inbound.timeout = args.timeout
            self.outbound = TransportOutbound
        elif urn.startswith("push.chan.mesh"):
            self.path = "/org.ppc.ptp.PrivateTransferTransport/push"
            self.h1path = "/v1/interconn/chan/push"
            self.inbound = PushInbound()
            self.inbound.topic = args.topic
            self.inbound.payload = args.payload
            if args.metadata:
                for k, v in args.metadata.items():
                    self.inbound.metadata[k] = v
            self.outbound = TransportOutbound
        elif urn.startswith("release.chan.mesh"):
            self.path = "/org.ppc.ptp.PrivateTransferTransport/release"
            self.h1path = "/v1/interconn/chan/release"
            self.inbound = ReleaseInbound()
            self.inbound.topic = args.topic
            self.inbound.timeout = args.timeout
            self.outbound = TransportOutbound
        else:
            self.path = ""
            self.inbound = None
            self.outbound = None

    def encode(self, _: bytes):
        return self.inbound.SerializeToString()

    def decode(self, buffer: bytes):
        out = self.outbound()
        out.ParseFromString(buffer)
        v = "{" + f'"code":"{out.code}","message":"{out.message}","content":"{tools.b64encode(out.payload)}"' + "}"
        return v.encode('utf-8')


def bfia_consume(address: str, binding: BfiaBinding, execution: Execution[Reference], inbound: bytes,
                 channel: GrpcChannels) -> Future:
    if 'grpc' != tool.get_bfia_proto():
        return bfia_consume_h11(address, binding, execution)

    urn = Mesh.context().get_urn()
    try:
        Mesh.context().rewrite_urn("")
        ct = Mesh.context().get_attribute(Mesh.TIMEOUT)
        timeout = (ct if ct else execution.schema().timeout) / 1000
        channel = channel.create_if_absent(address)
        conn: grpc.UnaryUnaryMultiCallable = channel.unary_unary(binding.path, binding.encode, binding.decode)
        outbound = conn(inbound, timeout, [], None, True, False)
        future = Future()
        future.set_result(outbound)
    finally:
        Mesh.context().rewrite_urn(urn)
    return future


session = requests.Session()


def bfia_consume_h11(address: str, binding: BfiaBinding, execution: Execution[Reference]) -> Future:
    urn = Mesh.context().get_urn()
    try:
        Mesh.context().rewrite_urn("")
        ct = Mesh.context().get_attribute(Mesh.TIMEOUT)
        timeout = (ct if ct else execution.schema().timeout) / 1000
        headers = {'content-type': 'application/x-protobuf'}
        for (k, v) in Mesh.context().get_attachments().items():
            if k.startswith('x-ptp-'):
                headers[k] = v
        rs = session.post(url=f"http://{address}{binding.h1path}", data=binding.encode(b''), timeout=timeout,
                          headers=headers)

        if rs.status_code != 200:
            raise MeshException(MeshCode.SYSTEM_ERROR.get_code(), f"{binding.h1path} {rs.status_code} {rs.text}")
        outbound = binding.decode(rs.content)
        future = Future()
        future.set_result(outbound)
    finally:
        Mesh.context().rewrite_urn(urn)
    return future
