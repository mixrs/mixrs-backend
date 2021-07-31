package models

import (
	"time"
)

type ChannelModel struct {
	ID          string `gorm:"primaryKey;column:id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
}

type CreateChannelDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type ChannelDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FetchChannelModel struct {
	ID          string    `gorm:"column:id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
