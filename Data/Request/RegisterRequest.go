package Request

import (
	"Api/Data/Enum"
)

type RegisterRequest struct {
	Name     string               `json:"name" binding:"required"`
	Document string               `json:"document" binding:"required"`
	Email    string               `json:"email" binding:"required,email"`
	Password string               `json:"password" binding:"required"`
	BornDate string               `json:"birthDate" binding:"required"`
	Type     Enum.AccountTypeEnum `json:"type" example:"CLINICAL"`
	Clinical *ClinicalData        `json:"clinical,omitempty"`
}

type ClinicalData struct {
	Name       string `json:"name"`
	Document   string `json:"document"`
	PixKey     string `json:"pixKey"`
	PixType    string `json:"pixType"`
	ZipCode    string `json:"zipCode"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	City       string `json:"city"`
	State      string `json:"state"`
}
