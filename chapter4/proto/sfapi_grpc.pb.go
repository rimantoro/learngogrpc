// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// StarfriendsClient is the client API for Starfriends service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StarfriendsClient interface {
	GetFilm(ctx context.Context, in *GetFilmRequest, opts ...grpc.CallOption) (*GetFilmResponse, error)
	ListFilms(ctx context.Context, in *ListFilmsRequest, opts ...grpc.CallOption) (*ListFilmsResponse, error)
}

type starfriendsClient struct {
	cc grpc.ClientConnInterface
}

func NewStarfriendsClient(cc grpc.ClientConnInterface) StarfriendsClient {
	return &starfriendsClient{cc}
}

func (c *starfriendsClient) GetFilm(ctx context.Context, in *GetFilmRequest, opts ...grpc.CallOption) (*GetFilmResponse, error) {
	out := new(GetFilmResponse)
	err := c.cc.Invoke(ctx, "/Starfriends/GetFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starfriendsClient) ListFilms(ctx context.Context, in *ListFilmsRequest, opts ...grpc.CallOption) (*ListFilmsResponse, error) {
	out := new(ListFilmsResponse)
	err := c.cc.Invoke(ctx, "/Starfriends/ListFilms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StarfriendsServer is the server API for Starfriends service.
// All implementations must embed UnimplementedStarfriendsServer
// for forward compatibility
type StarfriendsServer interface {
	GetFilm(context.Context, *GetFilmRequest) (*GetFilmResponse, error)
	ListFilms(context.Context, *ListFilmsRequest) (*ListFilmsResponse, error)
	mustEmbedUnimplementedStarfriendsServer()
}

// UnimplementedStarfriendsServer must be embedded to have forward compatible implementations.
type UnimplementedStarfriendsServer struct {
}

func (UnimplementedStarfriendsServer) GetFilm(context.Context, *GetFilmRequest) (*GetFilmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilm not implemented")
}
func (UnimplementedStarfriendsServer) ListFilms(context.Context, *ListFilmsRequest) (*ListFilmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFilms not implemented")
}
func (UnimplementedStarfriendsServer) mustEmbedUnimplementedStarfriendsServer() {}

// UnsafeStarfriendsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StarfriendsServer will
// result in compilation errors.
type UnsafeStarfriendsServer interface {
	mustEmbedUnimplementedStarfriendsServer()
}

func RegisterStarfriendsServer(s grpc.ServiceRegistrar, srv StarfriendsServer) {
	s.RegisterService(&Starfriends_ServiceDesc, srv)
}

func _Starfriends_GetFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarfriendsServer).GetFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Starfriends/GetFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarfriendsServer).GetFilm(ctx, req.(*GetFilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starfriends_ListFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarfriendsServer).ListFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Starfriends/ListFilms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarfriendsServer).ListFilms(ctx, req.(*ListFilmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Starfriends_ServiceDesc is the grpc.ServiceDesc for Starfriends service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Starfriends_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Starfriends",
	HandlerType: (*StarfriendsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFilm",
			Handler:    _Starfriends_GetFilm_Handler,
		},
		{
			MethodName: "ListFilms",
			Handler:    _Starfriends_ListFilms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sfapi.proto",
}
