package config

import (
	"database/sql"

	"github.com/andrei0230/newsflow-app/internal/storage"
	"github.com/andrei0230/newsflow-app/internal/users"
	"github.com/gin-gonic/gin"
)

func SetupApp() (*gin.Engine, *sql.DB, error) {
	app := gin.Default()
	db, err := storage.StartMySql()
	if err != nil {
		return app, db, err
	}
	userStorage := users.NewUserStorage(db)
	userController := users.NewUserController(userStorage)
	users.SetCORS(app)
	users.SetRoutes(app, userController)
	return app, db, nil
}

func StopDB(db *sql.DB) error {
	err := storage.StopSql(db)
	if err != nil {
		return err
	}
	return nil
}
