package Repositories

import (
	"Api/Data/Models"
	"Api/Data/Request"
	"gorm.io/gorm"
)

type SensorRepository struct {
	*gorm.DB
}

func (c SensorRepository) Update(req Request.UpdateSensorRequest) error {
	update := Models.Analytics{
		WeedId:  req.WeedId,
		Umidity: req.UmidityPercent,
	}

	return c.DB.Create(&update).Error
}
