package dbaccess

import (
	"errors"
	"fmt"
	"log"

	"github.com/Tak1za/mixr/pkg/models"
)

type CommentRepository interface {
	CreateComment(comment *models.CommentModel) (string, error)
	GetComments(postId string) ([]*models.FetchCommentModel, error)
}

func (e *Env) CreateComment(comment *models.CommentModel) (string, error) {
	if err := e.DB.Model(&models.CommentModel{}).Create(comment).Error; err != nil {
		log.Println(err.Error())
		return "", errors.New("error occured while creating comment")
	}

	return comment.ID, nil
}

func (e *Env) GetComments(postId string) ([]*models.FetchCommentModel, error) {
	var res []*models.FetchCommentModel

	if err := e.DB.Model(&models.CommentModel{}).Select("comment_models.id, comment_models.value, comment_models.created_at, comment_models.deleted_at, user_models.id as user_id, user_models.name as user_name, user_models.email as user_email").Joins("inner join user_models on comment_models.user_id = user_models.id").Order("created_at desc").Where("comment_models.post_id = ?", postId).Find(&res).Error; err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error occured while getting all comments on post with id: %s", postId)
	}

	return res, nil
}
