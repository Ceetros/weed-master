package main

import (
	"Api/Auth"
	"Api/Config"
	"Api/Controller"
	"Api/Data/Models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

// @title Vet API
// @version 1.0
// @description API para gerenciamento de clínica veterinária
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config.ConnectDatabase()

	Config.DB.AutoMigrate(
		&Models.User{},
		&Models.Guardian{},
		&Models.Patient{},
		&Models.Clinical{},
		&Models.Address{},
		&Models.ClinicalUser{},
		&Models.Treatment{},
		&Models.TreatmentNote{},
		&Models.Consultation{},
		&Models.Exam{},
		&Models.Drug{},
	)

	Auth.InitFirebase()

	r := gin.Default()
	Controller.RegisterControllers(r)
	r.Run(":8080")
}
