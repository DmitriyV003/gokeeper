CLI_BIN := "./bin/cli"
SERVER_BIN := "./bin/server"

build-cli:
	go build -a -o $(CLI_BIN) -ldflags "$(LDFlAGS)" cmd/cli/main.go

build-server:
	go build -v -o $(SERVER_BIN) cmd/server/main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/proto/$(name).proto