include apobuild.mk

PROJECT_NAME := elemental

clean: apoclean_vendor apoclean_apomock
init: apoinit
test: apotest
release:

ci: create_test_container run_test_container clean_test_container
