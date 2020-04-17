export GO111MODULE=on

dir=`pwd`
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)

# find or download swag
# download swag if necessary
swag-install:
ifeq (, $(shell which swag))
	@{ \
	set -e ;\
	SWAG_INSTALL_TMP_DIR=$$(mktemp -d) ;\
	cd $$SWAG_INSTALL_TMP_DIR ;\
	go mod init tmp ;\
	go get -u github.com/swaggo/swag/cmd/swag;\
	rm -rf $$SWAG_INSTALL_TMP_DIR ;\
	}
SWAG=$(GOPATH)/bin/swag
else
SWAG=$(shell which swag)
endif

gen-openapi: swag-install
	chmod +x ./scripts/gen-openapi.sh
	sh ./scripts/gen-openapi.sh $(SWAG)


test-dev: crd sdk-up

debug-crd:
	chmod +x ./scripts/debug-crd.sh
	sh ./scripts/debug-crd.sh

proto:
	for d in proto; do \
		for f in $$d/**/*.proto; do \
			protoc -I. --micro_out=. --go_out=. $$f; \
			docker run --rm -v ${PWD}:${PWD} -w ${PWD} znly/protoc --gogofast_out=plugins=grpc:. -I. $$f;\
			echo compiled: $$f; \
		done \
	done

check:
	chmod +x ./build/lint.sh
	sh ./build/lint.sh

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

dep:
	go mod vendor

build: dep check
	chmod +x ./scripts/build.sh
	sh ./scripts/build.sh
clean:
	rm -rf bin/*

run:
	docker-compose build
	docker-compose up

uninstall-crd:
	kubectl delete -f deploy/crds/

deploy: generate
	kubectl apply -f deploy/crds/

generate:
	operator-sdk generate k8s
	operator-sdk generate crds

up: generate
	operator-sdk up local

gen-crd-client:
	hack/generate-groups.sh all github.com/yametech/fuxi/pkg/client github.com/yametech/fuxi/pkg/apis fuxi:v1
