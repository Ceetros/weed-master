package Controller

import (
	"Api/Abstractions/Services"
	"Api/Data/Request"
	"Api/Discord"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type SesorController struct {
	ISesorService Services.ISesorService
}

// Register godoc
// @Summary Atualiza um sensor
// @Tags Sensor
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body Request.UpdateSensorRequest true "Dados do sensor"
// @Success 201 {object} bool
// @Failure 400 {object} bool
// @Failure 500 {object} bool
// @Router /api/v1/sesor/update [post]
func (controller *SesorController) Update(c *gin.Context) {
	var request Request.UpdateSensorRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Fatalf("Error binding JSON: %v", err)
		return
	}

	ret, state := controller.ISesorService.Update(request)
	if request.UmidityPercent <= 79 {
		bot := Discord.GetDiscord()
		bot.Bot.ChannelMessageSend(bot.NotificationChannel, fmt.Sprintf("**UMIDADE BAIXA**\nHumidade Relativa:%v%%\n@here\n@everyone", request.UmidityPercent))
	}
	c.JSON(ret, state)
}
