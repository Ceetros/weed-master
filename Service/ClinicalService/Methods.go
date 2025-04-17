package ClinicalService

import (
	"Api/Abstractions/Repositories"
	"Api/Data/Models"
	"Api/Data/Request"
	"github.com/gin-gonic/gin"
)

type ClinicalService struct {
	Repositories.IClinicalRepository
}

func (c ClinicalService) RegisterClinical(req Request.ClinicalRegisterRequest) (int, gin.H) {
	//TODO implement me
	panic("implement me")
}

func (c ClinicalService) GetClinicalByUser(user Models.User) (Models.User, gin.H) {
	//TODO implement me
	panic("implement me")
}

func (c ClinicalService) GetClinicalByDocument(document string) (Models.Clinical, gin.H) {
	//TODO implement me
	panic("implement me")
}
