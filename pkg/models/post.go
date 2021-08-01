package models

import (
	"time"
)

type PostModel struct {
	ID        string       `gorm:"primaryKey;column:id"`
	Title     string       `gorm:"column:title"`
	Content   string       `gorm:"column:content"`
	UserID    string       `gorm:"column:user_id"`
	ChannelID string       `gorm:"column:channel_id"`
	Channel   ChannelModel `gorm:"association_foreignKey:channel_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      UserModel    `gorm:"association_foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreatePostDTO struct {
	UserID  string `json:"userId" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostDTO struct {
	ID        string      `json:"id"`
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	UpdatedAt time.Time   `json:"updatedAt"`
	CreatedAt time.Time   `json:"createdAt"`
	Channel   *ChannelDTO `json:"channel"`
	User      *UserDTO    `json:"user"`
}

type FetchPostModel struct {
	ID                 string    `gorm:"column:id"`
	Title              string    `gorm:"column:title"`
	Content            string    `gorm:"column:content"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
	ChannelID          string    `gorm:"column:channel_id"`
	ChannelTitle       string    `gorm:"column:channel_title"`
	ChannelDescription string    `gorm:"column:channel_description"`
	UserID             string    `gorm:"column:user_id"`
	UserName           string    `gorm:"column:user_name"`
	UserEmail          string    `gorm:"column:user_email"`
	UserImage          []byte    `gorm:"column:user_image"`
}
