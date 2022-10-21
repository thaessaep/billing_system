package httpserver

import "github.com/thaessaep/billingSystem/internal/storage"

type Config struct {
	BindAddr string
	Storage  *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Storage:  storage.NewConfig(),
	}
}
