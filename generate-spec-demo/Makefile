.ONESHELL:
.DELETE_ON_ERROR:
export SHELL     := bash
export SHELLOPTS := pipefail:errexit
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rule

.PHONY: spec
spec:
	swagger generate spec -o ./swagger.yaml
	swagger generate spec -o ./swagger.json