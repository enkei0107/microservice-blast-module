package v1hostserver

import (
	"golang.org/x/net/context"
	pb "microservice-blast-module/internal/v1/proto-generated/github.com/enkei0107/microservice-blast-modolue/blast/v1"
)

type HostServer struct {
	pb.UnimplementedHostServiceServer
}

func (h *HostServer) CreateHost(_ context.Context, in *pb.CreateHostRequest) (*pb.CreateHostResponse, error) {
	return nil, nil
}
