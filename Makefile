
install-pb:
	./install-pb.sh

go-get:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u github.com/boltdb/bolt
	go get -u github.com/stretchr/testify
	go get -u github.com/rs/zerolog

proto-gen:
	rm -rf proto
	mkdir proto
	protoc -I/usr/local/include -I. \
  		-I$(GOPATH)/src \
  		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  		--go_out=plugins=grpc:proto \
  		turl.proto
	protoc -I/usr/local/include -I. \
  		-I$(GOPATH)/src \
  		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  		--grpc-gateway_out=logtostderr=true:proto \
  		turl.proto

build: go-get
	go build -o turl main.go
	cp turl /go/bin

run: go-get
	go run main.go

all: install-pb go-get proto-gen build

.PHONY: install-pb proto-gen build run all