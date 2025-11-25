PROTO_DIR := proto
PROTO_SRC := $(wildcard $(PROTO_DIR)/*.proto)
GO_OUT := .

.PHONY: generate-proto
generate-proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		$(PROTO_SRC)

.PHONY: startk8s
startk8s:
	kind create cluster --config=infra/kind/config.yaml

.PHONY: stopk8s
stopk8s:
	kind delete cluster --name microservices-go-starter
