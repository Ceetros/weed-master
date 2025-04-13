package Response

import (
	"Api/Data/Enum"
)

type LoginResponse struct {
	Token    string       `json:"token" example:"firebase-token"`
	User     UserInfo     `json:"user"`
	Clinical ClinicalInfo `json:"clinical"`
}

type UserInfo struct {
	Name  string        `json:"name" example:"John Doe"`
	Email string        `json:"email" example:"johndoe@example.com"`
	Role  Enum.RoleEnum `json:"role" example:"TUTOR"`
}

type ClinicalInfo struct {
	Name      string `json:"name" example:"John Doe's Clinical Company'"`
	ExpiresIn string `json:"expiresIn" example:"2025-04-08T21:45:00Z"`
	Expired   bool   `json:"expired" example:"false"`
}
