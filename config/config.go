package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type ConnectionGetter interface {
	Get() (*Config, error)
}

func GetConfig(name string) (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Sub(fmt.Sprintf(`database.%s`, name)).Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
