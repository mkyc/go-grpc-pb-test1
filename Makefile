
CMD := go
GET := $(CMD) get
BUILD := $(CMD) build
VET := $(CMD) vet
FMT := $(CMD) fmt
CLEAN := $(CMD) clean


BUILD_DIR := ./output/
SRV_DIR := /cmd/server
CLI_DIR := /cmd/client

APP_REPO := github.com/mkyc/go-grpc-pb-test1
SRV_APP_NAME := go-grpc-pb-test1-srv
CLI_APP_NAME := go-grpc-pb-test1-cli

all: clean get build

get:
	$(GET) -d -v ./...

build:
	$(VET) ./cmd/... ./pkg/...
	$(FMT) ./cmd/... ./pkg/...
	$(BUILD) -x -o $(BUILD_DIR)$(SRV_APP_NAME) $(APP_REPO)$(SRV_DIR)
	$(BUILD) -x -o $(BUILD_DIR)$(CLI_APP_NAME) $(APP_REPO)$(CLI_DIR)

clean:
	$(CLEAN) ./cmd/... ./pkg/...
	rm -rf $(BUILD_DIR)

proto:
	protoc --proto_path=api/proto/v1 --go_out=plugins=grpc:pkg/api/v1 communication-testing-service.proto