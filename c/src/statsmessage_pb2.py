# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: statsmessage.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)




DESCRIPTOR = _descriptor.FileDescriptor(
  name='statsmessage.proto',
  package='',
  serialized_pb='\n\x12statsmessage.proto\"7\n\x0cStatsRequest\x12\x0b\n\x03key\x18\x02 \x02(\t\x12\x0c\n\x04view\x18\x03 \x02(\t\x12\x0c\n\x04topn\x18\x04 \x01(\x05\"&\n\x07MsgCell\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05\x63ount\x18\x02 \x01(\x05\"n\n\nStatsReply\x12\x0b\n\x03key\x18\x02 \x01(\t\x12\r\n\x05maybe\x18\x03 \x01(\t\x12\r\n\x05value\x18\x04 \x01(\x02\x12\x18\n\x06result\x18\x05 \x03(\x0b\x32\x08.MsgCell\x12\x0c\n\x04name\x18\x06 \x03(\t\x12\r\n\x05\x63ount\x18\x07 \x03(\x05')




_STATSREQUEST = _descriptor.Descriptor(
  name='StatsRequest',
  full_name='StatsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='StatsRequest.key', index=0,
      number=2, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='view', full_name='StatsRequest.view', index=1,
      number=3, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='topn', full_name='StatsRequest.topn', index=2,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=22,
  serialized_end=77,
)


_MSGCELL = _descriptor.Descriptor(
  name='MsgCell',
  full_name='MsgCell',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='MsgCell.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='count', full_name='MsgCell.count', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=79,
  serialized_end=117,
)


_STATSREPLY = _descriptor.Descriptor(
  name='StatsReply',
  full_name='StatsReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='StatsReply.key', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='maybe', full_name='StatsReply.maybe', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=unicode("", "utf-8"),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='StatsReply.value', index=2,
      number=4, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='result', full_name='StatsReply.result', index=3,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='StatsReply.name', index=4,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='count', full_name='StatsReply.count', index=5,
      number=7, type=5, cpp_type=1, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  serialized_start=119,
  serialized_end=229,
)

_STATSREPLY.fields_by_name['result'].message_type = _MSGCELL
DESCRIPTOR.message_types_by_name['StatsRequest'] = _STATSREQUEST
DESCRIPTOR.message_types_by_name['MsgCell'] = _MSGCELL
DESCRIPTOR.message_types_by_name['StatsReply'] = _STATSREPLY

class StatsRequest(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _STATSREQUEST

  # @@protoc_insertion_point(class_scope:StatsRequest)

class MsgCell(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _MSGCELL

  # @@protoc_insertion_point(class_scope:MsgCell)

class StatsReply(_message.Message):
  __metaclass__ = _reflection.GeneratedProtocolMessageType
  DESCRIPTOR = _STATSREPLY

  # @@protoc_insertion_point(class_scope:StatsReply)


# @@protoc_insertion_point(module_scope)