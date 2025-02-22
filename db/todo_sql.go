package db

import (
	"fmt"
	"to_do_list/models"
)

type CreateTodoParams struct {
	UserID      uint
	Title       string
	Description string
}

func (p *Postgres) CreateTodo(arg CreateTodoParams) (*models.Todo, error) {
	todo := models.Todo{
		UserID:      arg.UserID,
		Title:       arg.Title,
		Description: arg.Description,
	}

	result := p.DB.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func (p *Postgres) GetTodoByID(id uint) (*models.Todo, error) {
	todo := models.Todo{}

	result := p.DB.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

type UpdateTodoParams struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
}

func (p *Postgres) UpdateTodo(arg UpdateTodoParams) (*models.Todo, error) {
	todo := models.Todo{}

	result := p.DB.Where("id = ? AND user_id = ?", arg.ID, arg.UserID).First(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	if todo.IsDone {
		return nil, fmt.Errorf("the activity allready done")
	}

	result = p.DB.Model(&todo).Updates(models.Todo{
		Title:       arg.Title,
		Description: arg.Description,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

type DeleteTodeParams struct {
	ID     uint
	UserID uint
}

func (p *Postgres) DeleteTodo(arg DeleteTodeParams) error {
	todo := models.Todo{}

	result := p.DB.Where("id = ? AND user_id = ?", arg.ID, arg.UserID).First(&todo)
	if result.Error != nil {
		return result.Error
	}

	result = p.DB.Delete(&todo)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

type GetTodosByIDParams struct {
	UserID uint
	Page   int
	Limit  int
}

func (p *Postgres) GetTodosByID(arg GetTodosByIDParams) ([]models.Todo, error) {
	var todos []models.Todo

	offset := (arg.Page - 1) * arg.Limit

	result := p.DB.Where("user_id = ?", arg.UserID).
		Order("id ASC").
		Limit(arg.Limit).
		Offset(offset).
		Find(&todos)
	
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil

}
