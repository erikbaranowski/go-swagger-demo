.ONESHELL:
.DELETE_ON_ERROR:
export SHELL     := bash
export SHELLOPTS := pipefail:errexit
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rule

.PHONY: design
design:
	goa gen goa-demo/design

.PHONY: cmd
cmd:
	goa example goa-demo/design