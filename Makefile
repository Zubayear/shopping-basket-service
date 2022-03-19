GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/ShoppingBasket.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o ShoppingBasket *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t ShoppingBasket:latest .

.PHONY: coupon
coupon:
	@protoc --proto_path=. --micro_out=. --go_out=:. coupon-client/Coupon.proto

.PHONY: gen_grpc_gateway
gen_grpc_gateway:
	@protoc --proto_path=./proto --micro_out=. --grpc-gateway_out=logtostderr=true,register_func_suffix=SB:. --openapiv2_out=./proto --openapiv2_opt=logtostderr=true --openapiv2_opt=use_go_templates=true --go_out=plugins=grpc:. ./proto/ShoppingBasket.proto

.PHONY: run_service
DEFAULT_PORT=60008
run_service:
	@go run main.go --server_address=localhost:$(DEFAULT_PORT)

.PHONY: run_gateway
run_gateway:
	@go run gateway/gateway.go