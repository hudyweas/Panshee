package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ADDRESS                string        `mapstructure:"PANSHEE_ACCOUNT_SERVICE_ADDRESS"`
	GRPC_ADDRESS           string        `mapstructure:"PANSHEE_ACCOUNT_SERVICE_GRPC_ADDRESS"`

	DB_HOST                string        `mapstructure:"PANSHEE_ACCOUNT_SERVICE_DB_HOST"`
	DB_PORT                string        `mapstructure:"PANSHEE_ACCOUNT_SERVICE_DB_PORT"`
	DB_USER                string        `mapstructure:"PANSHEE_ACCOUNT_SERVICE_DB_USER"`
	DB_PASSWORD            string        `mapstructure:"PANSHEE_ACCOUNT_SERVICE_DB_PASSWORD"`
	DB_DBNAME              string        `mapstructure:"PANSHEE_ACCOUNT_SERVICE_DB_DBNAME"`

	TOKEN_KEY              string        `mapstructure:"PANSHEE_ACCOUNT_SERVICE_TOKEN_KEY"`
	ACCESS_TOKEN_DURATION  time.Duration `mapstructure:"PANSHEE_ACCOUNT_SERVICE_ACCESS_TOKEN_DURATION"`
	REFRESH_TOKEN_DURATION time.Duration `mapstructure:"PANSHEE_ACCOUNT_SERVICE_REFRESH_TOKEN_DURATION"`
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
