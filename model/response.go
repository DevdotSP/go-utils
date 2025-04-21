package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type (
	Response struct {
		ResponseTime     string      `json:"responseTime"`
		Device           string      `json:"device"`
		RetCode          string      `json:"retCode"`
		Message          string      `json:"message"`
		ValidationErrors interface{} `json:"validationErrors,omitempty"` // Add this field
		Data             interface{} `json:"data,omitempty"`
		Error            interface{} `json:"error,omitempty"`
		Page             int         `json:"currentPage,omitempty"`
		PageSize         int         `json:"pageSize,omitempty"`
		TotalItem        int         `json:"totalItem,omitempty"`
		TotalPages       int         `json:"totalPages,omitempty"`
		JwtToken         string      `json:"jwt_token,omitempty"`
	}

	ResponsePageDetails struct {
		ResponseTime string      `json:"responseTime"`
		Device       string      `json:"device"`
		RetCode      string      `json:"retCode"`
		Message      string      `json:"message"`
		Data         interface{} `json:"data,omitempty"`
		Error        interface{} `json:"error,omitempty"`
		PageDetails  PageDetails `json:"pageDetails"`
	}

	EPResponse struct {
		ProcessTime string      `json:"processTime"`
		Response    interface{} `json:"response"`
	}

	UserClaims struct {
		jwt.RegisteredClaims
	}

	PageDetails struct {
		Page       int `json:"currentPage"`
		PageSize   int `json:"pageSize"`
		TotalItem  int `json:"totalItem"`
		TotalPages int `json:"totalPages"`
	}
)
