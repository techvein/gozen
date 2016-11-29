#!/bin/bash
if [ "$CIRCLECI" = true ] ; then
    mkdir -p ./gopath/src
    git clone https://github.com/techvein/gozen.git ./gopath/src/
else
    echo "CIRCLECI is false"
fi

