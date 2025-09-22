package singlebase

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUploadPresignedFileSuccess(t *testing.T) {
	// Mock server that always returns 204
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, _, err := r.FormFile("file")
		if err != nil {
			t.Errorf("Expected file field, got error: %v", err)
		}
		w.WriteHeader(http.StatusNoContent)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	tmpFile, err := os.CreateTemp("", "upload")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("hello")
	tmpFile.Close()

	ok, err := UploadPresignedFile(tmpFile.Name(), map[string]interface{}{
		"url":    server.URL,
		"fields": map[string]interface{}{"key": "test"},
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !ok {
		t.Errorf("Expected upload success")
	}
}

func TestUploadPresignedFileMissingUrl(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "upload")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = UploadPresignedFile(tmpFile.Name(), map[string]interface{}{
		"fields": map[string]interface{}{"key": "test"},
	})
	if err == nil {
		t.Errorf("Expected error for missing URL")
	}
}

func TestUploadPresignedFileFailStatus(t *testing.T) {
	// Mock server that returns 403
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.Copy(io.Discard, r.Body)
		w.WriteHeader(http.StatusForbidden)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	tmpFile, _ := os.CreateTemp("", "upload")
	defer os.Remove(tmpFile.Name())

	_, err := UploadPresignedFile(tmpFile.Name(), map[string]interface{}{
		"url":    server.URL,
		"fields": map[string]interface{}{},
	})
	if err == nil {
		t.Errorf("Expected error for failed upload")
	}
}
