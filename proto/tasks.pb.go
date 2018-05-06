// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tasks.proto

/*
	Package landlady is a generated protocol buffer package.

	It is generated from these files:
		tasks.proto

	It has these top-level messages:
		TaskProject
		TaksProjectList
		Task
		TaskList
		RequestTaskOption
*/
package landlady

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TaskProject struct {
	Name       string                      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Module     string                      `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	Desc       string                      `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Status     string                      `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CountTasks int64                       `protobuf:"varint,5,opt,name=countTasks,proto3" json:"countTasks,omitempty"`
	StartAt    *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=startAt" json:"startAt,omitempty"`
}

func (m *TaskProject) Reset()                    { *m = TaskProject{} }
func (m *TaskProject) String() string            { return proto.CompactTextString(m) }
func (*TaskProject) ProtoMessage()               {}
func (*TaskProject) Descriptor() ([]byte, []int) { return fileDescriptorTasks, []int{0} }

type TaksProjectList struct {
	Items []*TaskProject `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *TaksProjectList) Reset()                    { *m = TaksProjectList{} }
func (m *TaksProjectList) String() string            { return proto.CompactTextString(m) }
func (*TaksProjectList) ProtoMessage()               {}
func (*TaksProjectList) Descriptor() ([]byte, []int) { return fileDescriptorTasks, []int{1} }

type Task struct {
	ProjectName string            `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	Url         string            `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Labels      map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptorTasks, []int{2} }

type TaskList struct {
	Items []*Task `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *TaskList) Reset()                    { *m = TaskList{} }
