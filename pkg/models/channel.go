package models

import (
	"time"

	"gorm.io/gorm"
)

type ChannelModel struct {
	gorm.Model
	ID    string `gorm:"primaryKey;column:id"`
	Title string `gorm:"column:title"`
}

type CreateChannelDTO struct {
	Title string `json:"title" binding:"required"`
}

type ChannelDTO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type FetchChannelModel struct {
	ID        string    `gorm:"column:id"`
	Title     string    `gorm:"column:title"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
