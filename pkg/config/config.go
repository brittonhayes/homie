package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	File   string
	Google struct {
		Secrets string
		Sheet   struct {
			ID        string
			HeaderRow int
			Title     string
		}
	}
	Telegram struct {
		Token   string
		Allowed []string
	}
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Configuration, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".homie")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
