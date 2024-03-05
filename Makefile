IMAGE_NAME := infoblox/atlas-gentool

GO_PATH                 := /go
SRCROOT_ON_HOST         := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
SRCROOT_IN_CONTAINER    := $(GO_PATH)/src/github.com/sbhagate-infoblox/user.app
IMAGE_VERSION           ?= $(shell git describe --tags)

.PHONY: test test-gen test-check test-clean
test: test-gen test-check test-clean

test-gen:
	docker run --rm -v $(SRCROOT_ON_HOST):$(SRCROOT_IN_CONTAINER) \
	 infoblox/atlas-gentool:latest \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=logtostderr=true:. \
	--validate_out="lang=go:." \
	--gorm_out=. \
	--atlas-query-validate_out=. \
	--atlas-validate_out=. \
	--preprocess_out=. \
	--doc_out=. --doc_opt=markdown,user.md,source_relative \
	--openapiv2_out=. github.com/sbhagate-infoblox/user.app/pb/user.proto

test-check:
	test -e pb/user.pb.go
	test -e pb/user_grpc.pb.go
	test -e pb/user.pb.gw.go
	test -e pb/user.pb.gorm.go
	test -e pb/user.pb.atlas.query.validate.go
	test -e pb/user.pb.atlas.validate.go
	test -e pb/user.pb.validate.go
	test -e pb/user.swagger.json
	test -e pb/user.md
	go mod vendor && go test ./pb

test-clean:
	find pb -type f -name '*pb*.go' -delete
	rm -f pb/*.json
	rm -f pb/*.md