#!/bin/bash

SERVER_ADDR="localhost:50051"

for i in {1..500}
do
  ./bin/client --client_id=$i -server_addr=$SERVER_ADDR &
done
