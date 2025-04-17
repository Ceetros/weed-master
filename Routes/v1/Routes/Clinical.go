package Routes

import (
	"Api/Config"
	"Api/Middleware"
	"github.com/gin-gonic/gin"
)

func RegisterClinical(r *gin.RouterGroup) {
	clinical := r.Group("/clinical")
	clinical.Use(Middleware.AuthMiddleware())

	controller := Config.ServiceContainer().InjectClinicalController()
	clinical.POST("/register", controller.Register)
}
