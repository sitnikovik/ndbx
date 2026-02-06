#!/bin/bash
set -e

# Check required files
# Usage: ./file_exists.sh <file1> <file2> ... <fileN>
# Example: ./file_exists.sh lab/01/Makefile lab/01/docker-compose.yml

if [ $# -eq 0 ]; then
    echo "❌ No files specified to check"
    exit 1
fi

echo "Checking required files..."
for file in "$@"; do
    if [ ! -f "$file" ]; then
        echo "  ❌ $file not found"
        exit 1
    fi
    echo "  ✓ $file"
done
echo "✅ All required files exist"
