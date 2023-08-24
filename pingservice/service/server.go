package service

import (
	"context"

	"github.com/superjcd/micro_services_demo/pingservice/config"
	v1 "github.com/superjcd/micro_services_demo/pingservice/genproto/v1"
)

type Server struct {
	v1.UnimplementedPingServiceServer
	PingClient v1.PingServiceClient
	conf       *config.Config
}

func NewServer(conf *config.Config, PingClient v1.PingServiceClient) v1.PingServiceServer {
	return &Server{
		PingClient: PingClient,
		conf:       conf,
	}
}

func (s *Server) Ping(ctx context.Context, req *v1.PingRequest) (*v1.PingResponse, error) {
	return &v1.PingResponse{Msg: "response Ping msg:" + req.Msg}, nil
}
