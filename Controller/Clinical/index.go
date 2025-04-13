package Clinical

import (
	"Api/Middleware"
	"github.com/gin-gonic/gin"
)

func RegisterController(v1 *gin.RouterGroup) {
	clinical := v1.Group("/clinical")
	clinical.Use(Middleware.AuthMiddleware())
	clinical.POST("/register", Register)
}
