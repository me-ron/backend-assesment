package domain

type Loan struct{
	UserID string `json:"userid" bson:"userid"`
	Amount float64 `json:"amount" bson:"amount"`
	Status string `json:"status" bson:"status"`
}

type Status struct {
	Status string `json:"status" bson:"status"`
}