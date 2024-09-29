#!/bin/bash

SERVER_ADDR="localhost:50051"

for i in {1..500}
do
  go run cmd/client/main.go --client_id=$i -server_addr=$SERVER_ADDR &
done
