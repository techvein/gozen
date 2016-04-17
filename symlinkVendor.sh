#!/bin/bash

# To use glide within Intellij IDEA, it symlinks vendor/* to vendor/src.
#https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/1820#issuecomment-158954480
rm -rf vendor/src
mkdir vendor/src
FILES=vendor/*
for f in $FILES
do
  if [ "$f" != "vendor/src" ]
  then
    FULL=`pwd`/$f
    echo "Symlinking vendor source dir: $FULL"
    ln -s $FULL vendor/src
  fi
done
