package middlewares

import (
	"go-rest/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message":     "Token Required",
			"status_code": http.StatusUnauthorized,
		})
		return
	}
	token = strings.Replace(token, "Bearer ", "", 1)

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message":     "Not Authorized",
			"status_code": http.StatusUnauthorized,
		})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
