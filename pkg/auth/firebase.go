package auth

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App

var (
    ClientID = os.Getenv("CLIENT_ID")
    clientSecret = os.Getenv("CLIENT_SECRET")
    redirectURI  = "http://localhost:8000/auth/google-callback"
)

// Initialize Firebase
func InitFirebase() {
    opt := option.WithCredentialsFile("serviceAccountKey.json")
    config := &firebase.Config{ProjectID: "thefatbean-c667b"}
    app, err := firebase.NewApp(context.Background(), config, opt)
    if err != nil {
        log.Fatalf("error initializing Firebase: %v", err)
    }
    App = app
    log.Println("Firebase initialized successfully")
}

