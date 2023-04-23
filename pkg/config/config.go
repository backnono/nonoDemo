package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"nonoDemo/pkg/adapters/agin"
	"nonoDemo/pkg/adapters/grpc"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability/log"
	"nonoDemo/pkg/utils/observability/mysql"
	"nonoDemo/pkg/utils/observability/redis"
	"nonoDemo/pkg/utils/observability/tracing"
	"os"
	"strconv"
)

type Config struct {
	framework.VConfig
	Database mysql.Config   `yaml:"database"`
	Redis    redis.Config   `yaml:"redis"`
	Trace    tracing.Config `yaml:"trace"`
	Gin      agin.Options   `yaml:"gin"`
	Grpc     grpc.Options   `yaml:"grpc"`
	Log      log.Config     `yaml:"log"`
	Env      string
}

func (cfg *Config) Hook(v *viper.Viper) {
	// env
	env := v.GetString("app.app_env")
	if env == "" {
		env = "dev"
	}
	//Mysql
	cfg.Database.Host = os.Getenv("MYSQL_HOST")
	cfg.Database.User = os.Getenv("MYSQL_USER")
	cfg.Database.Port = os.Getenv("MYSQL_PORT")
	cfg.Database.Password = os.Getenv("MYSQL_PASSWORD")

	//Redis
	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	cfg.Redis.Port = port
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")

	cfg.Env = env
}

var CfgInstance = Config{}

/*func ProvideConfig(config framework.Configuration) Config {
	options := utils.LoadConfig(config, Config{}).(Config)
	ConfigInstance = options
	return ConfigInstance
}*/

func ProvideConfig() Config {
	return CfgInstance
}

var CfgProviderSet = wire.NewSet(
	wire.Bind(new(framework.Configuration), new(*Config)),
	wire.Value(&CfgInstance),
	ProvideConfig,
)
