package handler

import (
	"net/http"
	eu "praktikum/entity/user"
	uc "praktikum/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUC uc.UserUsecase
}

func New(userUC uc.UserUsecase) *UserHandler {
	return &UserHandler{
		userUC,
	}
}

func (h *UserHandler) CreateUserController(c echo.Context) error {
	user := eu.User{}
	c.Bind(&user)

	err := h.userUC.CreateUserController(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func (h *UserHandler) GetUsersController(c echo.Context) error {

	err, res := h.userUC.GetUsersController()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   res,
	})
}
func (h *UserHandler) LoginUserController(c echo.Context) error {
	user := eu.User{}
	c.Bind(&user)

	err, token := h.userUC.LoginUserController(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	userResponse := eu.UserResponse{ID: int(user.ID), Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Login",
		"user":    userResponse,
	})
}
