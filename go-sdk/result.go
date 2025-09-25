package singlebase

import (
	"fmt"
	"strings"
)

// Result represents the result of an API operation.
type Result struct {
	Data       map[string]any `json:"data"`
	Meta       map[string]any `json:"meta"`
	Ok         bool           `json:"ok"`
	Error      string         `json:"error,omitempty"`
	StatusCode int            `json:"status_code"`
}

// ToMap converts the result to a map.
func (r *Result) ToMap() map[string]any {
	return map[string]any{
		"data":        r.Data,
		"meta":        r.Meta,
		"ok":          r.Ok,
		"error":       r.Error,
		"status_code": r.StatusCode,
	}
}

// String returns a string representation of the result.
func (r *Result) String() string {
	return fmt.Sprintf("<Result ok=%t status=%d error=%q>", r.Ok, r.StatusCode, r.Error)
}

// GetData retrieves a value from Data using dot-notation path.
// If no path provided, returns full Data.
// If a key is missing, returns defaultVal.
// Returns error if traversal hits a non-map type.
func (r *Result) GetData(path string, defaultVal any) (any, error) {
	if path == "" {
		return r.Data, nil
	}

	current := any(r.Data)
	for _, part := range strings.Split(path, ".") {
		m, ok := current.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("cannot traverse %q — expected map, got %T", part, current)
		}
		val, exists := m[part]
		if !exists {
			return defaultVal, nil
		}
		current = val
	}
	return current, nil
}

// ResultOK represents a successful operation.
func ResultOK(data, meta map[string]any, status int) *Result {
	return &Result{
		Data:       data,
		Meta:       meta,
		Ok:         true,
		StatusCode: status,
	}
}

// ResultError represents a failed operation.
func ResultError(errorMsg string, status int) *Result {
	return &Result{
		Data:       map[string]any{},
		Meta:       map[string]any{},
		Ok:         false,
		Error:      errorMsg,
		StatusCode: status,
	}
}
