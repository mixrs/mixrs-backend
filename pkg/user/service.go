package user

import (
	"encoding/base64"

	"github.com/Tak1za/mixr/pkg/dbaccess"
	"github.com/Tak1za/mixr/pkg/models"
	uuid "github.com/satori/go.uuid"
)

type Operations interface {
	GetUser(id string) (*models.UserDTO, error)
	CreateUser(user *models.CreateUserDTO) (string, error)
	UpdateUser(user *models.UpdateUserDTO, id string) error
	DeleteUser(id string) error
}

type Service struct {
	Dbo dbaccess.Operations
}

func (s *Service) GetUser(id string) (*models.UserDTO, error) {
	fetchedUser, err := s.Dbo.GetUser(id)
	if err != nil {
		return nil, err
	}

	encodedImage := base64.StdEncoding.EncodeToString(fetchedUser.Image)

	return &models.UserDTO{
		Name:  fetchedUser.Name,
		ID:    fetchedUser.ID,
		Email: fetchedUser.Email,
		Image: encodedImage,
	}, nil
}

func (s *Service) CreateUser(toCreateUser *models.CreateUserDTO) (string, error) {
	newUser := &models.UserModel{
		ID:    uuid.NewV4().String(),
		Name:  toCreateUser.Name,
		Email: toCreateUser.Email,
		Image: toCreateUser.Image,
	}

	userId, err := s.Dbo.CreateUser(newUser)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (s *Service) UpdateUser(toUpdateUser *models.UpdateUserDTO, id string) error {
	if err := s.Dbo.UpdateUser(toUpdateUser.MapToUserModel(), id); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUser(id string) error {
	if err := s.Dbo.DeleteUser(id); err != nil {
		return err
	}

	return nil
}
