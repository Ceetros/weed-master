package Routes

import (
	"Api/Config"
	"github.com/gin-gonic/gin"
)

func RegisterAuth(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	authController := Config.ServiceContainer().InjectAuthController()
	auth.POST("/login", authController.Login)
	auth.POST("/register", authController.Register)
}
