lint:
	golangci-lint run

test:
	go test -v -race ./...

proto:
	
	protoc  --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative schema/grpc_websocket_api_gateway.proto

run:
	go run cmd/*.go

run_race:
	go run -race cmd/*.go