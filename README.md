# A gRPC performance test project

## Description

This is a small playground project to test gRPC streaming with a Go server and client apps.

## Prerequisite

**For macOS:**

```shell
#  Install Homebrew (if not already installed):
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install Protocol Buffer Compiler (protoc):
brew install protobuf

# Install Go gRPC and protobuf plugins:
brew install protoc-gen-go
brew install protoc-gen-go-grpc

# Install gRPC and protobuf packages for Go:
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go

# Install .NET SDK for C#:
brew install --cask dotnet-sdk

# Install gRPC C# tools:
dotnet tool install -g Grpc.Tools
```

**For Windows:**

```shell
# Install Chocolatey (if not already installed):
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))

# Install Protocol Buffer Compiler (protoc):
choco install protobuf

#  Install Go gRPC and protobuf plugins:
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go

# Install .NET SDK for C#:
choco install dotnet-sdk

# Install gRPC C# tools:
dotnet tool install -g Grpc.Tools
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
# run client in another terminal or another host to start 500 clients
./run_clients.sh
```

The script will spawn 500 client instances. Using the script you can change the number of client, clients ids and a server's host and port. By default it uses `localhost:50051`.
