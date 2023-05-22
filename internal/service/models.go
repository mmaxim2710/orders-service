package service

import (
	_ "github.com/caarlos0/env/v8"
)

type Config struct {
	Server
	DB
}

type DB struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
}

type Server struct {
	Port string `env:"SERVER_PORT" envDefault:":3000"`
}
