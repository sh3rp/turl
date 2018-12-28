FROM golang:1.11.4-stretch

EXPOSE 8080
EXPOSE 9090
EXPOSE 8888

RUN apt update && apt -y upgrade && apt install -y autoconf libtool wget unzip

COPY . /go/src/github.com/sh3rp/turl/
WORKDIR /go/src/github.com/sh3rp/turl
RUN make all
ENTRYPOINT ["/go/bin/turl"]