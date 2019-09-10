#!/bin/bash

set -e

rm -rf swagger.json
multi-file-swagger -o json ./spec/root.yaml > swagger.json
rm -rf restapi/
rm -rf models/
swagger generate server -A tantakatu-api -f swagger.json --model-package=models

sh ./scripts/copy_post_code_gen.sh

rm -rf client/
swagger generate client -f swagger.json -A tantakatu-api