# generate client definitions
go:
    protoc --go_out=pkg/go --go_opt=paths=source_relative --go-grpc_out=pkg/go --go-grpc_opt=paths=source_relative -I ./src ./src/*.proto
