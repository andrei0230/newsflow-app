package users

import (
	"net/http"

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
