#!/bin/bash

PROTOC="3.6.1"
INSTALLED_VERSION=$(protoc --version | awk '{print $2}')

if [ $PROTOC != $INSTALLED_VERSION ]; then
    echo "Updating protoc to version ${PROTOC}..."
    rm -rf tmp
    mkdir tmp
    cd tmp
    wget -O protobuf.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
    sudo unzip -d /usr/local protobuf.zip 
fi
