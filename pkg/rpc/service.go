package rpc

import (
	"log"
	"time"

	api "github.com/ivanbulyk/grpc_websocket_api_gateway/api/schema"
	"github.com/ivanbulyk/grpc_websocket_api_gateway/pkg/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	sources []types.Source
}

func NewService() *Service {
	return &Service{
		sources: make([]types.Source, 0),
	}
}

func (s *Service) AddSource(source types.Source) {
	s.sources = append(s.sources, source)
}

func (s *Service) GetL2OrderBook(req *api.L2OrderBookRequest, stream api.GrpcWebsocketApiGateway_GetL2OrderBookServer) error {

	log.Printf("client connected")

	if req.Size <= 0 {
		return status.Error(codes.InvalidArgument, "Invalid size")
	}
	if req.Interval <= 0 {
		return status.Error(codes.InvalidArgument, "Invalid interval")
	}
	var (
		stop bool
		l2   types.L2OrderBook
		err  error
	)
	for !stop {
		select {
		case <-time.After(time.Duration(req.Interval) * time.Millisecond):
			for _, source := range s.sources {
				l2, err = source.GetL2OrderBook(req.Symbol, int(req.Size))
				if err != nil {
					stop = true
					break
				}
				if err = stream.Send(ConvertToProtoL2(req.Symbol, l2)); err != nil {
					stop = true
				}
			}
		case <-stream.Context().Done():
			stop = true
		}
	}

	log.Printf("client disconnected")
	return nil
}
