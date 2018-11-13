#!/bin/bash

if [ "$1" = "" ] 
then
    func=test
else
    func=$1
fi

if [ "$2" = "" ]
then
    os=darwin
else
    os=$2
fi

cd seuxw
make build func=$func os=$os

./_output/local/bin/$func.x 
