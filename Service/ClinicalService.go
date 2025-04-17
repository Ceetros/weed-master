package Service

import (
	"Api/Abstractions/Repositories"
	"Api/Data/Models"
	"Api/Data/Request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClinicalService struct {
	IClinicalRepository Repositories.IClinicalRepository
}

func (c ClinicalService) Register(req Request.ClinicalRegisterRequest) (int, gin.H) {
	status, err := c.GetClinicalByDocument(req.Document)
	if status != http.StatusOK {
		return http.StatusInternalServerError, gin.H{"error": err}
	}

	if regerr := c.IClinicalRepository.RegisterClinical(req); regerr != nil {
		return http.StatusInternalServerError, gin.H{"error": regerr.Error()}
	}

	return http.StatusOK, gin.H{}
}

func (c ClinicalService) GetClinicalByUser(user Models.User) (int, gin.H) {
	ret, err := c.IClinicalRepository.GetClinicalByUser(user)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}

	return http.StatusOK, gin.H{
		"name":        ret.Name,
		"document":    ret.Document,
		"nextPayment": ret.NextPayment,
	}
}

func (c ClinicalService) GetClinicalByDocument(document string) (int, gin.H) {
	ret, err := c.IClinicalRepository.GetClinicalByDocument(document)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}

	return http.StatusOK, gin.H{
		"name":        ret.Name,
		"document":    ret.Document,
		"nextPayment": ret.NextPayment,
	}
}
