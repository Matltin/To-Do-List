package db

import (
	"testing"
	"to_do_list/models"
	"to_do_list/util"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *models.User {
	arg := CreateUserParams{
		Username: util.RandomUsername(),
		Password: util.RandomPassword(),
		Email:    util.RandomEmail(),
	}

	user, err := p.CreateUser(arg)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.NotNil(t, user.Username)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createUser := createRandomUser(t)

	user, err := p.GetUser(createUser.Username)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.NotZero(t, user.ID)
	
}


