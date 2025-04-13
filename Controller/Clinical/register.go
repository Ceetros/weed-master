package Clinical

import "github.com/gin-gonic/gin"

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
func Register(c *gin.Context) {
}
