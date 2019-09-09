#!/bin/bash

if [ -f docker/DockerFile ]; then
    cp docker/Dockerfile ./
fi

git archive -v -o tantakatu-api-bundle.zip --format=zip HEAD
