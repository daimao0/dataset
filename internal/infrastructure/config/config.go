package config

type DatabaseConfig struct {
	DSN string `yaml:"dsn"`

	MaxOpenConn int `yaml:"maxOpenConn"`

	MaxIdleConn int `yaml:"maxIdleConn"`

	MaxLifetime int `yaml:"maxLifetime"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

var AppConfig = &Config{}
