#!/bin/bash

DATENOW=`date -u +%Y%m%d.%H%M%S`

GOOS=windows GOARCH=amd64 go build -ldflags "-X stl3mf/cmd.VERSION=${DATENOW}" -o dist/stl3mf-win-amd64.exe
GOOS=darwin GOARCH=amd64 go build -ldflags "-X stl3mf/cmd.VERSION=${DATENOW}" -o dist/stl3mf-macos-amd64
GOOS=darwin GOARCH=arm64 go build -ldflags "-X stl3mf/cmd.VERSION=${DATENOW}" -o dist/stl3mf-macos-arm64
GOOS=linux GOARCH=amd64 go build -ldflags "-X stl3mf/cmd.VERSION=${DATENOW}" -o dist/stl3mf-linux-amd64
