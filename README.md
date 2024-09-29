# A gRPC performance test project

## Description


## Requirments

```shell
# OS X

brew install go
brew install grpc
brew install protoc-gen-go
```

## Build/Run

```shell
go mod tidy
go mod download
```

```shell
# Generate the Go code from the .proto file:
protoc --go_out=. --go-grpc_out=. proto/service.proto
```

```shell
# Run server
go run cmd/server/main.go
```

```shell
# run client in another terminal to start 500 clients
./run_clients.sh
```
