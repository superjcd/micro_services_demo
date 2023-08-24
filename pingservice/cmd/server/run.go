package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/superjcd/micro_services_demo/pingservice/config"
)

func Run(cfg string) {
	conf := config.NewConfig(cfg)

	grpcServer, err := InitServer(conf)

	if err != nil {
		panic("Initialize server failed")
	}

	RunGrpcServer(grpcServer, conf)

	RunHttpServer(conf)

	HandleGracefulShutDown(conf)

}

func HandleGracefulShutDown(conf *config.Config) {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conf.Grpc.Server.GracefulStop()

	if err := conf.Http.Server.Shutdown(ctx); err != nil {
		panic("shutdown service failed.[ERROR]=>" + err.Error())
	}
	<-ctx.Done()
	close(ch)
	fmt.Println("Graceful shutdown http & grpc server.")

}
