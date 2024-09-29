#!/bin/bash

# Create bin directory if it doesn't exist
mkdir -p bin

# Compile client application
go build -o bin/client cmd/client/main.go

# Compile server application
go build -o bin/server cmd/server/main.go

echo "Client and Server applications have been compiled successfully."
