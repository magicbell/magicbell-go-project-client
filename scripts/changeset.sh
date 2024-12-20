#!/usr/bin/env bash

# Ensure script runs relative to the repo root
set -euo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")/.."

# Trap signals for graceful exit
trap 'echo "Script interrupted" >&2; exit 2' INT TERM

# Run changeset with passed arguments
npx changeset "$@"
exitCode=$?

if [ $exitCode -ne 0 ]; then
  echo "Changeset finished with error" >&2
  exit $exitCode
fi

# Handle "post" scripts if arguments are provided
if [ $# -gt 0 ]; then
  command="$1"
  shift # Remove the first argument from $@
  script="changeset:post${command}"
  echo "Running post script: ${script}"
  npm run "${script}" --if-present -- "$@"
fi
