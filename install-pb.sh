#!/bin/bash

rm -rf tmp
mkdir tmp
cd tmp
git clone https://github.com/protocolbuffers/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
cd ../..
rm -rf tmp
sudo ldconfig