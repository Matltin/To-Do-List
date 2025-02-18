package db

import "to_do_list/models"

type DataBase interface {
	Connect(string) error
	Init() error
	CreateUser(arg CreateUserParams) (*models.User, error)
	GetUser(string) (*models.User, error)
	CreateTodo(arg CreateTodoParams) (*models.Todo, error)
	UpdateTodo(arg UpdateTodoParams) (*models.Todo, error)
	DeleteTodo(arg DeleteTodeParams) error
	GetTodoByID(id uint) (*models.Todo, error)
	GetTodosByID(arg GetTodosByIDParams) ([]models.Todo, error) 
}
