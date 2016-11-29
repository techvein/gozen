#!/bin/bash
if [ "$CIRCLECI" = true ] ; then
    mkdir -p $GOPATH/src
    git clone https://github.com/techvein/gozen.git $GOPATH/src/
else
    echo "CIRCLECI is false"
fi

