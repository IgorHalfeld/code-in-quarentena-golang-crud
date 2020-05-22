package main

import (
	"net/http"

	"github.com/igorhalfeld/code-in-quarentena-golang-crud/repositories"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepository := repositories.NewUserRepository()

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.POST("/users", userRepository.CreateUser)
	e.GET("/users/:id", userRepository.GetUser)
	e.PUT("/users/:id", userRepository.UpdateUser)
	e.DELETE("/users/:id", userRepository.DeleteUser)

	e.Start(":1323")
}
