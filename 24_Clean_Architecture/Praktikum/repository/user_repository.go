package repository

import (
	"errors"
	"net/http"
	"praktikum/config"
	eu "praktikum/entity/user"
	md "praktikum/middleware"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type UserRepo interface {
	CreateUserController(user *eu.User) error
	GetUsersController() (err error, res interface{})
	LoginUserController(user *eu.User) (error, string)
}

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *userRepo {
	return &userRepo{
		db,
	}
}

func (u *userRepo) CreateUserController(user *eu.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("please complete all fields")
	}
	if err := u.db.Save(&user); err != nil {
		return err.Error
	}
	return nil
}

func (u *userRepo) GetUsersController() (err error, res interface{}) {
	var users []eu.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), nil
	}
	return nil, users
}
func (u *userRepo) LoginUserController(user *eu.User) (error, string) {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "fail login",
		}), ""
	}

	token, err := md.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		}), ""
	}
	return nil, token
}
