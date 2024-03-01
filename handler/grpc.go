package handler

import (
	"context"

	tkgorm "github.com/infobloxopen/atlas-app-toolkit/gorm"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type contextKey string

const (
	DBSession contextKey = "dbSession"
)

func NewGRPCServer(db *gorm.DB) (*grpc.Server, error) {
	// var grpcServer grpc.Server
	// return &grpcServer, nil

	grpcServer := grpc.NewServer(
		// grpc.ChainStreamInterceptor(
		// 	DBStreamServerInterceptor(dbSession),
		// ),
		grpc.ChainUnaryInterceptor(
			DBUnaryServerInterceptor(db),
		),
	)
	return grpcServer, nil
}

func DBUnaryServerInterceptor(db *gorm.DB) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//return handler(context.WithValue(ctx, DBSession, session), req)
		return tkgorm.UnaryServerInterceptor(db)(ctx, req, info, handler)
	}
}

// func DBStreamServerInterceptor(session *gorm.DB) grpc.StreamServerInterceptor {
// 	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
// 		wrapped := grpc_middleware.WrapServerStream(stream)
// 		wrapped.WrappedContext = context.WithValue(stream.Context(), DBSession, session)
// 		return handler(srv, wrapped)
// 	}
// }
