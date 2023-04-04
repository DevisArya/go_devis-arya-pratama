package routes

import (
	"praktikum/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.Pre(middleware.AddTrailingSlash())
	gUser := e.Group("/users")
	gUser.GET("/", controller.GetUsersController)
	gUser.GET("/:id/", controller.GetUserController)
	gUser.POST("/", controller.CreateUserController)
	gUser.DELETE("/:id/", controller.DeleteUserController)
	gUser.PUT("/:id/", controller.UpdateUserController)

	gBook := e.Group("/books")
	gBook.GET("/", controller.GetBooksController)
	gBook.GET("/:id/", controller.GetBookController)
	gBook.POST("/", controller.CreateBookController)
	gBook.DELETE("/:id/", controller.DeleteBookController)
	gBook.PUT("/:id/", controller.UpdateBookController)

	gBlog := e.Group("/blogs")
	gBlog.GET("/", controller.GetBlogsController)
	gBlog.GET("/:id/", controller.GetBlogController)
	gBlog.POST("/", controller.CreateBlogController)
	gBlog.DELETE("/:id/", controller.DeleteBlogController)
	gBlog.PUT("/:id/", controller.UpdateBlogController)

	return e
}
