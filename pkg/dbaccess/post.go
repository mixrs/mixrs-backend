package dbaccess

import (
	"errors"
	"log"

	"github.com/Tak1za/mixr/pkg/models"
)

type PostRepository interface {
	CreatePost(post *models.PostModel) (string, error)
	GetPosts(channelId string) ([]*models.FetchPostModel, error)
}

func (e *Env) CreatePost(post *models.PostModel) (string, error) {
	if err := e.DB.Model(&models.PostModel{}).Create(post).Error; err != nil {
		log.Println(err.Error())
		return "", errors.New("error occured while creating post")
	}

	return post.ID, nil
}

func (e *Env) GetPosts(channelId string) ([]*models.FetchPostModel, error) {
	var res []*models.FetchPostModel

	if err := e.DB.Model(&models.ChannelModel{}).Where("channel_models.id = ?", channelId).Select("post_models.*, user_models.id as user_id, user_models.name as user_name, user_models.email as user_email, channel_models.id as channel_id, channel_models.title as channel_title").Joins("inner join post_models on channel_models.id = post_models.channel_id").Joins("inner join user_models on post_models.user_id = user_models.id").Scan(&res).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("error occured while getting all posts")
	}
	return res, nil
}