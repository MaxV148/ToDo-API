package db

import (
	"CheckToDoAPI/utils"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: utils.RandomOwner(),
		Password: utils.RandomPassword(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.Password, arg.Password)
	require.NotEmpty(t, user.CreatedAt)
	return user

}

func TestCreateAuthor(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByName(t *testing.T) {
	dbUser := createRandomUser(t)
	getUser, err := testQueries.GetUserByName(context.Background(), dbUser.Username)
	require.NoError(t, err)
	require.Equal(t, dbUser.ID, getUser.ID)
	require.Equal(t, dbUser.Username, getUser.Username)
	require.Equal(t, dbUser.Password, getUser.Password)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
	getUser, err := testQueries.GetUserByName(context.Background(), user.Username)
	require.Empty(t, getUser)
	require.ErrorIs(t, err, sql.ErrNoRows)
}
