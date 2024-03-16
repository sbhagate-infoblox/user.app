package main

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/sbhagate-infoblox/user.app/handler"
	"github.com/sbhagate-infoblox/user.app/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func TestMain(t *testing.T) {
	_, grpcServer := ReturnGRPCAndDBServer()
	ser, err := handler.NewUsersServer()
	if err != nil {
		log.Println(err)
	}

	pb.RegisterUsersServer(grpcServer, ser)
	listner, err := net.Listen("tcp", "localhost:5050")
	if err != nil {
		log.Println(err)
	}

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Println(err)
	}

	conn, err := grpc.Dial("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	defer listner.Close()

	cli := pb.NewUsersClient(conn)

	JWT1 := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiMTExIiwiY29tcGFydG1lbnRfaWQiOiJhYSIsInVzZXJfbmFtZSI6IlNha3NoYW0ifQ.1JYHFym3CzRqkixtuAbaXGK866HamuyZCZh7RHf_6Sc"
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authorization": "bearer " + JWT1}))
	_, err = cli.Create(ctx, &pb.CreateUserRequest{
		Payload: &pb.User{
			UserName: "Saksham",
		},
	})

	if err != nil {
		t.Fatalf("expected error state not returned in test. actual: %s", err)
	}

	// sqlStatement := `INSERT INTO users(account_id, compartment_id, user_name) VALUES ('11', 'CMPRT1', 'Saksham')`
	// _, err = db.Exec().Exec(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }

}

func IdentifierFromID(in string) *resource.Identifier {
	aname, rtype, rid := resource.ParseString(in)
	return &resource.Identifier{
		ApplicationName: aname,
		ResourceType:    rtype,
		ResourceId:      rid,
	}
}
