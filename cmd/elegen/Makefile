include ../../domingo.mk

build: domingo_write_versions build_linux build_darwin build_windows package

package:
	mkdir -p ./static
	go-bindata -pkg static ./templates && mv bindata.go ./static


build_darwin: package domingo_build_darwin
	mkdir -p build/darwin
	mv ./elegen build/darwin

build_linux: package domingo_build_linux
	mkdir -p build/linux
	mv ./elegen build/linux

build_windows:
	mkdir -p build/windows
	echo 'echo disabled' > build/windows/elegen
