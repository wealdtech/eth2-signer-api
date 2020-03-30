package v1

//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I /usr/local/include -I . --go_out=plugins=grpc:. accountmanager.proto
//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I /usr/local/include -I . --go_out=plugins=grpc:. eth2.proto
//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I /usr/local/include -I . --go_out=plugins=grpc:. lister.proto
//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I /usr/local/include -I . --go_out=plugins=grpc:. responsestate.proto
//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I /usr/local/include -I . --go_out=plugins=grpc:. signer.proto
//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I /usr/local/include -I . --go_out=plugins=grpc:. walletmanager.proto
