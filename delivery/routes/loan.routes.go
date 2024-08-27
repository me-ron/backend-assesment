package routes

import (
	controller "loan_tracker/delivery/controllers"
	"loan_tracker/domain"
	middleware "loan_tracker/infrastructure/middlewares"
	tokenservice "loan_tracker/infrastructure/token_service"
	"loan_tracker/repository"
	"loan_tracker/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func NewLoanRouter(router *gin.Engine, coll domain.CollectionInterface, user domain.CollectionInterface) {
	loanRepo := repository.NewLoanRepo(coll)
	loanUsecase := usecase.NewLoanUsecase(loanRepo)
	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == ""{
		log.Panic("No accesstoken")
	}
	
	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == ""{
		log.Panic("No refreshtoken")
	}

	verfication_secret := os.Getenv("VERIFICATIONTOKENSECRET")
	if verfication_secret == ""{
		log.Panic("No verificationtoken")
	}
	
	TokenSvc := tokenservice.NewTokenService(access_secret, refresh_secret, verfication_secret)
	LoggedIN := middleware.LoggedIn(TokenSvc)
	userrepo, _ := repository.NewUserRepo(user)
	mustbeAdmin := middleware.RoleBasedAuth(true, userrepo)

	lc := controller.NewLoanController(loanUsecase)
	lr := router.Group("/loan")
	{
		lr.POST("/", LoggedIN, lc.CreateLoan)
		lr.GET("/:id", LoggedIN, lc.GetLoanByID)
	}

	admin := router.Group("/admin/loans/")
	{
		admin.GET("/", LoggedIN, mustbeAdmin, lc.GetAllLoans)
		admin.PUT("/:id/status", LoggedIN, mustbeAdmin, lc.ChangeLoanStatus)
		admin.DELETE("/:id", LoggedIN, mustbeAdmin, lc.DeleteLoan)
	}
}