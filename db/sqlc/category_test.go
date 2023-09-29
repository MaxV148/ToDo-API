package db

import (
	"CheckToDoAPI/utils"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomCategory(t *testing.T, user User) Category {
	arg := CreateCategoryParams{
		Name: utils.RandomString(10),
		User: user.ID,
	}
	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.User, category.User)
	return category
}

func TestCreateCategory(t *testing.T) {
	user := createRandomUser(t)
	createRandomCategory(t, user)

}

func TestListCategoriesForUser(t *testing.T) {
	user := createRandomUser(t)
	cat1 := createRandomCategory(t, user)
	cat2 := createRandomCategory(t, user)
	cat3 := createRandomCategory(t, user)
	categories, err := testQueries.ListCategoriesForUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.True(t, len(categories) == 3)
	require.Equal(t, cat1, categories[0])
	require.Equal(t, cat2, categories[1])
	require.Equal(t, cat3, categories[2])
}

func TestDeleteCategory(t *testing.T) {
	user := createRandomUser(t)
	cat1 := createRandomCategory(t, user)
	require.NotEmpty(t, cat1)
	err := testQueries.DeleteCategory(context.Background(), cat1.ID)
	require.NoError(t, err)
	getCategories, err := testQueries.ListCategoriesForUser(context.Background(), user.ID)
	require.Empty(t, getCategories)

}

func TestUpdateCategory(t *testing.T) {
	user := createRandomUser(t)
	cat1 := createRandomCategory(t, user)
	require.NotEmpty(t, cat1)
	arg := UpdateCategoryParams{
		ID:   cat1.ID,
		Name: utils.RandomString(5),
	}
	updatedCategory, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Name, updatedCategory.Name)

}
