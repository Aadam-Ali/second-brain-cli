#!/usr/bin/env bash

[[ -d "${HOME}/.local/bin" ]] || mkdir -p "${HOME}/.local/bin"

go get .

go build -o "${HOME}/.local/bin/sb" main.go
