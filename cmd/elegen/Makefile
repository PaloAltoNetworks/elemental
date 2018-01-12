include ../../domingo.mk

build: build_linux build_darwin build_windows package

package:
	mkdir -p ./static
	go-bindata -pkg static ./templates && mv bindata.go ./static


build_darwin: package domingo_build_darwin
	mkdir -p build/darwin
	mv ./elegen build/darwin

build_linux: package domingo_build_linux
	mkdir -p build/linux
	mv ./elegen build/linux

build_windows: package domingo_build_windows
	mkdir -p build/windows
	mv ./elegen.exe build/windows
