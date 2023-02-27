package main

import (
	"log"

	"github.com/andrei0230/newsflow-app/internal/storage"
	"github.com/andrei0230/newsflow-app/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	db, err := storage.StartMySql()
	if err != nil {
		log.Panic(err)
	}
	defer storage.StopSql(db)
	userStorage := users.NewUserStorage(db)
	userController := users.NewUserController(userStorage)
	users.SetCORS(app)
	users.SetRoutes(app, userController)
	app.Run(":8080")
}
