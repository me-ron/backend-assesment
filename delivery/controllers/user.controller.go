package controller

import (
	"loan_tracker/domain"
	utils "loan_tracker/infrastructure/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    UserUsecase domain.UserUsecase
}

func NewUserController(UserUsecase domain.UserUsecase) *UserController {
    return &UserController{UserUsecase: UserUsecase}
}

func (ac *UserController) SignUp(c *gin.Context) {
    var input domain.InputReq
    if err := c.BindJSON(&input); err != nil {
        utils.BadRequest(c)
        return
    }

    user, err := ac.UserUsecase.RegisterUser(input)
    if err != nil {
        utils.Error(c)
        return
    }

    utils.SuccessWithDetailed(user, "User registered successfully", c)

}

func (ac *UserController) LogIn(c *gin.Context) {
    var input domain.InputReq
    if err := c.BindJSON(&input); err != nil {
        utils.BadRequest(c)
        return
    }

    user, accessToken, refreshToken, err := ac.UserUsecase.LoginUser(input.Email, input.Password)
    if err != nil {
        utils.Unauthorized(c)
        return
    }

    bools, _ := ac.UserUsecase.GetBools(user.ID)
    verified := bools.Verified
    if!verified{
        utils.CustomResponse(http.StatusUnauthorized, "you need to be verified", c)
        c.Abort()
        return
    }

    http.SetCookie(c.Writer, &http.Cookie{
        Name:     "refresh_token",
        Value:    refreshToken,
        Path:     "/",
        HttpOnly: true,
    })

    utils.SuccessWithData(gin.H{
        "user":          user,
        "access_token":  accessToken,
    }, c)
}



func (ac *UserController) LogOut(c *gin.Context) {
    http.SetCookie(c.Writer, &http.Cookie{
        Name:     "refresh_token",
        Value:    "",
        Path:     "/",
        HttpOnly: true,
    })

    utils.SuccessWithMessage("Logged out successfully", c)
}



func (ac *UserController) Refresh(c *gin.Context) {
    cookie, err := c.Request.Cookie("refresh_token")
    if err != nil {
        utils.Unauthorized(c)
        return
    }

    refreshToken := cookie.Value

    accessToken, newRefreshToken, err := ac.UserUsecase.RefreshTokens(refreshToken)
    if err != nil {
        utils.Unauthorized(c)
        return
    }

    http.SetCookie(c.Writer, &http.Cookie{
        Name:     "refresh_token",
        Value:    newRefreshToken,
        Path:     "/",
        HttpOnly: true,
    })

    utils.SuccessWithData(gin.H{
        "access_token":  accessToken,
    }, c)
}

func (controller *UserController) GetOneUser(c *gin.Context) {
    id := c.Param("id")

    user,err := controller.UserUsecase.GetOneUser(id)
    if err != nil {
        utils.NotFound(c)
        return
    }

    utils.SuccessWithData(user, c)

}


func (controller *UserController) GetUsers(c *gin.Context) {
    users,err := controller.UserUsecase.GetUsers()

    if err != nil {
        utils.NotFound(c)
        return
    }
    utils.SuccessWithData(users, c)
}

func (controller *UserController) DeleteUser(c *gin.Context) {
    id := c.Param("id")

    err := controller.UserUsecase.DeleteUser(id)
    if err != nil {
        utils.NotFound(c)
        return
    }

    utils.SuccessWithMessage("User deleted successfully", c)
}

func (controller *UserController) SendVerificationEmail(c *gin.Context) {
    var input domain.VerifyEmail
    if err := c.BindJSON(&input); err != nil {
        utils.BadRequest(c)
        return
    }

    err := controller.UserUsecase.SendVerifyEmail(input)
    if err != nil {
        utils.Error(c)
        return
    }

    utils.SuccessWithMessage("Verification email sent successfully", c)
}

func (controller *UserController) VerifyEmail(c *gin.Context) {
    id := c.Param("id")

    err := controller.UserUsecase.VerifyUser(id)
    if err != nil {
        utils.NotFound(c)
        return
    }

    utils.SuccessWithMessage("Email verified successfully", c)
}

