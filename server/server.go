package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sh3rp/turl/db"
	"github.com/sh3rp/turl/hash"
	"github.com/sh3rp/turl/logger"
	"github.com/sh3rp/turl/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Server is the base struct
type turlServer struct {
	db     db.TurlDB
	hasher hash.HashProvider
	logger logger.Log
}

func NewTurlServer(db db.TurlDB, hasher hash.HashProvider, logger logger.Log) turlServer {
	return turlServer{
		db:     db,
		hasher: hasher,
		logger: logger,
	}
}

func (s turlServer) StartHTTP(httpPort int) error {
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", httpPort),
		Handler:        s,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := srv.ListenAndServe()

	if err != nil {
		s.logger.Info("Error: %+v", err)
	} else {
		s.logger.Info("Proxy server listening on %d", httpPort)
	}

	return err
}

func (s turlServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Path = %s", r.URL.Path)
	urlObj := proto.URL{Hash: r.URL.Path[1:]}
	s.logger.Info("urlObj: %+v", urlObj)
	url, err := s.db.Get(urlObj)
	if err != nil && url.Url == "" {
		w.Write([]byte("Error"))
	} else {
		s.logger.Info("Redirecting to %s", url.Url)
		http.Redirect(w, r, url.Url, http.StatusMovedPermanently)
	}
}

//StartProxy starts the HTTP proxy to the grpc service
func (s turlServer) StartProxy(grpcPort, proxyPort int) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := proto.RegisterTurlServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", grpcPort), opts)

	if err != nil {
		return err
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", proxyPort), mux)

	if err != nil {
		s.logger.Info("Error: %+v", err)
	} else {
		s.logger.Info("Proxy server listening on %d", proxyPort)
	}

	return err
}

//StartGRPC starts the GRPC server
func (s turlServer) StartGRPC(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return err
	}

	server := grpc.NewServer()

	proto.RegisterTurlServiceServer(server, &s)
	reflection.Register(server)

	err = server.Serve(listener)

	if err != nil {
		s.logger.Info("Error: %+v", err)
	} else {
		s.logger.Info("GRPC server listening on %d", port)
	}

	return err
}

func (s *turlServer) NewURL(ctx context.Context, in *proto.URL) (*proto.URL, error) {
	s.logger.Info("NewURL called: %+v", *in)
	in.Hash = s.hasher.Hash(in.Url)
	newUrl, err := s.db.Create(*in)
	if err != nil {
		s.logger.Error("Error: %+v", err)
	} else {
		s.logger.Info("NewURL created: %+v", newUrl)
	}
	return &newUrl, err
}

func (s *turlServer) GetURL(ctx context.Context, in *proto.URL) (*proto.URL, error) {
	s.logger.Info("GetURL called: %+v", *in)
	newUrl, err := s.db.Get(*in)
	if err != nil {
		s.logger.Error("Error: %+v", err)
	} else {
		s.logger.Info("GetURL returns: %+v", newUrl)
	}
	return &newUrl, err
}
