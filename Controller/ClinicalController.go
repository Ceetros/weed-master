package Controller

import (
	"Api/Abstractions/Services"
	"github.com/gin-gonic/gin"
)

type ClinicalController struct {
	IClinicalService Services.IClinicalService
}

// Register godoc
// @Summary Registra uma nova clinica
// @Tags Clinical
// @Accept json
// @Produce json
// @Param data body Request.ClinicalRegisterRequest true "Dados da nova Clinica"
// @Success 201 {object} Response.ClinicalRegisterResponse
// @Failure 400 {object} Response.ErrorResponse "Requisição inválida"
// @Failure 500 {object} Response.ErrorResponse "Erro interno"
// @Router /api/v1/clinicl/register [post]
func (controller *ClinicalController) Register(c *gin.Context) {
}
