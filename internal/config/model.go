package config

type Config struct {
	Server Server
	DB     DB
}

type DB struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME" envDefault:"orders_service_dev"`
}

type Server struct {
	Port string `env:"SERVER_PORT" envDefault:":3000"`
}
