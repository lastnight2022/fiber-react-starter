: go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27

: protoc --version 23.2

protoc --go_out=. --go-grpc_out=. *.proto
