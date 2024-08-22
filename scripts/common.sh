#!/usr/bin/env bash

# Common utilities, variables and checks for all build scripts.
set -eEuo pipefail

# Unset CDPATH, having it set messes up with script import paths
unset CDPATH

# This will canonicalize the path
CC_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")"/.. && pwd -P)

source "${CC_ROOT}/scripts/lib/init.sh"

# The variable CC_CLIENT_SIDE_COMPONENTS is used to define cc client-side components.
# These components no need to installed as a service, but used as a command line.
declare -Ax CC_CLIENT_SIDE_COMPONENTS=(
  # ["ctl"]="cctl"
)

# The variable CC_SERVER_SIDE_COMPONENTS is used to define cc server-side components.
# These components need to installed as a service.
declare -Ax CC_SERVER_SIDE_COMPONENTS=(
  ["uc"]="cc-usercenter"
)

# The variable CC_ALL_COMPONENTS is used to define all cc components.
declare -Ax CC_ALL_COMPONENTS
for key in "${!CC_CLIENT_SIDE_COMPONENTS[@]}"; do
  CC_ALL_COMPONENTS["$key"]="${CC_CLIENT_SIDE_COMPONENTS[$key]}"
done
for key in "${!CC_SERVER_SIDE_COMPONENTS[@]}"; do
  CC_ALL_COMPONENTS["$key"]="${CC_SERVER_SIDE_COMPONENTS[$key]}"
done

