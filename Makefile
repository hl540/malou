gen:
	protoc -I ./proto --go_out=. ./proto/v1/*.proto
	protoc -I ./proto --go-grpc_out=. ./proto/v1/*.proto
	protoc -I ./proto --grpc-gateway_out=. ./proto/v1/*.proto
