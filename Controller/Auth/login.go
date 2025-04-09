package Auth

import (
	"Api/Data/Request"
	"Api/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary Login com email e senha
// @Description Autentica o usuário e retorna tokens de acesso
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body Request.LoginRequest true "Credenciais"
// @Success 200 {object} Response.LoginResponse
// @Failure 400 {object} Response.ErrorResponse
// @Failure 401 {object} Response.ErrorResponse
// @Router /api/v1/auth/login [post]
func Login(c *gin.Context) {
	var req Request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email e senha são obrigatórios"})
		return
	}

	status, ret := Services.LoginWithUser(req.Email, req.Password)
	c.JSON(status, ret)
}
