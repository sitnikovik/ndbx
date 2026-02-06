#!/bin/bash
set -e

VERSION="$1"
OUTPUT_FILE="$2"

if [ -z "$VERSION" ] || [ -z "$OUTPUT_FILE" ]; then
  echo "❌ Usage: $0 <version> <output_file>"
  exit 1
fi

# Extract changelog section for this version
CHANGELOG=$(sed -n "/## \[$VERSION\]/,/## \[/p" CHANGELOG.md | sed '$d' | tail -n +2)

if [ -z "$CHANGELOG" ]; then
  echo "❌ Version $VERSION not found in CHANGELOG.md"
  exit 1
fi

echo "$CHANGELOG" > "$OUTPUT_FILE"
echo "✅ Release notes extracted for v$VERSION"
