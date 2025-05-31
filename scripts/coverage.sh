#!/bin/bash

green() {
  echo -e "\033[32m$1\033[0m"
}

red() {
  echo -e "\033[31m$1\033[0m"
}

cyan() {
  echo -e "\033[36m$1\033[0m"
}

cyan "🔍 Code coverage analyzing..."
echo "----------------------------------------------------------------------------------"

COV_PATH=test/coverage

mkdir -p $COV_PATH
if ! go test -cover $(go list ./concurrency/... | grep -v /app) -coverprofile=$COV_PATH/coverage.out; then
  red "❌ Tests failed. Cannot generate coverage report."
  exit 1
fi

if [ ! -f "test/coverage/coverage.out" ]; then
  red "❌ Coverage profile not generated."
  exit 1
fi

go tool cover -html=$COV_PATH/coverage.out -o $COV_PATH/coverage.html
echo "----------------------------------------------------------------------------------"

total_coverage=$(go tool cover -func=test/coverage/coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
coverage_threshold=80.0
# Check if bc is available
if ! command -v bc &> /dev/null; then
  red "❌ 'bc' command not found. Please install bc for coverage threshold comparison."
  exit 1
fi

comparison=$(awk "BEGIN {print ($total_coverage >= $coverage_threshold) ? 1 : 0}")
if [ "$comparison" -eq 0 ]; then
  red "📈 Total coverage: $total_coverage%"
  red "❌ Code coverage $total_coverage% is below the threshold of $coverage_threshold%."
  exit 1
else
  green "📈 Total coverage: $total_coverage%"
  green "✅ Code coverage $total_coverage% meets the threshold of $coverage_threshold%."
fi
