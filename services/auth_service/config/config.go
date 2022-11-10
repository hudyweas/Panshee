package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	GRPC_ADDRESS           string        `mapstructure:"PANSHEE_AUTH_SERVICE_GRPC_ADDRESS"`

	DB_HOST                string        `mapstructure:"PANSHEE_AUTH_SERVICE_DB_HOST"`
	DB_PORT                string        `mapstructure:"PANSHEE_AUTH_SERVICE_DB_PORT"`
	DB_USER                string        `mapstructure:"PANSHEE_AUTH_SERVICE_DB_USER"`
	DB_PASSWORD            string        `mapstructure:"PANSHEE_AUTH_SERVICE_DB_PASSWORD"`
	DB_DBNAME              string        `mapstructure:"PANSHEE_AUTH_SERVICE_DB_DBNAME"`

	TOKEN_KEY              string        `mapstructure:"PANSHEE_AUTH_SERVICE_TOKEN_KEY"`
	REFRESH_TOKEN_KEY      string        `mapstructure:"PANSHEE_AUTH_SERVICE_REFRESH_TOKEN_KEY"`
	ACCESS_TOKEN_DURATION  time.Duration `mapstructure:"PANSHEE_AUTH_SERVICE_ACCESS_TOKEN_DURATION"`
	REFRESH_TOKEN_DURATION time.Duration `mapstructure:"PANSHEE_AUTH_SERVICE_REFRESH_TOKEN_DURATION"`
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
