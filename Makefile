GO_OUT = pkg

.PHONY: all
all:
	protoc \
		--go_out=$(GO_OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT) --go-grpc_opt=paths=source_relative \
		$$(find api -type f -name *.proto)

	go fmt ./...
	go mod tidy
