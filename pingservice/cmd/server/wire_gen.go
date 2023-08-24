// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/superjcd/micro_services_demo/pingservice/config"
	"github.com/superjcd/micro_services_demo/pingservice/genproto/v1"
	"github.com/superjcd/micro_services_demo/pingservice/service"
)

// Injectors from wire.go:

func InitServer(conf *config.Config) (v1.PingServiceServer, error) {
	pingServiceClient, err := service.NewClient(conf)
	if err != nil {
		return nil, err
	}
	pingServiceServer := service.NewServer(conf, pingServiceClient)
	return pingServiceServer, nil
}