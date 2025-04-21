package utils

import (
	"log"
	"math"

	"gorm.io/gorm"
)

type PaginatedResult struct {
	CurrentPage int         `json:"currentPage"`
	TotalPages  int         `json:"totalPages"`
	TotalCount  int64       `json:"totalCount"`
	Records     interface{} `json:"records"`
}

// Paginate fetches records dynamically with optional filters & preloads
func Paginate(db *gorm.DB, model interface{}, page, limit int, filters map[string]interface{}, preloads []string) (*PaginatedResult, error) {
	var totalCount int64
	offset := (page - 1) * limit

	// Base query
	query := db.Model(model)

	// Apply filters if provided
	if len(filters) > 0 {
		query = query.Where(filters)
	}

	// Count total records
	if err := query.Count(&totalCount).Error; err != nil {
		log.Printf("Error retrieving total count: %v", err)
		return nil, err
	}

	// Apply preloads if specified
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	// Fetch paginated records
	if err := query.Order("id DESC").Limit(limit).Offset(offset).Find(model).Error; err != nil {
		log.Printf("Error retrieving paginated records: %v", err)
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return &PaginatedResult{
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalCount:  totalCount,
		Records:     model,
	}, nil
}
