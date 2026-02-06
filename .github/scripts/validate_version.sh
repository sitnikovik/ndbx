#!/bin/bash
set -e

VERSION="$1"

if [ -z "$VERSION" ]; then
  echo "❌ Version is required"
  exit 1
fi

if ! echo "$VERSION" | grep -qE '^[0-9]+\.[0-9]+\.[0-9]+$'; then
  echo "❌ Invalid version format. Use: X.Y.Z (e.g., 1.0.0)"
  exit 1
fi

echo "✅ Version $VERSION is valid"
