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
	CheckUserExists(email string) (bool, error)
}

type CreateUserParams struct {
	Username string
	Password string
	Email    string
}

type CreateTodoParams struct {
	UserID      uint
	Title       string
	Description string
}

type UpdateTodoParams struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
}

type DeleteTodeParams struct {
	ID     uint
	UserID uint
}

type GetTodosByIDParams struct {
	UserID uint
	Page   int
	Limit  int
}