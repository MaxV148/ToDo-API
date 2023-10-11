package utils

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	DBSourceProd  string `mapstructure:"DB_SOURCE_PROD"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	APISecret     string `mapstructure:"API_SECRET"`
	TokenLifespan string `mapstructure:"TOKEN_HOUR_LIFESPAN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, use ENV")
			config.DBDriver = viper.GetString("DB_DRIVER")
			config.DBSource = viper.GetString("DB_SOURCE")
			config.DBSourceProd = viper.GetString("DB_SOURCE_PROD")
			config.ServerAddress = viper.GetString("SERVER_ADDRESS")
			config.APISecret = viper.GetString("API_SECRET")
			config.TokenLifespan = viper.GetString("TOKEN_HOUR_LIFESPAN")
			return config, nil
		} else {
			return
		}
	}
	log.Println("Config file found")

	err = viper.Unmarshal(&config)
	return

}
