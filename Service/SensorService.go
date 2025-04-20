package Service

import (
	"Api/Abstractions/Repositories"
	"Api/Data/Request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SensorService struct {
	ISensorRepository Repositories.ISesorRepository
}

func (c SensorService) Update(req Request.UpdateSensorRequest) (int, gin.H) {
	if er := c.ISensorRepository.Update(req); er != nil {
		return http.StatusBadGateway, gin.H{}
	}

	return http.StatusOK, gin.H{}
}
