#!/bin/bash

PROTOCVER="3.6.1"
PROTOCROOT="$(pwd)/tmp/protobuf"
PROTOCBIN="$PROTOCROOT/bin"

if [ ! -f "$PROTOCBIN/protoc" ]; then
    echo "Updating protoc to version ${PROTOC}..."
    rm -rf tmp
    mkdir tmp
    cd tmp
    wget -O protobuf.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
    unzip -d $PROTOCDIR protobuf.zip 
fi

go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway