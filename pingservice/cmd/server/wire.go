//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/superjcd/micro_services_demo/pingservice/config"
	v1 "github.com/superjcd/micro_services_demo/pingservice/genproto/v1"
	"github.com/superjcd/micro_services_demo/pingservice/service"
)

func InitServer(conf *config.Config) (v1.PingServiceServer, error) {
	wire.Build(
		service.NewPongClient,
		service.NewClient,
		service.NewServer,
	)

	return &service.Server{}, nil
}
