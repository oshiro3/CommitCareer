// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        (unknown)
// source: proto/commit_career.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FetchBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *FetchBranchRequest) Reset() {
	*x = FetchBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commit_career_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchBranchRequest) ProtoMessage() {}

func (x *FetchBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commit_career_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchBranchRequest.ProtoReflect.Descriptor instead.
func (*FetchBranchRequest) Descriptor() ([]byte, []int) {
	return file_proto_commit_career_proto_rawDescGZIP(), []int{0}
}

func (x *FetchBranchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type FetchCommitsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Branch string `protobuf:"bytes,1,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *FetchCommitsRequest) Reset() {
	*x = FetchCommitsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commit_career_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchCommitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchCommitsRequest) ProtoMessage() {}

func (x *FetchCommitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commit_career_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchCommitsRequest.ProtoReflect.Descriptor instead.
func (*FetchCommitsRequest) Descriptor() ([]byte, []int) {
	return file_proto_commit_career_proto_rawDescGZIP(), []int{1}
}

func (x *FetchCommitsRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

type CherryPickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base string   `protobuf:"bytes,1,opt,name=Base,json=base,proto3" json:"Base,omitempty"`
	Sha  []string `protobuf:"bytes,2,rep,name=Sha,json=sha,proto3" json:"Sha,omitempty"`
}

func (x *CherryPickRequest) Reset() {
	*x = CherryPickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commit_career_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CherryPickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CherryPickRequest) ProtoMessage() {}

func (x *CherryPickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commit_career_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CherryPickRequest.ProtoReflect.Descriptor instead.
func (*CherryPickRequest) Descriptor() ([]byte, []int) {
	return file_proto_commit_career_proto_rawDescGZIP(), []int{2}
}

func (x *CherryPickRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *CherryPickRequest) GetSha() []string {
	if x != nil {
		return x.Sha
	}
	return nil
}

type FetchBranchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name []string `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
}

func (x *FetchBranchResponse) Reset() {
	*x = FetchBranchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commit_career_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchBranchResponse) ProtoMessage() {}

func (x *FetchBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commit_career_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchBranchResponse.ProtoReflect.Descriptor instead.
func (*FetchBranchResponse) Descriptor() ([]byte, []int) {
	return file_proto_commit_career_proto_rawDescGZIP(), []int{3}
}

func (x *FetchBranchResponse) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

type FetchCommitsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commits []*FetchCommitsResponse_Commit `protobuf:"bytes,1,rep,name=commits,proto3" json:"commits,omitempty"`
}

func (x *FetchCommitsResponse) Reset() {
	*x = FetchCommitsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commit_career_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchCommitsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchCommitsResponse) ProtoMessage() {}

func (x *FetchCommitsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commit_career_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchCommitsResponse.ProtoReflect.Descriptor instead.
func (*FetchCommitsResponse) Descriptor() ([]byte, []int) {
	return file_proto_commit_career_proto_rawDescGZIP(), []int{4}
}

func (x *FetchCommitsResponse) GetCommits() []*FetchCommitsResponse_Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

type CherryPickResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (x *CherryPickResponse) Reset() {
	*x = CherryPickResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commit_career_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CherryPickResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CherryPickResponse) ProtoMessage() {}

func (x *CherryPickResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commit_career_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CherryPickResponse.ProtoReflect.Descriptor instead.
func (*CherryPickResponse) Descriptor() ([]byte, []int) {
	return file_proto_commit_career_proto_rawDescGZIP(), []int{5}
}

func (x *CherryPickResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CherryPickResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FetchCommitsResponse_Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha   string `protobuf:"bytes,1,opt,name=sha,proto3" json:"sha,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *FetchCommitsResponse_Commit) Reset() {
	*x = FetchCommitsResponse_Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_commit_career_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchCommitsResponse_Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchCommitsResponse_Commit) ProtoMessage() {}

func (x *FetchCommitsResponse_Commit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commit_career_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchCommitsResponse_Commit.ProtoReflect.Descriptor instead.
func (*FetchCommitsResponse_Commit) Descriptor() ([]byte, []int) {
	return file_proto_commit_career_proto_rawDescGZIP(), []int{4, 0}
}

func (x *FetchCommitsResponse_Commit) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

func (x *FetchCommitsResponse_Commit) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_proto_commit_career_proto protoreflect.FileDescriptor

var file_proto_commit_career_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x63,
	0x61, 0x72, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x2d,
	0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x39, 0x0a,
	0x11, 0x43, 0x68, 0x65, 0x72, 0x72, 0x79, 0x50, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x68, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x03, 0x73, 0x68, 0x61, 0x22, 0x29, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x1a, 0x30, 0x0a, 0x06, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x68, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x12,
	0x43, 0x68, 0x65, 0x72, 0x72, 0x79, 0x50, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xe8, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x43,
	0x61, 0x72, 0x65, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0d, 0x66, 0x65, 0x74, 0x63, 0x68, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x49, 0x0a, 0x0c, 0x66, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x12,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x68,
	0x65, 0x72, 0x72, 0x79, 0x50, 0x69, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x65, 0x72, 0x72, 0x79, 0x50, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x72, 0x72,
	0x79, 0x50, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x07, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_commit_career_proto_rawDescOnce sync.Once
	file_proto_commit_career_proto_rawDescData = file_proto_commit_career_proto_rawDesc
)

func file_proto_commit_career_proto_rawDescGZIP() []byte {
	file_proto_commit_career_proto_rawDescOnce.Do(func() {
		file_proto_commit_career_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_commit_career_proto_rawDescData)
	})
	return file_proto_commit_career_proto_rawDescData
}

var file_proto_commit_career_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_commit_career_proto_goTypes = []interface{}{
	(*FetchBranchRequest)(nil),          // 0: proto.FetchBranchRequest
	(*FetchCommitsRequest)(nil),         // 1: proto.FetchCommitsRequest
	(*CherryPickRequest)(nil),           // 2: proto.CherryPickRequest
	(*FetchBranchResponse)(nil),         // 3: proto.FetchBranchResponse
	(*FetchCommitsResponse)(nil),        // 4: proto.FetchCommitsResponse
	(*CherryPickResponse)(nil),          // 5: proto.CherryPickResponse
	(*FetchCommitsResponse_Commit)(nil), // 6: proto.FetchCommitsResponse.Commit
}
var file_proto_commit_career_proto_depIdxs = []int32{
	6, // 0: proto.FetchCommitsResponse.commits:type_name -> proto.FetchCommitsResponse.Commit
	0, // 1: proto.CommitCareer.fetchBranches:input_type -> proto.FetchBranchRequest
	1, // 2: proto.CommitCareer.fetchCommits:input_type -> proto.FetchCommitsRequest
	2, // 3: proto.CommitCareer.cherryPick:input_type -> proto.CherryPickRequest
	3, // 4: proto.CommitCareer.fetchBranches:output_type -> proto.FetchBranchResponse
	4, // 5: proto.CommitCareer.fetchCommits:output_type -> proto.FetchCommitsResponse
	5, // 6: proto.CommitCareer.cherryPick:output_type -> proto.CherryPickResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_commit_career_proto_init() }
func file_proto_commit_career_proto_init() {
	if File_proto_commit_career_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_commit_career_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchBranchRequest); i {
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
		file_proto_commit_career_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchCommitsRequest); i {
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
		file_proto_commit_career_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CherryPickRequest); i {
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
		file_proto_commit_career_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchBranchResponse); i {
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
		file_proto_commit_career_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchCommitsResponse); i {
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
		file_proto_commit_career_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CherryPickResponse); i {
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
		file_proto_commit_career_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchCommitsResponse_Commit); i {
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
			RawDescriptor: file_proto_commit_career_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_commit_career_proto_goTypes,
		DependencyIndexes: file_proto_commit_career_proto_depIdxs,
		MessageInfos:      file_proto_commit_career_proto_msgTypes,
	}.Build()
	File_proto_commit_career_proto = out.File
	file_proto_commit_career_proto_rawDesc = nil
	file_proto_commit_career_proto_goTypes = nil
	file_proto_commit_career_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommitCareerClient is the client API for CommitCareer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommitCareerClient interface {
	FetchBranches(ctx context.Context, in *FetchBranchRequest, opts ...grpc.CallOption) (*FetchBranchResponse, error)
	FetchCommits(ctx context.Context, in *FetchCommitsRequest, opts ...grpc.CallOption) (*FetchCommitsResponse, error)
	CherryPick(ctx context.Context, in *CherryPickRequest, opts ...grpc.CallOption) (*CherryPickResponse, error)
}

type commitCareerClient struct {
	cc grpc.ClientConnInterface
}

func NewCommitCareerClient(cc grpc.ClientConnInterface) CommitCareerClient {
	return &commitCareerClient{cc}
}

func (c *commitCareerClient) FetchBranches(ctx context.Context, in *FetchBranchRequest, opts ...grpc.CallOption) (*FetchBranchResponse, error) {
	out := new(FetchBranchResponse)
	err := c.cc.Invoke(ctx, "/proto.CommitCareer/fetchBranches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commitCareerClient) FetchCommits(ctx context.Context, in *FetchCommitsRequest, opts ...grpc.CallOption) (*FetchCommitsResponse, error) {
	out := new(FetchCommitsResponse)
	err := c.cc.Invoke(ctx, "/proto.CommitCareer/fetchCommits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commitCareerClient) CherryPick(ctx context.Context, in *CherryPickRequest, opts ...grpc.CallOption) (*CherryPickResponse, error) {
	out := new(CherryPickResponse)
	err := c.cc.Invoke(ctx, "/proto.CommitCareer/cherryPick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommitCareerServer is the server API for CommitCareer service.
type CommitCareerServer interface {
	FetchBranches(context.Context, *FetchBranchRequest) (*FetchBranchResponse, error)
	FetchCommits(context.Context, *FetchCommitsRequest) (*FetchCommitsResponse, error)
	CherryPick(context.Context, *CherryPickRequest) (*CherryPickResponse, error)
}

// UnimplementedCommitCareerServer can be embedded to have forward compatible implementations.
type UnimplementedCommitCareerServer struct {
}

func (*UnimplementedCommitCareerServer) FetchBranches(context.Context, *FetchBranchRequest) (*FetchBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchBranches not implemented")
}
func (*UnimplementedCommitCareerServer) FetchCommits(context.Context, *FetchCommitsRequest) (*FetchCommitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCommits not implemented")
}
func (*UnimplementedCommitCareerServer) CherryPick(context.Context, *CherryPickRequest) (*CherryPickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CherryPick not implemented")
}

func RegisterCommitCareerServer(s *grpc.Server, srv CommitCareerServer) {
	s.RegisterService(&_CommitCareer_serviceDesc, srv)
}

func _CommitCareer_FetchBranches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitCareerServer).FetchBranches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CommitCareer/FetchBranches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitCareerServer).FetchBranches(ctx, req.(*FetchBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommitCareer_FetchCommits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCommitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitCareerServer).FetchCommits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CommitCareer/FetchCommits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitCareerServer).FetchCommits(ctx, req.(*FetchCommitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommitCareer_CherryPick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CherryPickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitCareerServer).CherryPick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CommitCareer/CherryPick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitCareerServer).CherryPick(ctx, req.(*CherryPickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommitCareer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CommitCareer",
	HandlerType: (*CommitCareerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "fetchBranches",
			Handler:    _CommitCareer_FetchBranches_Handler,
		},
		{
			MethodName: "fetchCommits",
			Handler:    _CommitCareer_FetchCommits_Handler,
		},
		{
			MethodName: "cherryPick",
			Handler:    _CommitCareer_CherryPick_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/commit_career.proto",
}
