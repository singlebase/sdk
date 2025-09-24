# Singlebase SDKs

Official SDKs for interacting with the Singlebase API. Available in Python, JavaScript/TypeScript, PHP, and Go.

---

## Python SDK (singlebase)

### Install

`pip install singlebase`

### Example

```
from singlebase import Client

client = Client(api_key="your-api-key", endpoint_key="vector-db")
result = client.dispatch({"op": "ping"})

if result.ok:
    print("✅ Success:", result.data)
else:
    print("❌ Error:", result.error)
```

---

## JavaScript / TypeScript SDK (singlebase-js)

### Install

```
npm install @singlebase/singlebase-js
# or
yarn add @singlebase/singlebase-js
```

### Example

`import { Client } from "@singlebase/singlebase-js";`

```
const client = new Client({
  apiKey: "your-api-key",
  endpointKey: "vector-db",
});

const result = await client.dispatch({ op: "ping" });

if (result.ok) {
  console.log("✅ Success:", result.data);
} else {
  console.error("❌ Error:", result.error);
}
```

---

## PHP SDK (singlebase/singlebase-php)

### Install

`composer require singlebase/singlebase-php`

### Example

```
<?php

require 'vendor/autoload.php';

use Singlebase\Client;

$client = new Client(apiKey: "your-api-key", endpointKey: "vector-db");
$result = $client->dispatch([ "op" => "ping" ]);

if ($result->ok) {
    echo "✅ Success: " . print_r($result->data, true);
} else {
    echo "❌ Error: " . $result->error;
}

```

---

## Go SDK (singlebase-go)

### Install

`go get github.com/singlebase/singlebase-go@latest`

### Example

```
package main

import (
	"fmt"
	"github.com/you/singlebase-go"
)

func main() {
	client, err := singlebase.NewClient("your-api-key", "", "vector-db", nil)
	if err != nil {
		panic(err)
	}

	result := client.Dispatch(map[string]interface{}{"op": "ping"}, nil, "")
	if result.Ok {
		fmt.Println("✅ Success:", result.Data)
	} else {
		fmt.Println("❌ Error:", result.Error)
	}
}
```

📦 Features (all SDKs)

✅ Simple Client for API dispatchs

✅ Consistent Result / ResultOK / ResultError types

✅ Support for synchronous & asynchronous requests (Python/JS)

✅ Presigned file uploads helpers (Python, JS, PHP, Go)

✅ Built-in error handling

🤝 Contributing

Fork this repo and open a PR 🚀

Run tests before submitting (pytest, npm test, phpunit, go test ./...)

---

### CI/CD (per language)

Inside .github/workflows/, we can have:

python.yml → runs pytest, publishes to PyPI on release tag

node.yml → runs npm test, publishes to npm

php.yml → runs phpunit, auto-updates Packagist (webhook or action)

go.yml → runs go test ./..., release on git tag

Each can be triggered separately when tagging a release like python-v0.1.0, js-v0.1.0, etc.

### Publishing

Python → from /python using pyproject.toml

JS → from /js using npm publish

PHP → from /php using Packagist + Composer

Go → from /go (git tag)

## License

MIT 

© 2025++ Singlebase