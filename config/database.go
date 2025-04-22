package config

import (
	"context"
	"fmt"
	"log"

	"github.com/DevdotSP/go-utils/utils" // Update with your actual repo path
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
		utils.GetEnv("DB_PASSWORD", "@Light1114"),
		utils.GetEnv("DB_NAME", "portfolio"),
		utils.GetEnv("DB_PORT", "5432"),
	)

	pgxDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		utils.GetEnv("DB_USER", "postgres"),
		utils.GetEnv("DB_PASSWORD", "@Light1114"),
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

// UpdateRecord updates a record in the database using a transaction (tx).
func UpdateRecordTX[T any](tx *gorm.DB, model *T, column any, value any, updates map[string]interface{}) error {
	var valueStr string

	// Convert value to string or int for the query
	switch v := value.(type) {
	case string:
		valueStr = v
	case int:
		valueStr = fmt.Sprintf("%d", v)
	default:
		return fmt.Errorf("unsupported value type: %T", value)
	}

	// Convert column to string and ensure proper formatting for SQL query
	columnStr, ok := column.(string)
	if !ok {
		return fmt.Errorf("column must be a string, got %T", column)
	}

	// Attempt to find the record first
	result := tx.Model(model).Where(fmt.Sprintf("%s = ?", columnStr), valueStr).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("failed to update record: %w", result.Error)
	}

	// If no records were updated, insert a new one
	if result.RowsAffected == 0 {
		if err := tx.Create(model).Error; err != nil {
			return fmt.Errorf("failed to insert new record: %w", err)
		}
	}

	return nil
}

// UpdateRecord updates a record in the database based on a given condition.
func UpdateRecord[T any](model *T, column any, value any, updates map[string]interface{}) error {
	var valueStr string

	// Convert value to string or int for the query
	switch v := value.(type) {
	case string:
		valueStr = v
	case int:
		valueStr = fmt.Sprintf("%d", v)
	default:
		return fmt.Errorf("unsupported value type: %T", value)
	}

	// Convert column to string and ensure proper formatting for SQL query
	columnStr, ok := column.(string)
	if !ok {
		return fmt.Errorf("column must be a string, got %T", column)
	}

	// Perform the update query
	result := DB.Model(model).Where(fmt.Sprintf("%s = ?", columnStr), valueStr).Updates(updates)

	// If no rows were affected, insert a new record
	if result.RowsAffected == 0 {
		// Insert new record
		if err := DB.Create(model).Error; err != nil {
			return fmt.Errorf("failed to insert new record: %w", err)
		}
	}

	return result.Error
}
