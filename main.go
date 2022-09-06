package main

import (
	"database/sql"
	"fmt"

	"docker/models"
	"docker/routers"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	r := routers.RegisterRoutes()
	models.ConnectDatabase()

	fmt.Println("Successfully connected")
	r.Run("localhost:8080")
}
