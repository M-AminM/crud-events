package main

import (
	"crud-events/db"
	"crud-events/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
