package storage

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "user=postgres password=password host=db port=5432 dbname=postgres sslmode=disable",
	}
}
