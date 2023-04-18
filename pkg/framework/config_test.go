package framework

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"testing"
)

type testConfig struct {
	ServiceName string          `yaml:"service_name" mapstructure:"service_name"`
	ID          int             `yaml:"ID"`
	Array       []string        `yaml:"array"`
	Redis       testRedisConfig `yaml:"redis"`
	Env         string
	HookStr     string
}

type testRedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (cfg *testConfig) Hook(viper *viper.Viper) {
	cfg.HookStr = "hook"
}

func TestLoadConfiguration(t *testing.T) {
	config := `
service_name: "test"
ID: 1
array:
- a
- b
redis:
  host: localhost
  port: 3306
`
	configPath := "./config.yaml"
	_ = ioutil.WriteFile(configPath, []byte(config), 0644)
	_ = os.Setenv("TEST_ENV", "TEST")
	cfg := &testConfig{}
	err := LoadConfiguration(configPath, "TEST", cfg)
	_ = os.Remove(configPath)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ServiceName != "test" || cfg.Redis.Host != "localhost" || cfg.Redis.Port != "3306" {
		t.Fatal("load from config file error")
	}
	if cfg.Env != "TEST" {
		t.Fatal("load from env error")
	}
	if cfg.HookStr != "hook" {
		t.Fatal("invoke hook error")
	}
}
