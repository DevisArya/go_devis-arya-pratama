package usecase

import (
	eu "praktikum/entity/user"
	ur "praktikum/repository"
)

type UserUsecase interface {
	CreateUserController(user *eu.User) error
	GetUsersController() (err error, res interface{})
	LoginUserController(user *eu.User) (error, string)
}

type userUsecase struct {
	userRepository ur.UserRepo
}

func NewUserUsecase(userRepo ur.UserRepo) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (uc userUsecase) CreateUserController(user *eu.User) error {
	return uc.userRepository.CreateUserController(user)
}

func (uc userUsecase) GetUsersController() (err error, res interface{}) {
	return uc.userRepository.GetUsersController()
}
func (uc userUsecase) LoginUserController(user *eu.User) (error, string) {
	return uc.userRepository.LoginUserController(user)
}
