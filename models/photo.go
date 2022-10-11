package models

import "time"

type Photo struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title" validate:"required"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url" validate:"required"`
	User_id    int       `json:"user_id"`
	User       User      `json:"user" gorm:"foreignKey:User_id" select:"username,email"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type RequestPhoto struct {
	Title     string `json:"title" binding:"required"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url" binding:"required"`
}

type ResponseCreatePhoto struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `json:"created_at"`
}

type ResponseUpdatePhoto struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    int       `json:"user_id"`
	Updated_at time.Time `json:"updated_at"`
}

type ResponseGetAllPhoto struct {
	ID         int                  `json:"id"`
	Title      string               `json:"title"`
	Caption    string               `json:"caption"`
	Photo_url  string               `json:"photo_url"`
	User_id    int                  `json:"userId"`
	Created_at time.Time            `json:"created_at"`
	Updated_at time.Time            `json:"updated_at"`
	User       ResponseUserForPhoto `json:"user"`
}

type ResponsePhotoForComment struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_id   int    `json:"user_id"`
}
