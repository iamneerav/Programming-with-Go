#!/usr/bin/env bash
set -euo pipefail

total_projects=0
executed_projects=0
skipped_projects=0
current_project=""

trap 'echo "❌ Failure while running: ${current_project:-unknown project}"' ERR

find . -mindepth 1 -maxdepth 1 -type d \( -name 'hands-on-*' -o -name 'hands-on-understanding-*' \) -print0 |
while IFS= read -r -d '' dir; do
  current_project="${dir#./}"
  total_projects=$((total_projects + 1))

  echo "========================================"
  echo "Project [$total_projects]: $current_project"

  if [[ -f "$dir/go.mod" ]]; then
    echo "Mode: module (go test ./... + go build ./...)"
    (cd "$dir"; go test -v ./...; go build -v ./...)
    executed_projects=$((executed_projects + 1))
    echo "Status: completed"
  elif [[ -f "$dir/main.go" ]]; then
    echo "Mode: main (go run main.go)"
    (cd "$dir"; go run main.go)
    executed_projects=$((executed_projects + 1))
    echo "Status: completed"
  else
    echo "Skipped: no main.go / go.mod"
    skipped_projects=$((skipped_projects + 1))
  fi
done

echo "========================================"
echo "CI summary"
echo "Total folders scanned: $total_projects"
echo "Executed: $executed_projects"
echo "Skipped: $skipped_projects"