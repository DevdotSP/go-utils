package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ApiRequest struct stores request details
type ApiRequest struct {
	method  string
	url     string
	headers map[string]string
	payload map[string]interface{}
}

// NewApiRequest initializes a new request with default values
func NewApiRequest() *ApiRequest {
	return &ApiRequest{
		method:  "GET",
		headers: map[string]string{"Content-Type": "application/json"},
		payload: make(map[string]interface{}),
	}
}

// Method sets the HTTP method
func (r *ApiRequest) Method(method string) *ApiRequest {
	r.method = method
	return r
}

// URL sets the request URL
func (r *ApiRequest) URL(url string) *ApiRequest {
	r.url = url
	return r
}

// Payload sets the request body, supporting both maps and structs
func (r *ApiRequest) Payload(data interface{}) *ApiRequest {
	switch v := data.(type) {
	case map[string]interface{}:
		r.payload = v
	default:
		jsonData, err := json.Marshal(v)
		if err == nil {
			json.Unmarshal(jsonData, &r.payload)
		} else {
			fmt.Println("Invalid payload format:", err)
		}
	}
	return r
}

// Headers adds custom headers to the request
func (r *ApiRequest) Headers(headers map[string]string) *ApiRequest {
	for key, value := range headers {
		r.headers[key] = value
	}
	return r
}

// Token adds an Authorization Bearer token
func (r *ApiRequest) Token(token string) *ApiRequest {
	r.headers["Authorization"] = "Bearer " + token
	return r
}

// Send executes the HTTP request and returns the response as a map
func (r *ApiRequest) Send() (map[string]interface{}, error) {
	client := &http.Client{}
	var reqBody io.Reader

	// Convert payload to JSON
	if len(r.payload) > 0 {
		jsonData, err := json.Marshal(r.payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	// Create HTTP request
	req, err := http.NewRequest(r.method, r.url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return result, nil
}
