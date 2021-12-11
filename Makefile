# generate data access classes and gRPC stubs.
.PHONY: api
api:
	export PATH="$PATH:$(go env GOPATH)/bin" # add protoc-gen-go to shell path
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative    api/user.proto
