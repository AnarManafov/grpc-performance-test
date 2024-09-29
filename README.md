# A gRPC performance test project

## Description


## Requirments

```shell
# OS X

brew install go
brew install grpc
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```

## Build

Go gRPC files are included in the repo. But if there is a need to regenerate them, use:

```shell
# Generate the Go code from the .proto file:
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/service.proto
```

Download dependencies: 

```shell
# Download dependencies
go mod tidy
go mod download
./build.sh
```

```shell
# Run server
./bin/server
```

```shell
# run client in another terminal to start 500 clients
./run_clients.sh
```
The script will spawn 500 client instances. Using the script you can change the number of client, clients ids and a server's host and port. By default it uses `localhost:50051`.
