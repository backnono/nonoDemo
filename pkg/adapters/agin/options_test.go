package agin

import (
	"github.com/spf13/viper"
	"testing"
)

type config struct {
	Gin Options `yaml:"gin" json:"gin"`
}

func (c *config) Hook(viper *viper.Viper) {

}

func TestNewOptionsFromConfig(t *testing.T) {

	cfg := config{Gin: Options{
		ListenAddr:     ":8080",
		TraceEnabled:   true,
		MetricsEnabled: true,
		middlewares:    nil,
		handlers:       nil,
	}}
	opt := NewOptionsFromConfig(&cfg)

	if opt.ListenAddr != ":8080" {
		t.Fatal("listen addr get error")
	}
	if !opt.TraceEnabled {
		t.Fatal("TraceEnabled get error")
	}
	if !opt.MetricsEnabled {
		t.Fatal("MetricsEnabled get error")
	}
}
