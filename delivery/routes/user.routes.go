package routes

import (
	"loan_tracker/delivery/controllers"
	"loan_tracker/domain"
	middleware "loan_tracker/infrastructure/middlewares"
	passwordservice "loan_tracker/infrastructure/password_service"
	tokenservice "loan_tracker/infrastructure/token_service"
	"loan_tracker/repository"
	"loan_tracker/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(router *gin.Engine, userCollection domain.CollectionInterface) {
	UserRepo, err := repository.NewUserRepo(userCollection)
	if err != nil{
		log.Panic(err.Error())
	}

	access_secret := os.Getenv("ACCESSTOKENSECRET")
	if access_secret == ""{
		log.Panic("No accesstoken")
	}
	
	refresh_secret := os.Getenv("REFRESHTOKENSECRET")
	if refresh_secret == ""{
		log.Panic("No refreshtoken")
	}

	verfication_secret := os.Getenv("VERIFICATIONTOKEN")
	if verfication_secret == ""{
		log.Panic("No verificationtoken")
	}
	
	TokenSvc := tokenservice.NewTokenService(access_secret, refresh_secret, verfication_secret)
	PasswordSvc := &passwordservice.PasswordS{}


	UserUsecase := usecase.NewUserUsecase(UserRepo, PasswordSvc, TokenSvc)
	UserController := controller.NewUserController(UserUsecase)

	LoggedIN := middleware.LoggedIn(TokenSvc)
	mustbeAdmin := middleware.RoleBasedAuth(true, UserRepo)
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", UserController.SignUp)
		userRouter.POST("/login", UserController.LogIn)
		userRouter.POST("/verify-email", UserController.VerifyEmail)
		userRouter.GET("/profile", LoggedIN, UserController.GetOneUser)
		userRouter.POST("/token/refresh", LoggedIN, UserController.Refresh)
		userRouter.POST("/logout", LoggedIN, UserController.LogOut)
		userRouter.POST("/password-reset", LoggedIN, UserController.SendForgetPasswordEmail)
		userRouter.POST("/password-update", LoggedIN, UserController.ResetPassword)
	}

	admin := router.Group("/admin")
	{
		admin.GET("/users", LoggedIN, mustbeAdmin, UserController.GetUsers)
		admin.DELETE("/user/:id", LoggedIN, mustbeAdmin, UserController.DeleteUser)
	}
}