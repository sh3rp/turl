
install-pb:
	./install-pb.sh

go-get: proto-gen
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u github.com/boltdb/bolt

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
	go build -o turl ./...

run: go-get
	go run main.go

.PHONY: install-pb proto-gen build run