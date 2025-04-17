package Swagger

import (
	"Api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"os"
)

func RegisterController(r *gin.Engine) {
	var config = ginSwagger.Config{
		URL:                      "doc.json",
		DocExpansion:             "list",
		InstanceName:             swag.Name,
		Title:                    "Vet SAAS",
		DefaultModelsExpandDepth: 1,
		DeepLinking:              true,
		PersistAuthorization:     false,
		Oauth2DefaultClientID:    "",
	}
	docs.SwaggerInfo.Host = os.Getenv("API_URL") + ":" + os.Getenv("API_PORT")

	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&config, swaggerFiles.Handler))
}
