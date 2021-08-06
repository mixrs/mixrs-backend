package dbaccess

import (
	"errors"
	"fmt"
	"log"

	"github.com/Tak1za/mixr/pkg/models"
)

type PostRepository interface {
	CreatePost(post *models.PostModel) (*models.PostModel, error)
	GetPosts(channelId string) ([]*models.FetchPostModel, error)
	GetPostById(channelId, postId string) (*models.FetchPostModel, error)
}

func (e *Env) CreatePost(post *models.PostModel) (*models.PostModel, error) {
	if err := e.DB.Model(&models.PostModel{}).Create(post).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("error occured while creating post")
	}

	return post, nil
}

func (e *Env) GetPosts(channelId string) ([]*models.FetchPostModel, error) {
	var res []*models.FetchPostModel

	if err := e.DB.Model(&models.ChannelModel{}).Where("channel_models.id = ?", channelId).Select("post_models.*, user_models.id as user_id, user_models.name as user_name, user_models.email as user_email, user_models.image as user_image, channel_models.id as channel_id, channel_models.title as channel_title, channel_models.description as channel_description").Joins("inner join post_models on channel_models.id = post_models.channel_id").Joins("inner join user_models on post_models.user_id = user_models.id").Order("post_models.created_at desc").Scan(&res).Error; err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error occured while getting all posts on channel with id: %s", channelId)
	}
	return res, nil
}

func (e *Env) GetPostById(channelId, postId string) (*models.FetchPostModel, error) {
	var res *models.FetchPostModel
	if err := e.DB.Model(&models.ChannelModel{}).Where("channel_models.id = ?", channelId).Select("post_models.*, user_models.id as user_id, user_models.name as user_name, user_models.email as user_email, user_models.image as user_image, channel_models.id as channel_id, channel_models.title as channel_title, channel_models.description as channel_description").Joins("inner join post_models on channel_models.id = post_models.channel_id").Joins("inner join user_models on post_models.user_id = user_models.id").Where("post_models.id = ?", postId).Find(&res).Error; err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error occured while getting post with id: %s on channel with id: %s", postId, channelId)
	}

	return res, nil
}
