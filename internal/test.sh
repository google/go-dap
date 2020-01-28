#!/bin/bash
set -euo pipefail

if [[ $# -gt 0 ]]; then
  echo "usage: runchecks.sh" 1>&2
  exit 64
fi

# ensure_go_binary verifies that a binary exists in $PATH corresponding to the
# given go-gettable URI. If no such binary exists, it is fetched via `go get`.
ensure_go_binary() {
  local binary=$(basename $1)
  if ! [ -x "$(command -v $binary)" ]; then
    info "Installing: $1"
    # Run in a subshell for convenience, so that we don't have to worry about
    # our PWD.
    (set -x; cd && env GO111MODULE=on go get -u $1)
  fi
}

echo "**** Running Go tests"
go test -race -count=1 ./...

echo "**** Running staticcheck"
ensure_go_binary honnef.co/go/tools/cmd/staticcheck
staticcheck ./...

echo "**** Running go vet"
go vet ./...
