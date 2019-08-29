# This version-strategy uses git tags to set the version string
# VERSION := $(shell git describe --tags --always --dirty)

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

SWAGGER_YML=swagger.yaml

generate-client: ## [SWAGGER] Generate swagger binary
	oapi-codegen $(SWAGGER_YML) -generate client  > client.go

build: generate-client