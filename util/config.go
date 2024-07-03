package util

import "github.com/spf13/viper"

// Config stores all configuration that
// are read by viper from a env file
type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	ServerAddr string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig loads configurations from env file.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
