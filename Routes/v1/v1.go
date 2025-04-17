package v1

import (
	"Api/Routes/v1/Routes"
	"github.com/gin-gonic/gin"
)

func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	Routes.RegisterAuth(v1)
	Routes.RegisterClinical(v1)
}
