package main

import (
	"go-backend/config"
	"go-backend/api/router"
	"go-backend/db"
)

func main() {
	config.Load()
	db.InitDatabase()
	r := router.SetupRouter()

	// Listen and Server in localhost:3000
	r.Run(":3000")
}
