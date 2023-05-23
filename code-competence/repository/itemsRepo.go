package repository

import (
	"code-competence/config"
	"code-competence/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pborman/uuid"
)

func CreateItem(item *models.Items) error {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	item.Id = uuid
	if err := config.DB.Save(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func GetItem(id string) (err error, res interface{}) {
	var item models.Items
	if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "item not found",
		}), nil
	}
	return nil, item
}
func GetItemsByName(name string) (err error, res interface{}) {
	var item models.Items
	if err := config.DB.Where("name = ?", name).First(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "item not found",
		}), nil
	}
	return nil, item
}

func GetItems() (err error, res interface{}) {
	var items []models.Items

	if err := config.DB.Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}
	return nil, items
}
func GetItemsByIdCategory(id int) (err error, res interface{}) {
	var items []models.Items

	if err := config.DB.Where("category_id = ?", id).Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}
	return nil, items
}

func DeleteItem(id string) error {

	result := config.DB.Where("id = ?", id).Delete(&models.Items{})

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

func UpdateItem(updateData *models.Items, id string) error {
	result := config.DB.Model(&models.Items{}).Where("id = ?", id).Updates(&updateData)

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
