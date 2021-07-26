package dbaccess

import (
	"errors"
	"fmt"
	"log"

	"github.com/Tak1za/mixr/pkg/models"
	"gorm.io/gorm"
)

type ChannelRepository interface {
	CreateChannel(channel *models.ChannelModel) (string, error)
	GetChannel(id string) (*models.FetchChannelModel, error)
	GetChannels() ([]*models.FetchChannelModel, error)
}

func (e *Env) CreateChannel(channel *models.ChannelModel) (string, error) {
	if err := e.DB.Model(&models.ChannelModel{}).Create(channel).Error; err != nil {
		log.Println(err.Error())
		return "", errors.New("error occured while creating channel")
	}

	return channel.ID, nil
}

func (e *Env) GetChannel(id string) (*models.FetchChannelModel, error) {
	var dbChannel models.FetchChannelModel
	if err := e.DB.Model(&models.ChannelModel{}).Where("id = ?", id).First(&dbChannel).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			log.Println(err.Error())
			return nil, fmt.Errorf("no channel found with id: %s", id)
		default:
			log.Println(err.Error())
			return nil, fmt.Errorf("error occured while getting channel with id: %s", id)
		}
	}

	return &dbChannel, nil
}

func (e *Env) GetChannels() ([]*models.FetchChannelModel, error) {
	var res []*models.FetchChannelModel
	if err := e.DB.Model(&models.ChannelModel{}).Scan(&res).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("error occured while getting all channels")
	}

	return res, nil
}
