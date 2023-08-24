package server

import (
	"fmt"
	"log"
	"net"

	"github.com/superjcd/micro_services_demo/pingservice/config"
	v1 "github.com/superjcd/micro_services_demo/pingservice/genproto/v1"
	"google.golang.org/grpc"
)

func RunGrpcServer(server v1.PingServiceServer, conf *config.Config) {
	grpcServer := grpc.NewServer()
	v1.RegisterPingServiceServer(grpcServer, server)

	fmt.Println("listening grpc on:", conf.Grpc.Port)
	listener, err := net.Listen("tcp", conf.Grpc.Port)

	if err != nil {
		panic("listen grpc failed:" + err.Error())
	}

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalln("grpc serve failed, details:" + err.Error())
		}
	}()

	conf.Grpc.Server = grpcServer
}
