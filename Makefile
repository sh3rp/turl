
all: install-pb go-get proto-gen build docker-build

install-pb:
	@./install-pb.sh

go-get:
	@go mod tidy
	@rm -rf $(PWD)/tmp/grpc-gateway
	@git clone https://github.com/grpc-ecosystem/grpc-gateway $(PWD)/tmp/grpc-gateway

proto-gen:
	@rm -rf proto 
	@mkdir proto
	@protoc -I. \
		-Itmp/protobuf/include \
  		-I$(PWD)/tmp/grpc-gateway/third_party/googleapis \
  		--go_out=plugins=grpc:proto \
  		turl.proto
	@protoc -I. \
		-Itmp/protobuf/include \
  		-I$(PWD)/tmp/grpc-gateway/third_party/googleapis \
  		--grpc-gateway_out=logtostderr=true:proto \
  		turl.proto

build: go-get
	@go build -o turl main.go
	@rm -rf target
	@mkdir target
	@mv turl target

run: go-get
	@go run main.go

docker-build: build
	@docker build . -t turl:latest

.PHONY: install-pb go-get proto-gen build run all
