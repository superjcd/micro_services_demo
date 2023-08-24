package main

import (
	"flag"

	"github.com/superjcd/micro_services_demo/pingservice/cmd/server"
)

var cfg = flag.String("config", "config/config.yaml", "config file")

func main() {
	server.Run(*cfg)
}
