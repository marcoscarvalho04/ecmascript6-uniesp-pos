package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (s *ConfigurationDB) InitializeDatabase() error {
	if len(s.ConnectionString) == 0 {
		s.ConnectionString = "host=localhost user=postgres password=postgres dbname=todolist port=5432 sslmode=disable "
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: s.ConnectionString,
	}))

	if err != nil {
		return err
	}
	s.DatabaseConnection = db
	return nil
}
