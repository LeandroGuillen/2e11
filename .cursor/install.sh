#!/usr/bin/env bash
set -euo pipefail

# 2e11 is a legacy, pre-modules Go project: it has no go.mod and its packages
# import each other via the canonical path github.com/leandroguillen/2e11/...
# Building and testing therefore require GOPATH mode with the checkout exposed
# under $GOPATH/src/github.com/leandroguillen/2e11.

# Disable modules for this toolchain (persisted in the go env config file).
go env -w GO111MODULE=off

GOPATH="$(go env GOPATH)"
PKG_PARENT="$GOPATH/src/github.com/leandroguillen"
PKG_LINK="$PKG_PARENT/2e11"

mkdir -p "$PKG_PARENT"
# Point the canonical import path at this checkout so edits are picked up live.
ln -sfn "$PWD" "$PKG_LINK"

# Validate the toolchain end-to-end from within GOPATH.
cd "$PKG_LINK"
go build ./...
go vet ./...
go test ./...

echo "2e11 environment ready. Build/test from: $PKG_LINK"
