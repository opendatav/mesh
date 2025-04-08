# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transfer_outer.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14transfer_outer.proto\x12\x0borg.ppc.ptp\"\x81\x01\n\x07Inbound\x12\x34\n\x08metadata\x18\x01 \x03(\x0b\x32\".org.ppc.ptp.Inbound.MetadataEntry\x12\x0f\n\x07payload\x18\x02 \x01(\x0c\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xa2\x01\n\x08Outbound\x12\x35\n\x08metadata\x18\x01 \x03(\x0b\x32#.org.ppc.ptp.Outbound.MetadataEntry\x12\x0f\n\x07payload\x18\x02 \x01(\x0c\x12\x0c\n\x04\x63ode\x18\x03 \x01(\t\x12\x0f\n\x07message\x18\x04 \x01(\t\x1a/\n\rMetadataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x32\x8e\x01\n\x17PrivateTransferProtocol\x12<\n\ttransport\x12\x14.org.ppc.ptp.Inbound\x1a\x15.org.ppc.ptp.Outbound(\x01\x30\x01\x12\x35\n\x06invoke\x12\x14.org.ppc.ptp.Inbound\x1a\x15.org.ppc.ptp.Outboundb\x06proto3')



_INBOUND = DESCRIPTOR.message_types_by_name['Inbound']
_INBOUND_METADATAENTRY = _INBOUND.nested_types_by_name['MetadataEntry']
_OUTBOUND = DESCRIPTOR.message_types_by_name['Outbound']
_OUTBOUND_METADATAENTRY = _OUTBOUND.nested_types_by_name['MetadataEntry']
Inbound = _reflection.GeneratedProtocolMessageType('Inbound', (_message.Message,), {

  'MetadataEntry' : _reflection.GeneratedProtocolMessageType('MetadataEntry', (_message.Message,), {
    'DESCRIPTOR' : _INBOUND_METADATAENTRY,
    '__module__' : 'transfer_outer_pb2'
    # @@protoc_insertion_point(class_scope:org.ppc.ptp.Inbound.MetadataEntry)
    })
  ,
  'DESCRIPTOR' : _INBOUND,
  '__module__' : 'transfer_outer_pb2'
  # @@protoc_insertion_point(class_scope:org.ppc.ptp.Inbound)
  })
_sym_db.RegisterMessage(Inbound)
_sym_db.RegisterMessage(Inbound.MetadataEntry)

Outbound = _reflection.GeneratedProtocolMessageType('Outbound', (_message.Message,), {

  'MetadataEntry' : _reflection.GeneratedProtocolMessageType('MetadataEntry', (_message.Message,), {
    'DESCRIPTOR' : _OUTBOUND_METADATAENTRY,
    '__module__' : 'transfer_outer_pb2'
    # @@protoc_insertion_point(class_scope:org.ppc.ptp.Outbound.MetadataEntry)
    })
  ,
  'DESCRIPTOR' : _OUTBOUND,
  '__module__' : 'transfer_outer_pb2'
  # @@protoc_insertion_point(class_scope:org.ppc.ptp.Outbound)
  })
_sym_db.RegisterMessage(Outbound)
_sym_db.RegisterMessage(Outbound.MetadataEntry)

_PRIVATETRANSFERPROTOCOL = DESCRIPTOR.services_by_name['PrivateTransferProtocol']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _INBOUND_METADATAENTRY._options = None
  _INBOUND_METADATAENTRY._serialized_options = b'8\001'
  _OUTBOUND_METADATAENTRY._options = None
  _OUTBOUND_METADATAENTRY._serialized_options = b'8\001'
  _INBOUND._serialized_start=38
  _INBOUND._serialized_end=167
  _INBOUND_METADATAENTRY._serialized_start=120
  _INBOUND_METADATAENTRY._serialized_end=167
  _OUTBOUND._serialized_start=170
  _OUTBOUND._serialized_end=332
  _OUTBOUND_METADATAENTRY._serialized_start=120
  _OUTBOUND_METADATAENTRY._serialized_end=167
  _PRIVATETRANSFERPROTOCOL._serialized_start=335
  _PRIVATETRANSFERPROTOCOL._serialized_end=477
# @@protoc_insertion_point(module_scope)