func (m *TaskList) String() string            { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()               {}
func (*TaskList) Descriptor() ([]byte, []int) { return fileDescriptorTasks, []int{3} }

type RequestTaskOption struct {
	ProjectName string `protobuf:"bytes,1,opt,name=projectName,proto3" json:"projectName,omitempty"`
	Count       int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *RequestTaskOption) Reset()                    { *m = RequestTaskOption{} }
func (m *RequestTaskOption) String() string            { return proto.CompactTextString(m) }
func (*RequestTaskOption) ProtoMessage()               {}
func (*RequestTaskOption) Descriptor() ([]byte, []int) { return fileDescriptorTasks, []int{4} }

func init() {
	proto.RegisterType((*TaskProject)(nil), "landlady.TaskProject")
	proto.RegisterType((*TaksProjectList)(nil), "landlady.TaksProjectList")
	proto.RegisterType((*Task)(nil), "landlady.Task")
	proto.RegisterType((*TaskList)(nil), "landlady.TaskList")
	proto.RegisterType((*RequestTaskOption)(nil), "landlady.RequestTaskOption")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tasks service

type TasksClient interface {
	NewProject(ctx context.Context, in *TaskProject, opts ...grpc.CallOption) (*TaskProject, error)
	GetProject(ctx context.Context, in *TaskProject, opts ...grpc.CallOption) (*TaskProject, error)
	AddTasks(ctx context.Context, in *TaskList, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	RequestTasks(ctx context.Context, in *RequestTaskOption, opts ...grpc.CallOption) (*TaskList, error)
	CompleteTask(ctx context.Context, in *TaskList, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type tasksClient struct {
	cc *grpc.ClientConn
}

func NewTasksClient(cc *grpc.ClientConn) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) NewProject(ctx context.Context, in *TaskProject, opts ...grpc.CallOption) (*TaskProject, error) {
	out := new(TaskProject)
	err := grpc.Invoke(ctx, "/landlady.Tasks/NewProject", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) GetProject(ctx context.Context, in *TaskProject, opts ...grpc.CallOption) (*TaskProject, error) {
	out := new(TaskProject)
	err := grpc.Invoke(ctx, "/landlady.Tasks/GetProject", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) AddTasks(ctx context.Context, in *TaskList, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/landlady.Tasks/AddTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) RequestTasks(ctx context.Context, in *RequestTaskOption, opts ...grpc.CallOption) (*TaskList, error) {
	out := new(TaskList)
	err := grpc.Invoke(ctx, "/landlady.Tasks/RequestTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) CompleteTask(ctx context.Context, in *TaskList, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/landlady.Tasks/CompleteTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tasks service

type TasksServer interface {
	NewProject(context.Context, *TaskProject) (*TaskProject, error)
	GetProject(context.Context, *TaskProject) (*TaskProject, error)
	AddTasks(context.Context, *TaskList) (*google_protobuf.Empty, error)
	RequestTasks(context.Context, *RequestTaskOption) (*TaskList, error)
	CompleteTask(context.Context, *TaskList) (*google_protobuf.Empty, error)
}

func RegisterTasksServer(s *grpc.Server, srv TasksServer) {
	s.RegisterService(&_Tasks_serviceDesc, srv)
}

func _Tasks_NewProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).NewProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landlady.Tasks/NewProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).NewProject(ctx, req.(*TaskProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landlady.Tasks/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).GetProject(ctx, req.(*TaskProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_AddTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).AddTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landlady.Tasks/AddTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).AddTasks(ctx, req.(*TaskList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_RequestTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTaskOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).RequestTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landlady.Tasks/RequestTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).RequestTasks(ctx, req.(*RequestTaskOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_CompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).CompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landlady.Tasks/CompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).CompleteTask(ctx, req.(*TaskList))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tasks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "landlady.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewProject",
			Handler:    _Tasks_NewProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Tasks_GetProject_Handler,
		},
		{
			MethodName: "AddTasks",
			Handler:    _Tasks_AddTasks_Handler,
		},
		{
			MethodName: "RequestTasks",
			Handler:    _Tasks_RequestTasks_Handler,
		},
		{
			MethodName: "CompleteTask",
			Handler:    _Tasks_CompleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasks.proto",
}

func (m *TaskProject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskProject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTasks(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Module) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTasks(dAtA, i, uint64(len(m.Module)))
		i += copy(dAtA[i:], m.Module)
	}
	if len(m.Desc) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTasks(dAtA, i, uint64(len(m.Desc)))
		i += copy(dAtA[i:], m.Desc)
	}
	if len(m.Status) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTasks(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	if m.CountTasks != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintTasks(dAtA, i, uint64(m.CountTasks))
	}
	if m.StartAt != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintTasks(dAtA, i, uint64(m.StartAt.Size()))
		n1, err := m.StartAt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *TaksProjectList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaksProjectList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTasks(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Task) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Task) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProjectName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTasks(dAtA, i, uint64(len(m.ProjectName)))
		i += copy(dAtA[i:], m.ProjectName)
	}
	if len(m.Url) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTasks(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x1a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovTasks(uint64(len(k))) + 1 + len(v) + sovTasks(uint64(len(v)))
			i = encodeVarintTasks(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTasks(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintTasks(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *TaskList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTasks(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *RequestTaskOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestTaskOption) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProjectName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTasks(dAtA, i, uint64(len(m.ProjectName)))
		i += copy(dAtA[i:], m.ProjectName)
	}
	if m.Count != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTasks(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func encodeVarintTasks(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TaskProject) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	if m.CountTasks != 0 {
		n += 1 + sovTasks(uint64(m.CountTasks))
	}
	if m.StartAt != nil {
		l = m.StartAt.Size()
		n += 1 + l + sovTasks(uint64(l))
	}
	return n
}

func (m *TaksProjectList) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovTasks(uint64(l))
		}
	}
	return n
}

func (m *Task) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProjectName)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTasks(uint64(len(k))) + 1 + len(v) + sovTasks(uint64(len(v)))
			n += mapEntrySize + 1 + sovTasks(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *TaskList) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovTasks(uint64(l))
		}
	}
	return n
}

