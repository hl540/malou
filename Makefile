gen:
	protoc --proto_path=./proto/v1 --go_out=. --go-grpc_out=. ./proto/v1/*.proto
