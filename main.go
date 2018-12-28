package main

import (
	"os"

	"github.com/sh3rp/turl/db"
	"github.com/sh3rp/turl/hash"
	"github.com/sh3rp/turl/logger"
	"github.com/sh3rp/turl/server"
)

var grpcPort = 8080
var proxyPort = 9090
var httpPort = 8888

func main() {
	log := logger.NewZeroLogger()

	log.Info("Starting database")

	database, err := db.NewDB(os.TempDir(), log)

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Creating hashing")

	hasher := hash.NewAdler32Provider()

	log.Info("Creating servers")

	srv := server.NewTurlServer(database, hasher, log)

	log.Info("Starting GRPC server on port %d", grpcPort)

	go srv.StartGRPC(grpcPort)

	log.Info("Starting proxy server on port %d", proxyPort)

	go srv.StartProxy(grpcPort, proxyPort)

	log.Info("Starting http server on port %d", httpPort)

	srv.StartHTTP(httpPort)
}
