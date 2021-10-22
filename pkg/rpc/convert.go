package rpc

import (
	api "github.com/ivanbulyk/grpc_websocket_api_gateway/api/schema"
	"github.com/ivanbulyk/grpc_websocket_api_gateway/pkg/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertToProtoL2(symbol string, l2 types.L2OrderBook) *api.L2OrderBook {
	convertItem := func(item *types.L2OrderBookItem) *api.L2OrderBookItem {
		return &api.L2OrderBookItem{
			Price:  item.Price.String(),
			Volume: item.Volume,
		}
	}

	ret := &api.L2OrderBook{
		Symbol: symbol,
		Time:   timestamppb.Now(),
		Bid:    make([]*api.L2OrderBookItem, 0, len(l2.Bid)),
		Ask:    make([]*api.L2OrderBookItem, 0, len(l2.Ask)),
	}

	for _, item := range l2.Bid {
		ret.Bid = append(ret.Bid, convertItem(item))
	}
	for _, item := range l2.Ask {
		ret.Ask = append(ret.Ask, convertItem(item))
	}

	return ret
}
