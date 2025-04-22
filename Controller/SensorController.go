package Controller

import (
	"Api/Abstractions/Services"
	"Api/Data/Request"
	"Api/Discord"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
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
		bot.Bot.ChannelMessageSend(bot.NotificationChannel, fmt.Sprintf("**UMIDADE BAIXA**\nHumidade Relativa:%v%%\n@here\n@everyone", 10))
		channel, err := bot.Bot.UserChannelCreate(os.Getenv("DISCORD_OWNER"))
		if err != nil {
			// If an error occurred, we failed to create the channel.
			//
			// Some common causes are:
			// 1. We don't share a server with the user (not possible here).
			// 2. We opened enough DM channels quickly enough for Discord to
			//    label us as abusing the endpoint, blocking us from opening
			//    new ones.
			fmt.Println("error creating channel:", err)
			return
		}

		_, err = bot.Bot.ChannelMessageSend(channel.ID, fmt.Sprintf("**UMIDADE BAIXA**\nHumidade Relativa:%v%%\n@here\n@everyone", 10))
	}
	c.JSON(ret, state)
}
