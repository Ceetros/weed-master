package Repositories

import (
	"Api/Data/Models"
	"Api/Data/Request"
)

type IClinicalRepository interface {
	RegisterClinical(req Request.ClinicalRegisterRequest) error
	GetClinicalByUser(user Models.User) (Models.Clinical, error)
	GetClinicalByDocument(document string) (Models.Clinical, error)
}
