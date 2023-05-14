package config

type Config struct {
	port             int
	connectionString string
	dir              string
}

func New() *Config {
	return &Config{
		port:             8000,
		connectionString: "redis://redis:80",
		dir:              "home/",
	}
}

func (c *Config) GetPort() int {
	return c.port
}

func (c *Config) GetConnectionString() string {
	return c.connectionString
}

func (c *Config) GetDir() string {
	return c.dir
}
