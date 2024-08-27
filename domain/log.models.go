package domain

type Log struct{
	RequestID string `json:"requestid" bson:"requestid"`
	Status int `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
	Data interface{} `json:"data" bson:"data"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}