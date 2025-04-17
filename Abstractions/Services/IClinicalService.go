package Services

import (
	"Api/Data/Models"
	"Api/Data/Request"
	"github.com/gin-gonic/gin"
)

type IClinicalService interface {
	RegisterClinical(req Request.ClinicalRegisterRequest) (int, gin.H)
	GetClinicalByUser(user Models.User) (Models.User, gin.H)
	GetClinicalByDocument(document string) (Models.Clinical, gin.H)
}
