package config

type RedisConfig struct {
	Name string
}

func (c *RedisConfig) Get() (*Config, error) {
	cfg, err := GetConfig(c.Name)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
