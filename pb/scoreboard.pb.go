// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scoreboard.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type QueryHomeworkRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryHomeworkRequest) Reset()         { *m = QueryHomeworkRequest{} }
func (m *QueryHomeworkRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHomeworkRequest) ProtoMessage()    {}
func (*QueryHomeworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff350828fb116070, []int{0}
}

func (m *QueryHomeworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryHomeworkRequest.Unmarshal(m, b)
}
func (m *QueryHomeworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryHomeworkRequest.Marshal(b, m, deterministic)
}
func (m *QueryHomeworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHomeworkRequest.Merge(m, src)
}
func (m *QueryHomeworkRequest) XXX_Size() int {
	return xxx_messageInfo_QueryHomeworkRequest.Size(m)
}
func (m *QueryHomeworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHomeworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHomeworkRequest proto.InternalMessageInfo

func (m *QueryHomeworkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Homework struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Target               string        `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Runner               string        `protobuf:"bytes,3,opt,name=runner,proto3" json:"runner,omitempty"`
	Files                []*SourceFile `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`
	PenaltyTime          float64       `protobuf:"fixed64,5,opt,name=penalty_time,json=penaltyTime,proto3" json:"penalty_time,omitempty"`
	Cases                []string      `protobuf:"bytes,6,rep,name=cases,proto3" json:"cases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Homework) Reset()         { *m = Homework{} }
func (m *Homework) String() string { return proto.CompactTextString(m) }
func (*Homework) ProtoMessage()    {}
func (*Homework) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff350828fb116070, []int{1}
}

func (m *Homework) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Homework.Unmarshal(m, b)
}
func (m *Homework) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Homework.Marshal(b, m, deterministic)
}
func (m *Homework) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Homework.Merge(m, src)
}
func (m *Homework) XXX_Size() int {
	return xxx_messageInfo_Homework.Size(m)
}
func (m *Homework) XXX_DiscardUnknown() {
	xxx_messageInfo_Homework.DiscardUnknown(m)
}

var xxx_messageInfo_Homework proto.InternalMessageInfo

func (m *Homework) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Homework) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Homework) GetRunner() string {
	if m != nil {
		return m.Runner
	}
	return ""
}

func (m *Homework) GetFiles() []*SourceFile {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Homework) GetPenaltyTime() float64 {
	if m != nil {
		return m.PenaltyTime
	}
	return 0
}

func (m *Homework) GetCases() []string {
	if m != nil {
		return m.Cases
	}
	return nil
}

type SourceFile struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fallback             string   `protobuf:"bytes,2,opt,name=fallback,proto3" json:"fallback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceFile) Reset()         { *m = SourceFile{} }
func (m *SourceFile) String() string { return proto.CompactTextString(m) }
func (*SourceFile) ProtoMessage()    {}
func (*SourceFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff350828fb116070, []int{2}
}

func (m *SourceFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceFile.Unmarshal(m, b)
}
func (m *SourceFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceFile.Marshal(b, m, deterministic)
}
func (m *SourceFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceFile.Merge(m, src)
}
func (m *SourceFile) XXX_Size() int {
	return xxx_messageInfo_SourceFile.Size(m)
}
func (m *SourceFile) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceFile.DiscardUnknown(m)
}

var xxx_messageInfo_SourceFile proto.InternalMessageInfo

func (m *SourceFile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SourceFile) GetFallback() string {
	if m != nil {
		return m.Fallback
	}
	return ""
}

type SubmissionReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmissionReply) Reset()         { *m = SubmissionReply{} }
func (m *SubmissionReply) String() string { return proto.CompactTextString(m) }
func (*SubmissionReply) ProtoMessage()    {}
func (*SubmissionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff350828fb116070, []int{3}
}

func (m *SubmissionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmissionReply.Unmarshal(m, b)
}
func (m *SubmissionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmissionReply.Marshal(b, m, deterministic)
}
func (m *SubmissionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmissionReply.Merge(m, src)
}
func (m *SubmissionReply) XXX_Size() int {
	return xxx_messageInfo_SubmissionReply.Size(m)
}
func (m *SubmissionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmissionReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubmissionReply proto.InternalMessageInfo

func (m *SubmissionReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StoredSubmission struct {
	User                 string    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Results              []*Result `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StoredSubmission) Reset()         { *m = StoredSubmission{} }
func (m *StoredSubmission) String() string { return proto.CompactTextString(m) }
func (*StoredSubmission) ProtoMessage()    {}
func (*StoredSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff350828fb116070, []int{4}
}

