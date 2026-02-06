#!/bin/bash
set -e

VERSION="$1"
NOTES_FILE="$2"

if [ -z "$VERSION" ] || [ -z "$NOTES_FILE" ]; then
  echo "❌ Usage: $0 <version> <notes_file>"
  exit 1
fi

if [ ! -f "$NOTES_FILE" ]; then
  echo "❌ Release notes file not found: $NOTES_FILE"
  exit 1
fi

gh release create "v$VERSION" \
  --title "v$VERSION" \
  --notes-file "$NOTES_FILE"

echo "✅ GitHub Release v$VERSION created"
