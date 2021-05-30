// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protof

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatBotClient is the client API for ChatBot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatBotClient interface {
	// Sends a message
	Chat(ctx context.Context, in *ChatBotRequest, opts ...grpc.CallOption) (*ChatBotReply, error)
}

type chatBotClient struct {
	cc grpc.ClientConnInterface
}

func NewChatBotClient(cc grpc.ClientConnInterface) ChatBotClient {
	return &chatBotClient{cc}
}

func (c *chatBotClient) Chat(ctx context.Context, in *ChatBotRequest, opts ...grpc.CallOption) (*ChatBotReply, error) {
	out := new(ChatBotReply)
	err := c.cc.Invoke(ctx, "/protof.ChatBot/chat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatBotServer is the server API for ChatBot service.
// All implementations must embed UnimplementedChatBotServer
// for forward compatibility
type ChatBotServer interface {
	// Sends a message
	Chat(context.Context, *ChatBotRequest) (*ChatBotReply, error)
	mustEmbedUnimplementedChatBotServer()
}

// UnimplementedChatBotServer must be embedded to have forward compatible implementations.
type UnimplementedChatBotServer struct {
}

func (UnimplementedChatBotServer) Chat(context.Context, *ChatBotRequest) (*ChatBotReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatBotServer) mustEmbedUnimplementedChatBotServer() {}

// UnsafeChatBotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatBotServer will
// result in compilation errors.
type UnsafeChatBotServer interface {
	mustEmbedUnimplementedChatBotServer()
}

func RegisterChatBotServer(s grpc.ServiceRegistrar, srv ChatBotServer) {
	s.RegisterService(&ChatBot_ServiceDesc, srv)
}

func _ChatBot_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatBotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatBotServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protof.ChatBot/chat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatBotServer).Chat(ctx, req.(*ChatBotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatBot_ServiceDesc is the grpc.ServiceDesc for ChatBot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatBot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protof.ChatBot",
	HandlerType: (*ChatBotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "chat",
			Handler:    _ChatBot_Chat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protof/message.proto",
}