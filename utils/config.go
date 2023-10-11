package utils

import (
	"github.com/spf13/viper"
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

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
