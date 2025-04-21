package config

import (
	"context"
	"log"
	"sync"

	firebase "firebase.google.com/go/v4"
	"github.com/DevdotSP/go-utils/utils"
	"google.golang.org/api/option"
)

var (
	FirebaseApp  *firebase.App
	initFirebase sync.Once
)

func InitFirebase() *firebase.App {
	initFirebase.Do(func() {
		credentialsPath := utils.GetEnv("FIREBASE_CREDENTIALS", "config/credentials/firebase.json")

		opt := option.WithCredentialsFile(credentialsPath)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("❌ Firebase init error: %v", err)
		}
		FirebaseApp = app
		log.Println("✅ Firebase initialized")
	})
	return FirebaseApp
}
