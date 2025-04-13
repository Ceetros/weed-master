package AuthService

import (
	"Api/Auth"
	"Api/Config"
	"Api/Data/Enum"
	"Api/Data/Models"
	"Api/Data/Request"
	"Api/Repository/UserRepository"
	"Api/Utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func LoginWithUser(email string, password string) (int, gin.H) {
	user := UserRepository.GetUserByEmail(email)
	if user == nil {
		return http.StatusUnauthorized, gin.H{"error": "Usuário ou Senha inválidos"}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Usuário ou Senha inválidos"}
	}

	customToken, err := Auth.FirebaseAuth.CustomToken(context.Background(), user.ID.String())
	if err != nil {
		return http.StatusUnauthorized, gin.H{"error": "Erro interno: EF002"}
	}

	clinical := gin.H{}

	if user.Role != Enum.Tutor {
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

func RegisterUser(req Request.RegisterRequest) (int, gin.H) {
	parsedDate, ee := time.ParseInLocation("2006-01-02", req.BirthDate, time.UTC)
	if ee != nil {
		return http.StatusBadRequest, gin.H{}
	}

	if req.Document == "" || !Utils.ValidateDocument(req.Document) {
		return http.StatusBadRequest, gin.H{"error": "Documento inválido"}
	}

	if UserRepository.GetUserByEmail(req.Email) == nil {
		return http.StatusBadRequest, gin.H{"error": "Email já cadastrado"}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	err := Config.DB.Transaction(func(tx *gorm.DB) error {
		user := Models.User{
			ID:        uuid.New(),
			Name:      req.Name,
			Email:     req.Email,
			Password:  string(hashedPassword),
			Role:      Enum.Tutor,
			CPF:       req.Document,
			BirthDate: parsedDate.Local(),
		}
		if err := UserRepository.RegisterUser(user, tx); err != nil {
			tx.Rollback()
			return *err
		}

		return nil
	})

	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": "Erro ao registar"}
	}

	return LoginWithUser(req.Email, req.Password)
}
