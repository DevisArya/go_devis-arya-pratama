package controller

import (
	"net/http"
	"praktikum/config"
	m "praktikum/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []m.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	var book m.Book
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}
	if result := config.DB.Where("id = ?", id).First(&book); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get book",
		"book":     book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := m.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	var books m.Book
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if found := config.DB.Where("id = ?", id).First(&books); found.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "book not found",
		})
	}

	if err := config.DB.Delete(&m.Book{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	var books m.Book
	updateBook := new(m.Book)

	if err := c.Bind(updateBook); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	updateBook.ID = id

	if found := config.DB.Where("id = ?", id).First(&books); found.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "book not found",
		})
	}

	if err := config.DB.Model(&books).Updates(updateBook).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update book",
	})
}
