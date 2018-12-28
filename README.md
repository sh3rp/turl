# Turl

Turl is a URL shortening service written in Go.  It was intended as an exercise in
using GRPC + GRPC Gateway + other Go libraries.  The server runs on three ports:

* 8080 - the GRPC service port
* 9090 - the GRPC HTTP gateway service port
* 8888 - the HTTP service that does the lookup and redirect

## Quick Quick Start

Make sure you have Docker installed then run:

```
docker build . -t turl
docker run -it -p 8080:8080 -p 8888:8888 -p 9090:9090 turl
``` 

Then, you can post to port 9090 for the HTTP/REST API and 8888 for the URL ala:

```
> curl -XPOST -d "{'url':'http://www.cnn.com'}" http://localhost:9090/v1/turl
{"url":"http://www.cnn.com", "hash":"<hash>"}
> curl http://localhost:8888/<hash>
```

## Quick start

The project relies on the following dependencies to build:

* make
* autoconf
* go 1.10+

Once dependencies are installed, you can run

```
make run
```

To verify functionality.