// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: consensus/ibft2/proto/ibft.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HandshakeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *HandshakeResp) Reset() {
	*x = HandshakeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeResp) ProtoMessage() {}

func (x *HandshakeResp) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeResp.ProtoReflect.Descriptor instead.
func (*HandshakeResp) Descriptor() ([]byte, []int) {
	return file_consensus_ibft2_proto_ibft_proto_rawDescGZIP(), []int{0}
}

func (x *HandshakeResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type MessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*MessageReq_Preprepare
	//	*MessageReq_Prepare
	//	*MessageReq_Commit
	//	*MessageReq_RoundChange
	Message isMessageReq_Message `protobuf_oneof:"message"`
	From    string               `protobuf:"bytes,10,opt,name=from,proto3" json:"from,omitempty"`
	// commited seal if its a commit message
	Seal string `protobuf:"bytes,11,opt,name=seal,proto3" json:"seal,omitempty"`
	// signed signature
	Signature string `protobuf:"bytes,12,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *MessageReq) Reset() {
	*x = MessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReq) ProtoMessage() {}

func (x *MessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReq.ProtoReflect.Descriptor instead.
func (*MessageReq) Descriptor() ([]byte, []int) {
	return file_consensus_ibft2_proto_ibft_proto_rawDescGZIP(), []int{1}
}

func (m *MessageReq) GetMessage() isMessageReq_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *MessageReq) GetPreprepare() *Preprepare {
	if x, ok := x.GetMessage().(*MessageReq_Preprepare); ok {
		return x.Preprepare
	}
	return nil
}

func (x *MessageReq) GetPrepare() *Subject {
	if x, ok := x.GetMessage().(*MessageReq_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (x *MessageReq) GetCommit() *Subject {
	if x, ok := x.GetMessage().(*MessageReq_Commit); ok {
		return x.Commit
	}
	return nil
}

func (x *MessageReq) GetRoundChange() *Subject {
	if x, ok := x.GetMessage().(*MessageReq_RoundChange); ok {
		return x.RoundChange
	}
	return nil
}

func (x *MessageReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *MessageReq) GetSeal() string {
	if x != nil {
		return x.Seal
	}
	return ""
}

func (x *MessageReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type isMessageReq_Message interface {
	isMessageReq_Message()
}

type MessageReq_Preprepare struct {
	Preprepare *Preprepare `protobuf:"bytes,1,opt,name=preprepare,proto3,oneof"`
}

type MessageReq_Prepare struct {
	Prepare *Subject `protobuf:"bytes,2,opt,name=prepare,proto3,oneof"`
}

type MessageReq_Commit struct {
	Commit *Subject `protobuf:"bytes,3,opt,name=commit,proto3,oneof"`
}

type MessageReq_RoundChange struct {
	RoundChange *Subject `protobuf:"bytes,4,opt,name=roundChange,proto3,oneof"`
}

func (*MessageReq_Preprepare) isMessageReq_Message() {}

func (*MessageReq_Prepare) isMessageReq_Message() {}

func (*MessageReq_Commit) isMessageReq_Message() {}

func (*MessageReq_RoundChange) isMessageReq_Message() {}

type Preprepare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View     *View     `protobuf:"bytes,1,opt,name=view,proto3" json:"view,omitempty"`
	Proposal *Proposal `protobuf:"bytes,2,opt,name=proposal,proto3" json:"proposal,omitempty"`
}

func (x *Preprepare) Reset() {
	*x = Preprepare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Preprepare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preprepare) ProtoMessage() {}

func (x *Preprepare) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preprepare.ProtoReflect.Descriptor instead.
func (*Preprepare) Descriptor() ([]byte, []int) {
	return file_consensus_ibft2_proto_ibft_proto_rawDescGZIP(), []int{2}
}

func (x *Preprepare) GetView() *View {
	if x != nil {
		return x.View
	}
	return nil
}

func (x *Preprepare) GetProposal() *Proposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View   *View  `protobuf:"bytes,1,opt,name=view,proto3" json:"view,omitempty"`
	Digest string `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_consensus_ibft2_proto_ibft_proto_rawDescGZIP(), []int{3}
}

func (x *Subject) GetView() *View {
	if x != nil {
		return x.View
	}
	return nil
}

func (x *Subject) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

type View struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round    uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Sequence uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *View) Reset() {
	*x = View{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *View) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*View) ProtoMessage() {}

func (x *View) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use View.ProtoReflect.Descriptor instead.
func (*View) Descriptor() ([]byte, []int) {
	return file_consensus_ibft2_proto_ibft_proto_rawDescGZIP(), []int{4}
}

