package utils

import (
	"github.com/spf13/viper"
	"testing"
)

type testConfig struct {
	Cfg testCfg1
}

type testCfg1 struct {
	Name string
}

type testConfigWithMultiLevel struct {
	Cfg testConfig
}

func (t *testConfig) Hook(viper *viper.Viper) {
}

func (t *testConfigWithMultiLevel) Hook(viper *viper.Viper) {
}

func TestConfigLoad(t *testing.T) {
	cfg := LoadConfig(&testConfig{Cfg: testCfg1{Name: "abc"}}, testCfg1{})
	c := cfg.(testCfg1)
	if c.Name != "abc" {
		t.Errorf("Expected %s, got %s", "abc", c.Name)
	}
}

func TestConfigLoad_WillReturnNil(t *testing.T) {
	cfg := LoadConfig(&testConfig{Cfg: testCfg1{Name: "abc"}}, struct {
	}{})
	if cfg != nil {
		t.Fatalf("Expected nil, got %v", cfg)
	}
}

func TestLoadConfig_WithMultiLevel(t *testing.T) {
	cfg := LoadConfig(
		&testConfigWithMultiLevel{
			Cfg: testConfig{
				Cfg: testCfg1{
					Name: "abc",
				},
			},
		}, testCfg1{},
	)
	c := cfg.(testCfg1)
	if c.Name != "abc" {
		t.Errorf("Expected %s, got %s", "abc", c.Name)
	}
}

func TestLoadConfig_WithMultiLevel_ShouldReturnNil(t *testing.T) {
	cfg := LoadConfig(
		&testConfigWithMultiLevel{
			Cfg: testConfig{
				Cfg: testCfg1{
					Name: "abc",
				},
			},
		}, struct {
		}{},
	)
	if cfg != nil {
		t.Fatalf("Expected nil, got %v", cfg)
	}
}
