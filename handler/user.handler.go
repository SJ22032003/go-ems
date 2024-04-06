package handler

import (
	"net/http"

	model "github.com/SJ22032003/go-ems/models"
	service "github.com/SJ22032003/go-ems/services"
	util "github.com/SJ22032003/go-ems/utils"
	gin "github.com/gin-gonic/gin"
)

const (
	USER_ALREADY_EXISTS_ERR = "UNIQUE constraint failed: users.email"
)

type UserHandler struct {
	User model.User
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	user := u.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, Please check the request body and try again.",
			"error":   err.Error(),
		})
		return
	}

	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	newUser, err := service.CreateUserService(&user)
	if err != nil {
		if err.Error() == USER_ALREADY_EXISTS_ERR {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "User with this email already exists",
				"error":   err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
		"user":    newUser.UserDetailsWithoutPassword(),
	})

}

func (u *UserHandler) LoginUser(ctx *gin.Context) {
	type LoginUser struct {
		Email    string `binding:"required" json:"email"`
		Password string `binding:"required" json:"password"`
	}

	loginUser := LoginUser{}

	err := ctx.ShouldBindJSON(&loginUser)
	println(loginUser.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, Please check the request body and try again.",
			"error":   err.Error(),
		})
		return
	}

	user, err := service.FindOneUserByEmail(loginUser.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"error":   err.Error(),
		})
		return
	}

	isPasswordValid := util.CheckPasswordHash(loginUser.Password, user.Password)
	if !isPasswordValid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Password",
		})
		return
	}

	token, err := util.GetSignedToken(*user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error, Please try again later.",
			"error":   err.Error(),
		})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Successfully Logged In",
		"auth_token": token,
	})

}
