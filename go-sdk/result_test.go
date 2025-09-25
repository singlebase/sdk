package singlebase

import (
	"testing"
)

func TestGetDataSuccessAndDefault(t *testing.T) {
	r := ResultOK(map[string]any{
		"address": map[string]any{
			"city": map[string]any{
				"city_fullname": "San Francisco",
				"zipcode":       94107,
			},
		},
	}, nil, 200)

	// Full data
	data, err := r.GetData("", nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if data.(map[string]any)["address"] == nil {
		t.Errorf("Expected address key")
	}

	// Nested path
	val, _ := r.GetData("address.city.city_fullname", nil)
	if val != "San Francisco" {
		t.Errorf("Expected San Francisco, got %v", val)
	}

	// Missing key returns default
	val, _ = r.GetData("address.country", "USA")
	if val != "USA" {
		t.Errorf("Expected default USA, got %v", val)
	}
}

func TestGetDataTypeError(t *testing.T) {
	r := ResultOK(map[string]any{"user": map[string]any{"id": 123}}, nil, 200)
	_, err := r.GetData("user.id.value", nil)
	if err == nil {
		t.Errorf("Expected TypeError, got nil")
	}
}
