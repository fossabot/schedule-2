// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package router

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NoParameter struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoParameter) Reset()         { *m = NoParameter{} }
func (m *NoParameter) String() string { return proto.CompactTextString(m) }
func (*NoParameter) ProtoMessage()    {}
func (*NoParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

func (m *NoParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoParameter.Unmarshal(m, b)
}
func (m *NoParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoParameter.Marshal(b, m, deterministic)
}
func (m *NoParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoParameter.Merge(m, src)
}
func (m *NoParameter) XXX_Size() int {
	return xxx_messageInfo_NoParameter.Size(m)
}
func (m *NoParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_NoParameter.DiscardUnknown(m)
}

var xxx_messageInfo_NoParameter proto.InternalMessageInfo

type Response struct {
	Error                uint32   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type EntryOption struct {
	CronTime             string   `protobuf:"bytes,1,opt,name=cron_time,json=cronTime,proto3" json:"cron_time,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Headers              []byte   `protobuf:"bytes,3,opt,name=headers,proto3" json:"headers,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryOption) Reset()         { *m = EntryOption{} }
func (m *EntryOption) String() string { return proto.CompactTextString(m) }
func (*EntryOption) ProtoMessage()    {}
func (*EntryOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}

func (m *EntryOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryOption.Unmarshal(m, b)
}
func (m *EntryOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryOption.Marshal(b, m, deterministic)
}
func (m *EntryOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryOption.Merge(m, src)
}
func (m *EntryOption) XXX_Size() int {
	return xxx_messageInfo_EntryOption.Size(m)
}
func (m *EntryOption) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryOption.DiscardUnknown(m)
}

var xxx_messageInfo_EntryOption proto.InternalMessageInfo

func (m *EntryOption) GetCronTime() string {
	if m != nil {
		return m.CronTime
	}
	return ""
}

func (m *EntryOption) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *EntryOption) GetHeaders() []byte {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *EntryOption) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type EntryOptionWithTime struct {
	CronTime             string   `protobuf:"bytes,1,opt,name=cron_time,json=cronTime,proto3" json:"cron_time,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Headers              []byte   `protobuf:"bytes,3,opt,name=headers,proto3" json:"headers,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	NextDate             int64    `protobuf:"varint,5,opt,name=next_date,json=nextDate,proto3" json:"next_date,omitempty"`
	LastDate             int64    `protobuf:"varint,6,opt,name=last_date,json=lastDate,proto3" json:"last_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryOptionWithTime) Reset()         { *m = EntryOptionWithTime{} }
func (m *EntryOptionWithTime) String() string { return proto.CompactTextString(m) }
func (*EntryOptionWithTime) ProtoMessage()    {}
func (*EntryOptionWithTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}

func (m *EntryOptionWithTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryOptionWithTime.Unmarshal(m, b)
}
func (m *EntryOptionWithTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryOptionWithTime.Marshal(b, m, deterministic)
}
func (m *EntryOptionWithTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryOptionWithTime.Merge(m, src)
}
func (m *EntryOptionWithTime) XXX_Size() int {
	return xxx_messageInfo_EntryOptionWithTime.Size(m)
}
func (m *EntryOptionWithTime) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryOptionWithTime.DiscardUnknown(m)
}

var xxx_messageInfo_EntryOptionWithTime proto.InternalMessageInfo

func (m *EntryOptionWithTime) GetCronTime() string {
	if m != nil {
		return m.CronTime
	}
	return ""
}

func (m *EntryOptionWithTime) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *EntryOptionWithTime) GetHeaders() []byte {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *EntryOptionWithTime) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *EntryOptionWithTime) GetNextDate() int64 {
	if m != nil {
		return m.NextDate
	}
	return 0
}

func (m *EntryOptionWithTime) GetLastDate() int64 {
	if m != nil {
		return m.LastDate
	}
	return 0
}

type Information struct {
	Identity             string                          `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Start                bool                            `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	TimeZone             string                          `protobuf:"bytes,3,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	Entries              map[string]*EntryOptionWithTime `protobuf:"bytes,4,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Information) Reset()         { *m = Information{} }
func (m *Information) String() string { return proto.CompactTextString(m) }
func (*Information) ProtoMessage()    {}
func (*Information) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{4}
}

func (m *Information) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Information.Unmarshal(m, b)
}
func (m *Information) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Information.Marshal(b, m, deterministic)
}
func (m *Information) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Information.Merge(m, src)
}
func (m *Information) XXX_Size() int {
	return xxx_messageInfo_Information.Size(m)
}
func (m *Information) XXX_DiscardUnknown() {
	xxx_messageInfo_Information.DiscardUnknown(m)
}

