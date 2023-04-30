package config

import (
	"errors"
)

type MongodbConfig struct {
	Name string
}

type MongoEnv struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func (c *MongodbConfig) get() interface{} {
	return &MongoEnv{}
}

func (c *MongodbConfig) GetConfig() (*MongoEnv, error) {
	rawCfg, err := GetConfig(c.Name, c)
	if err != nil {
		return nil, err
	}
	cfg, ok := rawCfg.(*MongoEnv)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	return cfg, nil
}
