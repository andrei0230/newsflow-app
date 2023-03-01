package main

import (
	"github.com/andrei0230/newsflow-app/internal/config"
)

func main() {
	app, db, err := config.SetupApp()
	if err != nil {
		panic(err)
	}
	defer config.StopDB(db)
	app.Run(":8080")
}
