#!/usr/bin/env bash
set -euo pipefail

# Pin the Go toolchain to a current release. The base image ships an older
# apt-provided Go, so install the official binary distribution when the pinned
# version is not already present, then expose it ahead of the system Go.
GO_VERSION="1.27.1"
GO_ROOT="/usr/local/go"

current="$("$GO_ROOT/bin/go" version 2>/dev/null | awk '{print $3}' || true)"
if [ "$current" != "go${GO_VERSION}" ]; then
  tmp="$(mktemp -d)"
  curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" -o "$tmp/go.tar.gz"
  sudo rm -rf "$GO_ROOT"
  sudo tar -C /usr/local -xzf "$tmp/go.tar.gz"
  rm -rf "$tmp"
fi

# /usr/local/bin precedes /usr/bin on PATH, so these symlinks make the pinned
# toolchain the default `go`/`gofmt` without editing shell profiles.
sudo ln -sfn "$GO_ROOT/bin/go" /usr/local/bin/go
sudo ln -sfn "$GO_ROOT/bin/gofmt" /usr/local/bin/gofmt
hash -r 2>/dev/null || true

# Module-mode project (has go.mod): build and test straight from the checkout.
go version
go mod download
go build ./...
go vet ./...
go test ./...

echo "2e11 environment ready (Go ${GO_VERSION}, module mode)."
