package models

import "time"

type Comment struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	User_id    int       `json:"user_id"`
	User       User      `json:"user" gorm:"foreignkey:User_id"`
	Photo_id   int       `json:"photo_id"`
	Photo      Photo     `json:"photo" gorm:"foreignkey:Photo_id"`
	Message    string    `json:"message" validate:"required"`
	Created_at time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type RequestComment struct {
	Message  string `json:"message" binding:"required"`
	Photo_id int    `json:"photo_id"`
}

type ResponseComment struct {
	ID         int       `json:"id"`
	Message    string    `json:"message"`
	Photo_id   int       `json:"photo_id"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `json:"created_at"`
}

type ResponseGetComment struct {
	ID         int                     `json:"id"`
	Message    string                  `json:"message"`
	Photo_id   int                     `json:"photo_id"`
	User_id    int                     `json:"user_id"`
	Created_at time.Time               `json:"created_at"`
	Updated_at time.Time               `json:"updated_at"`
	User       ResponseUserForComment  `json:"user"`
	Photo      ResponsePhotoForComment `json:"photo"`
}

type RequestUpdateComment struct {
	Message string `json:"message" binding:"required"`
}
