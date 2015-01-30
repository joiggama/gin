package gin

type Config struct {
	data map[string]interface{}
}

func (c *Config) Get(key string) interface{} {
	return c.data[key]
}

func (c *Config) Set(key string, value interface{}) {
	c.data[key] = value
}

func NewConfig() *Config {
	return &Config{map[string]interface{}{}}
}
