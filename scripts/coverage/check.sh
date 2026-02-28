#!/usr/bin/env bash
#
# This script checks if the code coverage percentage meets a specified threshold.
# It requires the COVERAGE_THRESHOLD environment variable to be set.
# Usage: coverage_check.sh [path_to_coverage_percent_file]

PERCENT_FILE=${1:-tmp/coverage_total.out}
THRESHOLD=${COVERAGE_THRESHOLD}

# Ensure threshold is specified.
if [ -z "$THRESHOLD" ] || [ "$THRESHOLD" = "0" ] || [ "$THRESHOLD" = "" ]; then
    echo "❌ COVERAGE_THRESHOLD is not specified" >&2
    exit 1
fi

# Run coverage generation first even though this is a check script.
if ! make coverage >/dev/null 2>&1; then
    echo "❌ make coverage failed" >&2
    exit 1
fi
if [ ! -f "$PERCENT_FILE" ]; then
    echo "❌ Coverage percent file missing: $PERCENT_FILE" >&2
    exit 1
fi

# validate numeric (allow formats: 1, 1.23, .5, 0.5).
VAL=$(tr -d '\r\n' < "$PERCENT_FILE")
VAL=${VAL%\%}
if ! printf '%s' "$VAL" | awk '/^([0-9]*\.[0-9]+|[0-9]+)$/ { exit 0 } { exit 1 }'; then
    echo "❌ Invalid coverage value: $VAL" >&2
    exit 1
fi

# Print coverage summary.
unit_val=$(cat tmp/coverage_unit_percent.out 2>/dev/null || echo "n/a")
integration_val=$(cat tmp/coverage_integration_percent.out 2>/dev/null || echo "n/a")
echo "  Unit tests:        ${unit_val}%"
echo "  Integration tests: ${integration_val}%"
echo "  Total:             ${VAL}%"
echo ""

# Check against threshold.
awk -v p="$VAL" -v t="$THRESHOLD" '
BEGIN {
    p0 = p + 0;
    t0 = t + 0;
    if (p0 < t0) {
        printf("❌ Coverage %s%% is below %s%%\n", p, t);
        exit 1
    } else {
        printf("✅ Total coverage meets the threshold of %s%%\n", t);
    }
}
'
