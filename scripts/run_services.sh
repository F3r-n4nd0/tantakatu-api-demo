#!/bin/bash

swagger serve ./swagger.json --port=9191 --no-open&

./tantakatu-api --host=0.0.0.0 --port=9090