package config

type MongodbConfig struct {
	Name string
}

func (c *MongodbConfig) Get() (*Config, error) {
	cfg, err := GetConfig(c.Name)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
