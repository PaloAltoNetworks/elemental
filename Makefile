include domingo.mk

init: domingo_init
test: domingo_test

install_elegen:
	@cd cmd/elegen && make install
