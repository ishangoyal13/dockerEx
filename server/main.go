package main

import (
	"fmt"

	"github.com/ishangoyal13/dockerEx/server/models"
	"github.com/ishangoyal13/dockerEx/server/routers"
)

func main() {
	r := routers.RegisterRoutes()
	models.ConnectDatabase()

	fmt.Println("Successfully connected")
	r.Run("localhost:8080")
}
