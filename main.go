package main

import (
	"go-jwt/database"
	"go-jwt/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(os.Getenv("PORT"))
}
