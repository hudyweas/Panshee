package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GRPC_ADDRESS   string `mapstructure:"PANSHEE_EMAIL_SERVICE_GRPC_ADDRESS"`
	EMAIL_ADDRESS  string `mapstructure:"PANSHEE_EMAIL_SERVICE_EMAIL_ADDRESS"`
	EMAIL_PASSWORD string `mapstructure:"PANSHEE_EMAIL_SERVICE_EMAIL_PASSWORD"`
}

func LoadConfigFromFile(configFileName string) (config Config, err error) {
	viper.SetConfigFile(configFileName)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PANSHEE_EMAIL_SERVICE")

	err = viper.ReadInConfig()
	if err != nil {
		config = Config{
			GRPC_ADDRESS:   viper.GetString("GRPC_ADDRESS"),
			EMAIL_ADDRESS:  viper.GetString("EMAIL_ADDRESS"),
			EMAIL_PASSWORD: viper.GetString("EMAIL_PASSWORD"),
		}

		err = nil

		return
	}

	err = viper.Unmarshal(&config)

	return
}
