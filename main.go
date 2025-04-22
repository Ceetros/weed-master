package main

import (
	"Api/Config"
	"Api/Controller/Swagger"
	"Api/Data/Models"
	"Api/Discord"
	v1 "Api/Routes/v1"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

// @title Vet API
// @version 1.0
// @description API para gerenciamento de clínica veterinária
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config.ConnectDatabase()

	err = Config.DB.AutoMigrate(
		&Models.Analytics{},
	)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	session, _ := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	session.Identify.Intents = discordgo.IntentsGuildMessages

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		println("Bot inicializado!", r.User.ID)
	})

	gg, err := session.Guild(os.Getenv("DISCORD_GUILD"))
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = session.Channel(os.Getenv("DISCORD_NOTIFICATION_CHAT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	bot := Discord.Discord{Bot: session, MainGuild: gg, NotificationChannel: os.Getenv("DISCORD_NOTIFICATION_CHAT")}

	Discord.SetBot(bot)

	if err != nil {
		log.Fatal(err.Error())
	}
	err = session.Open()
	if err != nil {
		log.Fatalf("could not open session: %s", err)
	}

	r := gin.Default()
	v1.RegisterControllers(r)
	Swagger.RegisterController(r)
	err = r.Run(":" + os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal(err)
		return
	}
}
