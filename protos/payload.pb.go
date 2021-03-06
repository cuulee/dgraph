// Code generated by protoc-gen-gogo.
// source: payload.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BackupPayload_Status int32

const (
	BackupPayload_NONE      BackupPayload_Status = 0
	BackupPayload_SUCCESS   BackupPayload_Status = 1
	BackupPayload_DUPLICATE BackupPayload_Status = 2
	BackupPayload_FAILED    BackupPayload_Status = 3
)

var BackupPayload_Status_name = map[int32]string{
	0: "NONE",
	1: "SUCCESS",
	2: "DUPLICATE",
	3: "FAILED",
}
var BackupPayload_Status_value = map[string]int32{
	"NONE":      0,
	"SUCCESS":   1,
	"DUPLICATE": 2,
	"FAILED":    3,
}

func (x BackupPayload_Status) String() string {
	return proto.EnumName(BackupPayload_Status_name, int32(x))
}
func (BackupPayload_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorPayload, []int{1, 0}
}

type Payload struct {
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptorPayload, []int{0} }

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// BackupPayload is used both as a request and a response.
// When used in request, groups represents the list of groups that need to be backed up.
// When used in response, groups represent the list of groups that were backed up.
type BackupPayload struct {
	ReqId   uint64               `protobuf:"varint,1,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	GroupId uint32               `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Status  BackupPayload_Status `protobuf:"varint,3,opt,name=status,proto3,enum=protos.BackupPayload_Status" json:"status,omitempty"`
}

func (m *BackupPayload) Reset()                    { *m = BackupPayload{} }
func (m *BackupPayload) String() string            { return proto.CompactTextString(m) }
func (*BackupPayload) ProtoMessage()               {}
func (*BackupPayload) Descriptor() ([]byte, []int) { return fileDescriptorPayload, []int{1} }

func (m *BackupPayload) GetReqId() uint64 {
	if m != nil {
		return m.ReqId
	}
	return 0
}

func (m *BackupPayload) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *BackupPayload) GetStatus() BackupPayload_Status {
	if m != nil {
		return m.Status
	}
	return BackupPayload_NONE
}

