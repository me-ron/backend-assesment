package usecase

import "loan_tracker/domain"

type LoanUsecase struct {
	LoanRepo domain.LoanRepository
}

func NewLoanUsecase(lr domain.LoanRepository) *LoanUsecase {
	return &LoanUsecase{
		LoanRepo: lr,
	}
}

func (lu *LoanUsecase) CreateLoan(loan domain.Loan) (domain.Loan, error) {
	return lu.LoanRepo.CreateLoan(loan)
}

func (lu *LoanUsecase) GetLoanByID(id string) (domain.Loan, error) {
	return lu.LoanRepo.GetLoanByID(id)
}

func (lu *LoanUsecase) GetAllLoans() ([]domain.Loan, error) {
	return lu.LoanRepo.GetAllLoans()
}

func (lu *LoanUsecase) ChangeLoanStatus(id string, status string) (domain.Loan, error) {
	return lu.LoanRepo.ChangeLoanStatus(id, status)
}

func (lu *LoanUsecase) DeleteLoan(id string) error {
	return lu.LoanRepo.DeleteLoan(id)
}