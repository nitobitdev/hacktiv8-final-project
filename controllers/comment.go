package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"final-project/views"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentControllers struct {
	db *gorm.DB
}

func NewCommentController(db *gorm.DB) *CommentControllers {
	return &CommentControllers{db}
}

func (c *CommentControllers) CretaeComment(ctx *gin.Context) {
	requestComment := models.RequestComment{}

	fmt.Println(requestComment)

	userId := helpers.GetDataToken(ctx)

	if err := ctx.ShouldBindJSON(&requestComment); err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	comment := models.Comment{
		Message:  requestComment.Message,
		User_id:  int(userId),
		Photo_id: requestComment.Photo_id,
	}

	if err := c.db.Create(&comment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &views.SuccessResponse{
		Status: http.StatusCreated,
		Data: models.ResponseComment{
			ID:         comment.ID,
			Message:    comment.Message,
			Photo_id:   comment.Photo_id,
			User_id:    comment.User_id,
			Created_at: comment.Created_at,
		},
	})
}

func (c *CommentControllers) GetAllComment(ctx *gin.Context) {
	userId := helpers.GetDataToken(ctx)

	var comments []models.Comment

	if err := c.db.Preload("User").Preload("Photo").Where("user_id = ?", int(userId)).Find(&comments).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	var response []models.ResponseGetComment

	for _, comment := range comments {
		response = append(response, models.ResponseGetComment{
			ID:         comment.ID,
			Message:    comment.Message,
			Photo_id:   comment.Photo_id,
			User_id:    comment.User_id,
			Created_at: comment.Created_at,
			User: models.ResponseUserForComment{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: models.ResponsePhotoForComment{
				ID:        comment.Photo.ID,
				Title:     comment.Photo.Title,
				Caption:   comment.Photo.Caption,
				Photo_url: comment.Photo.Photo_url,
				User_id:   comment.Photo.User_id,
			},
		})
	}

	ctx.JSON(http.StatusOK, &views.SuccessResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

func (c *CommentControllers) UpdateComment(ctx *gin.Context) {
	comment := models.Comment{}
	userId := helpers.GetDataToken(ctx)
	commentId := ctx.Param("commentId")

	requestUpdateCommand := models.RequestUpdateComment{}

	if err := ctx.ShouldBindJSON(&requestUpdateCommand); err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	if err := c.db.Preload("Photo").Where("id = ? AND user_id = ?", commentId, int(userId)).First(&comment).Error; err != nil {
		ctx.JSON(http.StatusNotFound, &views.ErrorResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		})
		return
	}

	if err := c.db.Preload("Photo").Model(&comment).Updates(models.Comment{Message: requestUpdateCommand.Message}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &views.SuccessResponse{
		Status: http.StatusOK,
		Data: models.ResponseUpdatePhoto{
			ID:         comment.ID,
			Title:      comment.Photo.Title,
			Caption:    comment.Photo.Caption,
			Photo_url:  comment.Photo.Photo_url,
			User_id:    comment.Photo.User_id,
			Updated_at: comment.Updated_at,
		},
	})
}

func (c *CommentControllers) DeleteComment(ctx *gin.Context) {
	comment := models.Comment{}
	userId := helpers.GetDataToken(ctx)
	commentId := ctx.Param("commentId")

	if err := c.db.Where("id = ? AND user_id = ?", commentId, int(userId)).First(&comment).Error; err != nil {
		ctx.JSON(http.StatusNotFound, &views.ErrorResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		})
		return
	}

	if err := c.db.Delete(&comment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
