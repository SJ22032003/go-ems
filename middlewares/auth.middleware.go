package middleware

import (
	"net/http"
	"strings"

	util "github.com/SJ22032003/go-ems/utils"
	gin "github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorization := ctx.GetHeader("Authorization")
		if authorization == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "Authorization header is required",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.Split(authorization, " ")[1]
		token, err := util.VerifyToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "Invalid token claims",
			})
			ctx.Abort()
			return
		}

		id := claims["id"]
		if id == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "Invalid user id",
			})
			ctx.Abort()
			return
		} 

		ctx.Set("user", claims)
		ctx.Next()

	}
}
