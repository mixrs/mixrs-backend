package dbaccess

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Operations interface {
	UserRepository
	PostRepository
	ChannelRepository
}

type Env struct {
	DB *gorm.DB
}

func InitDB() (*Env, error) {
	dsn := "host=localhost user=postgres password=password dbname=mixrdb port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Env{
		DB: conn,
	}, nil
}
