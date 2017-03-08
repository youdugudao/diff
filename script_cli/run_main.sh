#!/usr/bin/env bash

var=$(pgrep contract)

cd /var/go/src/contract
if [ $var ]; then
    echo "old id: $var"
    kill -9 $var
else
    echo "pid not found"
fi
rm contract
go build
chmod +x contract
(nohup ./contract >/var/log/go/contract/day.log 2>/var/log/go/contract/error.log &)

var_new=$(pgrep contract)
if [ $var_new ]; then
    echo "new id: $var_new"
else
    echo "restart failed"
fi