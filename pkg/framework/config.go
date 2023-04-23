package framework

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ViperOption struct {
	CfPath    string
	EnvPrefix string
}

type Configuration interface {
	Hook(v *viper.Viper)
}

type VConfig struct{}

func (vc *VConfig) Hook(v *viper.Viper) {}

func LoadConfiguration(o *ViperOption, instance Configuration) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(o.CfPath)
	v.SetEnvPrefix(o.EnvPrefix)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, CommonInternalErr.Wrap(err, "load configuration error")
	}
	envs := os.Environ()
	for _, envPair := range envs {
		if envPair[0:len(o.EnvPrefix)] == o.EnvPrefix {
			_ = v.BindEnv(strings.ReplaceAll(strings.Split(envPair, "=")[0], o.EnvPrefix+"_", ""))
		}
	}
	if err := v.Unmarshal(instance); err != nil {
		return nil, CommonInternalErr.Wrap(err, "unmarshal configuration error")
	}

	instance.Hook(v)
	return nil, nil
}
