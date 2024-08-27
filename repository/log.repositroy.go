package repository

import "loan_tracker/domain"

type LogRepo struct {
	collection domain.CollectionInterface
}