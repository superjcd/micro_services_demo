package config

import (
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Config struct {
	Grpc   Grpc   `json:"grpc" yaml:"grpc"`
	Http   Http   `json:"http" yaml:"http"`
	Client Client `json:"client" yaml:"client"`
}

type Grpc struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Name   string `json:"name" yaml:"name"`
	Server *grpc.Server
}

type Http struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Name   string `json:"name" yaml:"name"`
	Server *http.Server
}

func NewConfig(cfgFile string) *Config {
	conf := &Config{}

	if cfgFile == "" {
		panic("no config file specified")
	}

	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		panic("read config failed, detail:" + err.Error())
	}

	if err := viper.Unmarshal(conf); err != nil {
		panic("unmarshal config failed, detail:" + err.Error())

	}
	return conf
}
