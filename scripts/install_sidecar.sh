#!/usr/bin/env bash
set -euo pipefail

GOOS=linux go build -o bin/sidecar ./sidecar/sidecar.go
