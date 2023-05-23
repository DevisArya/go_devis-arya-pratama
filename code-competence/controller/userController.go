package controller

import (
	md "code-competence/middleware"
	m "code-competence/models"
	"code-competence/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {

	user := m.User{}
	c.Bind(&user)

	if err := repository.GetUserEmail(user.Email); err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "Email Account Already Exists",
		})
	}

	valid := md.PostUserValidation(user)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	result := m.UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "succes create new user",
		"User":    result,
	})
}

func LoginUser(c echo.Context) error {
	user := m.User{}
	c.Bind(&user)

	if err := repository.LoginUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	token, err := md.CreateToken(int(user.Id), user.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	UserResponseLogin := m.UserResponseLogin{Id: user.Id, Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success login",
		"User":    UserResponseLogin,
	})
}
