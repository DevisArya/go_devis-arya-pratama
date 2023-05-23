package repository

import (
	"code-competence/config"
	"code-competence/models"
	"code-competence/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserEmail(email string) (err error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "user not found",
		})
	}
	return nil
}
func CreateUser(user *models.User) error {
	password := user.Password

	hash, err := utils.HashPassword(password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user.Password = hash
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func LoginUser(user *models.User) error {
	password := user.Password

	if err := config.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "user not found",
			"error":   err.Error(),
		})
	}
	if match := utils.CheckPasswordHash(password, user.Password); match == false {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "mismatch password",
		})
	}
	return nil
}
