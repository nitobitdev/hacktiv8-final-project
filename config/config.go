package config

import (
	"final-project-gin-go/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host    = "containers-us-west-67.railway.app"
	port    = "6092"
	user    = "postgres"
	pass    = "lTw5yk2P3YXgl7bdEOv8"
	db_name = "railway"
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
