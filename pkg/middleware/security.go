package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("TOKEN")
		tokenHeader := ctx.GetHeader("tokenPostman")

		if tokenHeader != token {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Usuario invalido",
			})
			return
		} else {
			ctx.Next()
		}

	}
}
