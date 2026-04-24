#!/usr/bin/env bash

set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$repo_root"

shopt -s nullglob

exercise_dirs=(hands-on-*)
exercise_dirs+=(hands-on-understanding-*)

if [ ${#exercise_dirs[@]} -eq 0 ]; then
  echo "No hands-on directories found"
  exit 1
fi

for dir in "${exercise_dirs[@]}"; do
  if [ ! -d "$dir" ]; then
    continue
  fi

  if [ -f "$dir/main.go" ]; then
    echo "==> go run ./$dir"
    go run "./$dir"
  fi
done
