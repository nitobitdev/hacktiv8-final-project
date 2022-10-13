package controllers

import (
	"final-project-gin-go/helpers"
	"final-project-gin-go/models"
	"final-project-gin-go/views"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoControllers struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoControllers {
	return &PhotoControllers{db}
}

func (p *PhotoControllers) GetAllPhoto(ctx *gin.Context) {
	var photos []models.Photo
	if err := p.db.Preload("User").Find(&photos).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	var response []models.ResponseGetAllPhoto

	for _, photo := range photos {
		response = append(response, models.ResponseGetAllPhoto{
			ID:         photo.ID,
			Title:      photo.Title,
			Photo_url:  photo.Photo_url,
			User_id:    photo.User_id,
			Created_at: photo.Created_at,
			Updated_at: photo.Updated_at,
			User: models.ResponseUserForPhoto{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}
	fmt.Println(response)

	ctx.JSON(http.StatusOK, &views.SuccessResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

func (p *PhotoControllers) CreatePhoto(ctx *gin.Context) {
	userId := helpers.GetDataToken(ctx)
	photo := models.RequestPhoto{}
	photoUser := models.Photo{}

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	photoUser.Title = photo.Title
	photoUser.Caption = photo.Caption
	photoUser.Photo_url = photo.Photo_url
	photoUser.User_id = int(userId)

	if err := p.db.Create(&photoUser).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &views.SuccessResponse{
		Status: http.StatusCreated,
		Data: models.ResponseCreatePhoto{
			ID:         photoUser.ID,
			Title:      photoUser.Title,
			Caption:    photoUser.Caption,
			Photo_url:  photoUser.Photo_url,
			User_id:    photoUser.User_id,
			Created_at: photoUser.Created_at,
		},
	})
}

func (p *PhotoControllers) UpdatePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	userId := helpers.GetDataToken(ctx)
	photo := models.Photo{}

	photoUpdate := models.RequestPhoto{}

	if err := ctx.ShouldBindJSON(&photoUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	if err := p.db.Where("id = ? AND user_id = ?", photoId, userId).First(&photo).Error; err != nil {
		ctx.JSON(http.StatusNotFound, &views.ErrorResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		})
		return
	}

	if photoUpdate.Caption != "" {
		photo.Caption = photoUpdate.Caption
	}

	if err := p.db.Model(&photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:      photoUpdate.Title,
		Caption:    photo.Caption,
		Photo_url:  photoUpdate.Photo_url,
		Updated_at: time.Time{},
	}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &views.SuccessResponse{
		Status: http.StatusOK,
		Data: models.ResponseUpdatePhoto{
			ID:         photo.ID,
			Title:      photo.Title,
			Caption:    photo.Caption,
			Photo_url:  photo.Photo_url,
			User_id:    photo.User_id,
			Updated_at: photo.Updated_at,
		},
	})
}

func (p *PhotoControllers) DeletePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	userId := helpers.GetDataToken(ctx)
	photo := models.Photo{}

	if err := p.db.Where("id = ? AND user_id = ?", photoId, userId).First(&photo).Error; err != nil {
		ctx.JSON(http.StatusNotFound, &views.ErrorResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		})
		return
	}

	if err := p.db.Delete(&photo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
