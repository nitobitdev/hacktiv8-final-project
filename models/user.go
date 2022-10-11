package models

import (
	"time"
)

type User struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username" validate:"required,unique" binding:"required"`
	Email      string    `json:"email" validate:"required,email,unique" binding:"required"`
	Password   string    `json:"password" validate:"required,minlength=6" binding:"required"`
	Age        int       `json:"age" validate:"required,min=8" binding:"required"`
	Created_at time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type RequestUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,min=8" `
}

type RequestUpdateUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type ResponseUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type ResponseUpdateUser struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Age        int       `json:"age"`
	Updated_at time.Time `json:"updatedAt"`
}

type RequestLoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ResponseLoginUser struct {
	Token string `json:"token"`
}

type ResponseUserForPhoto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponseUserForComment struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponseUserForSocialMedia struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	Profile_image_url string `json:"profile_image_url"`
}
