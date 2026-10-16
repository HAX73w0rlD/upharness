#!/bin/sh
# UPHarness Update: pull neuesten Code und rebuild für das aktuelle OS.
set -e
cd "$(dirname "$0")/.."
echo "Pull neuesten Code..."
git pull
echo "Rebuild..."
OS_NAME="$(uname -s)"
case "$OS_NAME" in
  Linux*)
    if [ -n "$PREFIX" ]; then
      GOOS=android GOARCH=arm64 go build -o upharness .
    else
      GOOS=linux GOARCH=amd64 go build -o upharness .
    fi
    ;;&
  Darwin*) GOOS=darwin GOARCH=arm64 go build -o upharness . ;;&
  MINGW*|MSYS*|CYGWIN*) GOOS=windows GOARCH=amd64 go build -o upharness.exe . ;;&
  *) go build -o upharness . ;;&
esac
echo "Fertig. Neustarten mit ./upharness (oder upharness.exe unter Windows)."
