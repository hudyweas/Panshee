package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GRPC_ADDRESS   string `mapstructure:"PANSHEE_AUTH_SERVICE_GRPC_ADDRESS"`
}

func LoadConfigFromFile(configFileName string) (config Config, err error) {
	viper.SetConfigFile(configFileName)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PANSHEE_AUTH_SERVICE")

	err = viper.ReadInConfig()
	if err != nil {
		config = Config{
			GRPC_ADDRESS:   viper.GetString("GRPC_ADDRESS"),
		}

		err = nil

		return
	}

	err = viper.Unmarshal(&config)

	return
}
