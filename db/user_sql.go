package db

import (
	"to_do_list/models"

	"gorm.io/gorm"
)

type CreateUserParams struct {
	Username string
	Password string
	Email    string
}

func (p *Postgres) CreateUser(arg CreateUserParams) (*models.User, error) {
	user := models.User{
		Username: arg.Username,
		Password: arg.Password,
		Email:    arg.Email,
	}

	result := p.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (p *Postgres) GetUser(email string) (*models.User, error) {
	var user models.User

	result := p.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (p *Postgres) CheckUserExists(email string) (bool, error) {
	var existedUser models.User

	err := p.DB.Where("email = ?", email).First(&existedUser).Error

	if err == nil {
		return true, nil // User exists
	} else if err == gorm.ErrRecordNotFound {
		return false, nil // User does not exist
	}
	return false, err // Some DB error occurred
}
