// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: domain/billing/payment/v1/event.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event int32

const (
	Event_EVENT_UNSPECIFIED      Event = 0
	Event_EVENT_PAYMENT_CREATED  Event = 1
	Event_EVENT_PAYMENT_APPROVED Event = 2
	Event_EVENT_PAYMENT_CLOSED   Event = 3
	Event_EVENT_PAYMENT_REJECTED Event = 4
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "EVENT_PAYMENT_CREATED",
		2: "EVENT_PAYMENT_APPROVED",
		3: "EVENT_PAYMENT_CLOSED",
		4: "EVENT_PAYMENT_REJECTED",
	}
	Event_value = map[string]int32{
		"EVENT_UNSPECIFIED":      0,
		"EVENT_PAYMENT_CREATED":  1,
		"EVENT_PAYMENT_APPROVED": 2,
		"EVENT_PAYMENT_CLOSED":   3,
		"EVENT_PAYMENT_REJECTED": 4,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_billing_payment_v1_event_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_domain_billing_payment_v1_event_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_domain_billing_payment_v1_event_proto_rawDescGZIP(), []int{0}
}

type EventPaymentCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status StatusPayment `protobuf:"varint,2,opt,name=status,proto3,enum=domain.billing.payment.v1.StatusPayment" json:"status,omitempty"`
	UserId string        `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *EventPaymentCreated) Reset() {
	*x = EventPaymentCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_billing_payment_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPaymentCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPaymentCreated) ProtoMessage() {}

func (x *EventPaymentCreated) ProtoReflect() protoreflect.Message {
	mi := &file_domain_billing_payment_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPaymentCreated.ProtoReflect.Descriptor instead.
func (*EventPaymentCreated) Descriptor() ([]byte, []int) {
	return file_domain_billing_payment_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventPaymentCreated) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EventPaymentCreated) GetStatus() StatusPayment {
	if x != nil {
		return x.Status
	}
	return StatusPayment_STATUS_PAYMENT_UNSPECIFIED
}

func (x *EventPaymentCreated) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type EventPaymentApproved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status StatusPayment `protobuf:"varint,1,opt,name=status,proto3,enum=domain.billing.payment.v1.StatusPayment" json:"status,omitempty"`
}

func (x *EventPaymentApproved) Reset() {
	*x = EventPaymentApproved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_billing_payment_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPaymentApproved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPaymentApproved) ProtoMessage() {}

func (x *EventPaymentApproved) ProtoReflect() protoreflect.Message {
	mi := &file_domain_billing_payment_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPaymentApproved.ProtoReflect.Descriptor instead.
func (*EventPaymentApproved) Descriptor() ([]byte, []int) {
	return file_domain_billing_payment_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventPaymentApproved) GetStatus() StatusPayment {
	if x != nil {
		return x.Status
	}
	return StatusPayment_STATUS_PAYMENT_UNSPECIFIED
}

type EventPaymentRejected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status StatusPayment `protobuf:"varint,1,opt,name=status,proto3,enum=domain.billing.payment.v1.StatusPayment" json:"status,omitempty"`
}

func (x *EventPaymentRejected) Reset() {
	*x = EventPaymentRejected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_billing_payment_v1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPaymentRejected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPaymentRejected) ProtoMessage() {}

func (x *EventPaymentRejected) ProtoReflect() protoreflect.Message {
	mi := &file_domain_billing_payment_v1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPaymentRejected.ProtoReflect.Descriptor instead.
func (*EventPaymentRejected) Descriptor() ([]byte, []int) {
	return file_domain_billing_payment_v1_event_proto_rawDescGZIP(), []int{2}
}

func (x *EventPaymentRejected) GetStatus() StatusPayment {
	if x != nil {
		return x.Status
	}
	return StatusPayment_STATUS_PAYMENT_UNSPECIFIED
}

type EventPaymentClosed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status StatusPayment `protobuf:"varint,1,opt,name=status,proto3,enum=domain.billing.payment.v1.StatusPayment" json:"status,omitempty"`
}

func (x *EventPaymentClosed) Reset() {
	*x = EventPaymentClosed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_billing_payment_v1_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPaymentClosed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPaymentClosed) ProtoMessage() {}

func (x *EventPaymentClosed) ProtoReflect() protoreflect.Message {
	mi := &file_domain_billing_payment_v1_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPaymentClosed.ProtoReflect.Descriptor instead.
func (*EventPaymentClosed) Descriptor() ([]byte, []int) {
	return file_domain_billing_payment_v1_event_proto_rawDescGZIP(), []int{3}
}

func (x *EventPaymentClosed) GetStatus() StatusPayment {
	if x != nil {
		return x.Status
	}
	return StatusPayment_STATUS_PAYMENT_UNSPECIFIED
}

var File_domain_billing_payment_v1_event_proto protoreflect.FileDescriptor

var file_domain_billing_payment_v1_event_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x27, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x13,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x58, 0x0a, 0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x58, 0x0a, 0x14,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x56, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x8b,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x52,
	0x4f, 0x56, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x42, 0x52, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a,
	0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_billing_payment_v1_event_proto_rawDescOnce sync.Once
	file_domain_billing_payment_v1_event_proto_rawDescData = file_domain_billing_payment_v1_event_proto_rawDesc
)

func file_domain_billing_payment_v1_event_proto_rawDescGZIP() []byte {
	file_domain_billing_payment_v1_event_proto_rawDescOnce.Do(func() {
		file_domain_billing_payment_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_billing_payment_v1_event_proto_rawDescData)
	})
	return file_domain_billing_payment_v1_event_proto_rawDescData
}

var file_domain_billing_payment_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_billing_payment_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_domain_billing_payment_v1_event_proto_goTypes = []interface{}{
	(Event)(0),                   // 0: domain.billing.payment.v1.Event
	(*EventPaymentCreated)(nil),  // 1: domain.billing.payment.v1.EventPaymentCreated
	(*EventPaymentApproved)(nil), // 2: domain.billing.payment.v1.EventPaymentApproved
	(*EventPaymentRejected)(nil), // 3: domain.billing.payment.v1.EventPaymentRejected
	(*EventPaymentClosed)(nil),   // 4: domain.billing.payment.v1.EventPaymentClosed
	(StatusPayment)(0),           // 5: domain.billing.payment.v1.StatusPayment
}
var file_domain_billing_payment_v1_event_proto_depIdxs = []int32{
	5, // 0: domain.billing.payment.v1.EventPaymentCreated.status:type_name -> domain.billing.payment.v1.StatusPayment
	5, // 1: domain.billing.payment.v1.EventPaymentApproved.status:type_name -> domain.billing.payment.v1.StatusPayment
	5, // 2: domain.billing.payment.v1.EventPaymentRejected.status:type_name -> domain.billing.payment.v1.StatusPayment
	5, // 3: domain.billing.payment.v1.EventPaymentClosed.status:type_name -> domain.billing.payment.v1.StatusPayment
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_domain_billing_payment_v1_event_proto_init() }
func file_domain_billing_payment_v1_event_proto_init() {
	if File_domain_billing_payment_v1_event_proto != nil {
		return
	}
	file_domain_billing_payment_v1_payment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_domain_billing_payment_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPaymentCreated); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_domain_billing_payment_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPaymentApproved); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_domain_billing_payment_v1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPaymentRejected); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_domain_billing_payment_v1_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPaymentClosed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_billing_payment_v1_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_billing_payment_v1_event_proto_goTypes,
		DependencyIndexes: file_domain_billing_payment_v1_event_proto_depIdxs,
		EnumInfos:         file_domain_billing_payment_v1_event_proto_enumTypes,
		MessageInfos:      file_domain_billing_payment_v1_event_proto_msgTypes,
	}.Build()
	File_domain_billing_payment_v1_event_proto = out.File
	file_domain_billing_payment_v1_event_proto_rawDesc = nil
	file_domain_billing_payment_v1_event_proto_goTypes = nil
	file_domain_billing_payment_v1_event_proto_depIdxs = nil
}
