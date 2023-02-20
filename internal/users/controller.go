package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	storage *UserStorage
}

func NewUserController(storage *UserStorage) *UserController {
	return &UserController{storage: storage}
}

func (u *UserController) getAll(c *gin.Context) {
	users, err := u.storage.geAllUsers()
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (u *UserController) getByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	user, err := u.storage.getUserByID(id)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, user)
}
