#!/bin/bash
echo "go build"
# git pull 
#go mod tidy
#go build -o go-admin main.go
echo "kill server service"
killall gvaServer # kill gvaServer service
