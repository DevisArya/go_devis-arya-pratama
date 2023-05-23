package middleware

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var secretKey = func() string {
	godotenv.Load()
	key := os.Getenv("SECRET_KEY")

	return key
}()
var IsloggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(secretKey),
})

func CreateToken(userId int, name string) (string, error) {
	payload := jwt.MapClaims{}
	payload["userId"] = userId
	payload["name"] = name
	payload["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRET_KEY")

	return token.SignedString([]byte(secretKey))
}

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${time_unix_milli}, latency=${latency_human}\n",
	}))
}
