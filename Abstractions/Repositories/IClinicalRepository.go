package Repositories

import (
	"Api/Data/Request"
)

type ISesorRepository interface {
	Update(req Request.UpdateSensorRequest) error
}
