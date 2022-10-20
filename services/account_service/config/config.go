package config

import (
	"time"

	"github.com/sirupsen/logrus"
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

	EMAIL_SERVICE_GRPC_ADDRESS string `mapstructure:"PANSHEE_ACCOUNT_SERVICE_EMAIL_SERVICE_GRPC_ADDRESS"`
}

func LoadConfigFromFile(configFileName string, log *logrus.Logger) (config Config, err error) {
	viper.SetConfigFile(configFileName)

	viper.AutomaticEnv()

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
			TOKEN_KEY:              viper.GetString("TOKEN_KEY"),
			ACCESS_TOKEN_DURATION:  time.Duration(viper.GetInt64("ACCESS_TOKEN_DURATION")),
			REFRESH_TOKEN_DURATION: time.Duration(viper.GetInt64("REFRESH_TOKEN_DURATION")),

			EMAIL_SERVICE_GRPC_ADDRESS: viper.GetString("EMAIL_SERVICE_GRPC_ADDRESS"),
		}

		log.Info("config created from env")

		err = nil

		return
	}

	err = viper.Unmarshal(&config)

	log.Info("config created from file")

	return
}
