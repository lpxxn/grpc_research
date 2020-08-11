#! /bin/bash
protoc --go_out=plugins=grpc:. helloworld.proto
# protoc -I=$GOPATH/src:. --go_out=plugins=grpc:. --govalidators_out=. helloworld.proto