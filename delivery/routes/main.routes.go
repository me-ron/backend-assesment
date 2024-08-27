package routes

import (
	"loan_tracker/config"
	"loan_tracker/database"

	"github.com/gin-gonic/gin"
)

func SetUp(router *gin.Engine) {
	var clinect config.ServerConnection
	clinect.Connect_could()

	userCollection := &database.MongoCollection{
		Collection: clinect.Client.Database("BlogPost").Collection("Users"),
	}

	NewUserRouter(router, userCollection)

}