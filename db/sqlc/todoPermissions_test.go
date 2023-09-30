package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGrantUserToToDo(t *testing.T) {

	user1 := createRandomUser(t) //author
	todo := createRandomToDo(t, user1)
	user2 := createRandomUser(t) // invite 1
	arg := GrantUserToToDoParams{
		UserID: user2.ID,
		TodoID: todo.ID,
	}
	perm1, err := testQueries.GrantUserToToDo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user2.ID, perm1.UserID)
	require.Equal(t, todo.ID, perm1.TodoID)

}

func TestListGrantedToDoForUser(t *testing.T) {
	user1 := createRandomUser(t) //author
	todo1 := createRandomToDo(t, user1)
	todo2 := createRandomToDo(t, user1)
	user2 := createRandomUser(t) // invited user
	// 1. invite
	arg := GrantUserToToDoParams{
		UserID: user2.ID,
		TodoID: todo1.ID,
	}
	perm1, err := testQueries.GrantUserToToDo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user2.ID, perm1.UserID)
	require.Equal(t, todo1.ID, perm1.TodoID)

	// 2. invite
	arg = GrantUserToToDoParams{
		UserID: user2.ID,
		TodoID: todo2.ID,
	}
	perm2, err := testQueries.GrantUserToToDo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user2.ID, perm2.UserID)
	require.Equal(t, todo2.ID, perm2.TodoID)

	todos, err := testQueries.ListGrantedToDoForUser(context.Background(), user2.ID)
	require.NoError(t, err)
	require.True(t, len(todos) == 2)

}