func (x *View) GetRound() uint64 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *View) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *any.Any `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_ibft2_proto_ibft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_consensus_ibft2_proto_ibft_proto_rawDescGZIP(), []int{5}
}

func (x *Proposal) GetBlock() *any.Any {
	if x != nil {
		return x.Block
	}
	return nil
}

var File_consensus_ibft2_proto_ibft_proto protoreflect.FileDescriptor

var file_consensus_ibft2_proto_ibft_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x69, 0x62, 0x66, 0x74,
	0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x62, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21,
	0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x90, 0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x30, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x70, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x12, 0x2f, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x6c, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x54, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x70, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x28, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x22, 0x3f, 0x0a, 0x07, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x04, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x36, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x12, 0x2a, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x32, 0x71, 0x0a,
	0x04, 0x49, 0x62, 0x66, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61,
	0x6b, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x18, 0x5a, 0x16, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x69,
	0x62, 0x66, 0x74, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_consensus_ibft2_proto_ibft_proto_rawDescOnce sync.Once
	file_consensus_ibft2_proto_ibft_proto_rawDescData = file_consensus_ibft2_proto_ibft_proto_rawDesc
)

func file_consensus_ibft2_proto_ibft_proto_rawDescGZIP() []byte {
	file_consensus_ibft2_proto_ibft_proto_rawDescOnce.Do(func() {
		file_consensus_ibft2_proto_ibft_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_ibft2_proto_ibft_proto_rawDescData)
	})
	return file_consensus_ibft2_proto_ibft_proto_rawDescData
}

var file_consensus_ibft2_proto_ibft_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_consensus_ibft2_proto_ibft_proto_goTypes = []interface{}{
	(*HandshakeResp)(nil), // 0: v1.HandshakeResp
	(*MessageReq)(nil),    // 1: v1.MessageReq
	(*Preprepare)(nil),    // 2: v1.Preprepare
	(*Subject)(nil),       // 3: v1.Subject
	(*View)(nil),          // 4: v1.View
	(*Proposal)(nil),      // 5: v1.Proposal
	(*any.Any)(nil),       // 6: google.protobuf.Any
	(*empty.Empty)(nil),   // 7: google.protobuf.Empty
}
var file_consensus_ibft2_proto_ibft_proto_depIdxs = []int32{
	2,  // 0: v1.MessageReq.preprepare:type_name -> v1.Preprepare
	3,  // 1: v1.MessageReq.prepare:type_name -> v1.Subject
	3,  // 2: v1.MessageReq.commit:type_name -> v1.Subject
	3,  // 3: v1.MessageReq.roundChange:type_name -> v1.Subject
	4,  // 4: v1.Preprepare.view:type_name -> v1.View
	5,  // 5: v1.Preprepare.proposal:type_name -> v1.Proposal
	4,  // 6: v1.Subject.view:type_name -> v1.View
	6,  // 7: v1.Proposal.block:type_name -> google.protobuf.Any
	7,  // 8: v1.Ibft.Handshake:input_type -> google.protobuf.Empty
	1,  // 9: v1.Ibft.Message:input_type -> v1.MessageReq
	0,  // 10: v1.Ibft.Handshake:output_type -> v1.HandshakeResp
	7,  // 11: v1.Ibft.Message:output_type -> google.protobuf.Empty
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_consensus_ibft2_proto_ibft_proto_init() }
func file_consensus_ibft2_proto_ibft_proto_init() {
	if File_consensus_ibft2_proto_ibft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consensus_ibft2_proto_ibft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeResp); i {
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
		file_consensus_ibft2_proto_ibft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReq); i {
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
		file_consensus_ibft2_proto_ibft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Preprepare); i {
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
		file_consensus_ibft2_proto_ibft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_consensus_ibft2_proto_ibft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*View); i {
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
		file_consensus_ibft2_proto_ibft_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
	file_consensus_ibft2_proto_ibft_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MessageReq_Preprepare)(nil),
		(*MessageReq_Prepare)(nil),
		(*MessageReq_Commit)(nil),
		(*MessageReq_RoundChange)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_consensus_ibft2_proto_ibft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consensus_ibft2_proto_ibft_proto_goTypes,
		DependencyIndexes: file_consensus_ibft2_proto_ibft_proto_depIdxs,
		MessageInfos:      file_consensus_ibft2_proto_ibft_proto_msgTypes,
	}.Build()
	File_consensus_ibft2_proto_ibft_proto = out.File
	file_consensus_ibft2_proto_ibft_proto_rawDesc = nil
	file_consensus_ibft2_proto_ibft_proto_goTypes = nil
	file_consensus_ibft2_proto_ibft_proto_depIdxs = nil
}
