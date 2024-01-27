// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: proto/grps.proto

package grps

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

// youtuba
type RequestGetData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdUser int32       `protobuf:"varint,1,opt,name=idUser,proto3" json:"idUser,omitempty"`
	Url    []*Arrayurl `protobuf:"bytes,2,rep,name=url,proto3" json:"url,omitempty"`
}

func (x *RequestGetData) Reset() {
	*x = RequestGetData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGetData) ProtoMessage() {}

func (x *RequestGetData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGetData.ProtoReflect.Descriptor instead.
func (*RequestGetData) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{0}
}

func (x *RequestGetData) GetIdUser() int32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *RequestGetData) GetUrl() []*Arrayurl {
	if x != nil {
		return x.Url
	}
	return nil
}

type ResponseStreamData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdUser   int32       `protobuf:"varint,1,opt,name=idUser,proto3" json:"idUser,omitempty"`
	Urlvideo []*Arrayurl `protobuf:"bytes,2,rep,name=urlvideo,proto3" json:"urlvideo,omitempty"`
}

func (x *ResponseStreamData) Reset() {
	*x = ResponseStreamData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseStreamData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseStreamData) ProtoMessage() {}

func (x *ResponseStreamData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseStreamData.ProtoReflect.Descriptor instead.
func (*ResponseStreamData) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseStreamData) GetIdUser() int32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *ResponseStreamData) GetUrlvideo() []*Arrayurl {
	if x != nil {
		return x.Urlvideo
	}
	return nil
}

type Arrayurl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url         string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Dateservice string `protobuf:"bytes,3,opt,name=dateservice,proto3" json:"dateservice,omitempty"`
	Datascan    string `protobuf:"bytes,4,opt,name=datascan,proto3" json:"datascan,omitempty"`
}

func (x *Arrayurl) Reset() {
	*x = Arrayurl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Arrayurl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Arrayurl) ProtoMessage() {}

func (x *Arrayurl) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Arrayurl.ProtoReflect.Descriptor instead.
func (*Arrayurl) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{2}
}

func (x *Arrayurl) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Arrayurl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Arrayurl) GetDateservice() string {
	if x != nil {
		return x.Dateservice
	}
	return ""
}

func (x *Arrayurl) GetDatascan() string {
	if x != nil {
		return x.Datascan
	}
	return ""
}

// telegram
type RequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdUser      int32    `protobuf:"varint,1,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
	NameChannel []string `protobuf:"bytes,2,rep,name=name_channel,json=nameChannel,proto3" json:"name_channel,omitempty"`
}

func (x *RequestData) Reset() {
	*x = RequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestData) ProtoMessage() {}

func (x *RequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestData.ProtoReflect.Descriptor instead.
func (*RequestData) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{3}
}

func (x *RequestData) GetIdUser() int32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *RequestData) GetNameChannel() []string {
	if x != nil {
		return x.NameChannel
	}
	return nil
}

type ResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdData int32        `protobuf:"varint,1,opt,name=id_data,json=idData,proto3" json:"id_data,omitempty"`
	Post   []*ArrayPost `protobuf:"bytes,2,rep,name=post,proto3" json:"post,omitempty"`
}

func (x *ResponseData) Reset() {
	*x = ResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseData) ProtoMessage() {}

func (x *ResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseData.ProtoReflect.Descriptor instead.
func (*ResponseData) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseData) GetIdData() int32 {
	if x != nil {
		return x.IdData
	}
	return 0
}

func (x *ResponseData) GetPost() []*ArrayPost {
	if x != nil {
		return x.Post
	}
	return nil
}

type ArrayPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel_Name string  `protobuf:"bytes,1,opt,name=Channel_Name,json=ChannelName,proto3" json:"Channel_Name,omitempty"`
	Text         string  `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	Img          []*IMGS `protobuf:"bytes,3,rep,name=img,proto3" json:"img,omitempty"`
	CreatedAt    string  `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	GetAt        string  `protobuf:"bytes,5,opt,name=GetAt,proto3" json:"GetAt,omitempty"`
	LogoLink     string  `protobuf:"bytes,6,opt,name=LogoLink,proto3" json:"LogoLink,omitempty"`
}

func (x *ArrayPost) Reset() {
	*x = ArrayPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayPost) ProtoMessage() {}

func (x *ArrayPost) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayPost.ProtoReflect.Descriptor instead.
func (*ArrayPost) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{5}
}

func (x *ArrayPost) GetChannel_Name() string {
	if x != nil {
		return x.Channel_Name
	}
	return ""
}

func (x *ArrayPost) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ArrayPost) GetImg() []*IMGS {
	if x != nil {
		return x.Img
	}
	return nil
}

func (x *ArrayPost) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ArrayPost) GetGetAt() string {
	if x != nil {
		return x.GetAt
	}
	return ""
}

func (x *ArrayPost) GetLogoLink() string {
	if x != nil {
		return x.LogoLink
	}
	return ""
}

type IMGS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlImg string `protobuf:"bytes,2,opt,name=url_img,json=urlImg,proto3" json:"url_img,omitempty"`
}

func (x *IMGS) Reset() {
	*x = IMGS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMGS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMGS) ProtoMessage() {}

func (x *IMGS) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMGS.ProtoReflect.Descriptor instead.
func (*IMGS) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{6}
}

func (x *IMGS) GetUrlImg() string {
	if x != nil {
		return x.UrlImg
	}
	return ""
}

type RequestDataReddit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdUser int32    `protobuf:"varint,1,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
	Link   []string `protobuf:"bytes,2,rep,name=link,proto3" json:"link,omitempty"`
}

