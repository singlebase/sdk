package singlebase

import (
	"errors"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// UploadPresignedFile uploads a file using presigned URL data.
//
// filepath: local file path to upload
// data: must include "url" and "fields" (map[string]string)
func UploadPresignedFile(filepathStr string, data map[string]interface{}) (bool, error) {
	url, ok := data["url"].(string)
	if !ok || url == "" {
		return false, errors.New("missing upload URL")
	}
	fields, ok := data["fields"].(map[string]interface{})
	if !ok {
		return false, errors.New("missing fields")
	}

	file, err := os.Open(filepathStr)
	if err != nil {
		return false, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add fields
	for k, v := range fields {
		writer.WriteField(k, v.(string))
	}

	// Add file
	part, err := writer.CreateFormFile("file", filepath.Base(filepathStr))
	if err != nil {
		return false, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return false, err
	}
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true, nil
	}
	return false, errors.New("upload failed with status " + resp.Status)
}
