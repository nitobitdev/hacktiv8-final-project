package main

import (
	"final-project-gin-go/config"
	"final-project-gin-go/controllers"
	"final-project-gin-go/routes"
	"os"
)

func main() {
	db := config.ConnectGorm()

	userControllers := controllers.NewUserController(db)
	photoControllers := controllers.NewPhotoController(db)
	commentControllers := controllers.NewCommentController(db)
	socialMediaControllers := controllers.NewSocialMediaController(db)

	router := routes.NewRouter(userControllers, photoControllers, commentControllers, socialMediaControllers)

	router.Start(":" + os.Getenv("PORT"))
}
