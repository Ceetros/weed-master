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
	clinical := Models.Clinical{
		Name:     req.Name,
		Document: req.Document,
	}

	return c.Create(&clinical).Error
}

func (c ClinicalRepository) GetClinicalByUser(user Models.User) (Models.Clinical, error) {
	var ret Models.Clinical
	err := c.Table("clinicals").Joins("INNER JOIN clinical_users cu ON clinicals.Id = cu.clinical_id").Where("cu.user_id = ?", user.ID).First(&ret).Error
	return ret, err
}

func (c ClinicalRepository) GetClinicalByDocument(document string) (Models.Clinical, error) {
	var ret Models.Clinical
	err := c.Table("clinicals").Where("document = ?", document).First(&ret).Error

	return ret, err
}
