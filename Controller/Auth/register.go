package Auth

import (
	"Api/Data/Request"
	"Api/Service/AuthService"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register godoc
// @Summary Registra um novo usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body Request.RegisterRequest true "Dados do novo usuário"
// @Success 201 {object} Response.LoginResponse
// @Failure 400 {object} Response.ErrorResponse "Requisição inválida"
// @Failure 500 {object} Response.ErrorResponse "Erro interno"
// @Router /api/v1/auth/register [post]
func Register(c *gin.Context) {
	var req Request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	status, ret := AuthService.RegisterUser(req)

	c.JSON(status, ret)
}
