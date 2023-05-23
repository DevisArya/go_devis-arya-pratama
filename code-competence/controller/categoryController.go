package controller

import (
	m "code-competence/models"
	"code-competence/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCategory(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	err, res := repository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":   "200",
		"Message":  "success get category",
		"Category": res,
	})

}

func GetCategorys(c echo.Context) error {

	err, res := repository.GetCategorys()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":    "200",
		"Message":   "success get all categorys",
		"Categorys": res,
	})
}

func CreateCategory(c echo.Context) error {

	category := m.Category{}
	c.Bind(&category)

	if err := repository.CreateCategory(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	result := m.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":   "200",
		"Message":  "succes create new category",
		"Category": result,
	})
}

func DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.DeleteCategory(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete category",
	})
}

func UpdateCategory(c echo.Context) error {
	updateData := m.Category{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	if err := repository.UpdateCategory(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update category",
	})
}
