BUILD_VERSION=dev-snapshot

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(dir $(MAKEFILE_PATH))

all: build

.PHONEY: build
build:
	cd cmd && go build -ldflags="-X main.BuildVersion=$(BUILD_VERSION)" -o ../bin/kubehound kubehound/*.go

.PHONY: test
test:
	cd pkg && go test ./...

.PHONY: system-test
system-test:
	cd test/system && go test ./... 

.PHONY: local-cluster-create
local-cluster-setup:
	cd test/setup && bash setup-cluster.sh && bash create-cluster-resources.sh

.PHONY: local-cluster-destroy
local-cluster-destroy:
	cd test/setup && bash destroy-cluster.sh