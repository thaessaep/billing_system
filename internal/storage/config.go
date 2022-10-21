package storage

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "user=postgres password=password host=localhost dbname=postgres sslmode=disable",
	}
}
