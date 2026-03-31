#!/usr/bin/env bash
set -euo pipefail

find . -mindepth 1 -maxdepth 1 -type d \( -name 'hands-on-*' -o -name 'hands-on-understanding-*' \) -print0 |
while IFS= read -r -d '' dir; do
  echo "Checking ${dir#./}"

  if [[ -f "$dir/go.mod" ]]; then
    (cd "$dir"; go test -v ./...; go build -v ./...)
  elif [[ -f "$dir/main.go" ]]; then
    (cd "$dir"; go run main.go)
  else
    echo "Skipped: no main.go / go.mod"
  fi
done