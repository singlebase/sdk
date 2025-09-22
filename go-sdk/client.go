package singlebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// Client for Singlebase API requests.
type Client struct {
	APIKey  string
	APIUrl  string
	Headers map[string]string
}

// Base API URL for Singlebase Cloud
const BaseAPIURL = "https://cloud.singlebaseapis.com/api"

// NewClient initializes a new Client.
// apiKey: required API key.
// apiUrl: full URL or empty if using endpointKey.
// endpointKey: appended to base URL if apiUrl is empty.
func NewClient(apiKey string, apiUrl string, endpointKey string, headers map[string]string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("MISSING_API_KEY")
	}
	if apiUrl == "" && endpointKey == "" {
		return nil, errors.New("MISSING_ENDPOINT_KEY")
	}
	if apiUrl == "" {
		apiUrl = BaseAPIURL + "/" + endpointKey
	}
	return &Client{APIKey: apiKey, APIUrl: apiUrl, Headers: headers}, nil
}

// Call makes a synchronous request to the API.
func (c *Client) Call(payload map[string]interface{}, headers map[string]string, bearerToken string) *Result {
	// Validate payload
	op, ok := payload["op"]
	if !ok || op == "" {
		return ResultError("INVALID_PAYLOAD: missing 'op'", 400)
	}

	// Build headers
	reqHeaders := map[string]string{
		"x-api-key":        c.APIKey,
		"x-sbc-sdk-client": "singlebase-go",
		"Content-Type":     "application/json",
	}
	for k, v := range c.Headers {
		reqHeaders[k] = v
	}
	for k, v := range headers {
		reqHeaders[k] = v
	}
	if bearerToken != "" {
		reqHeaders["Authorization"] = "Bearer " + bearerToken
	}

	// Encode payload
	body, err := json.Marshal(payload)
	if err != nil {
		return ResultError("EXCEPTION: "+err.Error(), 500)
	}

	// Build request
	req, err := http.NewRequest("POST", c.APIUrl, bytes.NewBuffer(body))
	if err != nil {
		return ResultError("EXCEPTION: "+err.Error(), 500)
	}
	for k, v := range reqHeaders {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return ResultError("EXCEPTION: "+err.Error(), 500)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var parsed map[string]interface{}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return ResultError("EXCEPTION: "+err.Error(), 500)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		data, _ := parsed["data"].(map[string]interface{})
		meta, _ := parsed["meta"].(map[string]interface{})
		return ResultOK(data, meta, resp.StatusCode)
	}
	errMsg, _ := parsed["error"].(string)
	if errMsg == "" {
		errMsg = "Unknown Error"
	}
	return ResultError(errMsg, resp.StatusCode)
}
