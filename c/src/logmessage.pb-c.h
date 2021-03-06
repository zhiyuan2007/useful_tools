/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: logmessage.proto */

#ifndef PROTOBUF_C_logmessage_2eproto__INCLUDED
#define PROTOBUF_C_logmessage_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1000000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1000000 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _LogMessage LogMessage;


/* --- enums --- */


/* --- messages --- */

struct  _LogMessage
{
  ProtobufCMessage base;
  char *id;
  int32_t date;
  int32_t port;
  char *cip;
  char *view;
  char *domain;
  char *cls;
  char *rtype;
  char *rcode;
  protobuf_c_boolean has_recur;
  int32_t recur;
  protobuf_c_boolean has_signed_;
  int32_t signed_;
  protobuf_c_boolean has_edns;
  int32_t edns;
  protobuf_c_boolean has_tcp;
  int32_t tcp;
  protobuf_c_boolean has_dnssec;
  int32_t dnssec;
  protobuf_c_boolean has_cd;
  int32_t cd;
};
#define LOG_MESSAGE__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&log_message__descriptor) \
    , NULL, 0, 0, NULL, NULL, NULL, NULL, NULL, NULL, 0,0, 0,0, 0,0, 0,0, 0,0, 0,0 }


/* LogMessage methods */
void   log_message__init
                     (LogMessage         *message);
size_t log_message__get_packed_size
                     (const LogMessage   *message);
size_t log_message__pack
                     (const LogMessage   *message,
                      uint8_t             *out);
size_t log_message__pack_to_buffer
                     (const LogMessage   *message,
                      ProtobufCBuffer     *buffer);
LogMessage *
       log_message__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   log_message__free_unpacked
                     (LogMessage *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*LogMessage_Closure)
                 (const LogMessage *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor log_message__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_logmessage_2eproto__INCLUDED */