func init() {
	proto.RegisterType((*Payload)(nil), "protos.Payload")
	proto.RegisterType((*BackupPayload)(nil), "protos.BackupPayload")
	proto.RegisterEnum("protos.BackupPayload_Status", BackupPayload_Status_name, BackupPayload_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Worker service

type WorkerClient interface {
	// Connection testing RPC.
	Echo(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error)
	// Data serving RPCs.
	AssignUids(ctx context.Context, in *Num, opts ...grpc.CallOption) (*List, error)
	Mutate(ctx context.Context, in *Mutations, opts ...grpc.CallOption) (*Payload, error)
	ServeTask(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Result, error)
	PredicateAndSchemaData(ctx context.Context, opts ...grpc.CallOption) (Worker_PredicateAndSchemaDataClient, error)
	Sort(ctx context.Context, in *SortMessage, opts ...grpc.CallOption) (*SortResult, error)
	RebuildIndex(ctx context.Context, in *RebuildIndexMessage, opts ...grpc.CallOption) (*Payload, error)
	Schema(ctx context.Context, in *SchemaRequest, opts ...grpc.CallOption) (*SchemaResult, error)
	// RAFT serving RPCs.
	RaftMessage(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error)
	JoinCluster(ctx context.Context, in *RaftContext, opts ...grpc.CallOption) (*Payload, error)
	UpdateMembership(ctx context.Context, in *MembershipUpdate, opts ...grpc.CallOption) (*MembershipUpdate, error)
	Backup(ctx context.Context, in *BackupPayload, opts ...grpc.CallOption) (*BackupPayload, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Echo(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/protos.Worker/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) AssignUids(ctx context.Context, in *Num, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := grpc.Invoke(ctx, "/protos.Worker/AssignUids", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Mutate(ctx context.Context, in *Mutations, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/protos.Worker/Mutate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) ServeTask(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/protos.Worker/ServeTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) PredicateAndSchemaData(ctx context.Context, opts ...grpc.CallOption) (Worker_PredicateAndSchemaDataClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Worker_serviceDesc.Streams[0], c.cc, "/protos.Worker/PredicateAndSchemaData", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerPredicateAndSchemaDataClient{stream}
	return x, nil
}

type Worker_PredicateAndSchemaDataClient interface {
	Send(*GroupKeys) error
	Recv() (*KV, error)
	grpc.ClientStream
}

type workerPredicateAndSchemaDataClient struct {
	grpc.ClientStream
}

func (x *workerPredicateAndSchemaDataClient) Send(m *GroupKeys) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerPredicateAndSchemaDataClient) Recv() (*KV, error) {
	m := new(KV)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workerClient) Sort(ctx context.Context, in *SortMessage, opts ...grpc.CallOption) (*SortResult, error) {
	out := new(SortResult)
	err := grpc.Invoke(ctx, "/protos.Worker/Sort", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) RebuildIndex(ctx context.Context, in *RebuildIndexMessage, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/protos.Worker/RebuildIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Schema(ctx context.Context, in *SchemaRequest, opts ...grpc.CallOption) (*SchemaResult, error) {
	out := new(SchemaResult)
	err := grpc.Invoke(ctx, "/protos.Worker/Schema", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) RaftMessage(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/protos.Worker/RaftMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) JoinCluster(ctx context.Context, in *RaftContext, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := grpc.Invoke(ctx, "/protos.Worker/JoinCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) UpdateMembership(ctx context.Context, in *MembershipUpdate, opts ...grpc.CallOption) (*MembershipUpdate, error) {
	out := new(MembershipUpdate)
	err := grpc.Invoke(ctx, "/protos.Worker/UpdateMembership", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Backup(ctx context.Context, in *BackupPayload, opts ...grpc.CallOption) (*BackupPayload, error) {
	out := new(BackupPayload)
	err := grpc.Invoke(ctx, "/protos.Worker/Backup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Worker service

type WorkerServer interface {
	// Connection testing RPC.
	Echo(context.Context, *Payload) (*Payload, error)
	// Data serving RPCs.
	AssignUids(context.Context, *Num) (*List, error)
	Mutate(context.Context, *Mutations) (*Payload, error)
	ServeTask(context.Context, *Query) (*Result, error)
	PredicateAndSchemaData(Worker_PredicateAndSchemaDataServer) error
	Sort(context.Context, *SortMessage) (*SortResult, error)
	RebuildIndex(context.Context, *RebuildIndexMessage) (*Payload, error)
	Schema(context.Context, *SchemaRequest) (*SchemaResult, error)
	// RAFT serving RPCs.
	RaftMessage(context.Context, *Payload) (*Payload, error)
	JoinCluster(context.Context, *RaftContext) (*Payload, error)
	UpdateMembership(context.Context, *MembershipUpdate) (*MembershipUpdate, error)
	Backup(context.Context, *BackupPayload) (*BackupPayload, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Echo(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_AssignUids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Num)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).AssignUids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/AssignUids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).AssignUids(ctx, req.(*Num))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Mutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mutations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Mutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Mutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Mutate(ctx, req.(*Mutations))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_ServeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).ServeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/ServeTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).ServeTask(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_PredicateAndSchemaData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServer).PredicateAndSchemaData(&workerPredicateAndSchemaDataServer{stream})
}

type Worker_PredicateAndSchemaDataServer interface {
	Send(*KV) error
	Recv() (*GroupKeys, error)
	grpc.ServerStream
}

type workerPredicateAndSchemaDataServer struct {
	grpc.ServerStream
}

func (x *workerPredicateAndSchemaDataServer) Send(m *KV) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerPredicateAndSchemaDataServer) Recv() (*GroupKeys, error) {
	m := new(GroupKeys)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Worker_Sort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Sort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Sort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Sort(ctx, req.(*SortMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_RebuildIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebuildIndexMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).RebuildIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/RebuildIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).RebuildIndex(ctx, req.(*RebuildIndexMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Schema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Schema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Schema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Schema(ctx, req.(*SchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_RaftMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).RaftMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/RaftMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).RaftMessage(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_JoinCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).JoinCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/JoinCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).JoinCluster(ctx, req.(*RaftContext))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_UpdateMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MembershipUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).UpdateMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/UpdateMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).UpdateMembership(ctx, req.(*MembershipUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Worker/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Backup(ctx, req.(*BackupPayload))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Worker_Echo_Handler,
		},
		{
			MethodName: "AssignUids",
			Handler:    _Worker_AssignUids_Handler,
		},
		{
			MethodName: "Mutate",
			Handler:    _Worker_Mutate_Handler,
		},
		{
			MethodName: "ServeTask",
			Handler:    _Worker_ServeTask_Handler,
		},
		{
			MethodName: "Sort",
			Handler:    _Worker_Sort_Handler,
		},
		{
			MethodName: "RebuildIndex",
			Handler:    _Worker_RebuildIndex_Handler,
		},
		{
			MethodName: "Schema",
			Handler:    _Worker_Schema_Handler,
		},
		{
			MethodName: "RaftMessage",
			Handler:    _Worker_RaftMessage_Handler,
		},
		{
			MethodName: "JoinCluster",
			Handler:    _Worker_JoinCluster_Handler,
		},
		{
			MethodName: "UpdateMembership",
			Handler:    _Worker_UpdateMembership_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _Worker_Backup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PredicateAndSchemaData",
			Handler:       _Worker_PredicateAndSchemaData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "payload.proto",
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *BackupPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackupPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReqId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.ReqId))
	}
	if m.GroupId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.GroupId))
	}
	if m.Status != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func encodeFixed64Payload(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Payload(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPayload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Payload) Size() (n int) {
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	return n
}

func (m *BackupPayload) Size() (n int) {
	var l int
	_ = l
	if m.ReqId != 0 {
		n += 1 + sovPayload(uint64(m.ReqId))
	}
	if m.GroupId != 0 {
		n += 1 + sovPayload(uint64(m.GroupId))
	}
	if m.Status != 0 {
		n += 1 + sovPayload(uint64(m.Status))
	}
	return n
}

func sovPayload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPayload(x uint64) (n int) {
	return sovPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackupPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackupPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackupPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqId", wireType)
			}
			m.ReqId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			m.GroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (BackupPayload_Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayload
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPayload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPayload
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPayload(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPayload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayload   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("payload.proto", fileDescriptorPayload) }

var fileDescriptorPayload = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xe1, 0x8e, 0xd2, 0x4c,
	0x14, 0x6d, 0x77, 0xd9, 0xb2, 0x7b, 0x81, 0xfd, 0xfa, 0x5d, 0x5d, 0x83, 0x55, 0x09, 0xe9, 0x2f,
	0x34, 0x86, 0xb8, 0xab, 0x46, 0x63, 0xa2, 0x09, 0x0b, 0xa8, 0xb8, 0x80, 0xd8, 0x2e, 0xfa, 0xd3,
	0x0c, 0xf4, 0x0a, 0x0d, 0xd0, 0x96, 0x99, 0xa9, 0x59, 0x5e, 0xc4, 0xf8, 0x1e, 0xbe, 0x84, 0x3f,
	0x7d, 0x04, 0x83, 0x2f, 0x62, 0xda, 0x52, 0x70, 0x37, 0x98, 0xf8, 0x8b, 0x39, 0xe7, 0x9e, 0x39,
	0xe7, 0x5e, 0xee, 0x14, 0x0a, 0x01, 0x5b, 0x4c, 0x7d, 0xe6, 0x54, 0x03, 0xee, 0x4b, 0x1f, 0xb5,
	0xf8, 0x47, 0x18, 0x79, 0x31, 0x1c, 0xd3, 0x8c, 0x25, 0xac, 0x01, 0x92, 0x89, 0x49, 0x72, 0x36,
	0xef, 0x40, 0xb6, 0x97, 0x5c, 0x41, 0x84, 0x4c, 0x83, 0x49, 0x56, 0x54, 0xcb, 0x6a, 0x25, 0x6f,
	0xc5, 0x67, 0xf3, 0x9b, 0x0a, 0x85, 0x53, 0x36, 0x9c, 0x84, 0x41, 0xaa, 0x3a, 0x02, 0x8d, 0xd3,
	0xfc, 0xa3, 0xeb, 0xc4, 0xba, 0x8c, 0xb5, 0xc7, 0x69, 0xde, 0x72, 0xf0, 0x26, 0xec, 0x8f, 0xb8,
	0x1f, 0x06, 0x51, 0x61, 0xa7, 0xac, 0x56, 0x0a, 0x56, 0x36, 0xc6, 0x2d, 0x07, 0x1f, 0x81, 0x26,
	0x24, 0x93, 0xa1, 0x28, 0xee, 0x96, 0xd5, 0xca, 0xe1, 0xc9, 0xed, 0x24, 0x5a, 0x54, 0x2f, 0x19,
	0x57, 0xed, 0x58, 0x63, 0xad, 0xb4, 0xe6, 0x33, 0xd0, 0x12, 0x06, 0xf7, 0x21, 0xd3, 0x7d, 0xdb,
	0x6d, 0xea, 0x0a, 0xe6, 0x20, 0x6b, 0xf7, 0xeb, 0xf5, 0xa6, 0x6d, 0xeb, 0x2a, 0x16, 0xe0, 0xa0,
	0xd1, 0xef, 0xb5, 0x5b, 0xf5, 0xda, 0x79, 0x53, 0xdf, 0x41, 0x00, 0xed, 0x65, 0xad, 0xd5, 0x6e,
	0x36, 0xf4, 0xdd, 0x93, 0x2f, 0x7b, 0xa0, 0x7d, 0xf0, 0xf9, 0x84, 0x38, 0xde, 0x83, 0x4c, 0x73,
	0x38, 0xf6, 0xf1, 0xbf, 0x34, 0x74, 0x15, 0x67, 0x5c, 0x25, 0x4c, 0x05, 0xef, 0x02, 0xd4, 0x84,
	0x70, 0x47, 0x5e, 0xdf, 0x75, 0x04, 0xe6, 0x52, 0x41, 0x37, 0x9c, 0x19, 0xf9, 0x14, 0xb4, 0x5d,
	0x21, 0x4d, 0x05, 0xab, 0xa0, 0x75, 0x42, 0xc9, 0x24, 0xe1, 0xff, 0x69, 0x25, 0xc6, 0xae, 0xef,
	0x89, 0x6d, 0xd6, 0xf7, 0xe1, 0xc0, 0x26, 0xfe, 0x99, 0xce, 0x99, 0x98, 0x60, 0x21, 0xad, 0xbf,
	0x0b, 0x89, 0x2f, 0x8c, 0xc3, 0x14, 0x5a, 0x24, 0xc2, 0x69, 0xe4, 0xfe, 0x1c, 0x6e, 0xf4, 0x38,
	0x39, 0xee, 0x90, 0x49, 0xaa, 0x79, 0x8e, 0x1d, 0x2f, 0x2f, 0xda, 0xc7, 0x26, 0xed, 0x55, 0xf4,
	0xe7, 0x9e, 0xd1, 0x42, 0x18, 0x90, 0x52, 0x67, 0xef, 0x4d, 0xa5, 0xa2, 0x3e, 0x50, 0xf1, 0x18,
	0x32, 0xb6, 0xcf, 0x25, 0x5e, 0x4b, 0x2b, 0x11, 0xea, 0x90, 0x10, 0x6c, 0x44, 0x06, 0xfe, 0x49,
	0xae, 0x13, 0x5f, 0x40, 0xde, 0xa2, 0x41, 0xe8, 0x4e, 0x9d, 0x96, 0xe7, 0xd0, 0x05, 0xde, 0xda,
	0xf4, 0xb4, 0x61, 0x53, 0x8b, 0x2d, 0xf3, 0x3d, 0x01, 0x2d, 0xe9, 0x12, 0x8f, 0xd6, 0xfe, 0x31,
	0xb6, 0x68, 0x1e, 0x92, 0x90, 0xc6, 0xf5, 0xab, 0xf4, 0x2a, 0xf8, 0x18, 0x72, 0x16, 0xfb, 0x94,
	0x76, 0xf7, 0x4f, 0x6b, 0x7a, 0x0c, 0xb9, 0x37, 0xbe, 0xeb, 0xd5, 0xa7, 0xa1, 0x90, 0xc4, 0x37,
	0x53, 0x46, 0x3e, 0x75, 0xdf, 0x93, 0x74, 0x21, 0xb7, 0x5d, 0x7b, 0x0d, 0x7a, 0x3f, 0x70, 0x98,
	0xa4, 0x0e, 0xcd, 0x06, 0xc4, 0xc5, 0xd8, 0x0d, 0xb0, 0xb8, 0x5e, 0xde, 0x9a, 0x4b, 0x34, 0xc6,
	0x5f, 0x2b, 0xa6, 0x82, 0x4f, 0x41, 0x4b, 0x9e, 0xee, 0x66, 0xd8, 0x4b, 0x4f, 0xd9, 0xd8, 0x4e,
	0x9b, 0xca, 0xa9, 0xfe, 0x7d, 0x59, 0x52, 0x7f, 0x2c, 0x4b, 0xea, 0xcf, 0x65, 0x49, 0xfd, 0xfa,
	0xab, 0xa4, 0x0c, 0x92, 0x2f, 0xf4, 0xe1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0x6e, 0xf6,
	0x02, 0xb9, 0x03, 0x00, 0x00,
}
