# Turl

Turl is a URL shortening service written in Go.  It was intended as an exercise in
using GRPC + GRPC Gateway + other Go libraries.  The server runs on three ports:

* 8080 - the GRPC service port
* 9090 - the GRPC HTTP gateway service port
* 8888 - the HTTP service that does the lookup and redirect

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