func (m *StoredSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoredSubmission.Unmarshal(m, b)
}
func (m *StoredSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoredSubmission.Marshal(b, m, deterministic)
}
func (m *StoredSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredSubmission.Merge(m, src)
}
func (m *StoredSubmission) XXX_Size() int {
	return xxx_messageInfo_StoredSubmission.Size(m)
}
func (m *StoredSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_StoredSubmission proto.InternalMessageInfo

func (m *StoredSubmission) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *StoredSubmission) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type UserSubmission struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// string secret = 2;
	Homework             string    `protobuf:"bytes,3,opt,name=homework,proto3" json:"homework,omitempty"`
	Results              []*Result `protobuf:"bytes,4,rep,name=results,proto3" json:"results,omitempty"`
	Code                 []byte    `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserSubmission) Reset()         { *m = UserSubmission{} }
func (m *UserSubmission) String() string { return proto.CompactTextString(m) }
func (*UserSubmission) ProtoMessage()    {}
func (*UserSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff350828fb116070, []int{5}
}

func (m *UserSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSubmission.Unmarshal(m, b)
}
func (m *UserSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSubmission.Marshal(b, m, deterministic)
}
func (m *UserSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSubmission.Merge(m, src)
}
func (m *UserSubmission) XXX_Size() int {
	return xxx_messageInfo_UserSubmission.Size(m)
}
func (m *UserSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_UserSubmission proto.InternalMessageInfo

func (m *UserSubmission) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserSubmission) GetHomework() string {
	if m != nil {
		return m.Homework
	}
	return ""
}

func (m *UserSubmission) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *UserSubmission) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type Result struct {
	Case                 string   `protobuf:"bytes,1,opt,name=case,proto3" json:"case,omitempty"`
	Passed               bool     `protobuf:"varint,2,opt,name=passed,proto3" json:"passed,omitempty"`
	Time                 float64  `protobuf:"fixed64,3,opt,name=time,proto3" json:"time,omitempty"`
	Verdict              string   `protobuf:"bytes,4,opt,name=verdict,proto3" json:"verdict,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff350828fb116070, []int{6}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCase() string {
	if m != nil {
		return m.Case
	}
	return ""
}

func (m *Result) GetPassed() bool {
	if m != nil {
		return m.Passed
	}
	return false
}

func (m *Result) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Result) GetVerdict() string {
	if m != nil {
		return m.Verdict
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryHomeworkRequest)(nil), "pb.QueryHomeworkRequest")
	proto.RegisterType((*Homework)(nil), "pb.Homework")
	proto.RegisterType((*SourceFile)(nil), "pb.SourceFile")
	proto.RegisterType((*SubmissionReply)(nil), "pb.SubmissionReply")
	proto.RegisterType((*StoredSubmission)(nil), "pb.StoredSubmission")
	proto.RegisterType((*UserSubmission)(nil), "pb.UserSubmission")
	proto.RegisterType((*Result)(nil), "pb.Result")
}

func init() { proto.RegisterFile("scoreboard.proto", fileDescriptor_ff350828fb116070) }

