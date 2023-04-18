package main

import (
	"nonoDemo/pkg/utils/observability/log"
	"nonoDemo/pkg/utils/observability/tracing"
	"nonoDemo/pkg/adapters/agin"
	"nonoDemo/pkg/adapters/grpc"
	"github.com/spf13/viper"
)

type Config struct {
	Trace tracing.Config `yaml:"trace"`
	Gin   agin.Options   `yaml:"gin"`
	Grpc  grpc.Options   `yaml:"grpc"`
	Log   log.Config     `yaml:"log"`
}

func (cfg *Config) Hook(v *viper.Viper) {
}
    