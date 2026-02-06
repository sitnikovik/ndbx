#!/bin/bash
set -e

VERSION="$1"

if [ -z "$VERSION" ]; then
  echo "❌ Version is required"
  exit 1
fi

if git rev-parse "v$VERSION" >/dev/null 2>&1; then
  echo "❌ Tag v$VERSION already exists"
  exit 1
fi

echo "✅ Tag v$VERSION does not exist"
