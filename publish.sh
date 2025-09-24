#!/usr/bin/env bash
set -euo pipefail

# Root script for publishing all SDKs
# Usage:
#   ./publish.sh python  # publish only Python SDK
#   ./publish.sh js      # publish only JS SDK
#   ./publish.sh php     # trigger Packagist update
#   ./publish.sh go      # push Go tag
#   ./publish.sh all     # publish all SDKs

publish_python() {
  echo "🐍 Publishing Python SDK..."
  cd python-sdk
  rm -rf dist build
  python -m pip install --upgrade build
  python -m build
  twine upload dist/*  # requires ~/.pypirc with PyPI token
  cd ..
}
  

publish_js() {
  echo "🟦 Publishing JS/TS SDK..."
  cd js-sdk
  npm version patch   # bump version automatically (patch release)
  npm publish --access public
  cd ..
}

publish_php() {
  echo "🐘 Publishing PHP SDK..."
  echo "PHP SDKs are distributed via Packagist."
  echo "Ensure composer.json is updated and repo is pushed to GitHub."
  echo "If Packagist webhook is configured, nothing more to do."
}

publish_go() {
  echo "🐹 Publishing Go SDK..."
  cd go-sdk
  VERSION="v$(date +'%Y.%m.%d.%H%M')"
  git tag $VERSION
  git push origin $VERSION
  cd ..
  echo "Tagged Go module as $VERSION"
}

publish_all() {
  publish_python
  publish_js
  publish_php
  publish_go
}

case "${1:-}" in
  python) publish_python ;;
  js) publish_js ;;
  php) publish_php ;;
  go) publish_go ;;
  all) publish_all ;;
  *)
    echo "Usage: $0 {python|js|php|go|all}"
    exit 1
    ;;
esac
