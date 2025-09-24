

# Singlebase Python SDK

Python SDK for interacting with Singlebase APIs.

## Installation

```bash
pip install singlebase
```

## Usage

```
from singlebase import Client

client = Client(api_key="your-api-key", endpoint_key="endpoint-key")

result = client.dispatch({"op": "ping"})

if result.ok:
    print("Success:", result.data)
else:
    print("Error:", result.error)

```

See [root README](../README.md) for full installation and usage instructions.

---


## Local Build & Install Locally
```bash
# Build package
python -m build

# Install locally
pip install dist/singlebase-0.1.0-py3-none-any.whl
```

```
pip install -e ".[dev]"
```