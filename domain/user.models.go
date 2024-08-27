package domain

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InputReq struct {
	Name 	 string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserInfo struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"Password"`
	VerifiedCode string `json:"verifiedCode,omitempty" bson:"VerifiedCode,omitempty"`
	Verified     bool   `json:"verified" bson:"Verified" default:"false"`
	IsAdmin      bool   `json:"isadmin"  bson:"isadmin"`
}

type UserResponse struct {
	ID    string `json:"_id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

type UserClaims struct{
	ID string
	jwt.StandardClaims
}

type UpdatePassword struct {
	Password string `json:"password" bson:"password"`
	Token   string `json:"token" bson:"token"`
}

type Bools struct{
	IsAdmin bool   `json:"isadmin"  bson:"isadmin"`
	Verified bool     `json:"verified" bson:"Verified"`
}

type VerifyEmail struct {
	Email string `json:"email" bson:"email"`
}

// from actual user model to response model to be done in usecase
func CreateResponseUser(user UserInfo) UserResponse {
	return UserResponse{
		ID:         user.ID.Hex(),
		Name:       user.Name,
		Email:          user.Email,
	}
}
