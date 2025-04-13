package Controller

import (
	"Api/Controller/Auth"
	"Api/Controller/Clinical"
	_ "Api/docs"
	"github.com/gin-gonic/gin"
)

func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	Auth.RegisterController(v1)
	Clinical.RegisterController(v1)
}
