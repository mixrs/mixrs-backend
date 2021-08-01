package post

import (
	"encoding/base64"

	"github.com/Tak1za/mixr/pkg/dbaccess"
	"github.com/Tak1za/mixr/pkg/models"
	uuid "github.com/satori/go.uuid"
)

type Operations interface {
	CreatePost(post *models.CreatePostDTO, channelId string) (string, error)
	GetPosts(channelId string) ([]*models.PostDTO, error)
	GetPostById(channelId, postId string) (*models.PostDTO, error)
}

type Service struct {
	Dbo dbaccess.Operations
}

func (s *Service) CreatePost(post *models.CreatePostDTO, channelId string) (string, error) {
	newPost := &models.PostModel{
		ID:        uuid.NewV4().String(),
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		ChannelID: channelId,
	}

	postId, err := s.Dbo.CreatePost(newPost)
	if err != nil {
		return "", err
	}

	return postId, nil
}

func (s *Service) GetPosts(channelId string) ([]*models.PostDTO, error) {
	posts, err := s.Dbo.GetPosts(channelId)
	if err != nil {
		return nil, err
	}

	if posts == nil {
		return []*models.PostDTO{}, nil
	}

	var res []*models.PostDTO
	for _, v := range posts {
		item := &models.PostDTO{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			UpdatedAt: v.UpdatedAt,
			CreatedAt: v.CreatedAt,
			Channel: &models.ChannelDTO{
				ID:    v.ChannelID,
				Title: v.ChannelTitle,
			},
			User: &models.UserDTO{
				ID:    v.UserID,
				Name:  v.UserName,
				Email: v.UserEmail,
				Image: base64.StdEncoding.EncodeToString(v.UserImage),
			},
		}

		res = append(res, item)
	}

	return res, nil
}

func (s *Service) GetPostById(channelId, postId string) (*models.PostDTO, error) {
	post, err := s.Dbo.GetPostById(channelId, postId)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return &models.PostDTO{}, nil
	}

	item := &models.PostDTO{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UpdatedAt: post.UpdatedAt,
		CreatedAt: post.CreatedAt,
		Channel: &models.ChannelDTO{
			ID:          post.ChannelID,
			Title:       post.ChannelTitle,
			Description: post.ChannelDescription,
		},
		User: &models.UserDTO{
			ID:    post.UserID,
			Name:  post.UserName,
			Email: post.UserEmail,
			Image: base64.StdEncoding.EncodeToString(post.UserImage),
		},
	}

	return item, nil
}
