package channel

import (
	"encoding/base64"
	"encoding/json"

	"github.com/Tak1za/mixr/pkg/dbaccess"
	"github.com/Tak1za/mixr/pkg/models"
	uuid "github.com/satori/go.uuid"
)

type Operations interface {
	CreateChannel(channel *models.CreateChannelDTO) (*models.ChannelDTO, error)
	GetChannel(id string) (*models.ChannelDTO, error)
	GetChannels() ([]*models.ChannelDTO, error)
}

type Service struct {
	Dbo dbaccess.Operations
}

func (s *Service) CreateChannel(channel *models.CreateChannelDTO) (*models.ChannelDTO, error) {
	serializedTags, _ := json.Marshal(channel.Tags)
	newChannel := &models.ChannelModel{
		ID:          uuid.NewV4().String(),
		Title:       channel.Title,
		Description: channel.Description,
		Image:       channel.Image,
		Tags:        serializedTags,
	}

	createdChannel, err := s.Dbo.CreateChannel(newChannel)
	if err != nil {
		return nil, err
	}

	var deserializedTags []string
	json.Unmarshal(createdChannel.Tags, &deserializedTags)

	return &models.ChannelDTO{
		ID:          createdChannel.ID,
		Title:       createdChannel.Title,
		Description: createdChannel.Description,
		Image:       base64.StdEncoding.EncodeToString(createdChannel.Image),
		Tags:        deserializedTags,
	}, nil
}

func (s *Service) GetChannel(id string) (*models.ChannelDTO, error) {
	fetchedChannel, err := s.Dbo.GetChannel(id)
	if err != nil {
		return nil, err
	}

	encodedImage := base64.StdEncoding.EncodeToString(fetchedChannel.Image)

	var deserializedTags []string
	json.Unmarshal(fetchedChannel.Tags, &deserializedTags)

	return &models.ChannelDTO{
		ID:          fetchedChannel.ID,
		Title:       fetchedChannel.Title,
		Description: fetchedChannel.Description,
		Image:       encodedImage,
		Tags:        deserializedTags,
	}, nil
}

func (s *Service) GetChannels() ([]*models.ChannelDTO, error) {
	channels, err := s.Dbo.GetChannels()
	if err != nil {
		return nil, err
	}

	if channels == nil {
		return []*models.ChannelDTO{}, nil
	}

	var res []*models.ChannelDTO
	for _, v := range channels {
		encodedImage := base64.StdEncoding.EncodeToString(v.Image)

		var deserializedTags []string
		json.Unmarshal(v.Tags, &deserializedTags)

		item := &models.ChannelDTO{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Image:       encodedImage,
			Tags:        deserializedTags,
		}

		res = append(res, item)
	}

	return res, nil
}
