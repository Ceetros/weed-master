package Middleware

import (
	"context"
	"net/http"
	"strings"

	"Api/Auth" // ajuste para seu path
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não encontrado"})
			c.Abort()
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		client, err := Auth.FirebaseApp.Auth(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter Firebase Auth"})
			c.Abort()
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("uid", token.UID)
		c.Next()
	}
}
