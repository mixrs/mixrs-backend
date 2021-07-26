package post

import (
	"github.com/Tak1za/mixr/pkg/dbaccess"
	"github.com/Tak1za/mixr/pkg/models"
	uuid "github.com/satori/go.uuid"
)

type Operations interface {
	CreatePost(post *models.CreatePostDTO) (string, error)
	GetPosts(channelId string) ([]*models.PostDTO, error)
}

type Service struct {
	Dbo dbaccess.Operations
}

func (s *Service) CreatePost(post *models.CreatePostDTO) (string, error) {
	newPost := &models.PostModel{
		ID:        uuid.NewV4().String(),
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		ChannelID: post.ChannelID,
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
			},
		}

		res = append(res, item)
	}

	return res, nil
}
