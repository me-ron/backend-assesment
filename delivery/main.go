package main

import (
	"loan_tracker/delivery/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetUp(router)
	router.Run("127.0.0.1:8080")
}