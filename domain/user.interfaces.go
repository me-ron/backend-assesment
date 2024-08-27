package domain

import "github.com/gin-gonic/gin"

type UserController interface {
	SignUp(gin.Context)
	// VerifyEmail(gin.Context)
	LogIn(gin.Context)
	// ForgotPassword(gin.Context)
	// ResetPassword(gin.Context)
	RefreshToken(gin.Context)
	GetUserProfile(gin.Context)
	ViewAllUsers(gin.Context)
	DeleteUser(gin.Context)
}

type UserUsecase interface{
	RegisterUser(InputReq) (UserResponse, error)
	LoginUser(string, string) (UserResponse, string, string, error)
	RefreshTokens(string) (string, string, error) 
	GetOneUser(string) (UserResponse , error) 
	GetUsers() ([]UserResponse , error)
	DeleteUser(string) (error)
	GetBools(string) (Bools, error)
}

type UserRepository interface{
	SaveUser(*UserInfo) error
	FindUserByEmail(string) (*UserInfo, error)
	GetUserDocumentByID(id string) (UserResponse, error)
	GetUserDocuments() ([]UserResponse , error)
	DeleteUserDocument(id string) (error)
	GetBools(string) (Bools, error)

}