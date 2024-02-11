#!/bin/bash
echo "go build"
# git pull 
#go mod tidy
#go build -o go-admin main.go
echo "kill server gvaServer"
killall gvaServer # kill go-admin service
nohup ./gvaServer -c config.online.yaml >> access.log 2>&1 & #后台启动服务将日志写入access.log文件
echo "run server success"
ps -aux | grep gavServer
