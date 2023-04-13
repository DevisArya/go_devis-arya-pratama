package repository

import (
	"errors"
	eu "praktikum/entity/user"

	mock "github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	Mock mock.Mock
}

func NewMockUserRepo() *MockUserRepository {
	return &MockUserRepository{}
}

var data = []eu.User{
	{
		ID:       1,
		Name:     "devis",
		Email:    "Devis@gmail.com",
		Password: "123",
	},
	{
		ID:       2,
		Name:     "devis2",
		Email:    "Devis2@gmail.com",
		Password: "12345"},
}

func (m *MockUserRepository) CreateUserController(user *eu.User) error {
	argument := m.Mock.Called(user)
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("please complete all fields")
	}

	if argument.Get(0) == nil {
		return errors.New("error")
	} else {
		return nil
	}
}
func (m *MockUserRepository) GetUsersController() (error, interface{}) {
	if data == nil {
		return errors.New("empty data"), nil
	} else {
		return nil, data
	}
}
func (*MockUserRepository) LoginUserController(user *eu.User) (error, string) {

	if user.Email == "" && user.Password == "" {
		return errors.New("please fill in the form"), ""
	}

	var idx int
	found := 0
	for i, val := range data {
		if val.Email == user.Email {
			idx = i
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found"), ""
	}

	if data[idx].Password != user.Password {
		return errors.New("fail login"), ""
	}

	*user = data[idx]

	return nil, ""
}