var xxx_messageInfo_Information proto.InternalMessageInfo

func (m *Information) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *Information) GetStart() bool {
	if m != nil {
		return m.Start
	}
	return false
}

func (m *Information) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *Information) GetEntries() map[string]*EntryOptionWithTime {
	if m != nil {
		return m.Entries
	}
	return nil
}

type GetParameter struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParameter) Reset()         { *m = GetParameter{} }
func (m *GetParameter) String() string { return proto.CompactTextString(m) }
func (*GetParameter) ProtoMessage()    {}
func (*GetParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{5}
}

func (m *GetParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParameter.Unmarshal(m, b)
}
func (m *GetParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParameter.Marshal(b, m, deterministic)
}
func (m *GetParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParameter.Merge(m, src)
}
func (m *GetParameter) XXX_Size() int {
	return xxx_messageInfo_GetParameter.Size(m)
}
func (m *GetParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParameter.DiscardUnknown(m)
}

var xxx_messageInfo_GetParameter proto.InternalMessageInfo

func (m *GetParameter) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

type GetResponse struct {
	Error                uint32       `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg                  string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *Information `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *GetResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetResponse) GetData() *Information {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListsParameter struct {
	Identity             []string `protobuf:"bytes,1,rep,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListsParameter) Reset()         { *m = ListsParameter{} }
func (m *ListsParameter) String() string { return proto.CompactTextString(m) }
func (*ListsParameter) ProtoMessage()    {}
func (*ListsParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{7}
}

func (m *ListsParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListsParameter.Unmarshal(m, b)
}
func (m *ListsParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListsParameter.Marshal(b, m, deterministic)
}
func (m *ListsParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListsParameter.Merge(m, src)
}
func (m *ListsParameter) XXX_Size() int {
	return xxx_messageInfo_ListsParameter.Size(m)
}
func (m *ListsParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListsParameter.DiscardUnknown(m)
}

var xxx_messageInfo_ListsParameter proto.InternalMessageInfo

func (m *ListsParameter) GetIdentity() []string {
	if m != nil {
		return m.Identity
	}
	return nil
}

type ListsResponse struct {
	Error                uint32         `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg                  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 []*Information `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListsResponse) Reset()         { *m = ListsResponse{} }
func (m *ListsResponse) String() string { return proto.CompactTextString(m) }
func (*ListsResponse) ProtoMessage()    {}
func (*ListsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{8}
}

func (m *ListsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListsResponse.Unmarshal(m, b)
}
func (m *ListsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListsResponse.Marshal(b, m, deterministic)
}
func (m *ListsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListsResponse.Merge(m, src)
}
func (m *ListsResponse) XXX_Size() int {
	return xxx_messageInfo_ListsResponse.Size(m)
}
func (m *ListsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListsResponse proto.InternalMessageInfo

func (m *ListsResponse) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *ListsResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ListsResponse) GetData() []*Information {
	if m != nil {
		return m.Data
	}
	return nil
}

type AllResponse struct {
	Error                uint32   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 []string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllResponse) Reset()         { *m = AllResponse{} }
func (m *AllResponse) String() string { return proto.CompactTextString(m) }
func (*AllResponse) ProtoMessage()    {}
func (*AllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{9}
}

func (m *AllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllResponse.Unmarshal(m, b)
}
func (m *AllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllResponse.Marshal(b, m, deterministic)
}
func (m *AllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllResponse.Merge(m, src)
}
func (m *AllResponse) XXX_Size() int {
	return xxx_messageInfo_AllResponse.Size(m)
}
func (m *AllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllResponse proto.InternalMessageInfo

func (m *AllResponse) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *AllResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AllResponse) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

type PutParameter struct {
	Identity             string                  `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	TimeZone             string                  `protobuf:"bytes,2,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	Start                bool                    `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	Entries              map[string]*EntryOption `protobuf:"bytes,4,rep,name=entries,proto3" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PutParameter) Reset()         { *m = PutParameter{} }
func (m *PutParameter) String() string { return proto.CompactTextString(m) }
func (*PutParameter) ProtoMessage()    {}
func (*PutParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{10}
}

func (m *PutParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutParameter.Unmarshal(m, b)
}
func (m *PutParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutParameter.Marshal(b, m, deterministic)
}
func (m *PutParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutParameter.Merge(m, src)
}
func (m *PutParameter) XXX_Size() int {
	return xxx_messageInfo_PutParameter.Size(m)
}
func (m *PutParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PutParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PutParameter proto.InternalMessageInfo

func (m *PutParameter) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *PutParameter) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *PutParameter) GetStart() bool {
	if m != nil {
		return m.Start
	}
	return false
}

func (m *PutParameter) GetEntries() map[string]*EntryOption {
	if m != nil {
		return m.Entries
	}
	return nil
}

type DeleteParameter struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteParameter) Reset()         { *m = DeleteParameter{} }
func (m *DeleteParameter) String() string { return proto.CompactTextString(m) }
func (*DeleteParameter) ProtoMessage()    {}
func (*DeleteParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{11}
}

func (m *DeleteParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteParameter.Unmarshal(m, b)
}
func (m *DeleteParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteParameter.Marshal(b, m, deterministic)
}
func (m *DeleteParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteParameter.Merge(m, src)
}
func (m *DeleteParameter) XXX_Size() int {
	return xxx_messageInfo_DeleteParameter.Size(m)
}
func (m *DeleteParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteParameter.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteParameter proto.InternalMessageInfo

func (m *DeleteParameter) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

type RunningParameter struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Running              bool     `protobuf:"varint,2,opt,name=running,proto3" json:"running,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunningParameter) Reset()         { *m = RunningParameter{} }
func (m *RunningParameter) String() string { return proto.CompactTextString(m) }
func (*RunningParameter) ProtoMessage()    {}
func (*RunningParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{12}
}

func (m *RunningParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunningParameter.Unmarshal(m, b)
}
func (m *RunningParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunningParameter.Marshal(b, m, deterministic)
}
func (m *RunningParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunningParameter.Merge(m, src)
}
func (m *RunningParameter) XXX_Size() int {
	return xxx_messageInfo_RunningParameter.Size(m)
}
func (m *RunningParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_RunningParameter.DiscardUnknown(m)
}

var xxx_messageInfo_RunningParameter proto.InternalMessageInfo

func (m *RunningParameter) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *RunningParameter) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func init() {
	proto.RegisterType((*NoParameter)(nil), "NoParameter")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*EntryOption)(nil), "EntryOption")
	proto.RegisterType((*EntryOptionWithTime)(nil), "EntryOptionWithTime")
	proto.RegisterType((*Information)(nil), "Information")
	proto.RegisterMapType((map[string]*EntryOptionWithTime)(nil), "Information.EntriesEntry")
	proto.RegisterType((*GetParameter)(nil), "GetParameter")
	proto.RegisterType((*GetResponse)(nil), "GetResponse")
	proto.RegisterType((*ListsParameter)(nil), "ListsParameter")
	proto.RegisterType((*ListsResponse)(nil), "ListsResponse")
	proto.RegisterType((*AllResponse)(nil), "AllResponse")
	proto.RegisterType((*PutParameter)(nil), "PutParameter")
	proto.RegisterMapType((map[string]*EntryOption)(nil), "PutParameter.EntriesEntry")
	proto.RegisterType((*DeleteParameter)(nil), "DeleteParameter")
	proto.RegisterType((*RunningParameter)(nil), "RunningParameter")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6f, 0xd4, 0x3c,
	0x10, 0xc6, 0x9b, 0x66, 0xbb, 0x9b, 0x4c, 0xb2, 0x6d, 0x5f, 0xbf, 0x3d, 0x84, 0x70, 0x09, 0x06,
	0x89, 0xa5, 0x02, 0x1f, 0xb6, 0x1c, 0x10, 0xb7, 0x4a, 0xad, 0xda, 0x4a, 0x08, 0x56, 0x16, 0x52,
	0x05, 0x97, 0x2a, 0x65, 0x87, 0x36, 0x90, 0xb5, 0x2b, 0xc7, 0x41, 0x94, 0xcf, 0xc4, 0x57, 0xe2,
	0xc4, 0x77, 0xe0, 0x8c, 0xec, 0x6c, 0x2a, 0x67, 0xc5, 0xc2, 0x56, 0x88, 0x5b, 0xe6, 0x8f, 0x67,
	0x1e, 0x3f, 0xfa, 0xc5, 0x10, 0x2b, 0x59, 0x6b, 0x54, 0xec, 0x4a, 0x49, 0x2d, 0xe9, 0x10, 0xa2,
	0x97, 0x72, 0x92, 0xab, 0x7c, 0x86, 0x1a, 0x15, 0x1d, 0x43, 0xc0, 0xb1, 0xba, 0x92, 0xa2, 0x42,
	0xb2, 0x03, 0x1b, 0xa8, 0x94, 0x54, 0x89, 0x97, 0x79, 0xa3, 0x21, 0x6f, 0x02, 0xb2, 0x0d, 0xfe,
	0xac, 0xba, 0x48, 0xd6, 0x33, 0x6f, 0x14, 0x72, 0xf3, 0x49, 0x3f, 0x40, 0x74, 0x28, 0xb4, 0xba,
	0x7e, 0x75, 0xa5, 0x0b, 0x29, 0xc8, 0x5d, 0x08, 0xdf, 0x29, 0x29, 0xce, 0x74, 0x31, 0x43, 0x7b,
	0x34, 0xe4, 0x81, 0x49, 0xbc, 0x2e, 0x66, 0x68, 0x4e, 0xd7, 0xaa, 0x6c, 0x4f, 0xd7, 0xaa, 0x24,
	0x09, 0x0c, 0x2e, 0x31, 0x9f, 0xa2, 0xaa, 0x12, 0x3f, 0xf3, 0x46, 0x31, 0x6f, 0x43, 0x42, 0xa0,
	0x77, 0x2e, 0xa7, 0xd7, 0x49, 0xcf, 0xa6, 0xed, 0x37, 0xfd, 0xea, 0xc1, 0xff, 0xce, 0xb2, 0xd3,
	0x42, 0x5f, 0xda, 0xb9, 0xff, 0x72, 0xa9, 0x19, 0x2e, 0xf0, 0xb3, 0x3e, 0x9b, 0xe6, 0x1a, 0x93,
	0x8d, 0xcc, 0x1b, 0xf9, 0x3c, 0x30, 0x89, 0x83, 0x5c, 0xdb, 0xcd, 0x65, 0x5e, 0xcd, 0x8b, 0xfd,
	0xa6, 0x68, 0x12, 0xa6, 0x48, 0xbf, 0x7b, 0x10, 0x9d, 0x88, 0xf7, 0x52, 0xcd, 0x72, 0xeb, 0x4d,
	0x0a, 0x41, 0x31, 0x45, 0xa1, 0x0b, 0x7d, 0xdd, 0xaa, 0x6c, 0x63, 0x63, 0x77, 0xa5, 0x73, 0xa5,
	0xad, 0xce, 0x80, 0x37, 0x81, 0x19, 0x6f, 0xee, 0x74, 0xf6, 0x45, 0x0a, 0xb4, 0x5a, 0x43, 0x1e,
	0x98, 0xc4, 0x5b, 0x29, 0x90, 0xec, 0xc1, 0x00, 0x85, 0x56, 0x05, 0x56, 0x49, 0x2f, 0xf3, 0x47,
	0xd1, 0xf8, 0x0e, 0x73, 0xb6, 0xb1, 0xc3, 0xa6, 0x66, 0xfd, 0xe2, 0x6d, 0x67, 0x3a, 0x81, 0xd8,
	0x2d, 0x18, 0x77, 0x3e, 0x62, 0x2b, 0xc7, 0x7c, 0x92, 0x5d, 0xd8, 0xf8, 0x94, 0x97, 0x35, 0x5a,
	0x25, 0xd1, 0x78, 0x87, 0xfd, 0xc2, 0x71, 0xde, 0xb4, 0x3c, 0x5f, 0x7f, 0xe6, 0xd1, 0x5d, 0x88,
	0x8f, 0x50, 0xdf, 0x40, 0xf4, 0xbb, 0x5b, 0xd2, 0x53, 0x88, 0x8e, 0x50, 0xdf, 0x96, 0x31, 0x92,
	0x41, 0x6f, 0x9a, 0xeb, 0xdc, 0x3a, 0x10, 0x8d, 0x63, 0xf7, 0x9a, 0xdc, 0x56, 0xe8, 0x63, 0xd8,
	0x7c, 0x51, 0x54, 0xba, 0x5a, 0x26, 0xc3, 0xef, 0xc8, 0x78, 0x03, 0x43, 0xdb, 0xfd, 0x17, 0x42,
	0xfc, 0x25, 0x42, 0x4e, 0x20, 0xda, 0x2f, 0xcb, 0x5b, 0x0f, 0x26, 0xce, 0xe0, 0x70, 0x3e, 0xea,
	0x9b, 0x07, 0xf1, 0xa4, 0x5e, 0xcd, 0xd9, 0x2e, 0x29, 0xeb, 0x0b, 0xa4, 0xdc, 0xc0, 0xe5, 0xbb,
	0x70, 0x3d, 0x5d, 0xe4, 0x27, 0x65, 0xee, 0xba, 0x25, 0x00, 0x1d, 0xff, 0x11, 0x20, 0xda, 0x05,
	0x28, 0x76, 0x01, 0x72, 0xc1, 0x79, 0x02, 0x5b, 0x07, 0x58, 0xa2, 0xc6, 0xd5, 0xd8, 0x39, 0x86,
	0x6d, 0x5e, 0x0b, 0x51, 0x88, 0x8b, 0xd5, 0x1c, 0x49, 0x60, 0xa0, 0x9a, 0xfe, 0xf9, 0x3f, 0xd5,
	0x86, 0xe3, 0x1f, 0x1e, 0xf4, 0xb9, 0x7d, 0x06, 0xc9, 0x03, 0xf0, 0x8f, 0x50, 0x93, 0x21, 0x73,
	0x11, 0x4e, 0x63, 0xe6, 0x50, 0x4a, 0xd7, 0xcc, 0x2f, 0x61, 0x79, 0x21, 0x5b, 0xac, 0x4b, 0x59,
	0xba, 0xc9, 0x3a, 0x20, 0xd1, 0x35, 0x72, 0x1f, 0xfc, 0xfd, 0xb2, 0x24, 0x31, 0x73, 0x1e, 0xd6,
	0x34, 0x66, 0x0e, 0x14, 0x74, 0x8d, 0xdc, 0x03, 0x7f, 0x52, 0x9b, 0xb5, 0xae, 0xe1, 0x69, 0xc8,
	0x9c, 0x96, 0x87, 0xd0, 0x6f, 0xdc, 0x21, 0xdb, 0x6c, 0xc1, 0xa6, 0x6e, 0xe3, 0x23, 0x18, 0xcc,
	0x7d, 0x21, 0xff, 0xb1, 0x45, 0x87, 0x3a, 0xad, 0xe7, 0x7d, 0xfb, 0xea, 0xef, 0xfd, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0xe7, 0xff, 0x83, 0xed, 0x05, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	Get(ctx context.Context, in *GetParameter, opts ...grpc.CallOption) (*GetResponse, error)
	Lists(ctx context.Context, in *ListsParameter, opts ...grpc.CallOption) (*ListsResponse, error)
	All(ctx context.Context, in *NoParameter, opts ...grpc.CallOption) (*AllResponse, error)
	Put(ctx context.Context, in *PutParameter, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *DeleteParameter, opts ...grpc.CallOption) (*Response, error)
	Running(ctx context.Context, in *RunningParameter, opts ...grpc.CallOption) (*Response, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Get(ctx context.Context, in *GetParameter, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/Router/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Lists(ctx context.Context, in *ListsParameter, opts ...grpc.CallOption) (*ListsResponse, error) {
	out := new(ListsResponse)
	err := c.cc.Invoke(ctx, "/Router/Lists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) All(ctx context.Context, in *NoParameter, opts ...grpc.CallOption) (*AllResponse, error) {
	out := new(AllResponse)
	err := c.cc.Invoke(ctx, "/Router/All", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Put(ctx context.Context, in *PutParameter, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Router/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Delete(ctx context.Context, in *DeleteParameter, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Router/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Running(ctx context.Context, in *RunningParameter, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Router/Running", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	Get(context.Context, *GetParameter) (*GetResponse, error)
	Lists(context.Context, *ListsParameter) (*ListsResponse, error)
	All(context.Context, *NoParameter) (*AllResponse, error)
	Put(context.Context, *PutParameter) (*Response, error)
	Delete(context.Context, *DeleteParameter) (*Response, error)
	Running(context.Context, *RunningParameter) (*Response, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) Get(ctx context.Context, req *GetParameter) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRouterServer) Lists(ctx context.Context, req *ListsParameter) (*ListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lists not implemented")
}
func (*UnimplementedRouterServer) All(ctx context.Context, req *NoParameter) (*AllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (*UnimplementedRouterServer) Put(ctx context.Context, req *PutParameter) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedRouterServer) Delete(ctx context.Context, req *DeleteParameter) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedRouterServer) Running(ctx context.Context, req *RunningParameter) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Running not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Get(ctx, req.(*GetParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Lists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListsParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Lists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router/Lists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Lists(ctx, req.(*ListsParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router/All",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).All(ctx, req.(*NoParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Put(ctx, req.(*PutParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Delete(ctx, req.(*DeleteParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Running_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunningParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Running(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Router/Running",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Running(ctx, req.(*RunningParameter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Router_Get_Handler,
		},
		{
			MethodName: "Lists",
			Handler:    _Router_Lists_Handler,
		},
		{
			MethodName: "All",
			Handler:    _Router_All_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Router_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Router_Delete_Handler,
		},
		{
			MethodName: "Running",
			Handler:    _Router_Running_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}
