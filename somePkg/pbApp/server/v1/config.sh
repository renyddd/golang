#!/bin/bash

GIT=$(which git)
# `bash generated_proto.sh`
# to avoid "$GOPATH/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis: warning: directory does not exist.
# google/api/annotations.proto: File not found." to checkout tag v1.16.0
GRPC_GATEWAY="github.com/grpc-ecosystem/grpc-gateway/"
cd $GOPATH"/src/github.com/grpc-ecosystem/grpc-gateway/" && $GIT checkout v1.16.0

# to generate server.pb.gp
protoc -I . --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis server.proto

# to generate server.pb.gw.go
protoc -I . --grpc-gateway_out ./ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis server.proto