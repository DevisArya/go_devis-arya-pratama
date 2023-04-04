package controller

import (
	"net/http"
	"praktikum/config"
	md "praktikum/middleware"
	m "praktikum/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []m.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var user m.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}
	if result := config.DB.Where("id = ?", id).First(&user); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"user":     user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := m.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var users m.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if found := config.DB.Where("id = ?", id).First(&users); found.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user not found",
		})
	}

	if err := config.DB.Delete(&m.User{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	var users m.User
	updateUser := new(m.User)

	if err := c.Bind(updateUser); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	updateUser.ID = uint(id)

	if found := config.DB.Where("id = ?", id).First(&users); found.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "user not found",
		})
	}

	if err := config.DB.Model(&users).Updates(updateUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
	})
}

func LoginUserController(c echo.Context) error {
	user := m.User{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := md.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	userResponse := m.UserResponse{ID: int(user.ID), Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    userResponse,
	})
}
