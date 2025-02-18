package db

import (
	"to_do_list/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Postgres struct {
	DB *gorm.DB
}

func NewPostgres() DataBase {
	return &Postgres{
		DB: &gorm.DB{},
	}
} 

func (p *Postgres) Connect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	p.DB = db
	return nil
}

func (p *Postgres) Init() error {
	return p.DB.AutoMigrate(&models.User{}, &models.Todo{})
}


