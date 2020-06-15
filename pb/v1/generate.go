package v1

//go:generate protoc -I . --go_out=plugins=grpc:. dkg.proto
//go:generate protoc -I . --go_out=. eth2.proto
//go:generate protoc -I . --go_out=. endpoint.proto
//go:generate protoc -I . --go_out=. responsestate.proto
//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I . --go_out=plugins=grpc:. accountmanager.proto
//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I . --go_out=plugins=grpc:. lister.proto
//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I . --go_out=plugins=grpc:. signer.proto
//go:generate protoc -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I . --go_out=plugins=grpc:. walletmanager.proto
