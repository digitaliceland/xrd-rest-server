#!/bin/bash

echo compling for Windows: release/windows/rest-server.exe
mkdir -p release/windows
GOOS=windows GOARCH=amd64 go build -o release/windows/rest-server.exe main.go

echo compling for MacOS: release/macos/rest-server
mkdir -p release/macos
GOOS=darwin GOARCH=amd64 go build -o release/macos/rest-server main.go

echo compling for Linux: release/linux/rest-server
mkdir -p release/linux
GOOS=linux GOARCH=amd64 go build -o release/linux/rest-server main.go

