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
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, users)
}

func (u *UserController) getByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, err := u.storage.getUserByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, user)
}

func (u *UserController) addUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := u.storage.createUser(newUser.Name, newUser.Email)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func (u *UserController) removeUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = u.storage.deleteUser(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
