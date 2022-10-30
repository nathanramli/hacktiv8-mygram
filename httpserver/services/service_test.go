package services

import (
	"context"
	"errors"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userService     UserSvc
	userRepo        repositories.UserRepo
	createUser      func(ctx context.Context, user *models.User) error
	updateUser      func(ctx context.Context, user *models.User) error
	findUserByID    func(ctx context.Context, id int) (*models.User, error)
	findUserByEmail func(ctx context.Context, email string) (*models.User, error)
	deleteUser      func(ctx context.Context, id int) error
)

type userRepositoryMock struct{}

func (r *userRepositoryMock) CreateUser(ctx context.Context, user *models.User) error {
	return createUser(ctx, user)
}

func (r *userRepositoryMock) UpdateUser(ctx context.Context, user *models.User) error {
	return updateUser(ctx, user)
}

func (r *userRepositoryMock) FindUserByID(ctx context.Context, id int) (*models.User, error) {
	return findUserByID(ctx, id)
}

func (r *userRepositoryMock) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return findUserByEmail(ctx, email)
}

func (r *userRepositoryMock) DeleteUser(ctx context.Context, id int) error {
	return deleteUser(ctx, id)
}

func TestUserService_CreateUser_Success(t *testing.T) {
	userRepo = &userRepositoryMock{}
	userService = NewUserSvc(userRepo)
	ctx := context.Background()

	paramsRegister := params.Register{
		Username: "myusername",
		Age:      20,
		Email:    "user@gmail.com",
		Password: "secret123",
	}

	expectedVal := views.Register{
		Id:       1,
		Age:      20,
		Email:    "user@gmail.com",
		Username: "myusername",
	}

	createUser = func(ctx context.Context, user *models.User) error {
		user.Id = 1
		return nil
	}

	resp := userService.Register(ctx, &paramsRegister)
	data := resp.Payload.(views.Register)
	assert.Equal(t, expectedVal.Id, data.Id)
	assert.Equal(t, expectedVal.Email, data.Email)
	assert.Equal(t, expectedVal.Username, data.Username)
	assert.Equal(t, expectedVal.Age, data.Age)
}

func TestUserService_CreateUser_ServerError(t *testing.T) {
	userRepo = &userRepositoryMock{}
	userService = NewUserSvc(userRepo)
	ctx := context.Background()

	paramsRegister := params.Register{
		Username: "myusername",
		Age:      20,
		Email:    "user@gmail.com",
		Password: "secret123",
	}

	serverError := errors.New("internal server error")
	createUser = func(ctx context.Context, user *models.User) error {
		return serverError
	}

	resp := userService.Register(ctx, &paramsRegister)
	assert.Equal(t, http.StatusInternalServerError, resp.Status)
	assert.Equal(t, serverError.Error(), resp.Error)
	assert.Equal(t, views.M_INTERNAL_SERVER_ERROR, resp.Message)
}
