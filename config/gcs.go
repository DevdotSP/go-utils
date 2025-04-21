package config

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/DevdotSP/go-utils/utils"
	"google.golang.org/api/option"
)

var (
	GoogleCloudStorageClient *storage.Client
	once                     sync.Once
)

func InitGoogleCloud() *storage.Client {
	once.Do(func() {
		credentialsPath := utils.GetEnv("GOOGLE_CLOUD_CREDENTIALS", "config/credentials/googlecloud.json")
		client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(credentialsPath))
		if err != nil {
			log.Fatalf("❌ Failed to initialize Google Cloud Storage: %v", err)
		}
		GoogleCloudStorageClient = client
		log.Println("✅ Google Cloud Storage client initialized")
	})
	return GoogleCloudStorageClient
}
