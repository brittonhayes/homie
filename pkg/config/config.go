package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	File     string   `yaml:"file"`
	Google   google   `yaml:"google"`
	Telegram telegram `yaml:"telegram"`
}

type google struct {
	Secrets string `yaml:"secrets"`
	Sheet   sheet  `yaml:"sheet"`
}

type sheet struct {
	ID        string `yaml:"id"`
	HeaderRow int    `yaml:"header_row"`
	Title     string `yaml:"title"`
	Link      string `yaml:"link"`
}

type telegram struct {
	Token   string   `yaml:"token"`
	Allowed []string `yaml:"allowed"`
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
