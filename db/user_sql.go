package db

import "to_do_list/models"

type CreateUserParams struct {
	Username string
	Password string
	Email    string
}

func (p *Postgres) CreateUser(arg CreateUserParams) (*models.User, error) {
	user := models.User {
		Username: arg.Username,
		Password: arg.Password,
		Email: arg.Email,
	}

	result := p.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (p *Postgres) GetUser(username string) (*models.User, error) {
	var user models.User

	result := p.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

