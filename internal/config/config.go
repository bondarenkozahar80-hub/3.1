package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/wb-go/wbf/zlog"
)

type Config struct {
	Postgres   Postgres   `mapstructure:"postgres"`
	HttpServer HttpServer `mapstructure:"http_server"`
	Redis      Redis      `mapstructure:"redis"`
	RabbitMq   RabbitMq   `mapstructure:"rabbitmq"`
}

type Postgres struct {
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type HttpServer struct {
	Address     string `mapstructure:"address"`
	Timeout     int    `mapstructure:"timeout"`
	IdleTimeout int    `mapstructure:"idle_timeout"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type RabbitMq struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

var path string = "internal/config/config.go"
var Cfg Config

func LoadConfig() *Config {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		zlog.Logger.Fatal().Err(err).Msg("Error reading config file")
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.MergeInConfig(); err == nil {
		zlog.Logger.Info().Msg("env file loaded successfully")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		zlog.Logger.Fatal().Err(err).Msg("Error unmarshaling config")
	}

	Cfg = cfg

	return &cfg
}
