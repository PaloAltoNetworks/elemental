
package:
	mkdir -p ./static
	go-bindata -pkg static ./templates && mv bindata.go ./static

build: package
	go build

install: package
	go install
