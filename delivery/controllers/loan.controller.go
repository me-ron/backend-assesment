package controller

import (
	"loan_tracker/domain"
	utils "loan_tracker/infrastructure/utilities"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	LoanUsecase domain.LoanUsecase
}

func NewLoanController(lu domain.LoanUsecase) *LoanController {
	return &LoanController{
		LoanUsecase: lu,
	}
}

func (lc *LoanController) CreateLoan(c *gin.Context) {
	var loan domain.Loan
	c.BindJSON(&loan)
	result, err := lc.LoanUsecase.CreateLoan(loan)
	if err != nil {
		utils.BadRequest(c)
		return
	}
	utils.SuccessWithData(result, c)
}

func (lc *LoanController) GetLoanByID(c *gin.Context) {
	id := c.Param("id")
	result, err := lc.LoanUsecase.GetLoanByID(id)
	if err != nil {
		utils.NotFound(c)
		return
	}
	utils.SuccessWithData(result, c)
}

func (lc *LoanController) GetAllLoans(c *gin.Context) {
	result, err := lc.LoanUsecase.GetAllLoans()
	if err != nil {
		utils.NotFound(c)
		return
	}
	utils.SuccessWithData(result, c)
}

func (lc *LoanController) ChangeLoanStatus(c *gin.Context) {
	id := c.Param("id")
	var status domain.Status
	c.BindJSON(&status)
	result, err := lc.LoanUsecase.ChangeLoanStatus(id, status.Status)
	if err != nil {
		utils.NotFound(c)
		return
	}
	utils.SuccessWithData(result, c)
}

func (lc *LoanController) DeleteLoan(c *gin.Context) {
	id := c.Param("id")
	err := lc.LoanUsecase.DeleteLoan(id)
	if err != nil {
		utils.NotFound(c)
		return
	}
	utils.Success(c)
}




