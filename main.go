package main

import (
	"final-project/config"
	"final-project/controllers"
	"final-project/routes"
)

func main() {
	db := config.ConnectGorm()

	userControllers := controllers.NewUserController(db)
	photoControllers := controllers.NewPhotoController(db)
	commentControllers := controllers.NewCommentController(db)
	socialMediaControllers := controllers.NewSocialMediaController(db)

	router := routes.NewRouter(userControllers, photoControllers, commentControllers, socialMediaControllers)

	router.Start(":4000")
}
