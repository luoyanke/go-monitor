#!/bin/bash
currentPath=$(pwd)
rm -f app/*
cp server/go-monitor-config.json app/

if [ $1 == "mac" ]
then
    go build -i -x -o $currentPath/app/go_monitor_server ./server/
    go build -i -x -o $currentPath/app/goMCli ./client/
elif [ $1 == "linux" ]
  then
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -i -x -o /Users/luoyanke/huacloud/git/go-monitor/app/go_monitor_server ./server/
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -i -x -o /Users/luoyanke/huacloud/git/go-monitor/app/goMCli ./client/
else
    echo "pls enter args 'mac' or 'linux' "
fi

#scp -r app root@172.16.121.87:/opt