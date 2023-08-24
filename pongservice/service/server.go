package service

import (
	"context"

	"github.com/superjcd/micro_services_demo/pongservice/config"
	v1 "github.com/superjcd/micro_services_demo/pongservice/genproto/v1"
)

type Server struct {
	v1.UnimplementedPongServiceServer
	pongClient v1.PongServiceClient
	conf       *config.Config
}

func NewServer(conf *config.Config, pongClient v1.PongServiceClient) v1.PongServiceServer {
	return &Server{
		pongClient: pongClient,
		conf:       conf,
	}
}

func (s *Server) Pong(ctx context.Context, req *v1.PongRequest) (*v1.PongResponse, error) {
	return &v1.PongResponse{Msg: "response pong msg:" + req.Msg}, nil
}
