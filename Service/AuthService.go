package Service

import (
	"Api/Abstractions/Repositories"
	"Api/Auth"
	"Api/Data/Enum"
	"Api/Data/Models"
	"Api/Data/Request"
	"Api/Utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserService struct {
	Repositories.IUserRepository
}

func (service *UserService) LoginWithUser(email string, password string) (int, gin.H) {
	user, err := service.GetUserByEmail(email)
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Usuário ou Senha inválidos"}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Usuário ou Senha inválidos"}
	}

	customToken, err := Auth.FirebaseAuth.CustomToken(context.Background(), user.ID.String())
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Erro interno: EF002"}
	}

	clinical := gin.H{}

	if user.Role != Enum.Tutor && user.Role != "" {
		clinical = gin.H{
			"name":      user.ClinicalUser.Clinical.Name,
			"expiresIn": user.ClinicalUser.Clinical.NextPayment.Format(time.RFC3339),
			"expired":   user.ClinicalUser.Clinical.NextPayment.Before(time.Now()),
		}
	}

	ret := gin.H{"token": customToken,
		"user": gin.H{
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
		"clinical": clinical,
	}

	return http.StatusOK, ret
}

func (service *UserService) RegisterUser(req Request.RegisterRequest) (int, gin.H) {
	parsedDate, ee := time.ParseInLocation("2006-01-02", req.BirthDate, time.UTC)
	if ee != nil {
		return http.StatusBadRequest, gin.H{"error": ee}
	}

	if req.Document == "" || !Utils.ValidateDocument(req.Document) {
		return http.StatusBadRequest, gin.H{"error": "Documento inválido"}
	}

	if _, err := service.GetUserByEmail(req.Email); err == nil {
		return http.StatusBadRequest, gin.H{"error": "Email já cadastrado"}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := Models.User{
		ID:        uuid.New(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CPF:       req.Document,
		BirthDate: parsedDate.Local(),
	}
	err := service.CreateUser(user)

	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": "Erro ao registar"}
	}

	return service.LoginWithUser(req.Email, req.Password)
}
