package models

type ChannelModel struct {
	ID          string `gorm:"primaryKey;column:id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Image       []byte `gorm:"column:image;type:bytea;"`
	Tags        []byte `gorm:"column:tags;type:bytea;"`
}

type CreateChannelDTO struct {
	Title       string
	Description string
	Image       []byte
	Tags        []string
}

type ChannelDTO struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Tags        []string `json:"tags"`
}

type FetchChannelModel struct {
	ID          string `gorm:"column:id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Image       []byte `gorm:"column:image;type:bytea;"`
	Tags        []byte `gorm:"column:tags;type:bytea;"`
}
