package Services

import (
	"Api/Data/Request"
	"github.com/gin-gonic/gin"
)

type ISesorService interface {
	Update(req Request.UpdateSensorRequest) (int, gin.H)
}
