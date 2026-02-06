#!/bin/bash
set -e

# Load environment variables from a specified .env file
# Usage: ./load_env.sh <env_file>
# Example: ./load_env.sh path/to/.env

if [ $# -ne 1 ]; then
    echo "❌ Usage: $0 <env_file>"
    exit 1
fi

ENV_FILE="$1"
if [ ! -f "$ENV_FILE" ]; then
    echo "❌ $ENV_FILE not found"
    exit 1
fi

echo "Loading environment from $ENV_FILE"
set -a
source "$ENV_FILE"
set +a
cat "$ENV_FILE" | grep -v '^#' | grep -v '^$' >> $GITHUB_ENV