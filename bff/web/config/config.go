package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	AppEnv          string `mapstructure:"APP_ENV"`
	AppName         string `mapstructure:"APP_NAME"`
	RestPort        string `mapstructure:"REST_PORT"`
	AppVersion      string `mapstructure:"APP_VERSION"`
	GracefulPeriod  int    `mapstructure:"GRACEFUL_PERIOD"`
	UserServiceHost string `mapstructure:"USER_SERVICE_HOST"`
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
