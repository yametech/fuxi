
GOPATH:=$(shell go env GOPATH)

.PHONY: run

run:
	go run *.go --registry etcd --registry_address fuxi.io:12379 --etcd_address fuxi.io:12379

swag:
	swag init

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/base/base.proto

.PHONY: build
build: proto
	go build -o base-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t base-srv:latest


