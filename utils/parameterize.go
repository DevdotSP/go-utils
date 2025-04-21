package utils

import (
	"fmt"
	"github.com/DevdotSP/go-utils/bloc"  // Correct import path
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

// GetGCSPATH fetches the 'api' value from Oasis parameters
func GetGCSPATH(db *gorm.DB, description string) (string, error) {
	result, err := bloc.FetchParam(db, "oasis.system_parameters", "description", description, []string{"api"})
	if err != nil {
		return "", err
	}

	api, err := GetParamValue(result, "api")
	if err != nil {
		return "", err
	}

	return api, nil
}

// GetAllAPI fetches the 'api' value from the API table
func GetAllAPI(db *gorm.DB, code string) (string, error) {
	result, err := bloc.FetchParam(db, "rose_v3.rose_api", "code", code, []string{"api"})
	if err != nil {
		return "", err
	}

	api, err := GetParamValue(result, "api")
	if err != nil {
		return "", err
	}

	return api, nil
}

// GetSystemParam fetches 'key' and 'value' from the system parameters table
func GetSystemParam(db *gorm.DB, code string) (string, string, error) {
	result, err := bloc.FetchParam(db, "rose_v3.system_parameter", "code", code, []string{"key", "value"})
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
