package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	utils "github.com/DevdotSP/go-utils/utils" // Update with your actual repo path
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
		utils.GetEnv("DB_PASSWORD", "password"),
		utils.GetEnv("DB_NAME", "portfolio"),
		utils.GetEnv("DB_PORT", "5432"),
	)

	pgxDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		utils.GetEnv("DB_USER", "postgres"),
		utils.GetEnv("DB_PASSWORD", "password"),
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_PORT", "5432"),
		utils.GetEnv("DB_NAME", "portfolio"),
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


// FetchParam fetches parameters dynamically from a given database connection.
func FetchParam(db *gorm.DB, tableName, columnName, columnValue string, columnToSelect []string) (map[string]interface{}, error) {
	var results []map[string]interface{}

	err := db.Table(tableName).
		Where(columnName+" = ?", columnValue).
		Select(columnToSelect).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results found")
	}

	return results[0], nil
}