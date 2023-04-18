package framework

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Configuration interface {
	Hook(v *viper.Viper)
}

func LoadConfiguration(path, envPrefix string, instance Configuration) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(path)
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return CommonInternalErr.Wrap(err, "load configuration error")
	}
	envs := os.Environ()
	for _, envPair := range envs {
		if envPair[0:len(envPrefix)] == envPrefix {
			_ = v.BindEnv(strings.ReplaceAll(strings.Split(envPair, "=")[0], envPrefix+"_", ""))
		}
	}
	if err := v.Unmarshal(instance); err != nil {
		return CommonInternalErr.Wrap(err, "unmarshal configuration error")
	}

	instance.Hook(v)
	return nil
}
