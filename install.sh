#!/usr/bin/env bash

[[ -d "${HOME}/.local/bin" ]] || mkdir -p "${HOME}/.local/bin"

go get .

go build -o "${HOME}/.local/bin/sb" -ldflags "-X github.com/aadam-ali/second-brain-cli/config.version=$(git rev-parse --short HEAD)" main.go
