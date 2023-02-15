package users

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, controller *UserController) {
	router.GET("/users", controller.getAll)
}