func (x *RequestDataReddit) Reset() {
	*x = RequestDataReddit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDataReddit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDataReddit) ProtoMessage() {}

func (x *RequestDataReddit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDataReddit.ProtoReflect.Descriptor instead.
func (*RequestDataReddit) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{7}
}

func (x *RequestDataReddit) GetIdUser() int32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *RequestDataReddit) GetLink() []string {
	if x != nil {
		return x.Link
	}
	return nil
}

type ResponseDataReddit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post []*ArrayPostReddit `protobuf:"bytes,1,rep,name=post,proto3" json:"post,omitempty"`
}

func (x *ResponseDataReddit) Reset() {
	*x = ResponseDataReddit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDataReddit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDataReddit) ProtoMessage() {}

func (x *ResponseDataReddit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDataReddit.ProtoReflect.Descriptor instead.
func (*ResponseDataReddit) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{8}
}

func (x *ResponseDataReddit) GetPost() []*ArrayPostReddit {
	if x != nil {
		return x.Post
	}
	return nil
}

type ArrayPostReddit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link      string  `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Text      string  `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	Img       []*IMGS `protobuf:"bytes,3,rep,name=img,proto3" json:"img,omitempty"`
	CreatedAt string  `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	GetAt     string  `protobuf:"bytes,5,opt,name=GetAt,proto3" json:"GetAt,omitempty"`
}

func (x *ArrayPostReddit) Reset() {
	*x = ArrayPostReddit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grps_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayPostReddit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayPostReddit) ProtoMessage() {}

func (x *ArrayPostReddit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grps_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayPostReddit.ProtoReflect.Descriptor instead.
func (*ArrayPostReddit) Descriptor() ([]byte, []int) {
	return file_proto_grps_proto_rawDescGZIP(), []int{9}
}

func (x *ArrayPostReddit) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *ArrayPostReddit) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ArrayPostReddit) GetImg() []*IMGS {
	if x != nil {
		return x.Img
	}
	return nil
}

func (x *ArrayPostReddit) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ArrayPostReddit) GetGetAt() string {
	if x != nil {
		return x.GetAt
	}
	return ""
}

var File_proto_grps_proto protoreflect.FileDescriptor

var file_proto_grps_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x22, 0x4e, 0x0a, 0x0e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x75, 0x72, 0x6c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x5c, 0x0a, 0x12,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x08, 0x75, 0x72,
	0x6c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4c,
	0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x75, 0x72, 0x6c,
	0x52, 0x08, 0x75, 0x72, 0x6c, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x6e, 0x0a, 0x08, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x63, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x63, 0x61, 0x6e, 0x22, 0x49, 0x0a, 0x0b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x50, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4c,
	0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0xb4, 0x01, 0x0a, 0x09, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x03,
	0x69, 0x6d, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x4c, 0x69, 0x76, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x4d, 0x47, 0x53, 0x52, 0x03, 0x69, 0x6d, 0x67, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x47, 0x65, 0x74, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x65, 0x74,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x1f,
	0x0a, 0x04, 0x49, 0x4d, 0x47, 0x53, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x72, 0x6c, 0x5f, 0x69, 0x6d,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x72, 0x6c, 0x49, 0x6d, 0x67, 0x22,
	0x40, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x22, 0x43, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x64, 0x64, 0x69, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77,
	0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x64, 0x69, 0x74,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x64, 0x64, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x4d, 0x47, 0x53, 0x52,
	0x03, 0x69, 0x6d, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x47, 0x65, 0x74, 0x41, 0x74, 0x32, 0x5a, 0x0a, 0x0e, 0x4c, 0x69, 0x76, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x48, 0x0a, 0x0e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x59, 0x6f, 0x75, 0x54, 0x75, 0x62, 0x65, 0x12, 0x18, 0x2e, 0x4c,
	0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1c, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x65,
	0x77, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x32, 0x44, 0x0a, 0x08, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d,
	0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x15, 0x2e, 0x4c, 0x69,
	0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x16, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x32, 0x54, 0x0a, 0x06, 0x52, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x12, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x64, 0x64, 0x69, 0x74, 0x47, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x64, 0x64,
	0x69, 0x74, 0x1a, 0x1c, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x64, 0x64, 0x69, 0x74,
	0x42, 0x15, 0x5a, 0x13, 0x4c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grps_proto_rawDescOnce sync.Once
	file_proto_grps_proto_rawDescData = file_proto_grps_proto_rawDesc
)

func file_proto_grps_proto_rawDescGZIP() []byte {
	file_proto_grps_proto_rawDescOnce.Do(func() {
		file_proto_grps_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grps_proto_rawDescData)
	})
	return file_proto_grps_proto_rawDescData
}

var file_proto_grps_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_grps_proto_goTypes = []interface{}{
	(*RequestGetData)(nil),     // 0: LiveView.RequestGetData
	(*ResponseStreamData)(nil), // 1: LiveView.ResponseStreamData
	(*Arrayurl)(nil),           // 2: LiveView.Arrayurl
	(*RequestData)(nil),        // 3: LiveView.RequestData
	(*ResponseData)(nil),       // 4: LiveView.ResponseData
	(*ArrayPost)(nil),          // 5: LiveView.ArrayPost
	(*IMGS)(nil),               // 6: LiveView.IMGS
	(*RequestDataReddit)(nil),  // 7: LiveView.RequestDataReddit
	(*ResponseDataReddit)(nil), // 8: LiveView.ResponseDataReddit
	(*ArrayPostReddit)(nil),    // 9: LiveView.ArrayPostReddit
}
var file_proto_grps_proto_depIdxs = []int32{
	2, // 0: LiveView.RequestGetData.url:type_name -> LiveView.Arrayurl
	2, // 1: LiveView.ResponseStreamData.urlvideo:type_name -> LiveView.Arrayurl
	5, // 2: LiveView.ResponseData.post:type_name -> LiveView.ArrayPost
	6, // 3: LiveView.ArrayPost.img:type_name -> LiveView.IMGS
	9, // 4: LiveView.ResponseDataReddit.post:type_name -> LiveView.ArrayPostReddit
	6, // 5: LiveView.ArrayPostReddit.img:type_name -> LiveView.IMGS
	0, // 6: LiveView.LiveViewReturn.ServiceYouTube:input_type -> LiveView.RequestGetData
	3, // 7: LiveView.Telegram.GetData:input_type -> LiveView.RequestData
	7, // 8: LiveView.Reddit.RedditGetData:input_type -> LiveView.RequestDataReddit
	1, // 9: LiveView.LiveViewReturn.ServiceYouTube:output_type -> LiveView.ResponseStreamData
	4, // 10: LiveView.Telegram.GetData:output_type -> LiveView.ResponseData
	8, // 11: LiveView.Reddit.RedditGetData:output_type -> LiveView.ResponseDataReddit
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_grps_proto_init() }
func file_proto_grps_proto_init() {
	if File_proto_grps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGetData); i {
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
		file_proto_grps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseStreamData); i {
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
		file_proto_grps_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Arrayurl); i {
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
		file_proto_grps_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestData); i {
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
		file_proto_grps_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseData); i {
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
		file_proto_grps_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayPost); i {
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
		file_proto_grps_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMGS); i {
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
		file_proto_grps_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDataReddit); i {
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
		file_proto_grps_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDataReddit); i {
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
		file_proto_grps_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayPostReddit); i {
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
			RawDescriptor: file_proto_grps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_proto_grps_proto_goTypes,
		DependencyIndexes: file_proto_grps_proto_depIdxs,
		MessageInfos:      file_proto_grps_proto_msgTypes,
	}.Build()
	File_proto_grps_proto = out.File
	file_proto_grps_proto_rawDesc = nil
	file_proto_grps_proto_goTypes = nil
	file_proto_grps_proto_depIdxs = nil
}
