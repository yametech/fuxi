mkdir -p bin
export GO111MODULE=on
rm -rf ./bin
# api
go build  -o ./bin/api-base ./cmd/base
go build  -o ./bin/api-gateway ./cmd/gateway
go build  -o ./bin/api-ops ./cmd/ops
# crd manager
go build  -o ./bin/crd-manager ./cmd/manager
# srv
go build  -o ./bin/srv-ops ./srv/ops
go build  -o ./bin/srv-ns ./srv/ns