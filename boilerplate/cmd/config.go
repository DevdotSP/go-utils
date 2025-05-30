package command

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateConfig(component string) {
	if component == "" {
		fmt.Println("❌ Usage: go run tool.go config [database|firebase|gcloud]")
		os.Exit(1)
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("❌ Failed to get current directory: %v\n", err)
		os.Exit(1)
	}

	// Build base path: {cwd}/package/config/
	basePath := filepath.Join(cwd, "package", "config")

	// Ensure directory exists
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		fmt.Printf("❌ Failed to create config folder: %v\n", err)
		os.Exit(1)
	}

	var filePath string
	var content string

	// Entry point for switch statement
	switch component {
	case "database":
		filePath = filepath.Join(basePath, "database.go")
		content = databaseConfigBoilerplate()
	case "firebase":
		filePath = filepath.Join(basePath, "firebase.go")
		content = firebaseConfigBoilerplate()
	case "gcloud":
		filePath = filepath.Join(basePath, "gcloud.go")
		content = googleCloudConfigBoilerplate()
	default:
		fmt.Printf("❌ Unknown component: %s\n", component)
		os.Exit(1)
	}

	// Check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("⚠️  %s already exists. Skipping...\n", filePath)
		return
	}

	// Write file
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		fmt.Printf("❌ Failed to write file: %v\n", err)
		return
	}

	fmt.Printf("✅ %s generated successfully\n", filePath)
}


func databaseConfigBoilerplate() string {
	return `package config

import (
	"context"
	"fmt"
	"log"

	"github.com/DevdotSP/go-utils/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	PgxPool *pgxpool.Pool
)

func PostgreSQLConnect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USER", "postgres"),
		utils.GetEnv("DB_PASSWORD", ""),
		utils.GetEnv("DB_NAME", "postgres"),
		utils.GetEnv("DB_PORT", "5432"),
	)

	pgxDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		utils.GetEnv("DB_USER", "postgres"),
		utils.GetEnv("DB_PASSWORD", ""),
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_PORT", "5432"),
		utils.GetEnv("DB_NAME", "postgres"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	log.Println("✅ Database connected")

	PgxPool, err = pgxpool.New(context.Background(), pgxDSN)
	if err != nil {
		log.Fatalf("❌ Failed to create pgx connection pool: %v", err)
	}
	log.Println("✅ pgx connection pool initialized")
}
`
}

func firebaseConfigBoilerplate() string {
	return `package config

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
`
}

func googleCloudConfigBoilerplate() string {
	return `package config

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
`
}
