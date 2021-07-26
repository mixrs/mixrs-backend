package channel

import (
	"github.com/Tak1za/mixr/pkg/dbaccess"
	"github.com/Tak1za/mixr/pkg/models"
	uuid "github.com/satori/go.uuid"
)

type Operations interface {
	CreateChannel(channel *models.CreateChannelDTO) (string, error)
	GetChannel(id string) (*models.ChannelDTO, error)
	GetChannels() ([]*models.ChannelDTO, error)
}

type Service struct {
	Dbo dbaccess.Operations
}

func (s *Service) CreateChannel(channel *models.CreateChannelDTO) (string, error) {
	newChannel := &models.ChannelModel{
		ID:    uuid.NewV4().String(),
		Title: channel.Title,
	}

	channelId, err := s.Dbo.CreateChannel(newChannel)
	if err != nil {
		return "", err
	}

	return channelId, nil
}

func (s *Service) GetChannel(id string) (*models.ChannelDTO, error) {
	fetchedChannel, err := s.Dbo.GetChannel(id)
	if err != nil {
		return nil, err
	}

	return &models.ChannelDTO{
		ID:    fetchedChannel.ID,
		Title: fetchedChannel.Title,
	}, nil
}

func (s *Service) GetChannels() ([]*models.ChannelDTO, error) {
	channels, err := s.Dbo.GetChannels()
	if err != nil {
		return nil, err
	}

	var res []*models.ChannelDTO
	for _, v := range channels {
		item := &models.ChannelDTO{
			ID:    v.ID,
			Title: v.Title,
		}

		res = append(res, item)
	}

	return res, nil
}
