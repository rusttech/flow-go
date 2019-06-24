// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inter/security.proto

package bamboo_proto

import (
	context "context"
	fmt "fmt"
	shared "github.com/dapperlabs/bamboo-node/grpc/shared"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

func init() { proto.RegisterFile("inter/security.proto", fileDescriptor_51b689a1b1b125c1) }

var fileDescriptor_51b689a1b1b125c1 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x85, 0x84, 0x38, 0x58, 0x15, 0x42, 0x86, 0x0d, 0x51, 0xd0, 0xc4, 0x40, 0x0c, 0x04,
	0x2c, 0x91, 0x40, 0x70, 0xe3, 0xb0, 0x55, 0x63, 0x70, 0x41, 0xd5, 0xca, 0x90, 0x38, 0x70, 0x70,
	0x92, 0x47, 0x62, 0xcd, 0xf1, 0x33, 0xb6, 0x53, 0x31, 0x2e, 0x7c, 0x0c, 0x3e, 0x21, 0xdf, 0x03,
	0xb5, 0xb6, 0xb7, 0xc6, 0x4d, 0x68, 0x76, 0xb4, 0xff, 0xbf, 0xf7, 0xff, 0x3f, 0xf7, 0xf5, 0x85,
	0xdc, 0xe1, 0xd2, 0x82, 0x4e, 0x0d, 0xe4, 0x8d, 0xe6, 0xf6, 0x3c, 0x51, 0x1a, 0x2d, 0xd2, 0x51,
	0xc6, 0xea, 0x0c, 0xd1, 0x9d, 0xc6, 0x5b, 0xa6, 0x62, 0x1a, 0x8a, 0xb4, 0x06, 0x63, 0x58, 0x09,
	0xc6, 0x5f, 0xdf, 0x2f, 0x11, 0x4b, 0x01, 0xe9, 0xf2, 0x94, 0x35, 0xdf, 0x53, 0xa8, 0x55, 0x70,
	0x78, 0xf5, 0x97, 0x90, 0xd1, 0xcc, 0x9b, 0x7e, 0xc2, 0x02, 0xe8, 0x3b, 0x72, 0x7d, 0xca, 0x65,
	0x49, 0xef, 0x25, 0xab, 0xde, 0xc9, 0xe2, 0xee, 0x04, 0x7e, 0x34, 0x60, 0xec, 0x78, 0xdc, 0x25,
	0x19, 0x85, 0xd2, 0x00, 0x2d, 0xc9, 0xf6, 0x41, 0x9e, 0x83, 0x31, 0xb3, 0x26, 0xab, 0xb9, 0x9d,
	0xa0, 0x10, 0x90, 0x5b, 0x8e, 0x92, 0x3e, 0x69, 0x57, 0x39, 0xea, 0x52, 0x0f, 0xe6, 0x7b, 0x9b,
	0x30, 0x1f, 0x74, 0x4a, 0x46, 0x53, 0x8d, 0x0a, 0x0d, 0x1c, 0x0a, 0xcc, 0xcf, 0xe8, 0x6e, 0xd4,
	0xd4, 0x8a, 0x16, 0xac, 0x1f, 0xfd, 0x0f, 0xf1, 0xb6, 0x15, 0xb9, 0x7d, 0xaa, 0x0a, 0x66, 0xc1,
	0xab, 0x85, 0x73, 0x7f, 0xda, 0x5f, 0xea, 0xf0, 0x90, 0xf1, 0x6c, 0x33, 0x78, 0xf1, 0x80, 0x9b,
	0xc7, 0x60, 0x97, 0xca, 0xe1, 0xf9, 0x07, 0x66, 0x2a, 0xfa, 0xb8, 0x5d, 0xdb, 0x56, 0x43, 0xc0,
	0x4e, 0x37, 0x74, 0x61, 0xfb, 0x95, 0xdc, 0x5a, 0x29, 0x04, 0x5e, 0x56, 0x36, 0xfe, 0xe9, 0x63,
	0x7d, 0xa8, 0xb5, 0x22, 0x5b, 0x53, 0x8d, 0x8b, 0x79, 0x9c, 0x80, 0x69, 0x84, 0x3d, 0x50, 0x4a,
	0xe3, 0x9c, 0x09, 0xfa, 0x7c, 0xed, 0xd1, 0xeb, 0x50, 0x08, 0x79, 0x31, 0x88, 0xf5, 0x89, 0xbf,
	0xc9, 0x83, 0x63, 0xb0, 0xef, 0xb9, 0x64, 0x82, 0xff, 0x82, 0x62, 0x66, 0x99, 0x85, 0xcf, 0x9a,
	0x49, 0xc3, 0x17, 0xff, 0x05, 0x43, 0x93, 0xb6, 0x59, 0x2f, 0x18, 0xc2, 0xd3, 0xc1, 0xbc, 0x6f,
	0xe0, 0x1b, 0xd9, 0xf1, 0x1d, 0x46, 0x88, 0x1b, 0x2c, 0x13, 0xf1, 0xd0, 0x66, 0xbc, 0x94, 0x6b,
	0x7e, 0xe3, 0xed, 0xc4, 0xed, 0x60, 0x12, 0x76, 0x30, 0x39, 0x5a, 0xec, 0x20, 0x3d, 0x23, 0xbb,
	0x7d, 0xf6, 0xa0, 0x98, 0x86, 0x2f, 0x68, 0x21, 0x7e, 0x64, 0x67, 0xc2, 0x0a, 0xdf, 0x1b, 0xc6,
	0xc9, 0xc3, 0xee, 0xb0, 0x09, 0xd6, 0x35, 0xb7, 0xcb, 0xac, 0xfd, 0x01, 0x59, 0x97, 0x78, 0x6f,
	0xd4, 0x9c, 0xdc, 0xf5, 0x51, 0x47, 0x3f, 0x21, 0x6f, 0xdc, 0xe2, 0xe6, 0xc0, 0x95, 0xa5, 0x2f,
	0x3b, 0xe7, 0x1f, 0x63, 0x61, 0x60, 0xfb, 0x03, 0x69, 0x3f, 0xae, 0x3f, 0xd7, 0xc8, 0x9e, 0xfb,
	0xf0, 0x7c, 0x94, 0x73, 0x26, 0x78, 0x11, 0x93, 0x93, 0x8a, 0x09, 0x01, 0xb2, 0x04, 0xfa, 0xa6,
	0xed, 0xbc, 0x89, 0x0f, 0x0d, 0xbd, 0xbd, 0x6a, 0x99, 0xeb, 0x2c, 0xbb, 0xb1, 0xe4, 0x5f, 0xff,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xf5, 0x8b, 0x23, 0xc8, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecurityNodeClient is the client API for SecurityNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecurityNodeClient interface {
	Ping(ctx context.Context, in *shared.PingRequest, opts ...grpc.CallOption) (*shared.PingResponse, error)
	// Receive a signed collection from an access node.
	AccessSubmitCollection(ctx context.Context, in *shared.AccessCollectionRequest, opts ...grpc.CallOption) (*shared.AccessCollectionResponse, error)
	// Notify the security node that another has proposed a block.
	ProposeBlock(ctx context.Context, in *shared.ProposeBlockRequest, opts ...grpc.CallOption) (*shared.ProposeBlockResponse, error)
	// Update a block proposal to add new signatures.
	UpdateProposedBlock(ctx context.Context, in *shared.ProposeBlockUpdateRequest, opts ...grpc.CallOption) (*shared.ProposeBlockUpdateResponse, error)
	// Returns a block by hash
	GetBlockByHash(ctx context.Context, in *shared.GetBlockByHashRequest, opts ...grpc.CallOption) (*shared.GetBlockResponse, error)
	// Returns a block by height
	GetBlockByHeight(ctx context.Context, in *shared.GetBlockByHeightRequest, opts ...grpc.CallOption) (*shared.GetBlockResponse, error)
	// Process result approval from access nodes.
	ProcessResultApproval(ctx context.Context, in *shared.ProcessResultApprovalRequest, opts ...grpc.CallOption) (*shared.ProcessResultApprovalResponse, error)
	// Returns the finalized state transitions at the requested heights.
	GetFinalizedStateTransitions(ctx context.Context, in *shared.FinalizedStateTransitionsRequest, opts ...grpc.CallOption) (*shared.FinalizedStateTransitionsResponse, error)
	// Process state transition proposal from other security node.
	ProcessStateTransitionProposal(ctx context.Context, in *shared.SignedStateTransition, opts ...grpc.CallOption) (*empty.Empty, error)
	// Process state transition prepare vote from other security node.
	ProcessStateTransitionPrepareVote(ctx context.Context, in *shared.SignedStateTransitionPrepareVote, opts ...grpc.CallOption) (*empty.Empty, error)
	// Process state transition commit vote from other security node.
	ProcessStateTransitionCommitVote(ctx context.Context, in *shared.SignedStateTransitionCommitVote, opts ...grpc.CallOption) (*empty.Empty, error)
	// Process execution result from execute nodes to propose block seals
	ProcessExecutionReceipt(ctx context.Context, in *shared.ProcessExecutionReceiptRequest, opts ...grpc.CallOption) (*shared.ProcessExecutionReceiptResponse, error)
	// Receive an execution receipt challenge.
	SubmitInvalidExecutionReceiptChallenge(ctx context.Context, in *shared.InvalidExecutionReceiptChallengeRequest, opts ...grpc.CallOption) (*shared.InvalidExecutionReceiptChallengeResponse, error)
}

type securityNodeClient struct {
	cc *grpc.ClientConn
}

func NewSecurityNodeClient(cc *grpc.ClientConn) SecurityNodeClient {
	return &securityNodeClient{cc}
}

func (c *securityNodeClient) Ping(ctx context.Context, in *shared.PingRequest, opts ...grpc.CallOption) (*shared.PingResponse, error) {
	out := new(shared.PingResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) AccessSubmitCollection(ctx context.Context, in *shared.AccessCollectionRequest, opts ...grpc.CallOption) (*shared.AccessCollectionResponse, error) {
	out := new(shared.AccessCollectionResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/AccessSubmitCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) ProposeBlock(ctx context.Context, in *shared.ProposeBlockRequest, opts ...grpc.CallOption) (*shared.ProposeBlockResponse, error) {
	out := new(shared.ProposeBlockResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/ProposeBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) UpdateProposedBlock(ctx context.Context, in *shared.ProposeBlockUpdateRequest, opts ...grpc.CallOption) (*shared.ProposeBlockUpdateResponse, error) {
	out := new(shared.ProposeBlockUpdateResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/UpdateProposedBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) GetBlockByHash(ctx context.Context, in *shared.GetBlockByHashRequest, opts ...grpc.CallOption) (*shared.GetBlockResponse, error) {
	out := new(shared.GetBlockResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/GetBlockByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) GetBlockByHeight(ctx context.Context, in *shared.GetBlockByHeightRequest, opts ...grpc.CallOption) (*shared.GetBlockResponse, error) {
	out := new(shared.GetBlockResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/GetBlockByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) ProcessResultApproval(ctx context.Context, in *shared.ProcessResultApprovalRequest, opts ...grpc.CallOption) (*shared.ProcessResultApprovalResponse, error) {
	out := new(shared.ProcessResultApprovalResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/ProcessResultApproval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) GetFinalizedStateTransitions(ctx context.Context, in *shared.FinalizedStateTransitionsRequest, opts ...grpc.CallOption) (*shared.FinalizedStateTransitionsResponse, error) {
	out := new(shared.FinalizedStateTransitionsResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/GetFinalizedStateTransitions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) ProcessStateTransitionProposal(ctx context.Context, in *shared.SignedStateTransition, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/ProcessStateTransitionProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) ProcessStateTransitionPrepareVote(ctx context.Context, in *shared.SignedStateTransitionPrepareVote, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/ProcessStateTransitionPrepareVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) ProcessStateTransitionCommitVote(ctx context.Context, in *shared.SignedStateTransitionCommitVote, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/ProcessStateTransitionCommitVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) ProcessExecutionReceipt(ctx context.Context, in *shared.ProcessExecutionReceiptRequest, opts ...grpc.CallOption) (*shared.ProcessExecutionReceiptResponse, error) {
	out := new(shared.ProcessExecutionReceiptResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/ProcessExecutionReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityNodeClient) SubmitInvalidExecutionReceiptChallenge(ctx context.Context, in *shared.InvalidExecutionReceiptChallengeRequest, opts ...grpc.CallOption) (*shared.InvalidExecutionReceiptChallengeResponse, error) {
	out := new(shared.InvalidExecutionReceiptChallengeResponse)
	err := c.cc.Invoke(ctx, "/bamboo.proto.SecurityNode/SubmitInvalidExecutionReceiptChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityNodeServer is the server API for SecurityNode service.
type SecurityNodeServer interface {
	Ping(context.Context, *shared.PingRequest) (*shared.PingResponse, error)
	// Receive a signed collection from an access node.
	AccessSubmitCollection(context.Context, *shared.AccessCollectionRequest) (*shared.AccessCollectionResponse, error)
	// Notify the security node that another has proposed a block.
	ProposeBlock(context.Context, *shared.ProposeBlockRequest) (*shared.ProposeBlockResponse, error)
	// Update a block proposal to add new signatures.
	UpdateProposedBlock(context.Context, *shared.ProposeBlockUpdateRequest) (*shared.ProposeBlockUpdateResponse, error)
	// Returns a block by hash
	GetBlockByHash(context.Context, *shared.GetBlockByHashRequest) (*shared.GetBlockResponse, error)
	// Returns a block by height
	GetBlockByHeight(context.Context, *shared.GetBlockByHeightRequest) (*shared.GetBlockResponse, error)
	// Process result approval from access nodes.
	ProcessResultApproval(context.Context, *shared.ProcessResultApprovalRequest) (*shared.ProcessResultApprovalResponse, error)
	// Returns the finalized state transitions at the requested heights.
	GetFinalizedStateTransitions(context.Context, *shared.FinalizedStateTransitionsRequest) (*shared.FinalizedStateTransitionsResponse, error)
	// Process state transition proposal from other security node.
	ProcessStateTransitionProposal(context.Context, *shared.SignedStateTransition) (*empty.Empty, error)
	// Process state transition prepare vote from other security node.
	ProcessStateTransitionPrepareVote(context.Context, *shared.SignedStateTransitionPrepareVote) (*empty.Empty, error)
	// Process state transition commit vote from other security node.
	ProcessStateTransitionCommitVote(context.Context, *shared.SignedStateTransitionCommitVote) (*empty.Empty, error)
	// Process execution result from execute nodes to propose block seals
	ProcessExecutionReceipt(context.Context, *shared.ProcessExecutionReceiptRequest) (*shared.ProcessExecutionReceiptResponse, error)
	// Receive an execution receipt challenge.
	SubmitInvalidExecutionReceiptChallenge(context.Context, *shared.InvalidExecutionReceiptChallengeRequest) (*shared.InvalidExecutionReceiptChallengeResponse, error)
}

func RegisterSecurityNodeServer(s *grpc.Server, srv SecurityNodeServer) {
	s.RegisterService(&_SecurityNode_serviceDesc, srv)
}

func _SecurityNode_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).Ping(ctx, req.(*shared.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_AccessSubmitCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.AccessCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).AccessSubmitCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/AccessSubmitCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).AccessSubmitCollection(ctx, req.(*shared.AccessCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_ProposeBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.ProposeBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).ProposeBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/ProposeBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).ProposeBlock(ctx, req.(*shared.ProposeBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_UpdateProposedBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.ProposeBlockUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).UpdateProposedBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/UpdateProposedBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).UpdateProposedBlock(ctx, req.(*shared.ProposeBlockUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_GetBlockByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.GetBlockByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).GetBlockByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/GetBlockByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).GetBlockByHash(ctx, req.(*shared.GetBlockByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_GetBlockByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.GetBlockByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).GetBlockByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/GetBlockByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).GetBlockByHeight(ctx, req.(*shared.GetBlockByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_ProcessResultApproval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.ProcessResultApprovalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).ProcessResultApproval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/ProcessResultApproval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).ProcessResultApproval(ctx, req.(*shared.ProcessResultApprovalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_GetFinalizedStateTransitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.FinalizedStateTransitionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).GetFinalizedStateTransitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/GetFinalizedStateTransitions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).GetFinalizedStateTransitions(ctx, req.(*shared.FinalizedStateTransitionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_ProcessStateTransitionProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.SignedStateTransition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).ProcessStateTransitionProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/ProcessStateTransitionProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).ProcessStateTransitionProposal(ctx, req.(*shared.SignedStateTransition))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_ProcessStateTransitionPrepareVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.SignedStateTransitionPrepareVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).ProcessStateTransitionPrepareVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/ProcessStateTransitionPrepareVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).ProcessStateTransitionPrepareVote(ctx, req.(*shared.SignedStateTransitionPrepareVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_ProcessStateTransitionCommitVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.SignedStateTransitionCommitVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).ProcessStateTransitionCommitVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/ProcessStateTransitionCommitVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).ProcessStateTransitionCommitVote(ctx, req.(*shared.SignedStateTransitionCommitVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_ProcessExecutionReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.ProcessExecutionReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).ProcessExecutionReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/ProcessExecutionReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).ProcessExecutionReceipt(ctx, req.(*shared.ProcessExecutionReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityNode_SubmitInvalidExecutionReceiptChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(shared.InvalidExecutionReceiptChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityNodeServer).SubmitInvalidExecutionReceiptChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bamboo.proto.SecurityNode/SubmitInvalidExecutionReceiptChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityNodeServer).SubmitInvalidExecutionReceiptChallenge(ctx, req.(*shared.InvalidExecutionReceiptChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecurityNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bamboo.proto.SecurityNode",
	HandlerType: (*SecurityNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SecurityNode_Ping_Handler,
		},
		{
			MethodName: "AccessSubmitCollection",
			Handler:    _SecurityNode_AccessSubmitCollection_Handler,
		},
		{
			MethodName: "ProposeBlock",
			Handler:    _SecurityNode_ProposeBlock_Handler,
		},
		{
			MethodName: "UpdateProposedBlock",
			Handler:    _SecurityNode_UpdateProposedBlock_Handler,
		},
		{
			MethodName: "GetBlockByHash",
			Handler:    _SecurityNode_GetBlockByHash_Handler,
		},
		{
			MethodName: "GetBlockByHeight",
			Handler:    _SecurityNode_GetBlockByHeight_Handler,
		},
		{
			MethodName: "ProcessResultApproval",
			Handler:    _SecurityNode_ProcessResultApproval_Handler,
		},
		{
			MethodName: "GetFinalizedStateTransitions",
			Handler:    _SecurityNode_GetFinalizedStateTransitions_Handler,
		},
		{
			MethodName: "ProcessStateTransitionProposal",
			Handler:    _SecurityNode_ProcessStateTransitionProposal_Handler,
		},
		{
			MethodName: "ProcessStateTransitionPrepareVote",
			Handler:    _SecurityNode_ProcessStateTransitionPrepareVote_Handler,
		},
		{
			MethodName: "ProcessStateTransitionCommitVote",
			Handler:    _SecurityNode_ProcessStateTransitionCommitVote_Handler,
		},
		{
			MethodName: "ProcessExecutionReceipt",
			Handler:    _SecurityNode_ProcessExecutionReceipt_Handler,
		},
		{
			MethodName: "SubmitInvalidExecutionReceiptChallenge",
			Handler:    _SecurityNode_SubmitInvalidExecutionReceiptChallenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inter/security.proto",
}
