
ifdef SystemRoot
	DATE := $(shell powershell -noprofile '(Get-Date).ToUniversalTime().ToString("yyMMddHHmmss")')
else
	DATE := $(shell date -u '+%y%m%d%H%M%S')
endif

PROJECT_NAMESPACE := gitlab.com/ondrej.smola/ctcgrpc

GIT_HASH := $(shell git rev-parse --short HEAD)
PROJECT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
IMAGE_NAME := ondrejsmola/ctcgrpc:$(PROJECT_BRANCH)-$(GIT_HASH)-$(DATE)

.PHONY: install-proto-dependencies
install-proto-dependencies:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

build: export GOARCH=amd64
build: export GOOS=linux

.PHONY: build
build:
	go build

.PHONY: image
image: build
	docker build -f Dockerfile -t $(IMAGE_NAME) .

image-push: image
	docker push $(IMAGE_NAME)

.PHONY: proto
proto:
	protoc $(PROTOC_INCLUDE) --go_out=. --go-grpc_out=. --go-grpc_opt=module=$(PROJECT_NAMESPACE) --go_opt=module=$(PROJECT_NAMESPACE) proto/api.proto

