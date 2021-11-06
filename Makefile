deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
.PHONY: deps

grpc: schema/calculator.proto schema/authenticator.proto schema/user.proto
	protoc --proto_path=. --go-grpc_out=. --go_out=. $^
.PHONY: grpc

api:
	go run ./cmd/api/main.go
.PHONY: api

authenticator:
	go run ./cmd/authenticator/main.go
.PHONY: authenticator

user:
	go run ./cmd/user/main.go
.PHONY: user
