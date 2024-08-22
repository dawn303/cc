#!/usr/bin/env bash

# Log an error but keep going.  Don't dump the stack or exit.
function cc::log::error() {
  timestamp=$(date +"[%m%d %H:%M:%S]")
  echo "!!! ${timestamp} ${1-}" >&2
  shift
  for message; do
    echo "    ${message}" >&2
  done
}
