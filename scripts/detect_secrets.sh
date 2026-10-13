#!/usr/bin/env bash
# Scan the repository with IBM detect-secrets and fail CI when new or
# confirmed secrets are present. This matches the GitHub Actions check used
# by other IBM Go providers and SDKs (for example IBM/go-sdk-core).
#
# Requires: pip install "git+https://github.com/ibm/detect-secrets.git@0.13.1+ibm.64.dss#egg=detect-secrets"
set -euo pipefail

# pip --user installs land here on some runners and local machines.
export PATH="${HOME}/.local/bin:${PATH}"

if ! command -v detect-secrets >/dev/null 2>&1; then
  echo "detect-secrets is not installed."
  echo "Install it with:"
  echo "  pip install --upgrade \"git+https://github.com/ibm/detect-secrets.git@0.13.1+ibm.64.dss#egg=detect-secrets\""
  exit 1
fi

# Refresh the baseline so newly introduced secrets are recorded, then audit
# against the committed allow-list of known false positives.
detect-secrets scan --update .secrets.baseline
detect-secrets -v audit --report --fail-on-unaudited --fail-on-live --fail-on-audited-real .secrets.baseline
