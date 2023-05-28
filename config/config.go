package config

import "github.com/caarlos0/env/v8"

type Config struct {
	Server Server
	DB     DB
	Log    Log
	JWT    JWT
}

type DB struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME" envDefault:"orders_service_dev"`
	SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

type Log struct {
	Level string `env:"LOG_LEVEL" envDefault:"debug"`
}

type Server struct {
	Port string `env:"SERVER_PORT" envDefault:":3000"`
}

type JWT struct {
	Secret      string `env:"JWT_SECRET"`
	ExpireHours string `env:"JWT_EXPIRE_HOURS" envDefault:"24"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
