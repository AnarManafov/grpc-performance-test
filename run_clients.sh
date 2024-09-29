#!/bin/bash

for i in {1..500}
do
  go run cmd/client/main.go &
done
