PROTO_DIR := proto
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
GO_OUT := .

PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)

BINARY := microhal

.PHONY: proto check-tools run build clean

check-tools:
	@if [ -z "$(PROTOC_GEN_GO)" ]; then \
		echo "protoc-gen-go not found. Run: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; exit 1; \
	fi
	@if [ -z "$(PROTOC_GEN_GO_GRPC)" ]; then \
		echo "protoc-gen-go-grpc not found. Run: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; exit 1; \
	fi

proto: check-tools
	protoc \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		$(PROTO_FILES)

run:
	go run main.go

build:
	go build -o $(BINARY) main.go

clean:
	rm -f $(BINARY)
