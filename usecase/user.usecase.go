package usecase

import (
	"errors"
	"fmt"
	"loan_tracker/domain"
	utils "loan_tracker/infrastructure/utilities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	UserRepo      domain.UserRepository
	PasswordSrv   domain.PasswordService
	TokenSrv      domain.TokenService
}

func NewUserUsecase(userrepo domain.UserRepository, passwordSrv domain.PasswordService, tokenSrv domain.TokenService) *UserUsecase {
	return &UserUsecase{
		UserRepo:    userrepo,
		PasswordSrv: passwordSrv,
		TokenSrv:    tokenSrv,
	}
}

func (u *UserUsecase) RegisterUser(input domain.InputReq) (domain.UserResponse, error) {
    var user domain.UserInfo

    // Hash the user's password
    hashedPassword, err := u.PasswordSrv.HashPassword(input.Password)
    if err != nil {
        return domain.CreateResponseUser(user), err
    }

    // Create the user model
    user = domain.UserInfo{
        ID:                primitive.NewObjectID(), 
        Name:          	   input.Name,
        Email:             input.Email,
        Password:          hashedPassword,
        Verified:        false,
    }

    err = u.UserRepo.SaveUser(&user)
    if err != nil {
        return domain.CreateResponseUser(user), err
    }

    return domain.CreateResponseUser(user), nil
}

func (u *UserUsecase) LoginUser(email, password string) (domain.UserResponse, string, string, error) {
    var user domain.UserInfo

    foundUser, err := u.UserRepo.FindUserByEmail(email)
    if err != nil {
        return domain.CreateResponseUser(user), "", "", err
    }   

    if foundUser == nil {
        return domain.CreateResponseUser(user), "", "", fmt.Errorf("user not found")
    }

    isMatch, err := u.PasswordSrv.ComparePassword(foundUser.Password, password)
    if err != nil {
        return domain.CreateResponseUser(user), "", "", err
    }

    if !isMatch {
        return domain.CreateResponseUser(user), "", "", fmt.Errorf("invalid password")
    }

    accessToken, err := u.TokenSrv.GenerateAccessToken(foundUser.ID.Hex())
    if err != nil {
        return domain.CreateResponseUser(user), "", "", err
    }

    refreshToken, err := u.TokenSrv.GenerateRefreshToken(foundUser.ID.Hex())
    if err != nil {
        return domain.CreateResponseUser(user), "", "", err
    }

    return domain.CreateResponseUser(*foundUser), accessToken, refreshToken, nil
}

func (u *UserUsecase) RefreshTokens(refreshToken string) (string, string, error) {
    userId, err := u.TokenSrv.ValidateRefreshToken(refreshToken)
    if err != nil {
        return "", "", err
    }

    newAccessToken, err := u.TokenSrv.GenerateAccessToken(userId)
    if err != nil {
        return "", "", err
    }

    newRefreshToken, err := u.TokenSrv.GenerateRefreshToken(userId)
    if err != nil {
        return "", "", err
    }

    return newAccessToken, newRefreshToken, nil
}

func (u *UserUsecase) GetOneUser(id string) (domain.UserResponse, error) {
	user,err := u.UserRepo.GetUserDocumentByID(id)
	if err != nil {
		return domain.UserResponse{},err
	}
	return user, nil
}

func (u *UserUsecase) GetUsers() ([]domain.UserResponse, error) {
	users, err := u.UserRepo.GetUserDocuments()
	if err != nil {
		return []domain.UserResponse{},err
	}
	return users, nil
}

func (u *UserUsecase) DeleteUser(id string) error {
	return u.UserRepo.DeleteUserDocument(id)
}

func (u *UserUsecase)GetBools(id string) (domain.Bools, error){
    return u.UserRepo.GetBools(id)
}

func (u *UserUsecase) SendVerifyEmail(id string , vuser domain.VerifyEmail) error {
    bools, _ := u.UserRepo.GetBools(id)

	if bools.Verified {
		return errors.New("user already verified")
	}

	token,err := u.TokenSrv.GenerateVerificationToken(id)
	if err != nil {
		return err
	}
	subject,body := utils.BodyVerify(token)

	err = utils.SendEmail(vuser.Email, subject , body)
	if err != nil {
		return err
	}
	
	return nil
}


func (u *UserUsecase) VerifyUser(token string) error {
	id,err := u.TokenSrv.ValidateVerificationToken(token)
	if err != nil {
		return err
	}
	return u.UserRepo.VerifyUser(id)
}

func (u *UserUsecase) SendForgretPasswordEmail(vuser domain.VerifyEmail) error {
    user,err := u.UserRepo.FindUserByEmail(vuser.Email)
    if err != nil {
        return err
    }
    if user == nil {
        return errors.New("user not found")
    }

    token,err := u.TokenSrv.GeneratePasswordToken(user.ID.Hex())
    if err != nil {
        return err
    }
    subject,body := utils.BodyForgetPassword(token)

    err = utils.SendEmail(vuser.Email, subject , body)
    if err != nil {
        return err
    }
    
    return nil
}


func (u *UserUsecase) ValidateForgetPassword(input domain.UpdatePassword) error {
    id, err := u.TokenSrv.ValidatePasswordToken(input.Token)
    if err != nil {
        return err
    }

    hashedPassword, err := u.PasswordSrv.HashPassword(input.Password)
    if err != nil {
        return err
    }

    return u.UserRepo.UpdatePassword(id, hashedPassword)
}