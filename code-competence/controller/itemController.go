package controller

import (
	md "code-competence/middleware"
	m "code-competence/models"
	"code-competence/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetItem(c echo.Context) error {

	id := c.Param("id")

	err, res := repository.GetItem(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get Item",
		"Item":    res,
	})

}
func GetItemsByIdCategory(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("category_id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}
	err, res := repository.GetItemsByIdCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get Item",
		"Item":    res,
	})

}
func GetItemByName(c echo.Context) error {

	name := c.QueryParam("keyword")

	err, res := repository.GetItemsByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get Item",
		"Item":    res,
	})

}

func GetItems(c echo.Context) error {

	err, res := repository.GetItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get all Items",
		"Items":   res,
	})
}

func CreateItem(c echo.Context) error {

	item := m.Items{}
	c.Bind(&item)

	valid := md.PostItemValidation(item)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.CreateItem(&item); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "succes create new Item",
		"Item":    item,
	})
}

func DeleteItem(c echo.Context) error {
	id := c.Param("id")

	if err := repository.DeleteItem(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete Item",
	})
}

func UpdateItem(c echo.Context) error {
	updateData := m.Items{}
	c.Bind(&updateData)

	id := c.Param("id")

	if err := repository.UpdateItem(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update Item",
	})
}
