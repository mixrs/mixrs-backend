package comment

import (
	"encoding/base64"

	"github.com/Tak1za/mixr/pkg/dbaccess"
	"github.com/Tak1za/mixr/pkg/models"
	uuid "github.com/satori/go.uuid"
)

type Operations interface {
	CreateComment(comment *models.CreateCommentDTO, postId string) (string, error)
	GetComments(postId string) ([]*models.CommentDTO, error)
}

type Service struct {
	Dbo dbaccess.Operations
}

func (s *Service) CreateComment(comment *models.CreateCommentDTO, postId string) (string, error) {
	newComment := &models.CommentModel{
		ID:     uuid.NewV4().String(),
		Value:  comment.Comment,
		UserID: comment.UserID,
		PostID: postId,
	}

	commentId, err := s.Dbo.CreateComment(newComment)
	if err != nil {
		return "", err
	}

	return commentId, nil
}

func (s *Service) GetComments(postId string) ([]*models.CommentDTO, error) {
	comments, err := s.Dbo.GetComments(postId)
	if err != nil {
		return nil, err
	}

	if comments == nil {
		return []*models.CommentDTO{}, nil
	}

	var res []*models.CommentDTO
	for _, v := range comments {
		if !v.DeletedAt.IsZero() {
			continue
		}

		item := &models.CommentDTO{
			ID:        v.ID,
			Comment:   v.Comment,
			CreatedAt: v.CreatedAt,
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
