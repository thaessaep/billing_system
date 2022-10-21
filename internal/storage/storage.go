package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	config                 *Config
	db                     *sql.DB
	userRepository         *UserRepository
	reserveBillsRepository *ReserveBillsRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	// ping for check to connect
	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		storage: s,
	}

	return s.userRepository
}

func (s *Storage) ReserveBills() *ReserveBillsRepository {
	if s.reserveBillsRepository != nil {
		return s.reserveBillsRepository
	}

	s.reserveBillsRepository = &ReserveBillsRepository{
		storage: s,
	}

	return s.reserveBillsRepository
}
