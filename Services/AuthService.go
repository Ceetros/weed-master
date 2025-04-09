package Services

import (
	"Api/Auth"
	"Api/Config"
	"Api/Data/Enum"
	"Api/Data/Models"
	"Api/Data/Request"
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
	var user Models.User
	if err := Config.DB.Preload("ClinicalUser.Clinical").Where("email = ?", email).First(&user).Error; err != nil {
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
	parsedDate, ee := time.ParseInLocation("2006-01-02", req.BornDate, time.UTC)
	if ee != nil {
		return http.StatusBadRequest, gin.H{}
	}

	if req.Type != Enum.Clinical && req.Type != Enum.User {
		return http.StatusBadRequest, gin.H{"error": "Tipo de usuário inválido"}
	}

	if req.Document == "" || !Utils.ValidateDocument(req.Document) {
		return http.StatusBadRequest, gin.H{"error": "Documento inválido"}
	}

	isClinical := req.Type == Enum.Clinical

	if isClinical && (req.Clinical.Document == "" || !Utils.ValidateDocument(req.Clinical.Document)) {
		return http.StatusBadRequest, gin.H{"error": "Documento da Clínica inválido"}
	}

	var existing Models.User
	if err := Config.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
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

		if isClinical {
			user.Role = Enum.Manager
		}

		var clinical Models.Clinical
		var address Models.Address

		if isClinical {
			clinical = Models.Clinical{
				Name:        req.Clinical.Name,
				Document:    req.Clinical.Document,
				PixKey:      req.Clinical.PixKey,
				NextPayment: time.Now(),
			}
			address = Models.Address{
				Street:     req.Clinical.Street,
				Number:     req.Clinical.Number,
				Complement: req.Clinical.Complement,
				City:       req.Clinical.City,
				State:      req.Clinical.State,
				ZipCode:    req.Clinical.ZipCode,
				ClinicalID: clinical.ID,
			}

			if err := tx.Create(&clinical).Error; err != nil {
				return err
			}
			if err := tx.Create(&address).Error; err != nil {
				return err
			}
		}

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		if isClinical {
			link := Models.ClinicalUser{
				ID:         uuid.New(),
				UserId:     user.ID,
				ClinicalId: clinical.ID,
			}
			if err := tx.Create(&link).Error; err != nil {
				return err
			}

			user.ClinicalUser = &link
			user.ClinicalUser.Clinical = clinical
		}

		return nil
	})

	if err != nil {
		return http.StatusInternalServerError, gin.H{"error": "Erro ao registar"}
	}

	return LoginWithUser(req.Email, req.Password)
}
