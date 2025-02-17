package db

import "to_do_list/models"

type DataBase interface {
	Connect(string) error
	Init() error
	CreateUser(arg CreateUserParams) (*models.User, error)
	GetUser(string) (*models.User, error)
}
