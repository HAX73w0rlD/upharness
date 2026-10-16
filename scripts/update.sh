#!/bin/sh
# UPHarness update: pull latest code and rebuild for the current OS.
set -e
cd "$(dirname "$0")/.."
echo "Pulling latest code..."
git pull
echo "Rebuilding..."
OS_NAME="$(uname -s)"
case "$OS_NAME" in
  Linux*)
    if [ -n "$PREFIX" ]; then
      GOOS=android GOARCH=arm64 go build -o upharness .
    else
      GOOS=linux GOARCH=amd64 go build -o upharness .
    fi
    ;;
  Darwin*) GOOS=darwin GOARCH=arm64 go build -o upharness . ;;
  MINGW*|MSYS*|CYGWIN*) GOOS=windows GOARCH=amd64 go build -o upharness.exe . ;;
  *) go build -o upharness . ;;
esac
echo "Done. Restart with ./upharness (or upharness.exe on Windows)."
