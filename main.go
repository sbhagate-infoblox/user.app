package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/logging"
	"github.com/infobloxopen/atlas-app-toolkit/requestinfo"
	"github.com/infobloxopen/atlas-app-toolkit/server"
	_ "github.com/lib/pq"
	"github.com/sbhagate-infoblox/user.app/handler"
	"github.com/sbhagate-infoblox/user.app/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "user_app"
)

func main() {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	dbSQL, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	defer dbSQL.Close()

	err = dbSQL.Ping()
	if err != nil {
		panic(err)
	}

	// sqlStatement := `INSERT INTO users(account_id, compartment_id, user_name) VALUES ('111', 'CMPRT1', 'Saksham')`
	// _, err = dbSQL.Exec(sqlStatement)
	// if err != nil {
	// 	panic(err)
	// }

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbSQL,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	// listner, err := net.Listen("tcp", "localhost:3000")
	// if err != nil {
	// 	log.Println(err)
	// }
	// //s := grpc.NewServer()

	grpcServer, err := handler.NewGRPCServer(db)
	if err != nil {
		log.Println(err)
	}

	ser, err := handler.NewUsersServer()
	if err != nil {
		log.Println(err)
	}

	pb.RegisterUsersServer(grpcServer, ser)

	s, err := server.NewServer(
		// register our grpc server
		server.WithGrpcServer(grpcServer),
		// register the gateway to proxy to the given server address with the service registration endpoints
		server.WithGateway(
			gateway.WithGatewayOptions(
				runtime.WithIncomingHeaderMatcher(CustomIncomingHeaderMatcher),
				runtime.WithMetadata(pb.AtlasValidateAnnotator),
				runtime.WithMetadata(requestinfo.MetadataAnnotator),
				runtime.WithMetadata(gateway.NewPresenceAnnotator("PATCH")),
				runtime.WithMetadata(logging.Annotator),
				runtime.WithForwardResponseOption(forwardResponseOption),
			),
			gateway.WithDialOptions(
				[]grpc.DialOption{grpc.WithInsecure(), grpc.WithUnaryInterceptor(
					grpc_middleware.ChainUnaryClient(
						[]grpc.UnaryClientInterceptor{
							gateway.ClientUnaryInterceptor,
							gateway.PresenceClientInterceptor()}...,
					),
				)}...,
			),
			gateway.WithServerAddress(fmt.Sprintf("%s:%s", "0.0.0.0", "3000")),
			gateway.WithEndpointRegistration("/v1/",
				pb.RegisterUsersHandlerFromEndpoint,
			),
		),
	)
	if err != nil {
		log.Println(err)
	}

	/*
		s, err := server.NewServer(
			// register our grpc server
			server.WithGrpcServer(grpcServer),
			// register the gateway to proxy to the given server address with the service registration endpoints
			server.WithGateway(
				gateway.WithGatewayOptions(),
				gateway.WithDialOptions(
					[]grpc.DialOption{grpc.WithInsecure(), grpc.WithUnaryInterceptor(
						grpc_middleware.ChainUnaryClient(
							[]grpc.UnaryClientInterceptor{
								gateway.ClientUnaryInterceptor,
								gateway.PresenceClientInterceptor()}...,
						),
					)}...,
				),
				gateway.WithServerAddress(fmt.Sprintf("%s:%s", "0.0.0.0", "4000")),
				gateway.WithEndpointRegistration("/v1/",
					pb.RegisterUsersHandlerFromEndpoint,
					// todo fix DEP dependency issue with new recyclebin
					// ***********
					// TODO 	master: Could not introduce github.com/Infoblox-CTO/atlas.recyclebin@master,
					//  as it has a dependency on github.com/grpc-ecosystem/grpc-gateway with constraint ^1.13.0,
					//  which has no overlap with existing constraint 1.9.6 from (root)
				),
			),
		)
		if err != nil {
			log.Println(err)
		}
	*/

	// open some listeners for our server and gateway
	grpcL, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", "3000"))
	if err != nil {
		log.Println(err)
	}
	gatewayL, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", "4000"))
	if err != nil {
		log.Println(err)
	}

	s.Serve(grpcL, gatewayL)

	// err = s.Serve(listner)
	// if err != nil {
	// 	log.Println(err)
	// }

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// grpcMux := runtime.NewServeMux()

	// err = pb.RegisterUsersHandlerServer(ctx, grpcMux, s.)
	// if err != nil {
	// 	log.Println(err)
	// }

	// mux := http.NewServeMux()
	// mux.Handle("/", grpcMux)

	// lis, err := net.Listen("tcp", "localhost:4000")
	// if err != nil {
	// 	log.Println(err)
	// }

	// err = http.Serve(lis, mux)
	// if err != nil {
	// 	log.Println(err)
	// }
}

func CustomIncomingHeaderMatcher(k string) (string, bool) {
	return k, true
}

func forwardResponseOption(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	return nil
}
