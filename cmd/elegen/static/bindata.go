// Code generated by go-bindata.
// sources:
// templates/README.md
// templates/identities_registry.gotpl
// templates/model.gotpl
// templates/relationships_registry.gotpl
// DO NOT EDIT!

package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz) // #nosec
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xcb\x11\x02\x21\x10\x45\xd1\x7d\x47\xf1\x2c\x53\x22\x81\xc6\xbe\x0a\xc5\x47\x6a\x60\x33\xd9\xcf\x79\x2b\x31\x56\xf7\xc3\x36\x4b\x85\x8d\xbe\xb5\xa3\x09\xa1\xf3\x57\x46\x8c\x4c\x04\xa1\x3a\x75\x0a\xca\x75\xfa\x75\xbf\xcc\x24\x69\x78\x43\xcb\x3f\xcd\x7f\xd8\x13\x00\x00\xff\xff\xaa\x97\xff\x85\x4d\x00\x00\x00")

func templatesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesReadmeMd,
		"templates/README.md",
	)
}

func templatesReadmeMd() (*asset, error) {
	bytes, err := templatesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1625075149, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\xdf\x6b\xe3\x38\x10\x7e\xf7\x5f\x31\x57\xca\x61\x43\xea\xbe\x1c\xf7\xd0\xa3\x0f\x65\xd9\x85\xc0\x75\x29\x2d\xdc\x4b\xe8\x83\xd6\x19\xa7\xe2\x64\xc9\x48\x72\x77\x83\xd1\xff\xbe\x48\x96\x7f\xc8\x71\x52\xa7\xed\x66\xa1\x7d\x69\x22\x69\xe6\x1b\x7d\xa3\xf9\x34\x4a\x49\xb2\xff\xc9\x06\xa1\xae\x21\x7d\x40\x9d\x7e\x12\x3c\xa7\x9b\x4a\x12\x4d\x05\x4f\xbf\x92\x02\xc1\x98\x28\xa2\x45\x29\xa4\x86\xb3\x8d\x48\x49\x29\x24\x6a\x91\x52\x71\x89\x0c\x0b\xe4\x9a\xb0\xb3\x28\x7a\x26\x12\xe2\x08\x00\x80\xae\x91\x6b\xaa\xb7\xd6\x58\xdd\x92\x12\xae\xa1\x20\xe5\x4a\x69\x49\xf9\xe6\xb1\xb3\x49\x97\x7e\x1d\xd4\xce\xcc\xfe\xd5\xf5\x05\x48\xc2\x37\xd8\x04\xf3\x50\x62\x46\x73\x9a\xb9\x60\x94\x0d\xa4\x5f\x08\x34\x07\xf5\x24\x2a\xb6\xbe\xc7\x0d\x55\x1a\x65\xb0\x1a\x52\x38\x4f\xef\xaa\x6f\x8c\x66\xb7\x62\x8d\xa1\xed\x05\x9c\xf7\x21\xc2\xd5\x35\xa4\x76\x0d\x4b\x3f\xf7\x83\x17\x03\x83\xb3\xba\xf6\x0b\xee\x51\x69\x3b\x6d\xcc\xd9\x95\x8d\x61\xe8\xc6\x98\x76\x43\x8b\x00\x0a\xf9\x7a\x8c\x3e\x18\x32\x51\xc0\x59\x46\x34\x6e\x84\xa4\x1f\x90\x38\x51\xc9\x0c\x7f\x09\x79\x84\x51\xa2\x7e\x21\x63\x17\x27\xa4\xac\xae\x7d\x54\xe7\x74\x01\xe7\x6e\x67\x03\xa3\x9b\x66\xa7\x10\x72\xdc\xae\x7b\x3f\x62\x0f\x1c\x54\xbe\xc6\x1f\x13\x5c\xaf\x1e\x57\x8f\xcd\xc7\x93\x73\x6c\x19\x18\xd5\x67\x47\x05\xcd\x81\x21\x87\x74\xd9\x84\x0d\xc6\x0c\xc2\x0b\x43\x74\x84\x67\xa2\x28\x45\xc5\xd7\x8e\xf3\xde\x28\x34\x19\x1a\x75\x06\xc6\xb8\x38\xec\xff\x45\x4f\x1d\x98\x83\xb4\x9b\x45\x5d\x03\x32\x65\x03\xe6\x94\x2d\x5e\x99\x99\x24\x8a\x2e\x2f\xc1\x51\xf0\x1f\x4a\x65\xe9\x92\xa8\x2b\xc9\x15\xe8\x27\x84\xac\x92\x12\xb9\x86\x67\x3f\x27\x72\x37\x5c\x38\xca\xa2\xbc\xe2\x59\x60\x1b\x27\x90\x33\x41\xf4\xdf\x7f\x41\xed\xfd\x74\xd7\xc3\xcd\xdd\x72\xc9\x73\x91\xb6\x30\x76\x87\x51\xa4\xb7\xa5\x77\x77\x4b\x38\xd9\xa0\x04\xa5\x65\x95\xe9\xda\x44\xce\x7d\x9c\x07\xb3\x09\xb4\x67\xf2\x8b\x14\x85\xcd\x57\xcc\x6d\xd2\x9a\xf3\x93\xc0\x64\xdd\xba\xad\xfa\x68\xc6\x37\xcc\xca\x9a\x3f\x46\x73\xd0\x3e\x35\x02\xbb\x8d\xbd\xd2\x6e\x8f\x47\x0d\x34\x7a\xd5\xfa\x99\x07\xef\x0a\x38\x6e\xca\x75\x36\x70\xaf\x6f\x2b\xf7\x71\x26\x14\xdf\xc6\x84\xf7\xfb\x8b\xe9\x04\x52\xd2\x42\xd1\x1c\x28\x5c\x43\x9e\xee\xa4\x86\xf0\x6d\xf2\x0f\xfc\x41\xd3\xa5\xfa\x5c\x94\x7a\x1b\x27\x83\x12\x6a\xa9\x09\x24\x62\xca\x55\xc7\xfb\xd1\xee\xfc\x58\xe8\xce\xf3\xc8\xb7\xc9\x0b\x5c\xe4\x94\x7c\x63\x18\xb7\xb9\x9b\xa4\x60\x3c\xd6\xd8\xb4\xcc\xa8\xef\x54\x67\x4f\x5d\xf6\x7d\xb4\x9d\x4e\x1f\x50\xb6\x57\xab\x5a\x46\x54\xd3\x91\xed\x5c\x15\xbd\x9c\x5f\x8d\x49\xfb\x8a\xdf\xf7\x98\xc4\x49\x34\x21\x1b\xa3\xaf\x6b\xcc\x49\xc5\xf4\x8e\x5b\x4e\x99\xcf\xc6\x3e\xa2\x1f\x4a\x22\x15\xbe\x86\xee\x5d\xcb\xdf\x48\xba\x37\xe4\x42\xb7\x24\x2e\xd5\xbd\x10\xfa\xad\x49\x69\x36\xf9\x96\xd4\xbc\x5b\xa6\xfc\x85\x76\x38\x3d\xc1\x4d\x1e\xe8\x5f\x77\xf5\xaf\x5a\x07\xee\x81\xf0\x92\x1c\x35\x99\xb5\x55\xfb\xe0\xdc\x06\xaa\x74\xb8\xf6\x46\xb5\xef\x4f\xd7\x48\x0a\x1a\x9d\x4b\xe6\x29\xc1\x0b\x9b\x9f\x0e\x47\xbd\xdb\xb1\xdc\x7b\xba\x4e\xab\x15\x7f\x4e\x19\xdc\xb1\x4a\x12\x06\xc6\xfc\x4b\x95\xbd\xba\x4f\x78\x30\x77\x85\x60\x76\x9e\x26\x4c\x3f\x5c\xb6\xf6\x4b\xc8\x6f\xcc\x59\x40\xf9\x51\xd5\xad\x0e\x96\xb7\x3a\xba\xbe\xef\x91\x35\xe9\x7b\xa2\xa5\x8a\x87\xa8\xc1\x4c\x93\x27\x39\xee\xae\xe4\xd4\x9a\x16\x0b\x76\xc0\x6e\x18\xf3\xd1\x51\xb4\x60\xab\xfd\xcf\x4c\x0f\x30\xb2\xb0\xae\x9f\x89\x84\xc2\xb7\xca\xd7\x01\x80\x6d\x99\x6d\x2f\xef\x27\x87\x6d\xbc\x5b\x36\xd8\xdc\xed\xc0\xac\x6d\xe3\x9b\x6f\x01\x07\xc3\x65\x7d\x43\xdf\xa2\x9b\xc8\xe1\x05\x31\x76\xa8\x84\x31\xc0\x1f\x54\x69\x7b\x19\xd0\x6e\xde\x83\xcd\x64\x62\x48\xc5\xd4\x92\xd3\x3f\xc9\x0f\x56\xdf\x71\xbf\x3f\x18\xcf\x9e\x6b\xd0\xbf\x08\xd9\xed\x7b\x48\xa1\x4d\x9e\xef\xe1\x21\x17\xd2\x7d\xdf\xd0\x67\xec\x9f\x14\x1d\xa3\x63\x3f\x2f\x5d\xd5\xe1\x45\xbd\x4f\xf0\x66\xd0\x7a\x5a\x45\x6b\x03\xaf\xdf\xfa\x7b\xc7\xe1\xd7\xf5\x0c\xfd\x0b\x1f\x19\x56\xf5\x4c\xf4\x33\x00\x00\xff\xff\x03\x1b\x0d\x41\x14\x15\x00\x00")

func templatesIdentities_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesIdentities_registryGotpl,
		"templates/identities_registry.gotpl",
	)
}

