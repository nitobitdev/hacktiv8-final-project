package config

import (
	"final-project/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host    = "localhost"
	port    = "5432"
	user    = "postgres"
	pass    = "root"
	db_name = "final_project"
)

func ConnectGorm() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, db_name, port)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

	return db
}
