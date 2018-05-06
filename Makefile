PWD := $(shell pwd)
APP := landlady
PKG := github.com/funxdata/$(APP)
CMS_PKG := ${PKG}/vendor/github.com/ckeyer/commons
GO := CGO_ENABLED=0 go
PROTOC := protoc
HASH := $(shell which sha1sum || which shasum)

OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)
VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_AT := $(shell date "+%Y-%m-%dT%H:%M:%SZ%z")
PACKAGE_NAME := $(APP)$(VERSION).$(OS)-$(ARCH)

LD_FLAGS := -X $(CMS_PKG)/version.version=$(VERSION) \
 -X $(CMS_PKG)/version.gitCommit=$(GIT_COMMIT) \
 -X $(CMS_PKG)/version.buildAt=$(BUILD_AT) -w

BUILD_IMAGE := ckeyer/go
IMAGE_NAME := ckeyer/$(APP):$(VERSION)

build:
	$(GO) build -v -ldflags="$(LD_FLAGS)" -o bundles/${APP} .

protoc:
	which protoc-gen-gogofast || ${GO} get github.com/gogo/protobuf/protoc-gen-gogofast
	${PROTOC} -I./proto \
	-I./vendor \
	--gogofast_out=\
	Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
	Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
	plugins=grpc:./proto proto/*.proto

run:
	 ${GO} run ./main.go

image:
	docker build -t $(IMAGE_NAME) .

test:
	${GO} test -ldflags="$(LD_FLAGS)" $$(go list ./... |grep -v "vendor")

release: clean build
	mkdir -p bundles/$(PACKAGE_NAME)
	mv bundles/$(APP) bundles/$(PACKAGE_NAME)
	cd bundles ;\
	 echo $(VERSION) > $(PACKAGE_NAME)/release.txt ;\
	 $(HASH) $(PACKAGE_NAME)/$(APP) > $(PACKAGE_NAME)/sha1.txt ;\
	 tar zcvf $(PACKAGE_NAME).tar.gz $(PACKAGE_NAME);

clean:
	rm -rf bundles/*

dev-server:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8089:8080 \
	 -v $(PWD):/go/src/$(PKG) \
	 -w /go/src/$(PKG) \
	 $(BUILD_IMAGE) bash

dev-client:
	docker run --rm -it \
	 --name $(APP)-dev-client \
	 -v /var/run/docker.sock:/var/run/docker.sock \
	 -v $(PWD):/go/src/$(PKG) \
	 -w /go/src/$(PKG) \
	 $(BUILD_IMAGE) bash
