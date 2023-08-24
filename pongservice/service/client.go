package service

import (
	"context"
	"time"

	"github.com/superjcd/micro_services_demo/pongservice/config"
	v1 "github.com/superjcd/micro_services_demo/pongservice/genproto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(conf *config.Config) (v1.PongServiceClient, error) {
	serverAddr := conf.Grpc.Host + conf.Grpc.Port
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return v1.NewPongServiceClient(conn), nil

}
