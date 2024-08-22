#
# These variables should not need tweaking.
#

# ==============================================================================
# Includes

# include the common make file
ifeq ($(origin CC_ROOT),undefined)
CC_ROOT :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
endif

# include $(CC_ROOT)/scripts/make-rules/common-versions.mk

# ==============================================================================
# Build options
#
PRJ_SRC_PATH :=github.com/dawn303/cc

ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(CC_ROOT)/_output
$(shell mkdir -p $(OUTPUT_DIR))
endif

# The OS must be linux when building docker images
PLATFORMS ?= linux_amd64 linux_arm64
# The OS can be linux/windows/darwin when building binaries
# PLATFORMS ?= darwin_amd64 windows_amd64 linux_amd64 linux_arm64

# Set a specific PLATFORM
ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
	# Use linux as the default OS when building images
	IMAGE_PLAT := linux_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
	IMAGE_PLAT := $(PLATFORM)
endif

MANIFESTS_DIR=$(CC_ROOT)/manifests
SCRIPTS_DIR=$(CC_ROOT)/scripts
APIROOT ?= $(CC_ROOT)/pkg/api
