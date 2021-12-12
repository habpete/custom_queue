.PHONY: install
install:
	go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: get
get:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/golang/protobuf/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: .generate
.generate:
	mkdir pkg
	protoc --go_out=./pkg --go_opt=paths=source_relative \
    --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
	./proto/service.proto

.PHONY: build
build:
	go build -o ./bin/main ./cmd

.PHONY: docker-build
docker-build:
	docker build ./docker/