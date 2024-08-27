package config

import (
	"loan_tracker/database"
	"loan_tracker/domain"
)

var LogCollection domain.CollectionInterface

func init() {
	var server ServerConnection
	server.Connect_could()
	LogCollection = &database.MongoCollection{
		Collection: server.Client.Database("LoanTracker").Collection("Log")}

}
