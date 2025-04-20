package Routes

import (
	"Api/Config"
	"github.com/gin-gonic/gin"
)

func RegisterClinical(r *gin.RouterGroup) {
	clinical := r.Group("/sesor")

	controller := Config.ServiceContainer().InjectClinicalController()
	clinical.POST("/update", controller.Update)
}
