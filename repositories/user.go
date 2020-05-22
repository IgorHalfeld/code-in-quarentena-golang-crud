package repositories

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct{}

var number = 1
var users = map[int]*User{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(c echo.Context) error {
	user := &User{
		ID: number,
	}

	if err := c.Bind(user); err != nil {
		return err
	}

	users[user.ID] = user
	number++
	return c.JSON(http.StatusCreated, user)
}

func (ur *UserRepository) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user := users[id]
	return c.JSON(http.StatusOK, user)
}

func (ur *UserRepository) UpdateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	users[id].Name = user.Name

	return c.JSON(http.StatusOK, users[id])
}

func (ur *UserRepository) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	delete(users, id)

	return c.NoContent(http.StatusNoContent)
}
