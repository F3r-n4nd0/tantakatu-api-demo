#!/bin/bash

if [ -f post_code_gen/configure_tantakatu_service_proxy.go ]; then
    cp post_code_gen/configure_tantakatu_service_proxy.go restapi/configure_tantakatu.go
fi

if [ -f post_code_gen/initialiser.go ]; then
    cp post_code_gen/initialiser.go restapi/initialiser.go
fi