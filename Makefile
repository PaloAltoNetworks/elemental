include domingo.mk

init: domingo_init
test: domingo_test

build:
	@cd cmd/elegen && make build
