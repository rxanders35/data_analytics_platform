SHELL := /bin/bash

.PHONY: \
	dev-up dev-validate dev-down \
	build build-api build-controller build-apis \
	test test-api test-controller test-apis \
	vet vet-api vet-controller vet-apis

dev-up:
	@bash hack/dev/bootstrap.sh

dev-validate:
	@bash hack/dev/validate.sh

dev-down:
	@bash hack/dev/teardown.sh

build: build-apis build-api build-controller

build-apis:
	@cd controlplane/pkg/apis && go build ./...

build-api:
	@cd controlplane/api-server && go build ./...

build-controller:
	@cd controlplane/controller-manager && go build ./...

test: test-apis test-api test-controller

test-apis:
	@cd controlplane/pkg/apis && go test ./...

test-api:
	@cd controlplane/api-server && go test ./...

test-controller:
	@cd controlplane/controller-manager && go test ./...

vet: vet-apis vet-api vet-controller

vet-apis:
	@cd controlplane/pkg/apis && go vet ./...

vet-api:
	@cd controlplane/api-server && go vet ./...

vet-controller:
	@cd controlplane/controller-manager && go vet ./...
