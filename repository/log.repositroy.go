package repository

import (
	"context"
	"loan_tracker/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type LogRepo struct {
	collection domain.CollectionInterface
}

func NewLogRepo(coll domain.CollectionInterface) *LogRepo{
	return &LogRepo{collection: coll}
}

func (lr *LogRepo) GetAllLogs() ([]domain.Log, error) {
	var logs []domain.Log
	cursor, err := lr.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return []domain.Log{}, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var log domain.Log
		cursor.Decode(&log)
		logs = append(logs, log)
	}
	return logs, nil
}