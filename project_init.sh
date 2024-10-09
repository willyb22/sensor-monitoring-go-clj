#!/bin/bash

set -e

dirs=(postgres-db go-backend clj-monitoring)

for dir in "${dirs[@]}"; do
    mkdir -p $dir/src
    touch $dir/{Dockerfile,.env,.dockerignore}
done