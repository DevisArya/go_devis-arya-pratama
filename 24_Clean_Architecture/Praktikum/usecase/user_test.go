package usecase

import (
	"errors"
	eu "praktikum/entity/user"
	"praktikum/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockUserRepository = &repository.MockUserRepository{Mock: mock.Mock{}}
var mockUserUsecase = userUsecase{userRepository: mockUserRepository}

func TestCreateUserController(t *testing.T) {
	data := eu.User{
		ID:       1,
		Name:     "devis",
		Email:    "Devis@gmail.com",
		Password: "123",
	}

	mockUserRepository.Mock.On("CreateUserController", &data).Return(data)

	err := mockUserUsecase.CreateUserController(&data)
	assert.Nil(t, err)

}
func TestCreateUserControllerInvalid(t *testing.T) {
	data := eu.User{}

	mockUserRepository.Mock.On("CreateUserController", &data).Return(errors.New("please complete all fields"))

	err := mockUserUsecase.CreateUserController(&data)
	assert.NotNil(t, err)
}

func TestGetUsersController(t *testing.T) {

	mockUserRepository.Mock.On("GetUsersController").Return()

	err, res := mockUserUsecase.GetUsersController()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
func TestLoginUserController(t *testing.T) {
	data := eu.User{
		Email:    "Devis@gmail.com",
		Password: "123",
	}

	mockUserRepository.Mock.On("LoginUserController", &data).Return(nil)

	err, token := mockUserUsecase.LoginUserController(&data)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}
func TestLoginUserControllerInvalid(t *testing.T) {
	data := eu.User{
		Email:    "Devis@gmail.com",
		Password: "",
	}

	mockUserRepository.Mock.On("LoginUserController", &data).Return(errors.New("user not found"))

	err, token := mockUserUsecase.LoginUserController(&data)
	assert.Equal(t, "", token)
	assert.NotNil(t, err)
}
