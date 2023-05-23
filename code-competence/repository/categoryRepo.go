package repository

import (
	"code-competence/config"
	"code-competence/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCategory(category *models.Category) error {
	if err := config.DB.Save(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func GetCategory(id int) (err error, res interface{}) {
	var category models.Category
	if err := config.DB.Preload("Items").Where("id = ?", id).First(&category).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "category not found",
		}), nil
	}
	return nil, category
}
func GetCategoryName(name string) (err error) {
	var area models.Category
	if err := config.DB.Where("name = ?", name).First(&area).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "area not found",
		})
	}
	return nil
}

func GetCategorys() (err error, res interface{}) {
	var categorys []models.Category

	if err := config.DB.Preload("Items").Find(&categorys).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}

	return nil, categorys
}

func DeleteCategory(id int) error {
	result := config.DB.Delete(&models.Category{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return nil
}

func UpdateCategory(updateData *models.Category, id int) error {
	result := config.DB.Model(&models.Category{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return nil
}
