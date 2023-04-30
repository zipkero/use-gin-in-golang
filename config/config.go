package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Getter interface {
	get() interface{}
}

func GetConfig(name string, configGetter Getter) (interface{}, error) {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config := configGetter.get()
	if err := viper.Sub(fmt.Sprintf(`database.%s`, name)).Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
