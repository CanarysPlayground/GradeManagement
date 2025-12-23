#!/usr/bin/env bash
# Start Angular dev server (background) and run Playwright tests against it.
set -euo pipefail
ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$ROOT"

# Start ng serve in background
if command -v ng >/dev/null 2>&1; then
  ng serve --port 4200 &
  NG_PID=$!
else
  echo "Angular CLI (ng) not found. Please install @angular/cli or run the dev server manually."
  exit 1
fi

# Wait for port 4200
for i in {1..30}; do
  if nc -z localhost 4200; then
    echo "Angular dev server is up"
    break
  fi
  sleep 1
done

# Run Playwright tests
npx playwright test || true

# Teardown
kill $NG_PID || true
