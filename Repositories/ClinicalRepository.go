package Repositories

import (
	"Api/Data/Models"
	"Api/Data/Request"
	"gorm.io/gorm"
)

type ClinicalRepository struct {
	*gorm.DB
}

func (c ClinicalRepository) RegisterClinical(req Request.ClinicalRegisterRequest) error {
	//TODO implement me
	panic("implement me")
}

func (c ClinicalRepository) GetClinicalByUser(user Models.User) (Models.Clinical, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClinicalRepository) GetClinicalByDocument(document string) (Models.Clinical, error) {
	//TODO implement me
	panic("implement me")
}
