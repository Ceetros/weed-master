package Auth

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var FirebaseAuth *auth.Client

func InitFirebase() {

	opt := option.WithCredentialsFile("firebase.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Erro ao inicializar Firebase: " + err.Error())
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		panic("Erro ao iniciar Firebase Auth: " + err.Error())
	}

	FirebaseApp = app
	FirebaseAuth = client
}
