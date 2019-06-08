FROM ubuntu:latest

EXPOSE 8080
EXPOSE 9090
EXPOSE 8888

RUN apt update && apt -y upgrade && apt install -y autoconf libtool wget unzip

COPY target/turl .
ENTRYPOINT ["/turl"]