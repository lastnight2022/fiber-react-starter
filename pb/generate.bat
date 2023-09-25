: go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

: protoc --version 23.2

protoc --go_out=. --go-grpc_out=. *.proto
