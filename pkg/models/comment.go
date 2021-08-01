package models

import (
	"time"

	"gorm.io/gorm"
)

type CommentModel struct {
	gorm.Model
	ID     string    `gorm:"primaryKey;column:id"`
	Value  string    `gorm:"column:value"`
	UserID string    `gorm:"column:user_id"`
	PostID string    `gorm:"column:post_id"`
	User   UserModel `gorm:"association_foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Post   PostModel `gorm:"association_foreignKey:post_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateCommentDTO struct {
	UserID  string `json:"userId" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}

type CommentDTO struct {
	ID        string    `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	User      *UserDTO  `json:"user"`
}

type FetchCommentModel struct {
	ID        string    `gorm:"column:id"`
	Comment   string    `gorm:"column:value"`
	CreatedAt time.Time `gorm:"column:created_at"`
	DeletedAt time.Time `gorm:"column:updated_at"`
	UserID    string    `gorm:"column:user_id"`
	UserName  string    `gorm:"column:user_name"`
	UserEmail string    `gorm:"column:user_email"`
}
