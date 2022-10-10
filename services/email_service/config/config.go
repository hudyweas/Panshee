package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GRPC_ADDRESS string `mapstructure:"PANSHEE_EMAIL_SERVICE_GRPC_ADDRESS"`
	EMAIL_ADDRESS string `mapstructure:"PANSHEE_EMAIL_SERVICE_EMAIL_ADDRESS"`
	EMAIL_PASSWORD string `mapstructure:"PANSHEE_EMAIL_SERVICE_EMAIL_PASSWORD"`
	EMAIL_SERVER string `mapstructure:"PANSHEE_EMAIL_SERVICE_EMAIL_SERVER"`
}

func LoadConfigFromFile(configFileName string) (config Config, err error) {
	viper.SetConfigFile(configFileName)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
