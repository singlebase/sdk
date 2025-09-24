package singlebase

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientDispatchSuccess(t *testing.T) {
	// Mock API server
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp := map[string]interface{}{
			"data": map[string]interface{}{"msg": "ok"},
			"meta": map[string]interface{}{"page": 1},
		}
		json.NewEncoder(w).Encode(resp)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client, _ := NewClient("abc", server.URL, "", nil)
	result := client.Dispatch(map[string]interface{}{"op": "ping"}, nil, "")

	if !result.Ok {
		t.Errorf("Expected Ok=true, got %v", result.Ok)
	}
	if result.Data["msg"] != "ok" {
		t.Errorf("Expected msg=ok, got %v", result.Data["msg"])
	}
}

func TestClientDispatchError(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		resp := map[string]interface{}{
			"error": "Bad Request",
		}
		json.NewEncoder(w).Encode(resp)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client, _ := NewClient("abc", server.URL, "", nil)
	result := client.Dispatch(map[string]interface{}{"op": "ping"}, nil, "")

	if result.Ok {
		t.Errorf("Expected Ok=false, got %v", result.Ok)
	}
	if result.Error != "Bad Request" {
		t.Errorf("Expected error=Bad Request, got %v", result.Error)
	}
}

func TestClientInvalidPayload(t *testing.T) {
	client, _ := NewClient("abc", "http://localhost", "", nil)
	result := client.Dispatch(map[string]interface{}{}, nil, "")
	if result.Ok {
		t.Errorf("Expected Ok=false for invalid payload")
	}
}
