import grpc
import requests

from protos import transfer_outer_pb2_grpc
from protos.transfer_inner_pb2 import PushInbound
from protos.transfer_outer_pb2 import Inbound


def push():
    conn = grpc.insecure_channel('127.0.0.1:7304')
    # conn = grpc.insecure_channel('10.10.20.14:32072')
    request = PushInbound()
    request.topic = "intersection-1742281765923"
    request.payload = b'1'
    data = request.SerializeToString()

    client = transfer_outer_pb2_grpc.PrivateTransferProtocolStub(channel=conn)

    req_data = Inbound()
    # 正确的方式：逐个添加键值对到 metadata 中
    req_data.metadata["x-ptp-version"] = "v1"
    req_data.metadata["x-ptp-tech-provider-code"] = "LX"
    req_data.metadata["x-ptp-trace-id"] = "AC1428EB2516155457700347904"
    req_data.metadata["x-ptp-token"] = "token_job_b0d8fd6011044ccc8c04d6d6d3922c20"
    req_data.metadata["x-ptp-uri"] = "https://ptp.cn/org.ppc.ptp.PrivateTransferTransport/push"
    req_data.metadata["x-ptp-session-id"] = "session_job_b0d8fd6011044ccc8c04d6d6d3922c20"
    req_data.metadata["x-ptp-source-node-id"] = "2025022116041625"
    req_data.metadata["x-ptp-target-node-id"] = "unionpay10000"
    req_data.metadata["x-ptp-source-inst-id"] = "2025022116041625"
    req_data.metadata["x-ptp-target-inst-id"] = "unionpay10000"
    req_data.payload = data
    response = client.invoke(
        req_data, 10, [
            ('x-ptp-trace-id', 'AC1428EB2516155457700347904'),
            ('x-ptp-target-node-id', 'LX0000000000000'),
            ('x-ptp-uri', 'https://ptp.cn/org.ppc.ptp.PrivateTransferTransport/push')
        ]
    )

    print(f"Response: {response}")

    pi = PushInbound()
    pi.topic = "321"
    pi.payload = b'1234567'
    inbound = Inbound()
    inbound.payload = pi.SerializeToString()
    response = requests.post(url='http://127.0.0.1:7304/v1/interconn/chan/invoke', headers={
        'Content-Type': 'application/x-protobuf',
        'x-ptp-tech-provider-code': 'LX',
        'x-ptp-trace-id': '1',
        'x-ptp-session-id': '1',
        'x-ptp-token': 'x',
        'x-ptp-source-node-id': '2',
        'x-ptp-source-inst-id': '2',
        'x-ptp-target-node-id': 'LX0000000000000',
        'x-ptp-target-inst-id': '',
        'x-ptp-uri': 'https://ptp.cn/v1/interconn/chan/push',
    }, data=inbound.SerializeToString())
    print(f"h1 {response.status_code} {response.content.decode('utf-8')}")


if __name__ == "__main__":
    push()
