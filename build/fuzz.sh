#!/bin/bash
#
# Run fuzz tests for cheat
# Usage: ./scripts/fuzz.sh [duration]
#
# Note: Go's fuzzer will fail immediately if it finds a known failing input
# in the corpus (testdata/fuzz/*). This is by design - it ensures you fix
# known bugs before searching for new ones. To see failing inputs:
#   ls internal/*/testdata/fuzz/*/
#

set -e

DURATION="${1:-15s}"

# Define fuzz tests: "TestName:Package:Description"
TESTS=(
    "FuzzParse:./internal/sheet:YAML frontmatter parsing"
    "FuzzValidateSheetName:./internal/cheatpath:sheet name validation (path traversal protection)"
    "FuzzSearchRegex:./internal/sheet:regex search operations"
    "FuzzSearchCatastrophicBacktracking:./internal/sheet:catastrophic backtracking"
    "FuzzTagged:./internal/sheet:tag matching with malicious input"
    "FuzzFilter:./internal/sheets:tag filtering operations"
    "FuzzTags:./internal/sheets:tag aggregation and sorting"
    "FuzzFindLocalCheatpath:./internal/config:recursive .cheat directory discovery"
    "FuzzFindLocalCheatpathNearestWins:./internal/config:nearest .cheat wins invariant"
)

echo "Running fuzz tests ($DURATION each)..."
echo

for i in "${!TESTS[@]}"; do
    IFS=':' read -r test_name package description <<< "${TESTS[$i]}"
    echo "$((i+1)). Testing $description..."
    go test -fuzz="^${test_name}$" -fuzztime="$DURATION" "$package"
    echo
done

echo "All fuzz tests passed!"