package models

type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}

type CreateUserDTO struct {
	Name  string
	Email string
	Image []byte
}

type UpdateUserDTO struct {
	Name string `json:"name" binding:"required"`
}

type UserModel struct {
	ID    string `gorm:"primary_key;column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email;unique;not null;"`
	Image []byte `gorm:"column:image;type:bytea;"`
}

type UpdateUserModel struct {
	Name string `gorm:"column:name"`
}

func (u *UpdateUserDTO) MapToUserModel() *UserModel {
	return &UserModel{
		Name: u.Name,
	}
}
