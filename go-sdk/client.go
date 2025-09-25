package singlebase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	apiKey   string
	apiUrl   string
	headers  map[string]string
}

const BaseApiUrl = "https://cloud.singlebaseapis.com/api"

func NewClient(apiKey string, apiUrl string, endpointKey string, headers map[string]string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("MISSING_API_KEY")
	}
	if apiUrl == "" && endpointKey == "" {
		return nil, fmt.Errorf("MISSING_ENDPOINT_KEY")
	}
	if headers == nil {
		headers = map[string]string{}
	}
	finalUrl := apiUrl
	if finalUrl == "" {
		finalUrl = fmt.Sprintf("%s/%s", BaseApiUrl, endpointKey)
	}

	return &Client{apiKey: apiKey, apiUrl: finalUrl, headers: headers}, nil
}

func (c *Client) validatePayload(payload map[string]any) error {
	if _, ok := payload["op"].(string); !ok {
		return fmt.Errorf("INVALID_PAYLOAD: missing 'op'")
	}
	return nil
}

// Dispatch a request synchronously.
func (c *Client) Dispatch(payload map[string]any, headers map[string]string, bearerToken string) *Result {
	if err := c.validatePayload(payload); err != nil {
		return ResultError(err.Error(), 400)
	}

	allHeaders := map[string]string{}
	for k, v := range c.headers {
		allHeaders[k] = v
	}
	for k, v := range headers {
		allHeaders[k] = v
	}
	allHeaders["x-api-key"] = c.apiKey
	allHeaders["x-sbc-sdk-client"] = "singlebase-go"
	allHeaders["Content-Type"] = "application/json"
	if bearerToken != "" {
		allHeaders["Authorization"] = "Bearer " + bearerToken
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", c.apiUrl, bytes.NewBuffer(body))
	if err != nil {
		return ResultError(err.Error(), 500)
	}
	for k, v := range allHeaders {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ResultError(err.Error(), 500)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var parsed map[string]any
	json.Unmarshal(respBody, &parsed)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		data, _ := parsed["data"].(map[string]any)
		meta, _ := parsed["meta"].(map[string]any)
		return ResultOK(data, meta, resp.StatusCode)
	}
	return ResultError(fmt.Sprint(parsed["error"]), resp.StatusCode)
}
