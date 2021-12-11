# generate data access classes and gRPC stubs.
.PHONY: api
api:
	export PATH="$PATH:$(go env GOPATH)/bin" # add protoc-gen-go to shell path
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative    api/user.proto

# dependencies
.PHONY: mod
mod:
	go mod tidy
	go mod download

# build binaries
.PHONY: build
build: build-server build-client
build-server:
	go build -o server ./cmd/account
build-client:
	go build -o client ./test/client
