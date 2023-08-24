package service

import (
	"context"
	"fmt"
	"time"

	"github.com/superjcd/micro_services_demo/pingservice/config"
	v1 "github.com/superjcd/micro_services_demo/pingservice/genproto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pongclientv1 "github.com/superjcd/micro_services_demo/pingservice/genproto/v1"
)

func NewClient(conf *config.Config) (v1.PingServiceClient, error) {
	serverAddr := conf.Grpc.Host + conf.Grpc.Port
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return v1.NewPingServiceClient(conn), nil

}

// NewPongClient New pong service client
func NewPongClient(conf *config.Config) (pongclientv1.PongServiceClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, conf.Client.Pong, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("dial auth server failed.[ERROR]=>" + err.Error())
		return nil, err
	}
	client := pongclientv1.NewPongServiceClient(conn)
	return client, nil

}
