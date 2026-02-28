SHELL := /bin/bash

.PHONY: dev-up dev-validate dev-down

dev-up:
	@bash hack/dev/bootstrap.sh

dev-validate:
	@bash hack/dev/validate.sh

dev-down:
	@bash hack/dev/teardown.sh
