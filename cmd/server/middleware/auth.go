package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func TokenMiddleware(ctx *gin.Context) {
	tokenEnvironment := os.Getenv("TOKEN")
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token vazio"})
	}
	if token != tokenEnvironment {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}
	ctx.Next()
}
