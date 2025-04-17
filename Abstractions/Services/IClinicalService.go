package Services

import (
	"Api/Data/Models"
	"Api/Data/Request"
	"github.com/gin-gonic/gin"
)

type IClinicalService interface {
	Register(req Request.ClinicalRegisterRequest) (int, gin.H)
	GetClinicalByUser(user Models.User) (int, gin.H)
	GetClinicalByDocument(document string) (int, gin.H)
}
