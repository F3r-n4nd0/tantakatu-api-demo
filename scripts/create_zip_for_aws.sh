#!/bin/bash

if [ -f docker/DockerFile ]; then
    cp docker/Dockerfile ./
fi

git archive -v -o tantakatu-api-bundle-$(date +"%d%m%y%HH%mm").zip --format=zip HEAD
