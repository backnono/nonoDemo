// Package redis @Author nono.he 2023/4/20 17:17:00
package redis

type Config struct {
	Host     string `yaml:"host" json:"host" mapstructure:"host"`
	Port     int    `yaml:"port" json:"port" mapstructure:"port"`
	Password string `yaml:"password" json:"password" mapstructure:"password"`
	DB       int    `yaml:"db" json:"db" mapstructure:"db"`
}
