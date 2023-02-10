package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ADDRESS                string        `mapstructure:"PANSHEE_API_SERVICE_ADDRESS"`
	GRPC_ADDRESS           string        `mapstructure:"PANSHEE_API_SERVICE_GRPC_ADDRESS"`

	DB_HOST                string        `mapstructure:"PANSHEE_API_SERVICE_DB_HOST"`
	DB_PORT                string        `mapstructure:"PANSHEE_API_SERVICE_DB_PORT"`
	DB_USER                string        `mapstructure:"PANSHEE_API_SERVICE_DB_USER"`
	DB_PASSWORD            string        `mapstructure:"PANSHEE_API_SERVICE_DB_PASSWORD"`
	DB_DBNAME              string        `mapstructure:"PANSHEE_API_SERVICE_DB_DBNAME"`

	EMAIL_SERVICE_GRPC_ADDRESS string `mapstructure:"PANSHEE_API_SERVICE_EMAIL_SERVICE_GRPC_ADDRESS"`
	AUTH_SERVICE_GRPC_ADDRESS string `mapstructure:"PANSHEE_API_SERVICE_AUTH_SERVICE_GRPC_ADDRESS"`
	WALLET_SERVICE_GRPC_ADDRESS string `mapstructure:"PANSHEE_API_SERVICE_WALLET_SERVICE_GRPC_ADDRESS"`
}

func LoadConfigFromFile(configFileName string) (config Config, err error) {
	viper.SetConfigFile(configFileName)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PANSHEE_API_SERVICE")

	err = viper.ReadInConfig()
	if err != nil {
		config = Config{
			ADDRESS:                viper.GetString("ADDRESS"),
			GRPC_ADDRESS:           viper.GetString("GRPC_ADDRESS"),

			DB_HOST:                viper.GetString("DB_HOST"),
			DB_PORT:                viper.GetString("DB_PORT"),
			DB_USER:                viper.GetString("DB_USER"),
			DB_PASSWORD:            viper.GetString("DB_PASSWORD"),
			DB_DBNAME:              viper.GetString("DB_DBNAME"),

			EMAIL_SERVICE_GRPC_ADDRESS: viper.GetString("EMAIL_SERVICE_GRPC_ADDRESS"),
		}

		err = nil

		return
	}

	err = viper.Unmarshal(&config)

	return
}
