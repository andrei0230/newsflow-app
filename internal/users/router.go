package users

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, controller *UserController) {
	router.GET("/users", controller.getAll)
	router.GET("/users/id/:id", controller.getByID)
	router.POST("/users/add", controller.addUser)
	router.DELETE("/users/delete/:id", controller.removeUser)
}

func SetCORS(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
}
