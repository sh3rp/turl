module github.com/sh3rp/turl

go 1.12

require (
	github.com/boltdb/bolt v1.3.1
	github.com/ghodss/yaml v1.0.0
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/grpc-gateway v1.9.0
	github.com/rogpeppe/fastuuid v1.1.0
	github.com/rs/zerolog v1.14.3
	github.com/stretchr/testify v1.3.0
	golang.org/x/net v0.0.0-20190607181551-461777fb6f67
	golang.org/x/sys v0.0.0-20190608050228-5b15430b70e3 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/genproto v0.0.0-20190605220351-eb0b1bdb6ae6
	google.golang.org/grpc v1.21.1
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.12.0
