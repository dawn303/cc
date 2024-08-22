# ==============================================================================
# Makefile helper functions for generate necessary files
#

SERVICES ?= $(filter-out tools,$(foreach service,$(wildcard ${CC_ROOT}/cmd/*),$(notdir ${service})))

.PHONY: gen.protoc
gen.protoc: ## Generate go source files from protobuf files.
	@protoc \
		--proto_path=$(APIROOT) \
		--proto_path=$(CC_ROOT)/third_party \
		--go_out=paths=source_relative:$(APIROOT) \
		--go-http_out=paths=source_relative:$(APIROOT) \
		--go-grpc_out=paths=source_relative:$(APIROOT) \
		--go-errors_out=paths=source_relative:$(APIROOT) \
		--validate_out=paths=source_relative,lang=go:$(APIROOT) \
		$(shell find $(APIROOT) -name *.proto)

.PHONY: gen.appconfig
gen.appconfig: $(addprefix gen.appconfig., $(SERVICES)) ## Generate all application configuration files.

.PHONY: gen.appconfig.%
gen.appconfig.%: ## Generate specified application configuration file.
	$(eval CC_ENV_FILE ?= $(MANIFESTS_DIR)/env.local)
	$(eval GENERATED_SERVICE_DIR := $(OUTPUT_DIR)/appconfig)
	$(eval SERVICE := $(lastword $(subst ., ,$*)))
	@echo "===========> Generating $(SERVICE) configuration file"
	@$(SCRIPTS_DIR)/gen-service-config.sh $(SERVICE) $(CC_ENV_FILE) \
		$(CC_ROOT)/configs/appconfig/$(SERVICE).config.tmpl.yaml $(GENERATED_SERVICE_DIR)
ifeq ($(V),1)
	echo "DBG: Generating $(SERVICE) application configuration file at $(GENERATED_SERVICE_DIR)/$(SERVICE)"
endif
