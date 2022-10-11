package models

import "time"

type SocialMedia struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name" validate:"required"`
	Social_media_url string    `json:"socialMediaUrl" validate:"required"`
	UserId           int       `json:"userId"`
	User             User      `json:"user" gorm:"foreignkey:UserId"`
	Created_at       time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Updated_at       time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type RequestSocialMedia struct {
	Name             string `json:"name" binding:"required"`
	Social_media_url string `json:"social_media_url" binding:"required"`
}

type ResponseSocialMedia struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	UserId           int       `json:"user_id"`
	Created_at       time.Time `json:"created_at"`
}

type ResponseUpdateSocialMedia struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	UserId           int       `json:"user_id"`
	Updated_at       time.Time `json:"updated_at"`
}

type ResponseGetAllSocialMedia struct {
	ID               int                        `json:"id"`
	Name             string                     `json:"name"`
	Social_media_url string                     `json:"social_media_url"`
	UserId           int                        `json:"userId"`
	Created_at       time.Time                  `json:"created_at"`
	Updated_at       time.Time                  `json:"updated_at"`
	User             ResponseUserForSocialMedia `json:"user"`
}
