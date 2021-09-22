package v1

//go:generate protoc -I . --go_out=. --go_opt=paths=source_relative eth2.proto
//go:generate protoc -I . --go_out=. --go_opt=paths=source_relative endpoint.proto
//go:generate protoc -I . --go_out=. --go_opt=paths=source_relative responsestate.proto
//go:generate protoc -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative dkg.proto
//go:generate protoc -I ../../third-party -I . --go_out=:. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative accountmanager.proto
//go:generate protoc -I ../../third-party -I . --go_out=:. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative lister.proto
//go:generate protoc -I ../../third-party -I . --go_out=:. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative signer.proto
//go:generate protoc -I ../../third-party -I . --go_out=:. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative walletmanager.proto
