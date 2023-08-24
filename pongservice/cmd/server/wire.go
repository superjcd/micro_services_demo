//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/superjcd/micro_services_demo/pongservice/config"
	v1 "github.com/superjcd/micro_services_demo/pongservice/genproto/v1"
	"github.com/superjcd/micro_services_demo/pongservice/service"
)

func InitServer(conf *config.Config) (v1.PongServiceServer, error) {
	wire.Build(
		service.NewClient,
		service.NewServer,
	)

	return &service.Server{}, nil
}
