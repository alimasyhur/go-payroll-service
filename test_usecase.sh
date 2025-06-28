#!/bin/bash
set -e

ROOT_DIR=$(pwd)

for dir in $(find internal/app/usecase -type f -name '*_test.go' ! -path '*/mocks/*' -exec dirname {} \; | sort -u); do
  echo "▶️ Running test in: $dir"
  go test -v "$ROOT_DIR/$dir"
done