func templatesIdentities_registryGotpl() (*asset, error) {
	bytes, err := templatesIdentities_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 5396, mode: os.FileMode(420), modTime: time.Unix(1625075149, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x69\x73\xdb\x38\x96\x9f\xa3\x5f\x81\x56\xb9\x13\x2a\x23\xd3\xf9\xac\x1e\x4f\x55\x62\x27\xbd\xae\xcd\xe1\x8a\x33\x3d\x55\x9b\x4a\x8d\x61\x12\x94\xd1\xa1\x00\x36\x00\xfa\x58\x15\xff\xfb\x16\x0e\x92\x00\x2f\x81\xb2\x94\x76\xef\x58\x1f\x12\x09\xc7\xc3\xbb\x00\xbc\xf7\x00\x3c\x67\x30\xfa\x0e\x97\x08\xac\xd7\x20\xbc\x40\x22\x3c\xa1\x24\xc1\xcb\x9c\x41\x81\x29\x09\x3f\xc2\x15\x02\x45\x31\x99\xe0\x55\x46\x99\x00\xc1\x04\x00\x00\xa6\xc9\x4a\x4c\xf5\xb7\x25\x0d\x61\x46\x19\x12\x34\xc4\xf4\x08\xa5\x68\x85\x88\x80\x69\x59\x8b\xc5\x75\x7e\x15\x46\x74\x75\xb4\x4c\xe9\x15\x4c\x39\x5e\x92\xa3\xd5\x92\x1e\x5d\x71\x4a\xda\x8d\x56\x58\x44\xd7\x28\x4d\xaf\x8f\x22\x9a\xdd\x73\xc1\xf2\x48\xe4\x0c\xe9\x86\xeb\xf5\x21\x60\x90\x2c\x11\x08\x2f\x32\x14\x85\x5f\xee\x33\x74\xce\xe8\x0d\x8e\x11\xe3\x12\x49\x05\x4d\xd2\x01\x8a\xa2\xee\x82\x48\x0c\x0e\x4d\x6d\x13\xc4\x6f\x30\xc5\xb1\xa2\xd4\x13\x50\x51\x4c\x66\x93\xc9\x7a\x0d\x0e\x52\x28\x10\x17\xbf\x21\xc6\x31\x25\x60\x71\x6c\x20\xbe\x57\xc5\xaf\x85\x60\xf8\x2a\x17\x88\x97\x0d\x24\x0f\xd7\x6b\x33\xf8\x01\x9e\x83\x03\x44\xf2\x95\xec\x77\x95\xe3\x34\x7e\x4b\xf2\x15\xd7\x20\x9a\xa0\x8b\x62\x72\x74\x24\xc5\xa3\x7a\x28\xaa\x41\x51\x00\x86\x32\x86\x38\x22\x82\x03\x71\x8d\x40\x46\x39\xc7\x57\x29\x02\x37\x30\xcd\x11\x07\x09\x65\x00\x96\x58\x28\x62\x74\xf7\x0a\x33\x23\xd9\x69\x38\x11\x12\x62\x0b\x3e\x17\x0c\x93\xe5\x64\x12\x51\xc2\x4b\xb9\xd7\xec\x3b\x20\x70\x85\xe6\xe0\x40\x8d\x26\xa9\xd0\x9d\x7f\xd3\x83\x1b\x16\x1a\xb4\x89\x1e\xa9\x89\xb1\xee\x2a\x1b\xe8\x6f\x45\x11\x9a\x41\xea\x2e\x2d\xac\x8e\x35\x29\x65\x8f\x52\x38\xb5\x6c\xea\xef\x13\x89\x2d\x4e\x00\xa1\xc2\xc8\xe6\x03\x8d\x51\x1a\x9e\x22\x01\xa3\x6b\x14\xd7\x8c\xb5\x6b\xdf\x12\x81\xc5\xbd\x61\xce\x59\x8c\xd4\xcf\x26\xea\x55\x39\x4d\xd4\x6f\x7a\xf5\x3b\x8a\x44\x38\xb9\x81\xcc\x0f\xde\x31\xa8\x66\x4a\x58\x15\xae\x15\x31\xb2\xe9\x02\x54\x1a\x68\x81\xfa\x8c\xb8\x90\xb5\x45\x31\x9d\xab\xa6\x27\x50\xa0\x25\x65\xf7\x8b\xae\xa6\x34\x67\x51\x25\x64\xdd\xfe\x5c\x4f\xf5\x45\x1b\xb4\xa9\xa9\x5b\x32\x7c\x03\x85\x6c\xd9\x6c\xa8\x2b\x8a\x62\x3e\x29\x46\xf2\x7a\xbd\xee\x6a\x71\xc6\x3f\x53\x2a\x36\xc9\xe2\x3c\xcd\x19\x4c\x41\x51\xbc\xc7\x5c\xd8\xd2\x80\x20\x95\x25\x34\xf1\xe8\x5b\x29\xba\xcf\x18\x5f\xbf\xbd\xec\x6d\x29\x09\x3e\x3a\x02\x96\x76\x88\x9c\x11\xad\x1a\xb8\x53\x35\x38\xc0\x44\xfd\x94\xd8\x86\x93\x24\x27\x11\x08\xa8\x27\x2e\xb3\x6a\xa4\x60\xd6\xad\x37\x4a\x66\x1a\x8b\x7e\x98\xb5\xfa\x4d\x34\xfe\x27\x34\xab\x71\x87\x20\xa3\x98\x08\xc4\x80\xa0\x00\x02\xb9\xfc\x2a\x84\xfd\x50\x1c\x4f\x92\x1c\xbc\x83\x9c\x04\xc3\xab\x14\xf1\x92\x26\x85\xc6\xe2\x18\xc0\x2c\x43\x24\x0e\xfc\x80\xaf\x8b\x39\xa0\x61\x18\xce\x6c\xb6\x3c\x97\xa0\x0c\xe1\xaf\x15\x34\x03\x94\x3b\x62\x12\x54\xfd\x84\x80\xa0\x5b\x3d\xba\x91\xe3\xbe\xf8\xa0\x71\x09\xca\xf1\xc3\x30\xec\x66\xc9\x46\x56\xd1\x5c\x3c\x90\x53\x72\xcb\xf8\xf7\x5c\xb2\x42\x02\xd2\xeb\x7c\x89\x97\x5e\x9b\xca\x71\xaa\x61\x68\x2e\x54\x87\x30\x18\x9a\x2d\x33\x0d\xbf\x70\xf4\x94\xe6\xc2\x88\x43\xcd\xb7\x88\x92\x1b\xc4\x84\x2d\x0d\xa5\x89\xa4\x8f\xee\xed\xd8\x2d\xff\xed\x57\x3b\x85\x89\xcb\xcf\x15\xfc\x8e\x82\x81\xe6\x73\x90\x22\x12\xd0\x59\xcd\x42\x2c\xbb\xbd\xfa\x05\x60\xf0\x77\x53\xf7\x0b\xc0\x7f\xfb\x9b\xcb\xc2\xaf\xf8\x1b\x38\x06\xf4\x2b\xfe\x36\xc8\x9a\x53\x94\xc0\x3c\x15\x9f\x58\x8c\x98\xb3\xcc\xc4\xba\x02\x50\x59\x83\xc9\x12\x24\x18\xa5\x31\x2f\xb5\x35\xa2\x44\x20\xb2\x05\x7f\xec\x01\x83\x19\xf8\xfa\x4d\x9b\x01\x8d\x35\xa6\x2c\xae\x49\xaa\x4c\x1b\x3d\x8a\x83\x77\x69\x7c\x39\x56\xd5\xdc\xee\x6a\x76\x11\xcd\x09\x4d\xf9\x17\x7a\x91\x41\xc6\x91\x43\xb5\xe7\xda\x6d\x74\x09\xc5\x52\x83\x34\x18\xdf\xe9\x7b\x74\x04\x3e\xb5\x57\x6c\x70\x8b\xd3\x14\x50\x92\xde\x2b\xce\x42\x53\xb5\xc4\x37\x88\x18\xce\x87\xe0\x23\xd5\x5f\xc1\x0a\x41\xc2\x81\xd4\x13\x86\x4c\x11\x47\x5b\xc8\xa2\x64\x41\x60\x64\x1b\x86\xa1\x66\xbb\xef\x5a\xa0\x74\x77\x0c\xfd\x0f\x56\xe6\xb0\x81\xb3\x5c\x5b\xc2\xe0\xe5\x06\x1c\x40\x51\x0c\xaf\x10\xa5\x29\x6c\xeb\xc2\x8d\x29\x7b\xa8\xc6\x1b\xd8\xc1\x0c\x60\x22\xba\xf6\x52\x24\xc2\xd7\xe7\x67\x67\x24\xa1\xa1\x65\x92\x6b\x73\xde\x28\xae\xe5\x1d\xa8\xef\x07\x98\xbf\x25\x11\xbb\xcf\x84\x14\x8b\x64\x61\x02\x53\x8e\x3c\x2c\xce\xa6\xa5\xb9\x92\x4d\x24\x8d\xb0\xd9\xad\xb4\x06\xcb\xf1\xf5\xf4\x53\xed\x4f\xe8\x4a\x6a\xc7\xbb\x14\x2e\xb9\x3b\xd4\x9d\x40\x44\x52\xc0\x2d\x54\x1a\x04\x0c\x1b\x49\xc6\x33\xc8\x23\xc9\xaa\xa6\x37\x55\x7b\x3d\x4d\x1f\xe6\xb0\x28\x8c\x85\x18\x1a\xc6\x28\x9b\xb0\x83\x57\xc7\x40\x30\x65\xda\x5b\x28\xad\xd7\xca\x95\xf9\x42\xdf\xa9\xc9\x74\x20\x65\x62\x38\xda\xc4\x5e\x49\x4f\x21\x5d\x0e\x2d\xa5\x7a\xf9\x3b\xa7\x64\x31\x3d\x9c\x82\x15\x5f\x4a\x87\x57\x7d\xbf\x52\x85\xff\x56\x2c\x33\xda\x34\xbd\x34\x1a\xf7\x11\xdd\x6e\x10\x53\x69\x36\x49\x43\xa1\x7f\xf3\x93\x38\x29\x8d\xdc\x00\x30\x98\x0d\x03\x69\x28\xe6\xf3\xa1\xb6\xf5\xdc\xb4\x19\xb1\x18\xd0\xe6\xf9\xc4\x5a\x8d\x0f\x6d\x27\x55\xf2\x5d\xea\x2f\xa7\xcc\x72\x6a\x41\xb0\x41\xe0\x33\x67\xd1\x37\xa2\xe7\xd7\x34\x4f\xe3\x7f\x31\x2c\xd0\x19\xc1\x02\xc3\x14\xff\x2f\x62\x52\x9c\xca\xeb\x95\x43\xe9\x78\x43\x43\x79\x0e\xc2\xf3\xfc\x2a\xc5\x91\xa4\x06\x34\xc0\x1e\x60\x82\xd5\x5a\x77\xdb\x01\x16\x09\x07\x78\xb3\x2f\x4e\x4c\x77\xa7\xbc\xab\xec\xd0\xde\xa1\xfc\x8a\xcc\x0a\xe1\xe3\x81\x76\xfa\x11\x7d\x2e\x66\xb9\xbc\x0d\x6a\xcb\x8e\x3c\x06\xe0\xb8\x0c\xcd\x45\xce\x97\xb0\x04\x37\x6c\x17\xed\xbc\x3b\x74\xbd\xe0\x20\x27\xf8\x8f\xbc\xf4\x9f\x64\x9f\x91\xb4\xca\x2e\xc1\x0c\xb8\xf6\x8a\xf6\x39\x75\x5f\x0b\x9b\x52\x39\xcb\x8d\x26\xac\x06\xa8\x1b\x85\x27\xa5\x15\x51\xce\xe3\x4a\xca\x72\xe1\x69\x80\x98\xb6\xc2\x44\xdb\x30\xec\x02\x09\x0b\x4b\x8e\xc4\x7e\x18\xe6\x0c\x13\xe0\x18\x94\x66\x85\x1f\xd7\xfc\xd8\x05\x8e\x01\x8e\x87\x99\x72\x74\x04\x7e\x45\xe2\xcd\xc5\xa7\x8f\x00\xaf\x32\xad\xa6\x9a\x62\xb9\x34\x83\x15\x64\xfc\x1a\xa6\x52\x9c\xca\x33\x4d\x60\x84\x94\x85\xf6\xe5\x1a\x73\x80\x39\xc8\xb9\x36\xf1\x04\x83\x84\x67\x90\x21\x22\xb4\x85\x26\x11\x01\x67\xa7\xb2\xee\x03\x25\x4b\x7a\xfa\xe6\xec\x14\x40\x0e\x3e\x5d\xa1\x48\x9c\x9d\x7a\x33\xca\x60\x17\xcc\x40\x50\x61\x20\x7d\x26\xc4\x18\x65\x15\xbb\x70\x02\x28\x38\x3e\x06\x04\xa7\x96\x5d\x64\x14\x83\xe0\x74\x2e\xff\xb1\xed\x1b\x2e\x17\xac\xe7\x2b\x89\x59\xbd\x82\x0e\xae\xe8\xa5\xf2\xf9\x6f\xb7\x13\x6b\x95\x0b\x2f\x04\x65\x28\x06\x8d\x52\x4b\xb4\xa6\x46\x52\xa2\x84\xdb\x12\xe6\x4f\xc7\x60\x3a\xb5\xa8\xe3\xdd\xcd\x8e\x95\xe4\x42\x6d\x42\x9f\xc5\xff\x85\xee\x82\x6e\x80\xa5\xbd\xe7\xcc\x29\x83\x45\x2f\xec\x6e\x50\x4d\x1d\x1b\xfe\x69\x4f\x5a\xae\x25\xa3\x35\xf1\xe2\x51\x6b\xa2\xc1\x2e\x60\xf0\x56\xb3\xf8\x33\xbc\x9d\x69\x3d\xf4\x55\xc3\x5d\x68\x20\x4e\xe4\x98\x3a\x3c\x70\x1b\xfe\x93\x18\xc6\x04\x7c\xf6\x8b\xaa\xf8\xa9\x67\x78\xc4\x98\x23\xf0\x3d\xeb\x71\x8f\x12\x1f\xf7\xa8\x56\x28\xf5\x74\xd6\xa9\x8b\x23\x21\x6d\xaf\x8b\x46\x11\xfd\xf6\x89\x2e\x7f\xe8\x1a\xb2\x38\xa2\x31\x8a\x9b\x9e\x91\xb2\x6f\xbd\x15\x6d\x6b\x77\xa8\xb1\xb0\xbf\x49\xd1\x0d\x52\x51\xfb\xe6\x84\x92\x15\xe1\x49\x0a\x39\xd7\x32\x3b\xab\x67\x94\x27\x8e\x15\xec\xd6\x7e\x5f\xee\xc6\xfd\xbe\xd2\xd4\x9f\xcb\xbd\xc1\x97\x32\xe4\xdc\x13\x84\xf1\xa2\x43\x11\xf2\x28\xa2\x2d\x23\xcd\x14\xdb\x38\x28\x6b\x79\xc4\x70\x26\xea\xc3\xaa\x53\x1a\xb9\xd1\x2a\x1a\xe5\xca\x06\x55\x6d\x12\xca\x2c\x4b\xc6\x57\xe8\xa7\x34\xea\x13\xf7\xa5\x24\x8a\x47\x6f\x60\xf4\x5d\xe0\xe8\x3b\x1f\xc0\xee\xd2\x39\xb6\x18\x49\xba\xef\x5a\xad\x70\xec\x43\x36\x59\x89\xf0\x22\x63\x98\x88\x24\x98\xfe\xfd\x67\xbe\xf8\x99\xff\x63\x3a\x07\x34\xac\x4d\x76\xe5\x05\xd5\x45\xda\xb2\x9d\xb5\x44\xe5\xb9\x88\x56\x32\xd3\xfe\xd7\xaf\x88\x20\x06\x05\xfa\x15\x09\x81\x18\x08\x5b\xee\x95\xb6\xca\x3a\x97\xbd\x66\x2c\xae\xd5\xc0\x2c\x39\x0c\x45\x08\xdf\x34\x2d\xd2\x83\x61\x4b\xab\x0b\x60\x30\x73\xc7\x29\x8f\x01\x5d\x96\xf6\xd8\x05\x8d\x10\x4d\x9b\x05\x17\x03\x2c\xb8\xe8\x61\x41\x65\x94\x67\x8c\x66\x88\x49\x67\xca\x83\x11\x20\xe7\x52\x13\xea\xa0\xa1\x32\xe9\xfd\xd9\xd3\x83\x8d\x0a\xf3\x7f\xac\xcf\x4a\x5b\x8c\xaa\x6c\xd4\xde\x7d\xcc\x82\xd0\x98\x1a\xcd\x99\x01\x49\x0c\x82\xbe\xe9\x31\x6b\x57\xe9\x93\xbd\x99\xe1\x67\x67\x3c\x97\xeb\xa2\xee\x0d\x4b\x99\x57\x65\x7b\x14\x97\x87\x02\x3b\x0d\xc5\x6e\x98\xc9\x5e\x11\x58\xdd\xc4\x8e\xc3\x5a\x16\x59\x8a\x88\xe9\x3c\x93\xb6\xd9\x2b\xcb\x34\x3a\x3a\x02\x84\xa6\x98\x88\x05\x58\x52\x7d\xbd\x82\x37\xed\xa6\xe7\x9b\x23\xa7\x16\x44\xd0\x71\xc3\x61\x70\x5d\x68\x76\xc4\x09\xb8\x86\xfc\x9c\xa1\x04\xdf\x81\x40\xc7\xdc\x94\x26\x39\x21\xb7\x19\x98\xbe\x9c\xb6\xbb\xb7\xd5\x6b\xd1\xa3\x76\xf3\xd6\xc0\xb6\xc9\x35\x0c\xf1\xb9\x37\x48\x37\x3c\xd3\x53\x5c\x38\x56\x71\xa6\xcc\x62\x1f\x9e\x17\xf6\x89\x59\x52\x9f\x97\x19\x45\xb1\x1c\xa5\x5b\x2c\xa2\x6b\x90\xec\x4a\x4c\x11\xe4\xfa\x3a\x47\x39\x6b\xa7\x0b\xa7\x7e\x07\xa2\xd4\xac\x68\xb3\x79\x93\x0f\xe6\x23\xd4\x41\xd8\xcf\x07\x9d\xc5\x1d\x08\xb8\xf4\xfb\x32\xcb\x08\x6c\xc4\x9f\xf5\x6a\x65\x4a\x2c\xa9\x20\x5d\xa2\xd7\x2d\x58\x97\xaf\x20\xfb\x8e\x62\xe9\xd2\x5d\xa2\x32\xb2\x7d\xd9\x5a\xee\xcb\x2a\xff\x18\x4d\x0b\x83\xa0\x82\x61\xad\x3d\x55\x75\x19\x55\x67\x33\x10\x48\x47\xac\x8a\x50\x80\x31\xfe\x56\xc3\xb1\xb2\x43\xf5\x03\xf1\x81\x42\x47\x44\xc0\xb1\x45\xa6\xe9\x6a\x4c\xa1\xce\x4e\x1b\x5c\x46\x69\x27\xbd\x95\x54\x24\xc1\x34\x27\x4a\x36\x82\x96\x23\x58\x57\x9b\x5e\x74\x80\x7e\xa1\x66\xe6\x8b\xc1\x4d\xf5\x05\x08\x7e\xe6\xb3\x05\xf8\x99\x4f\x9b\xa6\x96\x22\xa7\x15\xa1\xf0\x75\xe1\x94\xe7\xd0\x54\x9f\x18\x3d\x44\x7d\x4c\xef\x11\xea\xd3\xc2\xe0\xaf\xa5\x3e\x06\xfd\x9d\xab\x8f\x61\xe4\x63\x56\x1f\xbb\x56\xea\xd2\x39\x94\xfb\x07\xcc\xb2\x54\x5f\xc8\x21\x54\x35\xad\x83\xc2\x10\x78\x9c\xaf\x96\x17\x5b\x46\x1e\x23\xa8\xc1\x03\x63\xa6\x0d\x99\x3c\x56\xec\xf8\x50\x9f\xb6\xf0\xfa\x4e\xa4\xaf\xc6\x98\x7e\xb5\xb6\xfc\xa4\x47\xb6\x5d\xa3\x33\xfe\xf6\x8f\x1c\xa6\x81\xed\x2f\xcd\x2c\xf1\x67\x90\xe0\x28\x98\x46\x90\x48\x7b\x34\x53\xcc\x4b\x18\x5d\x01\x08\x34\x15\xb7\x58\x5c\x83\x18\x27\x09\x62\x88\x88\xea\xbe\xd6\xd4\x39\x81\xe6\x54\x1d\x7a\xe9\xd1\xfd\xce\xaf\x27\xee\xb6\xde\xa2\x85\xf7\x47\x56\x5d\x05\x7e\xc0\xee\xdd\x1f\xad\xda\xb0\x6d\x77\x6d\xd7\xbd\xc0\x5e\x7a\x41\xb3\x83\x0c\xad\xb2\xa1\x23\x81\x53\x84\xb2\xc6\xd5\xb4\x18\xa1\x4c\xdf\xc6\xc2\x1b\x6e\x63\xa9\x5b\xa4\xde\x6b\xa4\x1e\xc8\xf7\xec\xd5\x0d\xb0\x3e\x7b\x66\xcd\xdb\x67\xc5\x64\xf2\xcc\xdc\xba\x18\x3e\x9b\x2d\x26\xcf\x68\x58\x8e\x7c\x46\x04\x0d\x68\x2e\x66\x93\xc9\xb3\x8e\xbb\x3f\x75\x23\x49\x3c\x46\xdc\xf5\x29\x31\x31\x93\x5a\x6f\x12\x83\x34\x8c\x66\x4a\x89\xda\xa6\xf6\xeb\xc9\xe4\x99\x80\x6c\x89\xc4\xbc\x0c\x0d\x3b\x57\xb7\x43\xc5\x61\x3a\x9b\x3c\x33\xb1\xe3\x9f\x6a\x06\xea\xb9\xea\x04\x44\xfe\x69\x2d\xd5\x28\x53\x22\x1f\x1a\xdf\xac\xbf\x72\xbd\x9d\x69\x21\xbc\xd4\xf7\xd3\x5e\x6a\x9c\x86\xee\xa5\xa9\x59\x6b\xee\x97\xe8\x6b\xe0\xea\xa4\x0d\xc7\x86\xcf\x51\xce\xf4\x0a\x41\x12\xca\x56\x3a\x74\xc5\x75\x00\xba\xe2\x7c\x4d\xa6\x77\x7c\xd5\x0c\x15\x34\xa2\xf7\xea\x87\x5a\x33\xeb\x65\x56\xed\x5f\x7c\x5d\x1e\x34\xfe\x91\x63\x86\xe2\xb7\x43\x0d\xb7\xdc\xaf\xab\xc0\xd7\x17\x06\x09\xc7\x92\x6a\xa7\x2e\x7c\x7b\x97\x51\x8e\xea\x2d\xcb\x14\x7f\x36\x38\xb9\xad\xd1\x1f\x40\xdf\xd7\x9e\x6a\x67\x79\x6a\xad\x82\x46\x45\x6a\xd4\x4b\x7e\x94\xa0\xcc\x96\xef\x78\x38\xf3\x9e\xb5\xa8\xdf\x04\x70\x58\x75\xdc\x28\x08\xcd\x9d\xcb\xd6\x2e\xed\xec\xca\x9a\x16\xca\x40\x50\xd3\x83\x48\xbe\x9a\xce\x1e\x4e\x0e\xef\xb5\x6b\x24\x55\x3f\x80\xac\x9a\x24\x81\x57\x68\x94\x80\xbe\xe0\x15\x7a\xac\xe2\xb9\x13\x88\x11\x98\x4e\x67\x76\x69\x8a\xb9\x98\xce\x46\x50\xf8\xd6\x80\x79\x34\x54\xd6\xb4\x60\x22\xd0\x12\xb1\x51\x02\x3b\x23\xe2\x11\x52\x92\xa4\x14\x8a\x51\x74\xbc\x93\x3d\x1e\x07\x25\x43\x84\x31\x94\x4c\xbd\xce\xd3\x5d\xd4\x6a\xba\x3f\x23\x8e\x84\x39\xd2\x79\x47\xd9\xff\x20\x46\xf5\xb3\x9a\x8d\xd1\x91\x9a\x8b\xdd\x2d\xc3\x7a\xf3\xe9\x61\x90\xb5\x13\x1d\x9b\x2f\x2d\xa6\x00\x2b\xaa\xe2\x39\x31\x19\x4a\xde\xab\x59\xd8\x28\xfc\x00\xb3\x7a\x39\x35\xc1\x34\x9e\x5f\x59\xd7\xcf\xbb\xb9\xb7\xb6\x49\x96\x1d\x5a\xc7\xde\x40\x3d\x1e\x20\x02\x93\x1c\x35\xb0\xf6\xe5\x36\xcf\xaf\xba\x58\xcb\xf3\xab\x1f\xc8\xc7\xf0\x75\x9a\xd2\x5b\x14\x9f\x5c\x53\x1c\x21\xee\x33\x5f\xf4\x96\x73\x46\xd4\x55\xf7\x51\x1b\xcf\xbc\x3e\x6a\x94\xfd\x7e\xa7\x98\xb4\x10\xb8\x9c\xce\xc1\xf4\x52\x42\x2b\xe6\xca\x34\x7b\x9d\x0b\xba\x34\x07\x2a\xf1\xc0\xdc\xdb\xc4\x0e\x2f\x26\x40\xe6\xc5\x82\x73\x28\xe4\x12\xee\xb7\x58\xcc\xd5\xf9\x61\x73\x8c\xcb\x8e\xe2\x0f\x88\x73\xb8\x44\xba\x56\x56\x5a\xf6\xcf\x1e\xc8\x5e\x0a\x10\x7e\x80\x77\xef\x11\x59\x8a\x6b\xf0\xca\x87\xf0\x0f\xf0\x0e\xaf\xf2\x95\xee\xe2\x4b\xbe\x2c\xad\xc7\x91\x25\xca\xbf\xdc\x17\x45\x98\x8c\xa2\x08\x93\x2d\x29\xaa\xc6\xd9\x3b\x45\xf0\x4e\x2d\x19\xe0\x55\xf8\xaa\xcf\x12\xf6\xdf\xee\x8c\x08\x47\xec\x76\x95\x04\x7f\x33\xaf\x22\x77\x46\xae\x89\x08\xf8\xe2\xec\x6d\x69\xcc\xa5\x07\x15\x34\xb0\x9e\xed\x58\x4a\x9b\xb4\x70\x97\x32\xd3\x4a\x3a\x5e\x66\x25\x16\xbb\x97\x99\x27\xca\xdb\x88\xac\x46\xfa\xc7\x88\xac\xba\x84\x1e\x82\xd6\x33\x6e\x75\x49\x5d\x3d\xa3\x1e\x7a\xcb\x5d\x33\x42\x82\x5b\x95\xc4\x2a\xca\xad\x7b\xe7\x35\xf9\xba\xd0\xd7\xac\xf4\x25\xf3\xb0\xc7\x8e\xf4\x60\x82\x43\xef\x4d\x45\xa9\x42\xac\x8a\xb3\xea\x80\x43\xcd\x07\xfb\x21\xf6\x49\xce\x05\x5d\x95\x87\xe8\x35\x84\xb0\x8e\xda\xae\x60\x96\x61\xb2\x54\xaf\xb9\xd5\x3d\xaf\x1a\xd2\x07\x5d\x15\x9a\xff\xc1\xb4\x7e\xe8\xdf\x42\xa7\x11\xd3\x2d\xa1\x76\x8b\xc2\xc0\x2d\x05\x42\x77\xc6\xe2\x2e\x8e\x9b\xf3\x78\xd7\xe8\x9f\x81\x7f\x58\xc7\xf2\x26\x0a\xe7\x36\x31\x23\xd8\x30\x50\x47\x5f\xfb\xb6\x63\xa3\xd7\x36\x97\xfc\x64\x0d\x4e\x70\xa4\x38\xfb\x8e\xb2\x2a\x8e\xe3\x5c\xa1\xa8\x4a\x9d\xe6\xd5\x1d\x2b\x1d\x1a\xac\x8f\x3b\xd4\xc3\xfa\xef\xe8\xbe\x8c\x57\x6d\xba\xca\xd4\x87\x43\xa0\x00\xb5\x2f\x43\xf4\xa0\x53\x47\x50\x6f\xe6\x80\x7e\x37\xe2\xef\x1d\xb8\x0e\x59\x7d\x80\xd9\x57\x39\xd4\xb7\x5f\x64\xb7\x16\xa7\x6f\x6c\x26\x1f\x1d\x81\x7f\x21\x10\xd1\x3c\x8d\x15\x6f\x13\x4c\x62\x80\xc5\x1c\x70\x0a\x52\x24\x5e\x70\x10\x5d\xa3\xe8\x3b\xa0\xe6\x61\x1f\xbd\x45\x4c\x9f\xa7\x63\x12\xa3\x3b\x14\x03\x9e\xa1\x08\xac\x60\x66\xcb\x6c\x08\xcf\xf7\x12\xc4\x09\xe4\xa8\x03\xe1\xf2\xad\x71\x27\x43\xb8\x23\xc3\x24\x4f\x53\x4b\x46\xdc\x6d\xb9\x82\x99\xa7\xb4\x7a\xc6\x0a\x66\x12\xc6\x57\x2d\xac\x6f\xbe\xb2\xf2\x20\xdf\xa1\xba\x0e\xa5\xe6\xa8\x57\x5b\xf5\xa1\x55\x8f\x72\x3a\x37\xaa\x21\xb8\x41\xec\x1e\xc0\xf8\x06\x92\x08\xc5\x40\x32\x40\xa1\x27\xae\xa1\x00\xf7\x34\x37\x77\xb9\x94\xa4\x09\x42\x31\xb8\xca\x05\xc0\x04\x70\xba\x42\x12\x90\xea\x5e\xb2\x12\xe4\x1c\x29\x51\xfb\x5d\xce\x34\x81\x5a\x97\x10\x57\xe5\xad\xf7\x00\x25\xc7\xcc\x55\x0f\xd5\x6c\xdd\x58\xb7\x3d\x43\xb1\x43\xb7\x3b\x86\x2f\xbb\x75\x2c\x7f\x7d\xa7\xd3\xde\x22\x6d\x3d\x46\x84\x99\x3a\x70\xac\x44\x2b\x05\x39\x7c\xea\xb0\x29\x23\x86\x3b\xde\xf1\x18\x45\x5d\x37\xf7\xc6\x51\xe1\x6e\xeb\x31\x5a\xd5\x43\xa2\xd0\xb8\x0c\x78\x68\x67\x83\x69\x32\x7d\xba\x68\x1f\xdb\x75\xfa\xaa\xf2\x63\x97\x2f\xda\xbe\xa5\xf4\x29\x1d\x58\x8d\x7b\x2c\xae\x1b\xbe\xe8\x09\x0f\x1c\x16\xc5\x28\x17\xbe\x36\x18\xab\x6e\x45\x65\x7b\xcc\xdb\xb4\x35\x7c\xfd\x1a\x3b\xbb\x62\xd1\x19\x17\x18\xa4\xae\xfb\x82\xbf\xfc\xbc\xb9\xf8\xf4\x51\x3d\xfb\xfc\xa8\xf2\xa2\x4c\xcd\x5b\x50\xa7\xb8\x7d\xc7\xb9\x77\x80\xde\x33\x4b\xa7\x62\xd1\x23\x6f\xbf\x21\x18\x52\xfa\xf9\x89\xa4\xf7\xce\x08\x56\xb9\x66\x51\xa3\xa5\x17\x74\x13\xaa\x2a\x2d\xf0\xaa\xde\x2e\x57\xd0\xd5\x83\x48\xa7\xb5\xf3\x22\x32\xf4\x1f\x30\x63\x28\x6a\x0a\xbc\x2e\xd5\xa4\x38\xad\x3c\xe1\x3a\xb7\xc6\x6b\xc0\x55\xf1\xa2\xe3\x62\x77\xe3\x36\xb7\xd7\x48\xad\x7b\x28\xf2\x53\x15\x6a\xfc\xed\x36\x7e\x40\xeb\x13\xb1\x0a\xa4\x2e\x32\x00\xab\x7a\x2f\x70\xef\x70\x2a\x10\x2b\x6f\xa0\x95\xb5\x75\xa9\x06\xea\xb4\xf2\x83\x4b\x19\xc2\x4b\xf2\xdf\xc8\x51\xc5\xba\xd4\xc0\xb5\x5b\x79\xc1\x35\xb7\xc7\xad\x1a\x5d\xa2\xe1\x55\xb5\x5e\xb0\xda\xef\x77\xe4\xa7\x2e\xd5\x30\x9d\x56\x5e\x70\xed\x98\x56\x55\x59\x15\x2e\xda\x71\x2f\x4f\xa0\xad\xb9\x57\x96\x2d\x5a\x81\x18\x2f\x88\x56\xa0\xaa\x06\x59\x16\x2e\xda\xc1\x2c\x4f\xa0\x6d\x34\x4d\xd9\xa2\x15\x7b\xf0\x81\xd8\x5c\x30\xad\x75\x72\xd4\xf2\xa8\x5e\xa9\x34\x15\xbd\x2a\xd4\xb8\xd9\x6d\xbc\x80\x9e\x33\xbc\x82\xec\xbe\xa1\xe6\x75\xa9\x06\xeb\xb4\xf2\x82\xfb\x19\xc1\xb8\xb9\x8e\x97\x65\x0b\x13\x02\xae\x5a\x78\x42\x74\x8f\xcc\x35\x44\x5d\xb6\x68\x06\x95\xfd\xf6\x4c\x14\x31\xe4\x3c\x58\xd7\x25\xe5\x83\x7f\x53\xeb\x09\xab\x39\xad\x2f\xac\x69\x7d\x31\x6a\x5a\x5f\xe0\x25\x71\xe9\xd4\x25\x06\x56\x59\xbb\xa5\x5d\xa0\x4b\x0c\xac\xb2\xd6\x0f\x56\x7e\x65\x5e\x47\xd4\xc0\x74\x51\x99\x47\xad\x6a\xe0\xa7\xd1\xad\x4b\x13\xf2\x53\x15\x6a\x14\xed\x36\x7e\x40\x1b\x28\x5a\xf8\x6d\x44\xce\x8c\xd0\x1f\x0d\xd9\xec\x07\x74\x7b\xb6\x3f\xc0\x21\xe8\x19\xf8\x71\x7b\x06\x3a\x94\x10\x3e\xb9\x05\x4f\x6e\xc1\x93\x5b\xf0\xe4\x16\x3c\xb9\x05\x4f\x6e\xc1\x93\x5b\xf0\xe4\x16\x3c\xb9\x05\x4f\x6e\xc1\x5f\xcb\x2d\x58\xb7\x13\x22\x3c\xe0\xd9\xb7\x3e\xc5\xf4\x4f\x3f\xd9\x9d\xd8\xd9\x17\x82\xce\x5c\x38\x6a\xbc\xaf\xdf\x36\x3d\xed\xd9\x59\xaa\xe7\x31\x78\xfd\xa9\x09\x9f\xc7\xe5\x4b\xdd\x8e\xbc\xad\x92\x3f\x8f\x19\x62\x2f\x29\xa0\x7f\x04\x67\xf6\x94\x0e\x7a\x7b\xde\x3d\x2c\x29\xf4\xc6\xd9\xf5\x03\x52\x43\x8f\x13\xc0\x7f\x6a\x82\xe8\x71\x5c\x7a\x14\x89\x8b\x74\x5a\x91\xf3\x14\x62\x37\x13\xd6\xa8\x2d\xc0\xc9\x15\xbd\xdf\xb9\x6d\x70\x7d\x5c\xda\x15\x56\x58\x0d\xea\xd9\x36\x19\x98\xc7\x71\x67\xfb\x3c\xcc\x9b\xad\x8c\x8e\xc4\xca\xed\xd4\x33\x43\x19\x96\xc3\xdd\xa7\x58\xf6\x40\xb9\x4a\xb4\xdc\x31\x75\x7c\x92\xe6\x75\x64\x4d\x56\xc9\x95\xc3\x0e\xdb\x4f\xfd\xde\x4d\xe2\x64\x2f\x61\xd8\xe9\x93\x3d\x78\x11\x56\x59\x94\x37\xb7\x0d\x66\x5e\x2f\xe2\xd7\x8e\x6d\xb0\xb9\xc3\xba\x54\x36\xaf\x7c\xbd\x46\xbf\x5a\xef\xed\x3d\x9e\x92\xef\x23\x79\xaf\x6f\x3a\x5e\x07\xed\x0d\x49\x66\xfd\x49\xd9\x26\x37\x6f\xf5\x76\x6a\x73\xa2\xd9\x9e\x14\x10\x26\x27\x6f\x61\xb3\xeb\xe5\x9e\x72\xfd\xfa\x66\xef\xdd\x35\x7f\x1f\x98\xca\x17\x27\x00\xc7\xad\xe4\xae\xde\x09\x7e\x9f\x9b\x0c\xbf\x85\xe6\xd3\x16\x10\xea\xbc\xa4\x9d\x2c\xfd\xb3\xd3\x03\xfb\x88\xe0\x47\x26\x09\xf6\x5a\xa5\xb6\xdc\x2d\x76\x98\x2a\xd8\x25\xd2\x3f\x57\xf0\xcb\xf1\xc9\x82\xfb\x72\xca\xb8\x68\x6c\x93\x52\xb8\xa5\x92\xc3\x3f\xed\x15\xe2\x51\x25\x16\xf6\x5c\x46\xf6\x9e\x5e\xd8\x5b\x77\xff\x22\x49\x86\x71\xac\x9e\x83\x6e\xca\x28\xdc\x9b\x46\xe5\xb9\x9d\x1e\xdd\xd5\xe8\x0e\xa0\x4d\x8d\xde\x26\x31\xf1\x03\x34\xba\x56\xe7\x07\xe5\x1e\xf6\xd1\xc5\x07\x39\x02\x5d\xfe\x60\xa6\x4a\x1a\x98\x99\x8d\x78\x1b\x04\xbb\xfc\x38\x55\xd2\x95\x03\xd2\xf8\x71\x1b\xff\xfc\x47\xc5\xfc\x51\x37\x16\x36\xad\x7e\x3b\x4a\x2b\x94\x8b\x2d\xd6\xcf\x2e\xdd\xde\x00\xad\x67\xfd\xef\x0d\x68\x0f\x69\xac\xf1\x5d\xd7\x7f\x72\x96\x3f\x1f\x85\x7a\xac\xb9\xfe\xba\xc4\xb1\x39\xd9\x5f\x77\xaf\xfd\x67\xfb\xab\x38\xfd\xff\x37\xe7\x9f\x8f\x32\x3d\xd6\xcc\x7f\xbe\xca\xe4\xa6\xfe\xdb\xa1\x32\x8d\xca\xfd\xf7\x27\x2a\xd3\xfa\xaf\x9b\xe4\xdb\x83\x6b\x43\xa9\xbe\x55\xbe\xb1\xf5\x50\x1a\xeb\x7e\xaf\xa3\xc7\x32\xed\x08\x6a\xf6\x6d\x32\x8f\x33\x5d\x38\x8c\x63\x86\x78\x15\xb7\xef\xce\x1e\xee\xc5\xf7\xfd\xe5\x10\x7f\xbe\xde\x9c\x44\xdc\x33\xb1\x9f\x77\x24\xd0\x7f\x39\xac\x93\xfc\x79\x05\x05\x3b\x9c\x9d\xfe\x54\x7f\x5e\x3e\xcd\xde\x12\xfe\xed\x8d\x59\x75\xf2\x3f\x9f\x5e\xfb\x4f\x01\xb8\x19\x0b\x8f\x44\x80\x3e\x49\x3c\x1d\x95\x55\x61\xfa\x11\x7f\x23\xa7\x19\xaf\x1f\xf1\xc7\x11\x07\x1d\x50\x13\xc6\x57\x6e\x7e\xcf\x5f\x40\x1c\xde\x63\x76\x72\x99\xa3\x8b\x1b\xe3\xcf\x30\xf6\xcd\x93\xde\xf3\x8d\x16\x4b\xac\x1f\xff\x17\x00\x00\xff\xff\xf5\xac\xd8\x35\x03\x81\x00\x00")

func templatesModelGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelGotpl,
		"templates/model.gotpl",
	)
}

