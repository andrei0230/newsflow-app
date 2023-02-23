package main

import (
	"github.com/andrei0230/newsflow-app/internal/storage"
	"github.com/andrei0230/newsflow-app/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	db, err := storage.StartMySql()
	if err != nil {
		panic(err)
	}
	defer storage.StopSql(db)
	userStorage := users.NewUserStorage(db)
	userController := users.NewUserController(userStorage)
	users.SetRoutes(app, userController)
	app.Run()
}
