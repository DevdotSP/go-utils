package bloc

import (
	"fmt"
	"gorm.io/gorm"
)

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
