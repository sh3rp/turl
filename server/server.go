package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sh3rp/turl/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Server is the base struct
type TurlServer struct{}

//StartProxy starts the HTTP proxy to the grpc service
func (s TurlServer) StartProxy(grpcPort, httpPort int) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := proto.RegisterTurlServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", grpcPort), opts)

	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}

//StartGRPC starts the GRPC server
func (s TurlServer) StartGRPC(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return err
	}

	server := grpc.NewServer()

	proto.RegisterTurlServiceServer(server, &s)
	reflection.Register(server)

	return server.Serve(listener)
}

func (s *TurlServer) NewURL(ctx context.Context, in *proto.URL) (*proto.URL, error) {
	return nil, nil
}

func (s *TurlServer) GetURL(ctx context.Context, in *proto.URL) (*proto.URL, error) {
	return nil, nil
}
