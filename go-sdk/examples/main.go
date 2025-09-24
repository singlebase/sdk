package main

import (
	"fmt"
	"github.com/yourname/singlebase-go"
)

func main() {
	client, err := singlebase.NewClient("my-api-key", "", "vector-db", nil)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"op": "ping",
	}
	result := client.Dispatch(payload, nil, "")
	if result.Ok {
		fmt.Println("✅ Success:", result.Data)
	} else {
		fmt.Println("❌ Error:", result.Error)
	}

	// File upload (example)
	ok, err := singlebase.UploadPresignedFile("test.txt", map[string]interface{}{
		"url": "https://s3.amazonaws.com/bucket",
		"fields": map[string]interface{}{
			"key": "uploads/test.txt",
		},
	})
	if err != nil {
		fmt.Println("Upload failed:", err)
	} else {
		fmt.Println("Upload success:", ok)
	}
}
