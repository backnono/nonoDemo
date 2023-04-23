// Package mysql @Author nono.he 2023/4/20 17:11:00
package mysql

type Config struct {
	Host     string `yaml:"host" json:"host" mapstructure:"host"`
	Port     string `yaml:"port" json:"port" mapstructure:"port"`
	User     string `yaml:"user" json:"user" mapstructure:"user"`
	Password string `yaml:"password" json:"password" mapstructure:"password"`
	Name     string `yaml:"name" json:"name" mapstructure:"name"`
	ShowSql  bool   `yaml:"showsql" json:"showsql" mapstructure:"showsql"`
}
