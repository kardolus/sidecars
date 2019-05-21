#!/usr/bin/env bash
set -euo pipefail

mkdir -p bin
GOOS=linux go build -o bin/sidecar ./sidecar/sidecar.go
