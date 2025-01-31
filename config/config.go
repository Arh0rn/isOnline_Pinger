package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBMS    string `mapstructure:"dbms"`
	SSLMode string `mapstructure:"ssl_mode"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
