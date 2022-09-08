package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/ishangoyal13/dockerEx/models"
	"github.com/ishangoyal13/dockerEx/routers"
)

func main() {
	r := routers.RegisterRoutes()
	models.ConnectDatabase()

	r.Use(static.Serve("/", static.LocalFile("./build", true)))

	fmt.Println("Successfully connected")
	port := os.Getenv("PORT")
	//defaultPort := "8080"
	r.Run(":" + port)
}
