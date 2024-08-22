.DEFAULT_GOAL := help

include scripts/make-rules/common.mk # make sure include common.mk at the first include line
include scripts/make-rules/all.mk

# ==============================================================================
# Usage

define USAGE_OPTIONS

\033[35mOptions:\033[0m
  BINS             The binaries to build. Default is all of cmd.
                   This option is available when using: make build/build.multiarch
                   Example: make build BINS="cc-apiserver cc-miner-controller"
endef
export USAGE_OPTIONS

.PHONY: protoc
protoc: ## Generate api proto files.
	$(MAKE) gen.protoc

.PHONY: build
build: tidy ## Build source code for host platform.
	$(MAKE) go.build

.PHONY: tidy
tidy:
	@$(GO) mod tidy

.PHONY: help
help: ## Display this help info.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<TARGETS> <OPTIONS>\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' Makefile #$(MAKEFILE_LIST)
	@echo "$$USAGE_OPTIONS"
