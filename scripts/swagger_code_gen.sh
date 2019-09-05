#!/bin/bash

set -e

rm -rf swagger.json
multi-file-swagger -o yaml ./spec/root.yaml > swagger.yaml
rm -rf restapi/
rm -rf models/
swagger generate server -A tantakatu-api -f swagger.yaml --model-package=models

sh ./scripts/copy_post_code_gen.sh

rm -rf client/
swagger generate client -f swagger.yaml -A tantakatu-api