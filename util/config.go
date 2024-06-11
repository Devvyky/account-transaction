package util

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config stores all app configurations
// The values are read by viper from a config file or env variable
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configuration from file or env variable
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Config file not found, falling back to environment variables")
		config.DBDriver = os.Getenv("DB_DRIVER")
		config.DBSource = os.Getenv("DB_SOURCE")
		config.ServerAddress = os.Getenv("SERVER_ADDRESS")
		return config, nil
	} else {
		err = viper.Unmarshal(&config)
	}
	return
}
