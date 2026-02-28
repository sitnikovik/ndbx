#!/usr/bin/env bash
#
# This script combines unit and integration coverage profiles and calculates total coverage.
# It generates:
# - tmp/coverage.out - combined coverage profile
# - tmp/coverage_unit_percent.out - unit test coverage percentage (without % sign)
# - tmp/coverage_integration_percent.out - integration test coverage percentage (without % sign)
# - tmp/coverage_total.out - total coverage percentage (without % sign)
# - tmp/uncovered.out - list of uncovered lines

set -e

# Check that required coverage files exist
if [ ! -f "tmp/coverage_unit.out" ]; then
	echo "Error: tmp/coverage_unit.out not found" >&2
	exit 1
fi

if [ ! -f "tmp/coverage_integration.out" ]; then
	echo "Error: tmp/coverage_integration.out not found" >&2
	exit 1
fi

# Calculate unit test coverage
./scripts/coverage/calc.sh \
	"tmp/coverage_unit.out" \
	"tmp/coverage_unit_percent.out" \
	"Unit test coverage"

# Calculate integration test coverage
./scripts/coverage/calc.sh \
	"tmp/coverage_integration.out" \
	"tmp/coverage_integration_percent.out" \
	"Integration test coverage"

# Combine coverage files (excluding main packages)
echo "mode: count" > tmp/coverage.out
tail -n +2 tmp/coverage_unit.out | grep -v '/cmd/' >> tmp/coverage.out
tail -n +2 tmp/coverage_integration.out | grep -v '/cmd/' >> tmp/coverage.out

# Calculate total coverage percentage
./scripts/coverage/calc.sh \
	"tmp/coverage.out" \
	"tmp/coverage_total.out" \
	"Total test coverage"

# Generate list of uncovered lines (excluding main packages)
awk 'NR>1 {
	n=NF;
	if($n==0){
		split($1,a,":");
		file=a[1];
		split(a[2],b,",");
		split(b[1],c,"\\.");
		start=c[1];
		split(b[2],d,"\\.");
		end=d[1];
		print file ":" start "-" end
	}
}' tmp/coverage.out \
	| sed 's#^github.com/sitnikovik/ndbx/##' \
	| sort -u > tmp/uncovered.out
