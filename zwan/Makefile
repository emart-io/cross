# Variables
SERVICE=zwan
IMG_HUB?=docker.io/jmzwcn
PROTOC=PATH=${shell go env GOPATH}/bin:$(PATH) ${shell which protoc}
# Version information
VERSION=1.0.0
REVISION=${shell git rev-parse --short HEAD}
RELEASE=production
BUILD_HASH=${shell git rev-parse HEAD}
BUILD_TIME=${shell date "+%Y-%m-%d@%H:%M:%SZ%z"}
LD_FLAGS:=-X main.Version=$(VERSION) -X main.Revision=$(REVISION) -X main.Release=$(RELEASE) -X main.BuildHash=$(BUILD_HASH) -X main.BuildTime=$(BUILD_TIME)
TAG?=${shell date "+%Y%m%d-%H%M"}

ifeq (${shell uname -s}, Darwin)
	SED=gsed
else
	SED=sed
endif

prepare:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
#	@-docker swarm init
#	@-docker network create --driver=overlay devel

generate: generate-go generate-js

generate-go:
	@$(PROTOC) -I./service --go_out=. --go-grpc_out=. service/*.proto
	@echo Generate-go completely.

# generate-go-v1:
# 	@$(PROTOC) -I./service --gogofaster_out=\
# 	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
# 	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
# 	Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
# 	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
# 	Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
# 	Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
# 	plugins=grpc:./service/go service/*.proto
# 	@$(SED) -i '/google\/api/d' service/go/*.pb.go
# 	@echo Generate-go successfully.

generate-js:
	@$(PROTOC) -I./service service/*.proto \
	--js_out=import_style=commonjs:service/js \
	--grpc-web_out=import_style=commonjs+dts,mode=grpcweb:service/js
	cp -rf service/js/* ../pwa/src/sdk
	@echo Generate-js completely.

generate-descriptor:
	@$(PROTOC) -I./service --include_imports --include_source_info \
	--descriptor_set_out=service/descriptor.pb service/*.proto
	@echo Generate-descriptor completely.

build:generate
	CGO_ENABLED=0 go build -ldflags='$(LD_FLAGS)' -o bundles/$(SERVICE) internal/cmd/*.go

image:build
	docker build -t $(IMG_HUB)/$(SERVICE):$(TAG) .

push:image
	docker push $(IMG_HUB)/$(SERVICE):$(TAG)

run:image
	-docker service rm $(SERVICE) > /dev/null 2>&1  || true	
	@docker service create --name $(SERVICE) --network devel --mount type=bind,source=/home/daniel/.emart_uploads,destination=/uploads $(IMG_HUB)/$(SERVICE):$(TAG)

envoy:generate-descriptor
	docker build -t $(IMG_HUB)/envoy:$(TAG) -f envoy.Dockerfile .
	docker push $(IMG_HUB)/envoy:$(TAG)
#	docker service create --name envoy --network devel -p 80:80 $(IMG_HUB)/envoy:$(TAG)

mysql:
	-docker service create --name mysql_emart --network devel --mount type=bind,source=/home/daniel/.emart_mysqldata,destination=/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=emart mysql:5.7.24

test:
	go test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt