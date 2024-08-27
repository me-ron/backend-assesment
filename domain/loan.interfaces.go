package domain

import "github.com/gin-gonic/gin"

type LoanController interface {
	CreateLoan(c *gin.Context)
	GetLoanByID(c *gin.Context)
	GetAllLoans(c *gin.Context)
	ChangeLoanStatus(c *gin.Context)
	DeleteLoan(c *gin.Context)
}


type LoanUsecase interface {
	CreateLoan(loan Loan) (Loan, error)
	GetLoanByID(id string) (Loan, error)
	GetAllLoans() ([]Loan, error)
	ChangeLoanStatus(id string, status string) (Loan, error)
	DeleteLoan(id string) error
}

type LoanRepository interface {
	CreateLoan(loan Loan) (Loan, error)
	GetLoanByID(id string) (Loan, error)
	GetAllLoans() ([]Loan, error)
	ChangeLoanStatus(id string, status string) (Loan, error)
	DeleteLoan(id string) error
}

