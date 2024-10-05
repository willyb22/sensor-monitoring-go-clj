#!/bin/bash

set -e

dirs=(go-sensor go-backend clj-monitoring)

for dir in "${dirs[@]}"; do
    mkdir -p $dir/src
    touch $dir/{Dockerfile,.env,.dockerignore}
done