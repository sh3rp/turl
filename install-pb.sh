#!/bin/bash

rm -rf tmp
mkdir tmp
cd tmp
wget -O protobuf.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
unzip -d /usr/local protobuf.zip 