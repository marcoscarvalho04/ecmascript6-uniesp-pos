package db

import "gorm.io/gorm"

type ConfigurationDB struct {
	DatabaseConnection *gorm.DB
	ConnectionString   string
	Repository         []Repository
}

type Repository interface {
	Migrate() error
	Save() error
	Update() error
	Delete() error
}