func templatesModelGotpl() (*asset, error) {
	bytes, err := templatesModelGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.gotpl", size: 33027, mode: os.FileMode(420), modTime: time.Unix(1631045103, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xef\x6f\xf2\x36\x10\xc7\xdf\xe7\xaf\x38\xa1\x68\x6a\x27\x96\xad\xed\x3b\xa4\xbe\x98\xca\x54\xf5\x45\xab\x8a\xfd\x78\x83\x78\xe1\x91\x03\xbc\x19\x27\x35\x0e\x15\x8a\xf2\xbf\x4f\x26\x09\x84\x60\xc7\xe9\x36\x1e\x3d\x7d\x74\x57\xa9\x12\xf6\xdd\xf7\xce\xf6\xd9\x7c\x44\xca\xe6\x7f\xb3\x25\x42\x9e\x43\xf4\x2b\xea\xe8\x21\x91\x0b\xbe\xcc\x14\xd3\x3c\x91\xd1\x0b\x5b\x23\x14\x45\x10\xf0\x75\x9a\x28\x0d\x83\x65\x12\xb1\x34\x51\xa8\x93\x88\x27\x3f\xa2\xc0\x35\x4a\xcd\xc4\x20\x08\xb6\x4c\x81\x42\xb1\x8f\xdb\xac\x78\xba\x99\xe0\x92\x6f\xb4\xda\xc1\xc1\x2b\x9a\xd8\xe6\x83\x60\x91\xc9\x39\x70\xc9\xf5\xd5\x35\xe4\x41\x00\x00\x0e\xa5\x7b\x9f\x56\x5e\x94\xe1\x79\x0e\x8a\xc9\x25\x42\x88\x52\x73\xbd\x33\xeb\x18\x42\x58\xab\xc2\xe8\xbe\x5c\xed\x89\x88\x59\x68\x19\xfc\x03\xf0\x05\x6c\x56\x49\x26\xe2\x52\x19\x55\xd3\x13\x42\x13\xdc\xd4\x86\x30\x7a\xcd\xfe\x14\x7c\xfe\x9c\xc4\x58\xc9\x58\x97\x30\xcd\xf3\x93\xb8\xa2\x78\x8a\xcb\x8f\x33\xb8\x87\xef\xec\xcb\xcb\xf7\x7a\x8d\xd2\x96\x1a\xae\x04\xca\xe3\x82\xa2\x07\x85\x4c\xe3\x35\xfc\x54\x2f\xc2\x58\x39\x38\x82\x35\x4b\xa7\x1b\xad\xb8\x5c\xce\xbe\xb7\x67\x78\x92\x8b\x04\x8e\x69\xea\x54\xd5\x1e\xa6\x4c\xa1\xd4\x43\x08\xd9\xbc\xde\xbd\x76\xe6\x66\x5a\xe7\x1e\x3e\x49\x69\xdf\xc8\x32\xc1\xc9\x26\xb6\x05\x07\x66\xe3\x2a\xbf\xa2\x18\xc0\xe8\xb4\xda\x46\xce\xaa\xca\x68\x8c\xa9\xc2\x39\xd3\x18\xb7\xb5\x8c\x1d\x67\x47\xa0\x55\x86\x43\xab\x1c\x4a\x6b\x70\x2b\xd3\x2b\x53\x6c\x8d\x1a\xd5\x18\x17\xa6\x8d\xcd\x1e\xb9\xa3\x0e\x87\xe7\x8e\x8e\x26\xf8\x96\x71\x85\x71\xeb\x40\x6b\xab\xa7\x0f\xa1\x9b\x51\xe3\x66\xbc\xe0\xfb\x71\xa2\x72\x35\x53\x57\x67\x3a\xc6\xa6\x33\xf3\x57\xb6\xc7\xf9\x9e\x36\x6b\xaf\xba\x81\x0f\x21\x14\x37\xfb\x26\xe8\xb1\x02\x5b\xf9\xae\x0d\x11\x37\x8e\xf5\x1e\x02\x9c\x33\xb6\x1a\x6f\xf7\x35\x8a\x9b\x2e\x45\x6b\x19\xb7\x9e\x32\xfc\xa5\xd8\xca\xb9\x2b\xcb\xb9\xf5\x29\x43\xdd\xed\xe2\xce\x74\xfa\x79\x63\xda\x12\xa1\x8c\x3d\xb2\x45\xb7\x50\x3f\x11\xbf\x57\x47\x1a\x7f\x70\xb7\x87\x45\xfa\xfa\xdf\x5c\xdb\x3e\x17\xf0\x17\xa9\x15\xc7\x8d\xa3\x11\x9a\xf7\x6e\x3a\x3b\xde\x3c\x8b\x92\xfd\x99\x6a\xf4\x45\xea\xbb\x49\x55\x29\xce\xb6\x71\x77\xa2\xf9\x86\x19\x55\x2f\x67\xfd\x55\xde\xd1\x4f\xbf\xed\xd2\xa3\xbb\xf9\xd0\xed\x5e\x3f\x82\x69\x34\xc6\x05\xcb\x84\xfe\x83\x89\xec\xec\xdd\x6e\x5a\xd3\xef\x90\xa8\x15\xec\x49\xe8\xef\x1f\xbe\x00\x7c\x3b\xac\x60\x80\x32\x5b\x0f\xba\x8a\xfa\x59\x88\xe4\x1d\xe3\x87\x55\xc2\xe7\xb8\x3f\x4f\xdf\x5b\x08\xa7\x87\xf8\xd7\x10\xc2\xed\xfe\x10\xd3\xe8\x54\xcc\x77\xd3\xf7\x3b\xb0\xf5\x5f\xf2\x8e\x96\xae\xcd\x7f\xed\x7a\x3c\xc4\x61\x1a\x3d\x67\x42\xf3\x54\x74\x1e\x63\xed\xe3\xfa\xe2\xec\x99\xd8\x52\x72\x47\xc4\xc7\xbc\x1d\x53\x2d\x11\x87\x97\x65\xb8\x11\x68\x99\x75\x52\xd9\xef\x69\x7c\x4e\x65\xe5\xe0\x85\xa9\xac\x4c\xf2\xe5\xa9\xcc\xf2\x1a\x11\x96\x11\x96\x11\x96\xb9\x8c\xb0\xec\xdb\xc4\x32\xa2\xb2\x76\x49\x44\x65\x1f\x8c\x26\x2a\xf3\x0c\x37\x02\x5f\x99\x9e\xaf\x88\xa9\x88\xa9\x88\xa9\x7c\x71\xc4\x54\xc4\x54\xc4\x54\xc4\x54\xc4\x54\xc4\x54\xff\xdb\x2f\x5d\x63\x14\x78\xf6\x4b\x57\x39\x78\x61\x2a\x2b\x93\x10\x95\x11\x95\x11\x95\x75\x97\x43\x54\x46\x54\x06\x44\x65\x44\x65\xb5\x11\x95\xf5\x8f\xf8\x94\x54\xf6\x88\xba\xf5\x80\x4c\xd0\xdc\xe5\xed\xa5\xa1\xec\x11\x35\x11\x19\x11\x19\x11\x59\x77\x39\x44\x64\x44\x64\x40\x44\x46\x44\x56\x1b\x11\x59\xff\x88\xcf\x4a\x64\xcf\x4c\xee\x1c\x54\x66\xa6\x2e\x4f\x66\x26\x0b\xd1\x19\xd1\x19\xd1\x59\x77\x39\x44\x67\x44\x67\x40\x74\x46\x74\x56\x1b\xd1\x59\xff\x88\xaf\x86\xce\x0c\x2b\x11\x53\x11\x53\x11\x53\x11\x53\x11\x53\x11\x53\x11\x53\x9d\x1b\x31\x15\x31\xd5\x7f\xf8\xc5\xab\xfc\xdf\x1a\xcc\xf3\xfa\x53\x11\xfc\x13\x00\x00\xff\xff\x08\x39\x9d\x04\x65\x44\x00\x00")

func templatesRelationships_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRelationships_registryGotpl,
		"templates/relationships_registry.gotpl",
	)
}

func templatesRelationships_registryGotpl() (*asset, error) {
	bytes, err := templatesRelationships_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 17509, mode: os.FileMode(420), modTime: time.Unix(1625075149, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/README.md": templatesReadmeMd,
	"templates/identities_registry.gotpl": templatesIdentities_registryGotpl,
	"templates/model.gotpl": templatesModelGotpl,
	"templates/relationships_registry.gotpl": templatesRelationships_registryGotpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": {nil, map[string]*bintree{
		"README.md": {templatesReadmeMd, map[string]*bintree{}},
		"identities_registry.gotpl": {templatesIdentities_registryGotpl, map[string]*bintree{}},
		"model.gotpl": {templatesModelGotpl, map[string]*bintree{}},
		"relationships_registry.gotpl": {templatesRelationships_registryGotpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

