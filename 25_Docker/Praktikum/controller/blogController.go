package controller

import (
	"net/http"
	"praktikum/config"
	m "praktikum/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []m.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	var blog m.Blog
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}
	if result := config.DB.Where("id = ?", id).First(&blog); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "blog not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get blog",
		"blog":     blog,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := m.Blog{}
	c.Bind(&blog)

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	var blogs m.Blog
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if found := config.DB.Where("id = ?", id).First(&blogs); found.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "blog not found",
		})
	}

	if err := config.DB.Delete(&m.Blog{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete blog",
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	var blogs m.Blog
	updateBlog := new(m.Blog)

	if err := c.Bind(updateBlog); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	updateBlog.ID = id

	if found := config.DB.Where("id = ?", id).First(&blogs); found.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"messages": "blog not found",
		})
	}

	if err := config.DB.Model(&blogs).Updates(updateBlog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update blog",
	})
}
