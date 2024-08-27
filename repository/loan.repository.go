package repository

import (
	"context"
	"loan_tracker/domain"

	"go.mongodb.org/mongo-driver/bson"
)

type LoanRepo struct {
	Collections domain.CollectionInterface
}

func NewLoanRepo(coll domain.CollectionInterface) *LoanRepo{
	return &LoanRepo{
		Collections: coll,
	}
}

func (lr *LoanRepo) CreateLoan(loan domain.Loan) (domain.Loan, error) {
	_, err := lr.Collections.InsertOne(context.TODO(), loan)
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}

func (lr *LoanRepo) GetLoanByID(id string) (domain.Loan, error) {
	var loan domain.Loan
	err := lr.Collections.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&loan)
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}

func (lr *LoanRepo) GetAllLoans() ([]domain.Loan, error) {
	var loans []domain.Loan
	cursor, err := lr.Collections.Find(context.TODO(), bson.M{})
	if err != nil {
		return []domain.Loan{}, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var loan domain.Loan
		cursor.Decode(&loan)
		loans = append(loans, loan)
	}
	return loans, nil
}

func (lr *LoanRepo) ChangeLoanStatus(id string, status string) (domain.Loan, error) {
	var loan domain.Loan
	err := lr.Collections.FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status}}).Decode(&loan)
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}

func (lr *LoanRepo) DeleteLoan(id string) error {
	_, err := lr.Collections.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}


