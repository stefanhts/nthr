gen:
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
