// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/superjcd/micro_services_demo/pongservice/config"
	"github.com/superjcd/micro_services_demo/pongservice/genproto/v1"
	"github.com/superjcd/micro_services_demo/pongservice/service"
)

// Injectors from wire.go:

func InitServer(conf *config.Config) (v1.PongServiceServer, error) {
	pongServiceClient, err := service.NewClient(conf)
	if err != nil {
		return nil, err
	}
	pongServiceServer := service.NewServer(conf, pongServiceClient)
	return pongServiceServer, nil
}
