package onederx

import "github.com/shopspring/decimal"

func GetWsL2SubscribeRequest() interface{} {
	subscribeRequest := struct {
		Type    string
		Payload struct {
			Subscriptions []struct {
				Channel string
				Params  struct {
					Symbol string
				}
			}
		}
	}{}
	subscribeRequest.Type = "subscribe"
	subscribeRequest.Payload.Subscriptions = make([]struct {
		Channel string
		Params  struct {
			Symbol string
		}
	}, 1)
	subscribeRequest.Payload.Subscriptions[0].Channel = "l2"
	subscribeRequest.Payload.Subscriptions[0].Params.Symbol = "BTCUSD_P"

	return subscribeRequest
}

type WsL2Item struct {
	Price     decimal.Decimal `json:"price"`
	Volume    uint64          `json:"volume,string"`
	Side      string          `json:"side"`
	Timestamp int64           `json:"timestamp"`
}

type WsL2Update struct {
	Params struct {
		Symbol string `json:"symbol"`
	} `json:"params"`
	Payload WsL2Item `json:"payload"`
}

type WsL2Snapshot struct {
	Params struct {
		Symbol string `json:"symbol"`
	} `json:"params"`
	Payload struct {
		Snapshot []*WsL2Item `json:"snapshot"`
		Updates  []*WsL2Item `json:"updates"`
	}
}
