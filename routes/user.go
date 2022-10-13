package routes

import (
	"final-project-gin-go/controllers"
	"final-project-gin-go/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	user        *controllers.UserControllers
	photo       *controllers.PhotoControllers
	comment     *controllers.CommentControllers
	socialmedia *controllers.SocialMediaControllers
}

func NewRouter(
	user *controllers.UserControllers, photo *controllers.PhotoControllers, comment *controllers.CommentControllers, socialmedia *controllers.SocialMediaControllers) *Router {
	return &Router{user, photo, comment, socialmedia}
}

func (r *Router) Start(port string) {
	router := gin.Default()

	router.POST("/users/register", r.user.CreateUsers)
	router.POST("/users/login", r.user.LoginUser)

	router.PUT("/users/:userId", middlewares.Authentication(), r.user.UpdateUser)
	router.DELETE("/users", middlewares.Authentication(), r.user.DeleteUser)

	routerPhoto := router.Group("/photos")
	routerPhoto.Use(middlewares.Authentication())
	routerPhoto.POST("/", r.photo.CreatePhoto)
	routerPhoto.GET("/", r.photo.GetAllPhoto)
	routerPhoto.PUT("/:photoId", r.photo.UpdatePhoto)
	routerPhoto.DELETE("/:photoId", r.photo.DeletePhoto)

	routerComment := router.Group("/comments")
	routerComment.Use(middlewares.Authentication())
	routerComment.POST("/", r.comment.CretaeComment)
	routerComment.GET("/", r.comment.GetAllComment)
	routerComment.PUT("/:commentId", r.comment.UpdateComment)
	routerComment.DELETE("/:commentId", r.comment.DeleteComment)

	routerSocialMedia := router.Group("/socialmedias")
	routerSocialMedia.Use(middlewares.Authentication())
	routerSocialMedia.POST("/", r.socialmedia.CreateSocialMedia)
	routerSocialMedia.GET("/", r.socialmedia.GetAllSocialMedia)
	routerSocialMedia.PUT("/:socialMediaId", r.socialmedia.UpdateSocialMedia)
	routerSocialMedia.DELETE("/:socialMediaId", r.socialmedia.DeleteSocialMedia)

	router.Run(port)
}
