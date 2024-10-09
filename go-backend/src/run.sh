#!/bin/sh

set -e
# run fresh in the background
# fresh &

go build -a -o bin/go-backend -mod=vendor ./cmd/go-backend
bin/go-backend &

# run the simulation
go build -o bin/go-sensor -mod=vendor ./cmd/go-sensor
bin/go-sensor