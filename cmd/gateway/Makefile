.PHONY: run
run:
	go run main.go --registry etcd --registry_address fuxi.my:42379 --etcd_address fuxi.my:42379 api --handler=http

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	CGO_ENABLED=0 GOOS=linux go build -o micro main.go
	docker build . -t micro-gateway:latest
