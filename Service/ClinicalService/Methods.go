package ClinicalService

import (
	"Api/Abstractions/Repositories"
	"Api/Data/Models"
	"Api/Data/Request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClinicalService struct {
	clinicalRepository Repositories.IClinicalRepository
}

func (c ClinicalService) Register(req Request.ClinicalRegisterRequest) (int, gin.H) {
	_, err := c.GetClinicalByDocument(req.Document)
	if err == nil {
		return http.StatusUnauthorized, gin.H{"error": "Usuário ou Senha inválidos"}
	}

	if regerr := c.clinicalRepository.RegisterClinical(req); regerr != nil {
		return http.StatusInternalServerError, gin.H{"error": regerr.Error()}
	}

	return http.StatusOK, gin.H{}
}

func (c ClinicalService) GetClinicalByUser(user Models.User) (int, gin.H) {
	ret, err := c.clinicalRepository.GetClinicalByUser(user)
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
	ret, err := c.clinicalRepository.GetClinicalByDocument(document)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}

	return http.StatusOK, gin.H{
		"name":        ret.Name,
		"document":    ret.Document,
		"nextPayment": ret.NextPayment,
	}
}
