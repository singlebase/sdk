package singlebase

import "fmt"

// Result represents the result of an API operation.
type Result struct {
	Data       map[string]interface{} `json:"data"`
	Meta       map[string]interface{} `json:"meta"`
	Ok         bool                   `json:"ok"`
	Error      string                 `json:"error,omitempty"`
	StatusCode int                    `json:"statusCode"`
}

// ToMap converts the Result to a map[string]interface{}.
func (r *Result) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data":       r.Data,
		"meta":       r.Meta,
		"ok":         r.Ok,
		"error":      r.Error,
		"statusCode": r.StatusCode,
	}
}

// String implements fmt.Stringer for debugging.
func (r *Result) String() string {
	return fmt.Sprintf("<Result ok=%t status=%d error=%s>", r.Ok, r.StatusCode, r.Error)
}

// ResultOK returns a success result.
func ResultOK(data map[string]interface{}, meta map[string]interface{}, status int) *Result {
	return &Result{Data: data, Meta: meta, Ok: true, StatusCode: status}
}

// ResultError returns an error result.
func ResultError(err string, status int) *Result {
	return &Result{Ok: false, Error: err, StatusCode: status}
}