func (m *RequestTaskOption) Size() (n int) {
	var l int
	_ = l
	l = len(m.ProjectName)
	if l > 0 {
		n += 1 + l + sovTasks(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovTasks(uint64(m.Count))
	}
	return n
}

func sovTasks(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTasks(x uint64) (n int) {
	return sovTasks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TaskProject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
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
			return fmt.Errorf("proto: TaskProject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskProject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountTasks", wireType)
			}
			m.CountTasks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountTasks |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartAt == nil {
				m.StartAt = &google_protobuf1.Timestamp{}
			}
			if err := m.StartAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
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
func (m *TaksProjectList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
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
			return fmt.Errorf("proto: TaksProjectList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaksProjectList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &TaskProject{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
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
func (m *Task) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
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
			return fmt.Errorf("proto: Task: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTasks
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTasks
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTasks
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTasks
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTasks
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTasks(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTasks
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
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
func (m *TaskList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
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
			return fmt.Errorf("proto: TaskList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Task{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
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
func (m *RequestTaskOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTasks
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
			return fmt.Errorf("proto: RequestTaskOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestTaskOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTasks
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTasks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTasks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTasks
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
func skipTasks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTasks
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
					return 0, ErrIntOverflowTasks
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
					return 0, ErrIntOverflowTasks
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
				return 0, ErrInvalidLengthTasks
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTasks
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
				next, err := skipTasks(dAtA[start:])
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
	ErrInvalidLengthTasks = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTasks   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tasks.proto", fileDescriptorTasks) }

var fileDescriptorTasks = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xbb, 0x71, 0x12, 0xc2, 0xb8, 0x82, 0xb2, 0x2a, 0x95, 0xe5, 0x4a, 0xc6, 0xb2, 0x38,
	0x44, 0x42, 0x38, 0x28, 0x70, 0x28, 0x08, 0x21, 0x05, 0x54, 0x71, 0xa0, 0x2a, 0xc8, 0xca, 0x0b,
	0x6c, 0xe2, 0xc5, 0x98, 0xd8, 0x5e, 0xe3, 0x1d, 0x83, 0xf2, 0x48, 0x3c, 0x03, 0x12, 0xe7, 0x1e,
	0x79, 0x04, 0xc8, 0x93, 0xa0, 0xdd, 0xb5, 0x15, 0x87, 0xe6, 0x00, 0xdc, 0x66, 0x66, 0xbf, 0xdd,
	0x99, 0xff, 0xf7, 0x18, 0x6c, 0x64, 0x72, 0x25, 0xc3, 0xb2, 0x12, 0x28, 0xe8, 0x28, 0x63, 0x45,
	0x9c, 0xb1, 0x78, 0xed, 0x9e, 0x26, 0x42, 0x24, 0x19, 0x9f, 0xe8, 0xfa, 0xa2, 0x7e, 0x3f, 0xe1,
	0x79, 0x89, 0x6b, 0x83, 0xb9, 0xf7, 0xfe, 0x3c, 0xc4, 0x34, 0xe7, 0x12, 0x59, 0x5e, 0x36, 0xc0,
	0xc3, 0x24, 0xc5, 0x0f, 0xf5, 0x22, 0x5c, 0x8a, 0x7c, 0x92, 0x88, 0x44, 0x6c, 0x49, 0x95, 0xe9,
	0x44, 0x47, 0x06, 0x0f, 0xbe, 0x11, 0xb0, 0xe7, 0x4c, 0xae, 0xde, 0x55, 0xe2, 0x23, 0x5f, 0x22,
	0xa5, 0xd0, 0x2f, 0x58, 0xce, 0x1d, 0xe2, 0x93, 0xf1, 0xcd, 0x48, 0xc7, 0xf4, 0x04, 0x86, 0xb9,
	0x88, 0xeb, 0x8c, 0x3b, 0x3d, 0x5d, 0x6d, 0x32, 0xc5, 0xc6, 0x5c, 0x2e, 0x1d, 0xcb, 0xb0, 0x2a,
	0x56, 0xac, 0x44, 0x86, 0xb5, 0x74, 0xfa, 0x86, 0x35, 0x19, 0xf5, 0x00, 0x96, 0xa2, 0x2e, 0x50,
	0xf5, 0x92, 0xce, 0xc0, 0x27, 0x63, 0x2b, 0xea, 0x54, 0xe8, 0x13, 0xb8, 0x21, 0x91, 0x55, 0x38,
	0x43, 0x67, 0xe8, 0x93, 0xb1, 0x3d, 0x75, 0x43, 0xa3, 0x34, 0x6c, 0xe7, 0x0f, 0xe7, 0xad, 0xd2,
	0xa8, 0x45, 0x83, 0x17, 0x70, 0x7b, 0xce, 0x56, 0xb2, 0x19, 0xfe, 0x22, 0x95, 0x48, 0x1f, 0xc0,
	0x20, 0x45, 0x9e, 0x4b, 0x87, 0xf8, 0xd6, 0xd8, 0x9e, 0xde, 0x0d, 0x5b, 0x5f, 0xc3, 0x8e, 0xcc,
	0xc8, 0x30, 0xc1, 0x57, 0x02, 0x7d, 0x55, 0xa6, 0x3e, 0xd8, 0xa5, 0x39, 0xba, 0xdc, 0xaa, 0xef,
	0x96, 0xe8, 0x11, 0x58, 0x75, 0x95, 0x35, 0x0e, 0xa8, 0x90, 0x4e, 0x61, 0x98, 0xb1, 0x05, 0xcf,
	0xa4, 0x63, 0xe9, 0x56, 0xee, 0x6e, 0xab, 0xf0, 0x42, 0x1f, 0x9e, 0x17, 0x58, 0xad, 0xa3, 0x86,
	0x74, 0x9f, 0x82, 0xdd, 0x29, 0xab, 0x47, 0x57, 0x7c, 0xdd, 0xb4, 0x53, 0x21, 0x3d, 0x86, 0xc1,
	0x67, 0x96, 0xd5, 0xad, 0xd5, 0x26, 0x79, 0xd6, 0x3b, 0x23, 0xc1, 0x23, 0x18, 0xa9, 0x67, 0xb5,
	0xc8, 0xfb, 0xbb, 0x22, 0x6f, 0xed, 0x76, 0x6e, 0xd5, 0xbd, 0x81, 0x3b, 0x11, 0xff, 0x54, 0x73,
	0xa9, 0x3d, 0x7e, 0x5b, 0x62, 0x2a, 0x8a, 0xbf, 0x50, 0x7a, 0x0c, 0x03, 0xfd, 0x61, 0xf4, 0x08,
	0x56, 0x64, 0x92, 0xe9, 0xf7, 0x1e, 0x0c, 0xcc, 0xa7, 0x7a, 0x0e, 0x70, 0xc9, 0xbf, 0xb4, 0x0b,
	0xb3, 0xdf, 0x60, 0x77, 0x7f, 0x39, 0x38, 0x50, 0xb7, 0x5f, 0x73, 0xfc, 0xdf, 0xdb, 0x67, 0x30,
	0x9a, 0xc5, 0xb1, 0x99, 0x83, 0xee, 0x42, 0xca, 0x18, 0xf7, 0xe4, 0xda, 0xd6, 0x9c, 0xab, 0x9f,
	0x27, 0x38, 0xa0, 0x33, 0x38, 0xec, 0x98, 0x21, 0xe9, 0xe9, 0xf6, 0xf6, 0x35, 0x93, 0xdc, 0x3d,
	0x4f, 0xeb, 0xd1, 0x0f, 0x5f, 0x89, 0xbc, 0xcc, 0x38, 0x72, 0xbd, 0x34, 0xff, 0x34, 0xc0, 0xcb,
	0xa3, 0xab, 0x5f, 0xde, 0xc1, 0xd5, 0xc6, 0x23, 0x3f, 0x36, 0x1e, 0xf9, 0xb9, 0xf1, 0xc8, 0x62,
	0xa8, 0x99, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xbe, 0x0b, 0x82, 0x08, 0x04, 0x00,
	0x00,
}
