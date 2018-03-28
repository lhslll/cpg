#!/bin/bash
echo "Begin build rcapi"

cd ~/go/src/github.com/cpg/rcapi
go build ~/go/src/github.com/cpg/rcapi/rcapi.go

echo "Begin build rcportal"

cd ~/go/src/github.com/cpg/rcportal
go build ~/go/src/github.com/cpg/rcportal/rcportal.go

echo "Begin build rcwechat"

cd ~/go/src/github.com/cpg/rcwechat
go build ~/go/src/github.com/cpg/rcwechat/rcwechat.go