var fileDescriptor_ff350828fb116070 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x6d, 0x76, 0xb3, 0x69, 0x3a, 0x5d, 0x4a, 0x65, 0x2a, 0x64, 0xe5, 0x14, 0xac, 0x1e, 0x22,
	0x90, 0xf6, 0xd0, 0x9e, 0x90, 0x38, 0x23, 0x0e, 0x5c, 0x70, 0xe0, 0x8c, 0x9c, 0x64, 0x5a, 0xa2,
	0x26, 0x71, 0xf0, 0x38, 0xa0, 0x45, 0xfc, 0x25, 0xfe, 0x23, 0xb2, 0xf3, 0x51, 0x56, 0x5a, 0xd4,
	0xdb, 0xbc, 0x37, 0xe3, 0x19, 0xcf, 0x9b, 0x07, 0x97, 0x54, 0x6a, 0x83, 0x85, 0x56, 0xa6, 0xda,
	0xf5, 0x46, 0x5b, 0xcd, 0x56, 0x7d, 0x21, 0x5e, 0xc3, 0xd5, 0xa7, 0x01, 0xcd, 0xfe, 0x83, 0x6e,
	0xf1, 0xa7, 0x36, 0x0f, 0x12, 0xbf, 0x0f, 0x48, 0x96, 0x31, 0x08, 0x3b, 0xd5, 0x22, 0x0f, 0xd2,
	0x20, 0x3b, 0x93, 0x3e, 0x16, 0x7f, 0x02, 0x88, 0xe7, 0xba, 0x63, 0x05, 0xec, 0x25, 0x44, 0x56,
	0x99, 0x7b, 0xb4, 0x7c, 0xe5, 0xd9, 0x09, 0x39, 0xde, 0x0c, 0x5d, 0x87, 0x86, 0xaf, 0x47, 0x7e,
	0x44, 0xec, 0x1a, 0x36, 0x77, 0x75, 0x83, 0xc4, 0xc3, 0x74, 0x9d, 0x9d, 0xdf, 0x5c, 0xec, 0xfa,
	0x62, 0x97, 0xeb, 0xc1, 0x94, 0xf8, 0xbe, 0x6e, 0x50, 0x8e, 0x49, 0xf6, 0x0a, 0xb6, 0x3d, 0x76,
	0xaa, 0xb1, 0xfb, 0xaf, 0xb6, 0x6e, 0x91, 0x6f, 0xd2, 0x20, 0x0b, 0xe4, 0xf9, 0xc4, 0x7d, 0xae,
	0x5b, 0x64, 0x57, 0xb0, 0x29, 0x15, 0x21, 0xf1, 0x28, 0x5d, 0x67, 0x67, 0x72, 0x04, 0xe2, 0x1d,
	0xc0, 0x63, 0xb7, 0xa3, 0x1f, 0x4e, 0x20, 0xbe, 0x53, 0x4d, 0x53, 0xa8, 0xf2, 0x61, 0xfa, 0xf2,
	0x82, 0xc5, 0x1b, 0x78, 0x9e, 0x0f, 0x45, 0x5b, 0x13, 0xd5, 0xba, 0x93, 0xd8, 0x37, 0x7b, 0xc6,
	0xe1, 0xb4, 0x45, 0x22, 0x75, 0x3f, 0x77, 0x99, 0xa1, 0xf8, 0x08, 0x97, 0xb9, 0xd5, 0x06, 0xab,
	0xc7, 0x27, 0x6e, 0xe0, 0x40, 0x68, 0xe6, 0x81, 0x2e, 0x66, 0xd7, 0x70, 0x6a, 0x90, 0x86, 0xc6,
	0x12, 0x5f, 0xf9, 0x9d, 0xc1, 0xed, 0x2c, 0x3d, 0x25, 0xe7, 0x94, 0xf8, 0x05, 0x17, 0x5f, 0x08,
	0xcd, 0x13, 0xbd, 0x12, 0x88, 0xbf, 0x4d, 0xd7, 0x98, 0x74, 0x5d, 0xf0, 0xbf, 0x73, 0xc2, 0xff,
	0xce, 0x71, 0x5d, 0x4b, 0x5d, 0x8d, 0x8a, 0x6e, 0xa5, 0x8f, 0x45, 0x01, 0xd1, 0x58, 0xe6, 0xb3,
	0x8a, 0x16, 0xc1, 0x5c, 0xec, 0x2e, 0xd9, 0x2b, 0x22, 0xac, 0xbc, 0x5c, 0xb1, 0x9c, 0x90, 0xab,
	0xf5, 0xb7, 0x59, 0xfb, 0xdb, 0xf8, 0xd8, 0xa9, 0xf5, 0x03, 0x4d, 0x55, 0x97, 0x96, 0x87, 0xa3,
	0x5a, 0x13, 0xbc, 0xf9, 0x0d, 0x90, 0x2f, 0x66, 0x64, 0xb7, 0x10, 0xf9, 0x4d, 0x2d, 0x63, 0xee,
	0x93, 0x87, 0x9b, 0x27, 0x2f, 0xbc, 0x29, 0x0e, 0x0f, 0x21, 0x4e, 0xd8, 0x5b, 0x78, 0x76, 0xe0,
	0x5b, 0xc6, 0x5d, 0xdd, 0x31, 0x2b, 0x27, 0x5b, 0x97, 0x99, 0x49, 0x71, 0x52, 0x44, 0xde, 0xfd,
	0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xf5, 0x03, 0x6f, 0x11, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScoreboardClient is the client API for Scoreboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScoreboardClient interface {
	Submit(ctx context.Context, in *UserSubmission, opts ...grpc.CallOption) (*SubmissionReply, error)
	QueryHomework(ctx context.Context, in *QueryHomeworkRequest, opts ...grpc.CallOption) (*Homework, error)
}

type scoreboardClient struct {
	cc *grpc.ClientConn
}

func NewScoreboardClient(cc *grpc.ClientConn) ScoreboardClient {
	return &scoreboardClient{cc}
}

func (c *scoreboardClient) Submit(ctx context.Context, in *UserSubmission, opts ...grpc.CallOption) (*SubmissionReply, error) {
	out := new(SubmissionReply)
	err := c.cc.Invoke(ctx, "/pb.Scoreboard/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreboardClient) QueryHomework(ctx context.Context, in *QueryHomeworkRequest, opts ...grpc.CallOption) (*Homework, error) {
	out := new(Homework)
	err := c.cc.Invoke(ctx, "/pb.Scoreboard/QueryHomework", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoreboardServer is the server API for Scoreboard service.
type ScoreboardServer interface {
	Submit(context.Context, *UserSubmission) (*SubmissionReply, error)
	QueryHomework(context.Context, *QueryHomeworkRequest) (*Homework, error)
}

func RegisterScoreboardServer(s *grpc.Server, srv ScoreboardServer) {
	s.RegisterService(&_Scoreboard_serviceDesc, srv)
}

func _Scoreboard_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSubmission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreboardServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Scoreboard/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreboardServer).Submit(ctx, req.(*UserSubmission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scoreboard_QueryHomework_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHomeworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreboardServer).QueryHomework(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Scoreboard/QueryHomework",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreboardServer).QueryHomework(ctx, req.(*QueryHomeworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scoreboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Scoreboard",
	HandlerType: (*ScoreboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _Scoreboard_Submit_Handler,
		},
		{
			MethodName: "QueryHomework",
			Handler:    _Scoreboard_QueryHomework_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scoreboard.proto",
}
