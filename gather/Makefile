create:
	protoc --proto_path=proto proto/*.proto --go_out=pkg/gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=pkg/gen/

clean:
	rm gen/proto/*.go