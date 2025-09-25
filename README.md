# Singlebase SDKs

Official SDKs for interacting with the [Singlebase](https://singlebase.cloud) API.  

Available in:
- **Python**
- **JavaScript/TypeScript**
- **PHP**
- **Go**

These SDKs provide:

- Simple `Client` interface for making API requests  
- Standardized `Result` object with `ok`, `data`, `meta`, `error`, `statusCode`  
- Dot-notation access with `get_data` / `getData` / `GetData`  
- Presigned file upload helpers (Python, JS, PHP, Go)  

---

## Python SDK (`singlebase`)

### Install
```bash
pip install singlebase
```

### Usage
```python
from singlebase import Client

client = Client(api_key="your-api-key", endpoint_key="endpoint-key")

result = client.dispatch({"op": "ping"})

if result.ok:
    print("Success:", result.data)
else:
    print("Error:", result.error)
```

### Async
```python
import asyncio
from singlebase import Client

async def main():
    client = Client(api_key="your-api-key", endpoint_key="endpoint-key")
    result = await client.dispatch_async({"op": "ping"})
    print(result)

asyncio.run(main())
```

### Result Object
```python
print(result.to_dict())
```

| Attribute | Type | Description |
|-----------|------|-------------|
| `ok` | `bool` | True if request succeeded |
| `data` | `dict` | Response payload |
| `meta` | `dict` | Extra metadata (pagination, etc.) |
| `error` | `str` | Error message if failed |
| `status_code` | `int` | HTTP status code |

### Dot Notation Access
```python
val = result.get_data("address.city.city_fullname", default="N/A")
print(val)
```

---

## JavaScript / TypeScript SDK (`singlebase-js`)

### Install
```bash
npm install singlebase-js
# or
yarn add singlebase-js
```

### Usage
```typescript
import { Client } from "singlebase-js";

const client = new Client({
  apiKey: "your-api-key",
  endpointKey: "endpoint-key"
});

const result = await client.dispatch({ op: "ping" });

if (result.ok) {
  console.log("Success:", result.data);
} else {
  console.error("Error:", result.error);
}
```

### Result Object
```typescript
console.log(result.toObject());
```

| Attribute | Type | Description |
|-----------|------|-------------|
| `ok` | `boolean` | True if request succeeded |
| `data` | `object` | Response payload |
| `meta` | `object` | Extra metadata |
| `error` | `string` | Error message if failed |
| `statusCode` | `number` | HTTP status code |

### Dot Notation Access
```typescript
console.log(result.getData("address.city.city_fullname")); // "San Francisco"
console.log(result.getData("address.country", "USA"));     // "USA"
```

---

## PHP SDK (`singlebase-php`)

### Install
```bash
composer require singlebase/singlebase-php
```

### Usage
```php
<?php
require 'vendor/autoload.php';

use Singlebase\Client;

$client = new Client(apiKey: "your-api-key", endpointKey: "endpoint-key");
$result = $client->dispatch([ "op" => "ping" ]);

if ($result->ok) {
    echo "Success: " . print_r($result->data, true);
} else {
    echo "Error: " . $result->error;
}
```

### Result Object
```php
print_r($result->toArray());
```

| Attribute | Type | Description |
|-----------|------|-------------|
| `ok` | `bool` | True if request succeeded |
| `data` | `array` | Response payload |
| `meta` | `array` | Extra metadata |
| `error` | `string` | Error message if failed |
| `statusCode` | `int` | HTTP status code |

### Dot Notation Access
```php
echo $result->getData("address.city.city_fullname"); // "San Francisco"
echo $result->getData("address.country", "USA");     // "USA"
```

---

## Go SDK (`singlebase-go`)

### Install
```bash
go get github.com/singlebase/sdk/go-sdk@latest
```

### Usage
```go
package main

import (
    "fmt"
    "github.com/singlebase/sdk/go-sdk"
)

func main() {
    client, err := singlebase.NewClient("your-api-key", "", "endpoint-key", nil)
    if err != nil {
        panic(err)
    }

    result := client.Dispatch(map[string]any{"op": "ping"}, nil, "")
    if result.Ok {
        fmt.Println("Success:", result.Data)
    } else {
        fmt.Println("Error:", result.Error)
    }
}
```

### Result Object
```go
fmt.Println(result.ToMap())
```

| Field | Type | Description |
|-------|------|-------------|
| `Ok` | `bool` | True if request succeeded |
| `Data` | `map[string]any` | Response payload |
| `Meta` | `map[string]any` | Extra metadata |
| `Error` | `string` | Error message if failed |
| `StatusCode` | `int` | HTTP status code |

### Dot Notation Access
```go
val, _ := result.GetData("address.city.city_fullname", "N/A")
fmt.Println(val) // "San Francisco"
```

---

## File Uploads (Presigned URL)

All SDKs provide helpers to upload files to a presigned URL (e.g., AWS S3).
Examples differ slightly by language:

### Python
```python
from singlebase import upload_presigned_file
upload_presigned_file("myfile.txt", {"url": "...", "fields": {"key": "uploads/myfile.txt"}})
```

### JavaScript/TypeScript
```typescript
import { uploadPresignedFile } from "singlebase-js";
await uploadPresignedFile("myfile.txt", { url: "...", fields: { key: "uploads/myfile.txt" } });
```

### PHP
```php
use Singlebase\Upload;
Upload::presignedFile("myfile.txt", [ "url" => "...", "fields" => ["key" => "uploads/myfile.txt"] ]);
```

### Go
```go
ok, err := singlebase.UploadPresignedFile("myfile.txt", map[string]any{
    "url": "https://bucket.s3.amazonaws.com",
    "fields": map[string]any{"key": "uploads/myfile.txt"},
})
```

---

## Result Helper Methods

Every SDK implements these:

- `to_dict()` / `toObject()` / `toArray()` / `ToMap()` → serialize to native dict/object/map
- `get_data(path, default)` / `getData(path, default)` / `GetData(path, default)` → dot-notation nested access with default fallback
- `__repr__` / `toString()` / `__toString__` / `String()` → human-readable summary