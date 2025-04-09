package Controller

import (
	"Api/Controller/Auth"
	_ "Api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	auth := v1.Group("/auth")
	//auth.Use(Middleware.AuthMiddleware())
	auth.POST("/login", Auth.Login)
	auth.POST("/register", Auth.Register)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
