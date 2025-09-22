package singlebase

import "testing"

func TestResultToMapAndString(t *testing.T) {
	r := &Result{Data: map[string]interface{}{"foo": "bar"}, Ok: true, StatusCode: 201}
	m := r.ToMap()

	if m["data"].(map[string]interface{})["foo"] != "bar" {
		t.Errorf("Expected foo=bar, got %v", m["data"])
	}

	if r.String() == "" {
		t.Errorf("Expected non-empty string repr")
	}
}

func TestResultOkAndError(t *testing.T) {
	ok := ResultOK(map[string]interface{}{"success": true}, nil, 200)
	if !ok.Ok {
		t.Errorf("Expected Ok=true")
	}

	err := ResultError("Something failed", 400)
	if err.Ok {
		t.Errorf("Expected Ok=false")
	}
	if err.StatusCode != 400 {
		t.Errorf("Expected status=400, got %d", err.StatusCode)
	}
}
