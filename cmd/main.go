package main

import (
	"context"
	"log"
	"net"

	api "github.com/ivanbulyk/grpc_websocket_api_gateway/api/schema"
	"github.com/ivanbulyk/grpc_websocket_api_gateway/pkg/onederx"
	"github.com/ivanbulyk/grpc_websocket_api_gateway/pkg/rpc"
	"google.golang.org/grpc"
)

func main() {
	onederx := onederx.NewSource()
	onederx.Start(context.Background())

	service := rpc.NewService()
	service.AddSource(onederx)
	server := grpc.NewServer()
	api.RegisterGrpcWebsocketApiGatewayServer(server, service)

	lsn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
