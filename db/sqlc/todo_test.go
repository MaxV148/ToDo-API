package db

import (
	"CheckToDoAPI/utils"
	"context"
	"database/sql"
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

	todos, err := testQueries.ListToDoForUser(context.Background(), ListToDoForUserParams{UserID: user.ID, SortingOrder: "TITLE_ASC"})
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
	todo, err := testQueries.DeleteToDo(context.Background(), DeleteToDoParams{ID: todo1.ID, CreatedBy: user.ID})
	require.NoError(t, err)
	require.NotEmpty(t, todo)
	getToDo, err := testQueries.ListToDoForUser(context.Background(), ListToDoForUserParams{UserID: user.ID, SortingOrder: "TITLE_ASC"})
	require.Empty(t, getToDo)
}

func TestDeleteToDoForOtherUser(t *testing.T) {
	user := createRandomUser(t)
	user1 := createRandomUser(t)
	todo1 := createRandomToDo(t, user)
	require.NotEmpty(t, todo1.ID)
	todo, err := testQueries.DeleteToDo(context.Background(), DeleteToDoParams{ID: todo1.ID, CreatedBy: user1.ID})
	require.Empty(t, todo)
	require.ErrorIs(t, err, sql.ErrNoRows)

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

func TestToggleDone(t *testing.T) {
	user := createRandomUser(t)
	todo1 := createRandomToDo(t, user)
	require.NotEmpty(t, todo1)
	// set to done to true
	doneTodo, err := testQueries.ToggleToDoDone(context.Background(), todo1.ID)
	require.NoError(t, err)
	require.Equal(t, !todo1.Done, doneTodo.Done)

	// set to done to false
	falseTodo, err := testQueries.ToggleToDoDone(context.Background(), todo1.ID)
	require.NoError(t, err)
	require.Equal(t, !doneTodo.Done, falseTodo.Done)

}

func TestDeleteCascadeOnPermissions(t *testing.T) {
	t.Skip()
	author := createRandomUser(t)
	user1 := createRandomUser(t)
	todo := createRandomToDo(t, author)
	// link todo to user
	perm, err := testQueries.GrantUserToToDo(context.Background(), GrantUserToToDoParams{
		UserID: user1.ID,
		TodoID: todo.ID,
	})
	require.NoError(t, err)
	require.Equal(t, user1.ID, perm.UserID)
	require.Equal(t, todo.ID, perm.TodoID)
	// delete todo
	_, err = testQueries.DeleteToDo(context.Background(), DeleteToDoParams{ID: todo.ID, CreatedBy: author.ID})
	require.NoError(t, err)
	// check if entry in permissions table is also deleted
	perm, err = testQueries.GetPermByUser(context.Background(), user1.ID)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, perm)

}
