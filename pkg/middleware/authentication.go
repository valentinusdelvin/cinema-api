package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (m *Middleware) Authentication(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "token is required",
		})
		ctx.Abort()
		return
	}

	token := strings.Split(bearerToken, " ")[1]
	userID, isAdmin, isSubscribed, err := m.jwt.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed to validate token",
		})
		ctx.Abort()
		return
	}

	ctx.Set("userID", userID)
	ctx.Set("isAdmin", isAdmin)
	ctx.Set("isSubscribed", isSubscribed)

	ctx.Next()
}
