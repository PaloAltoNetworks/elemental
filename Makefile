include domingo.mk

PROJECT_NAME := elemental

ci: domingo_contained_build

init: domingo_init
test: domingo_test
release:
