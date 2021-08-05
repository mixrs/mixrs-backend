package models

type ChannelModel struct {
	ID          string `gorm:"primaryKey;column:id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Image       []byte `gorm:"column:image;type:bytea;"`
}

type CreateChannelDTO struct {
	Title       string
	Description string
	Image       []byte
}

type ChannelDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type FetchChannelModel struct {
	ID          string `gorm:"column:id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Image       []byte `gorm:"column:image;type:bytea;"`
}
