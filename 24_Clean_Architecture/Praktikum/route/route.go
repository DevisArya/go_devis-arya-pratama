package route

import (
	"log"
	"os"
	"praktikum/config"

	userHandler "praktikum/handler"
	userRepo "praktikum/repository"
	userUsecase "praktikum/usecase"

	m "praktikum/middleware"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func StartRoute() *echo.Echo {
	e := echo.New()
	userRepo := userRepo.New(config.DB)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler := userHandler.New(userUsecase)

	e.Pre(mid.AddTrailingSlash())
	m.LogMiddleware(e)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := os.Getenv("SECRET_KEY")
	Jwt := mid.JWT([]byte(secretKey))

	gUser := e.Group("/users")
	gUser.GET("/", userHandler.GetUsersController, Jwt)
	gUser.POST("/login/", userHandler.LoginUserController)
	gUser.POST("/", userHandler.CreateUserController)

	return e
}
