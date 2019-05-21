#!/usr/bin/env bash
set -euo pipefail

cf v3-create-app sidecar-dependent-app
cf v3-apply-manifest -f manifest.yml
cf v3-push sidecar-dependent-app

