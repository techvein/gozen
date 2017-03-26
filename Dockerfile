FROM golang:1.8

ENV WORK $GOPATH/src/github.com/techvein/gozen
WORKDIR  $WORK
RUN mkdir -p  $WORK

ADD . $WORK
