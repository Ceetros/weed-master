package Auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterController(v1 *gin.RouterGroup) {
	auth := v1.Group("/auth")

	auth.POST("/login", Login)
	auth.POST("/register", Register)
}
