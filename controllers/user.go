package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"final-project/views"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserControllers struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserControllers {
	return &UserControllers{db}
}

var validate *validator.Validate

func (u *UserControllers) CreateUsers(ctx *gin.Context) {
	users := models.RequestUser{}
	user := models.User{}

	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	u.db.Where("email = ?", users.Email).First(&user)

	if len := len(user.Email); len > 0 {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  "Email already exists",
		})
		return
	}

	u.db.Where("username = ?", users.Username).First(&user)

	if len := len(user.Username); len > 0 {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  "Username already exists",
		})
		return
	}

	password := helpers.HashPass(users.Password)

	newUser := models.User{
		Username: users.Username,
		Password: password,
		Email:    users.Email,
		Age:      users.Age,
	}

	if err := u.db.Create(&newUser).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, &views.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &views.SuccessResponse{
		Status: http.StatusCreated,
		Data: models.ResponseUser{
			ID:       newUser.ID,
			Username: newUser.Username,
			Email:    newUser.Email,
			Age:      newUser.Age,
		},
	})
	return

}

func (u *UserControllers) LoginUser(ctx *gin.Context) {
	user := models.RequestLoginUser{}
	userLogin := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	if err := u.db.Where("email = ?", user.Email).First(&userLogin).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "Invalid Email/Password",
			Error:   "Unauthorized",
		})
		return
	}

	checkPassword := helpers.ComparePass([]byte(userLogin.Password), []byte(user.Password))
	if !checkPassword {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "Invalid Email/Password",
			Error:   "Unauthorized",
		})
		return
	}

	token := helpers.GenerateToken(userLogin.ID, userLogin.Email)

	ctx.JSON(http.StatusOK, &views.SuccessResponse{
		Status: http.StatusOK,
		Data: models.ResponseLoginUser{
			Token: token,
		},
	})
}

func (u *UserControllers) UpdateUser(ctx *gin.Context) {
	user := models.User{}
	userUpdate := models.RequestUpdateUser{}

	if err := ctx.ShouldBindJSON(&userUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(ctx.Param("userId"))

	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	u.db.Table("users").Where("id = ?", id).Updates(models.User{
		Username:   userUpdate.Username,
		Email:      userUpdate.Email,
		Updated_at: time.Time{},
	})

	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &views.SuccessResponse{
		Status: http.StatusOK,
		Data: models.ResponseUpdateUser{
			ID:         user.ID,
			Username:   userUpdate.Username,
			Email:      userUpdate.Email,
			Age:        user.Age,
			Updated_at: user.Updated_at,
		},
	})
}

func (u *UserControllers) DeleteUser(ctx *gin.Context) {

	id := helpers.GetDataToken(ctx)

	if err := u.db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, &views.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfuly deleted",
	})
}
