package utils

import (
	"fmt"

	"github.com/DevdotSP/go-utils/fetchparam"
	"gorm.io/gorm"
)

// GetParamValue is a helper function to fetch a single value from a parameter result.
func GetParamValue(result map[string]interface{}, key string) (string, error) {
	value, ok := result[key].(string)
	if !ok {
		return "", fmt.Errorf("%s not found or invalid type", key)
	}
	return value, nil
}

//example of columnToSelect []string{"api"} 

// GetGCSPATH fetches the 'api' value from Oasis parameters
func GetGCSPATH(db *gorm.DB, tableName, columnName, columnValue string, columnToSelect []string) (string, error) {
	result, err := fetchparam.FetchParam(db, tableName,  columnName, columnValue, columnToSelect)
	if err != nil {
		return "", err
	}

	api, err := GetParamValue(result, "api")
	if err != nil {
		return "", err
	}

	return api, nil
}

//example of columnToSelect []string{"api"} 

// GetAllAPI fetches the 'api' value from the API table
func GetAllAPI(db *gorm.DB, tableName, columnName, columnValue string, columnToSelect []string) (string, error) {
	result, err := fetchparam.FetchParam(db, tableName, columnName, columnValue, columnToSelect)
	if err != nil {
		return "", err
	}

	api, err := GetParamValue(result, "api")
	if err != nil {
		return "", err
	}

	return api, nil
}

//example of columnToSelect []string{"key", "value"} 

// GetSystemParam fetches 'key' and 'value' from the system parameters table
func GetSystemParam(db *gorm.DB, tableName, columnName, columnValue string, columnToSelect []string) (string, string, error) {
	result, err := fetchparam.FetchParam(db, tableName, columnName, columnValue, columnToSelect )
	if err != nil {
		return "", "", err
	}

	key, err := GetParamValue(result, "key")
	if err != nil {
		return "", "", err
	}

	value, err := GetParamValue(result, "value")
	if err != nil {
		return "", "", err
	}

	return key, value, nil
}
