package main

import (
	"Api/Auth"
	"Api/Config"
	"Api/Controller/Swagger"
	"Api/Data/Models"
	v1 "Api/Routes/v1"
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

	err = Config.DB.AutoMigrate(
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
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	Auth.InitFirebase()

	r := gin.Default()
	v1.RegisterControllers(r)
	Swagger.RegisterController(r)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
