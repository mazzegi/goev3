#!/bin/bash

echo 'build ...'
GOOS=linux GOARCH=arm GOARM=5 go build
echo 'build ... done'
echo 'deploy ...'
scp ev3lua robot@192.168.0.107:
echo 'deploy ... done'