#!/bin/bash
goctl api plugin -plugin goctl-swagger="swagger -filename user.json -host 127.0.0.1:9999 -schemes http" -api test.api -dir ../static/swagger-ui/
# docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/user.json -v $PWD:/foo swaggerapi/swagger-ui