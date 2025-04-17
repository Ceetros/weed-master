package Controller

import (
	"Api/Abstractions/Services"
	"Api/Data/Request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	Services.IUserService
}

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
func (controller *AuthController) Login(c *gin.Context) {
	var req Request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email e senha são obrigatórios"})
		return
	}

	status, ret := controller.LoginWithUser(req.Email, req.Password)
	c.JSON(status, ret)
}

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
func (controller *AuthController) Register(c *gin.Context) {
	var req Request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	status, ret := controller.RegisterUser(req)

	c.JSON(status, ret)
}
