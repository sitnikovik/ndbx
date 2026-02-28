#!/usr/bin/env bash
#
# This script calculates coverage percentage from a coverage file.
# Usage: calculate_coverage.sh <coverage_file> <output_file> <label>

set -e

coverage_file=$1
output_file=$2
label=$3

if [ -z "$coverage_file" ] || [ -z "$output_file" ] || [ -z "$label" ]; then
	echo "Usage: $0 <coverage_file> <output_file> <label>" >&2
	exit 1
fi

# Filter out main packages and calculate coverage percentage.
filtered_file="${coverage_file}.filtered"
head -n1 "${coverage_file}" > "${filtered_file}"
grep -v '/cmd/' "${coverage_file}" | tail -n +2 >> "${filtered_file}"

cd autograder
percent=$(go tool cover -func="../${filtered_file}" 2>/dev/null \
	| tail -n1 \
	| awk '{print $NF}')
cd ..

rm -f "${filtered_file}"

percent_no_pct=${percent%\%}
printf '%s' "$percent_no_pct" > "$output_file"
echo "${label}: $percent"
