#!/bin/bash

# Generate the service stub
protoc --proto_path=proto \
       -I/usr/local/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --go_out=plugins=grpc:requester/ \
       proto/*.proto
