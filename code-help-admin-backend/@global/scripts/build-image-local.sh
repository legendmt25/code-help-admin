#!/bin/bash

cd ../..

apps=("hello")

for app in ${apps[@]} ; do
    docker build -t $app:latest -f deployment/DockerfileLocal --build-arg="SOURCE_DIRECTORY=$app"  .
done
