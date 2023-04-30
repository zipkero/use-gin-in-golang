package config

import "errors"

type RedisConfig struct {
	Name string
}

type RedisEnv struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

func (c *RedisConfig) get() interface{} {
	return &RedisEnv{}
}

func (c *RedisConfig) GetConfig() (*RedisEnv, error) {
	rawCfg, err := GetConfig(c.Name, c)
	if err != nil {
		return nil, err
	}
	cfg, ok := rawCfg.(*RedisEnv)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	return cfg, nil
}
