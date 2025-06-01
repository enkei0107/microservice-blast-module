package handler

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"microservice-blast-module/config/environment"
	v1hostserver "microservice-blast-module/internal/v1/host/server"
	pb "microservice-blast-module/internal/v1/proto-generated/github.com/enkei0107/microservice-blast-modolue/blast/v1"
	"net"
)

func StartGRPCServer(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()

	// Register your gRPC services here
	pb.RegisterHostServiceServer(grpcServer, &v1hostserver.HostServer{})

	if environment.ServerMode != "production" {
		log.Println("reflection on")
		reflection.Register(grpcServer)
	}

	log.Println("starting grpc server on", port)
	return grpcServer.Serve(lis)
}
