package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	AppEnv             string `mapstructure:"APP_ENV"`
	AppName            string `mapstructure:"APP_NAME"`
	GrpcPort           string `mapstructure:"GRPC_PORT"`
	AppVersion         string `mapstructure:"APP_VERSION"`
	GracefulPeriod     int    `mapstructure:"GRACEFUL_PERIOD"`
	PostgresHost       string `mapstructure:"POSTGRES_HOST"`
	PostgresPort       string `mapstructure:"POSTGRES_PORT"`
	PostgresDb         string `mapstructure:"POSTGRES_DB"`
	PostgresUsername   string `mapstructure:"POSTGRES_USERNAME"`
	PostgresPassword   string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresSslEnabled bool   `mapstructure:"POSTGRES_SSL_ENABLED"`
}

func Load(filename string) (*EnvConfig, error) {
	var envCfg EnvConfig

	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&envCfg); err != nil {
		return nil, err
	}

	return &envCfg, nil
}
