package db

import (
	"testing"
	"to_do_list/models"
	"to_do_list/util"

	"github.com/stretchr/testify/require"
)

func createRandomTodo(t *testing.T, id uint) *models.Todo {
	arg := CreateTodoParams{
		UserID: id,
		Title:       util.RandomTitle(),
		Description: util.RandomDescription(),
	}

	todo, err := p.CreateTodo(arg)
	require.NoError(t, err)
	require.NotNil(t, todo)
	require.NotZero(t, todo.ID)
	require.False(t, todo.IsDone)

	return todo
}

func TestCreateTodo(t *testing.T) {

	createRandomTodo(t, createRandomUser(t).ID)
}

func TestDeleteTodo(t *testing.T) {
	user := createRandomUser(t)
	todo := createRandomTodo(t, user.ID)

	arg := DeleteTodeParams{
		ID:     todo.ID,
		UserID: todo.UserID,
	}

	err := p.DeleteTodo(arg)
	require.NoError(t, err)

	deletedTodo, err := p.GetTodoByID(todo.ID)
	require.Error(t, err)
	require.Nil(t, deletedTodo)
}

func TestUpdateTodo(t *testing.T) {
	user := createRandomUser(t)
	todo := createRandomTodo(t, user.ID)

	arg := UpdateTodoParams{
		ID:          todo.ID,
		UserID:      todo.UserID,
		Title:       util.RandomTitle(),
		Description: util.RandomDescription(),
	}

	updatedTodo, err := p.UpdateTodo(arg)
	require.NoError(t, err)
	require.NotNil(t, updatedTodo)
	require.Equal(t, arg.Title, updatedTodo.Title)
	require.Equal(t, arg.Description, updatedTodo.Description)
	require.False(t, updatedTodo.IsDone)
}

func TestGetTodosByID(t *testing.T) {
	user := createRandomUser(t)

	todo1 := createRandomTodo(t, user.ID)
	todo2 := createRandomTodo(t, user.ID)
	todo3 := createRandomTodo(t, user.ID)
	todo4 := createRandomTodo(t, user.ID)

	arg := GetTodosByIDParams {
		UserID: user.ID,
		Page: 1,
		Limit: 2,
	}

	todos, err := p.GetTodosByID(arg)
	require.NoError(t, err)
	require.Len(t, todos, 2)
	require.Equal(t, todos[0].Title, todo1.Title)
	require.Equal(t, todos[1].Title, todo2.Title)

	arg = GetTodosByIDParams {
		UserID: user.ID,
		Page: 2,
		Limit: 2,
	}

	todos, err = p.GetTodosByID(arg)
	require.NoError(t, err)
	require.Len(t, todos, 2)
	require.Equal(t, todos[0].Title, todo3.Title)
	require.Equal(t, todos[1].Title, todo4.Title)

	arg = GetTodosByIDParams {
		UserID: user.ID,
		Page: 3,
		Limit: 2,
	}

	todos, err = p.GetTodosByID(arg)
	require.NoError(t, err)
	require.Len(t, todos, 0)

}