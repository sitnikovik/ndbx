#!/bin/bash
set -e

VERSION="$1"

if [ -z "$VERSION" ]; then
  echo "❌ Version is required"
  exit 1
fi

git config user.name "github-actions[bot]"
git config user.email "github-actions[bot]@users.noreply.github.com"
git tag -a "v$VERSION" -m "Release v$VERSION"
git push origin "v$VERSION"

echo "✅ Tag v$VERSION created and pushed"
