package controllers

import (
	"final-project-gin-go/helpers"
	"final-project-gin-go/models"
	"final-project-gin-go/views"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialMediaControllers struct {
	db *gorm.DB
}

func NewSocialMediaController(db *gorm.DB) *SocialMediaControllers {
	return &SocialMediaControllers{db}
}

func (s *SocialMediaControllers) CreateSocialMedia(ctx *gin.Context) {
	userId := helpers.GetDataToken(ctx)
	socialMedia := models.SocialMedia{}
	requestSocialMedia := models.RequestSocialMedia{}

	if err := ctx.ShouldBindJSON(&requestSocialMedia); err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	socialMedia.Name = requestSocialMedia.Name
	socialMedia.Social_media_url = requestSocialMedia.Social_media_url
	socialMedia.UserId = int(userId)

	if err := s.db.Create(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &views.SuccessResponse{
		Status: http.StatusCreated,
		Data: models.ResponseSocialMedia{
			ID:               socialMedia.ID,
			Name:             socialMedia.Name,
			Social_media_url: socialMedia.Social_media_url,
			UserId:           socialMedia.UserId,
			Created_at:       socialMedia.Created_at,
		},
	})
}

func (s *SocialMediaControllers) GetAllSocialMedia(ctx *gin.Context) {
	var socialMedia []models.SocialMedia
	userId := helpers.GetDataToken(ctx)

	if err := s.db.Preload("User").Where("user_id = ?", userId).Find(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	var responseSocialMedia []models.ResponseGetAllSocialMedia

	for _, value := range socialMedia {
		photo := models.Photo{}
		if err := s.db.Where("user_id = ?", value.UserId).First(&photo).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			})
			return
		}
		responseSocialMedia = append(responseSocialMedia, models.ResponseGetAllSocialMedia{
			ID:               value.ID,
			Name:             value.Name,
			Social_media_url: value.Social_media_url,
			UserId:           value.UserId,
			Created_at:       value.Created_at,
			Updated_at:       value.Updated_at,
			User: models.ResponseUserForSocialMedia{
				ID:                value.User.ID,
				Username:          value.User.Username,
				Profile_image_url: photo.Photo_url,
			},
		})
	}

	ctx.JSON(http.StatusOK, &views.SuccessResponse{
		Status: http.StatusOK,
		Data:   responseSocialMedia},
	)
}

func (s *SocialMediaControllers) UpdateSocialMedia(ctx *gin.Context) {
	socialMedia := models.SocialMedia{}
	requestSocialMedia := models.RequestSocialMedia{}

	if err := ctx.ShouldBindJSON(&requestSocialMedia); err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	userId := helpers.GetDataToken(ctx)
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	if err := s.db.Where("id = ? AND user_id = ?", socialMediaId, int(userId)).First(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	socialMedia.Name = requestSocialMedia.Name
	socialMedia.Social_media_url = requestSocialMedia.Social_media_url

	if err := s.db.Model(&socialMedia).Where("id = ? AND user_id = ?", socialMediaId, int(userId)).Updates(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &views.SuccessResponse{
		Status: http.StatusOK,
		Data: models.ResponseUpdateSocialMedia{
			ID:               socialMedia.ID,
			Name:             socialMedia.Name,
			Social_media_url: socialMedia.Social_media_url,
			UserId:           socialMedia.UserId,
			Updated_at:       socialMedia.Updated_at,
		},
	})
}

func (s *SocialMediaControllers) DeleteSocialMedia(ctx *gin.Context) {
	socialMedia := models.SocialMedia{}

	userId := helpers.GetDataToken(ctx)
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	if err := s.db.Where("id = ? AND user_id = ?", socialMediaId, int(userId)).First(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	if err := s.db.Delete(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
