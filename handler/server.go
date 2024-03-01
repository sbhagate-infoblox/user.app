package handler

import (
	"context"
	"log"

	"github.com/sbhagate-infoblox/user.app/pb"
)

func NewUsersServer() (pb.UsersServer, error) {
	return &UsersServer{
		UsersDefaultServer: &pb.UsersDefaultServer{},
	}, nil
}

type UsersServer struct {
	*pb.UsersDefaultServer
	pb.UnimplementedUsersServer
}

// Create implements pb.UsersServer.
func (u *UsersServer) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Println("***Create***")
	res, err := u.UsersDefaultServer.Create(ctx, in)
	return res, err
}

// Delete implements pb.UsersServer.
func (u *UsersServer) Delete(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Println("***Delete***")
	res, err := u.UsersDefaultServer.Delete(ctx, in)
	return res, err
}

// DeleteSet implements pb.UsersServer.
func (u *UsersServer) DeleteSet(ctx context.Context, in *pb.DeleteUserSetRequest) (*pb.DeleteUserSetResponse, error) {
	log.Println("***DeleteSet***")
	res, err := u.UsersDefaultServer.DeleteSet(ctx, in)
	return res, err
}

// Read implements pb.UsersServer.
func (u *UsersServer) Read(ctx context.Context, in *pb.ReadUserRequest) (*pb.ReadUserResponse, error) {
	log.Println("***Read***")
	res, err := u.UsersDefaultServer.Read(ctx, in)
	return res, err
}

// Update implements pb.UsersServer.
func (u *UsersServer) Update(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	log.Println("***Update***")
	res, err := u.UsersDefaultServer.Update(ctx, in)
	return res, err
}
