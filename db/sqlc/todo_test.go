package db

import (
	"CheckToDoAPI/utils"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomToDo(t *testing.T, user User) Todo {
	category := createRandomCategory(t, user)
	arg := CreateToDoParams{
		Title:     utils.RandomString(15),
		Content:   utils.RandomString(20),
		CreatedBy: user.ID,
		Category:  category.ID,
	}
	todo, err := testQueries.CreateToDo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Title, todo.Title)
	require.Equal(t, arg.Content, todo.Content)
	require.Equal(t, arg.Category, todo.Category)
	require.Equal(t, arg.CreatedBy, todo.CreatedBy)
	require.NotEmpty(t, todo.ID)
	require.NotEmpty(t, todo.CreatedAt)
	return todo

}

func TestCreateTodo(t *testing.T) {
	user := createRandomUser(t)
	createRandomToDo(t, user)
}

func TestListToDoForUser(t *testing.T) {
	user := createRandomUser(t)
	todo1 := createRandomToDo(t, user)
	todo2 := createRandomToDo(t, user)
	todo3 := createRandomToDo(t, user)

	todos, err := testQueries.ListToDoForUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.True(t, len(todos) == 3)
	require.Equal(t, todos[0].CreatedBy, todo1.CreatedBy)
	require.Equal(t, todos[1].CreatedBy, todo2.CreatedBy)
	require.Equal(t, todos[2].CreatedBy, todo3.CreatedBy)
}

func TestDeleteToDo(t *testing.T) {
	user := createRandomUser(t)
	todo1 := createRandomToDo(t, user)
	require.NotEmpty(t, todo1.ID)
	err := testQueries.DeleteToDo(context.Background(), todo1.ID)
	require.NoError(t, err)
	getToDo, err := testQueries.ListToDoForUser(context.Background(), user.ID)
	require.Empty(t, getToDo)
}

func TestUpdateToDo(t *testing.T) {
	user := createRandomUser(t)
	todo1 := createRandomToDo(t, user)
	require.NotEmpty(t, todo1.ID)
	arg := UpdateToDoParams{
		ID:      todo1.ID,
		Title:   utils.RandomString(5),
		Content: utils.RandomString(10),
	}
	updatedTodo, err := testQueries.UpdateToDo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Title, updatedTodo.Title)
	require.Equal(t, arg.Content, updatedTodo.Content)

}
