package routes

import (
	c "code-competence/controller"
	m "code-competence/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	gUser := e.Group("/user")
	gUser.POST("/", c.CreateUser)
	gUser.POST("/login/", c.LoginUser)

	gItems := e.Group("/items")
	gItems.GET("/", c.GetItems, m.IsloggedIn)
	gItems.GET("/:id/", c.GetItem, m.IsloggedIn)
	gItems.POST("/", c.CreateItem, m.IsloggedIn)
	gItems.PUT("/:id/", c.UpdateItem, m.IsloggedIn)
	gItems.DELETE("/:id/", c.DeleteItem, m.IsloggedIn)
	gItems.GET("/category/:category_id/", c.GetItemsByIdCategory, m.IsloggedIn)
	gItems.GET("", c.GetItemByName, m.IsloggedIn)

	gCategory := e.Group("/category")
	gCategory.GET("/", c.GetCategorys, m.IsloggedIn)
	gCategory.GET("/:id/", c.GetCategory, m.IsloggedIn)
	gCategory.POST("/", c.CreateCategory, m.IsloggedIn)
	gCategory.PUT("/:id/", c.UpdateCategory, m.IsloggedIn)
	gCategory.DELETE("/:id/", c.DeleteCategory, m.IsloggedIn)
	return e
}
