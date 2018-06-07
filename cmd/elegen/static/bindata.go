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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1515709395, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x56\xcf\x6b\xeb\x46\x10\xbe\xeb\xaf\x98\x1a\x53\x24\x78\x96\x2f\xa5\x87\x14\x1f\x42\x78\x01\x43\x5f\x08\x0e\xf4\x62\x72\xd8\xc8\x23\x79\xe8\x6a\x57\xec\xae\x92\x0a\xa1\xff\xbd\xac\xb4\x92\x25\x45\xf1\x8f\x34\x49\xe1\xe5\x14\xaf\xe6\xd7\x7e\xf3\xcd\xb7\x93\xb1\xe8\x6f\x96\x20\x94\x25\x84\x0f\x68\xc2\x1b\x29\x62\x4a\x72\xc5\x0c\x49\x11\xde\xb1\x14\xa1\xaa\x3c\x8f\xd2\x4c\x2a\x03\xb3\x84\xcc\x3e\x7f\x0a\x23\x99\x2e\x59\x26\x15\x1a\xb9\x20\x11\x2d\x91\x63\x8a\xc2\x30\x3e\xf3\xbc\x67\xa6\xc0\xf7\x00\x00\x68\x87\xc2\x90\x29\x6c\x14\xfd\x83\x65\xb0\x82\x94\x65\x5b\x6d\x14\x89\xe4\xb1\xf3\x09\xd7\xce\x0e\xca\xda\xcd\xfe\x95\xe5\x02\x14\x13\x09\x36\x55\x3d\x64\x18\x51\x4c\x51\x5d\x95\xb6\x15\x1d\x0c\x81\x62\xd0\x7b\x99\xf3\xdd\x06\x13\xd2\x06\xd5\xc0\x1a\x42\x98\x87\xf7\xf9\x13\xa7\xe8\x87\xdc\xe1\xd0\x77\x01\xf3\x43\x89\x70\xb5\x82\xd0\xda\xf0\xf0\xfb\xe1\x70\xd1\x73\x98\x95\xa5\x33\xd8\xa0\x36\xf6\x73\x55\xcd\xae\x6c\x0d\xfd\x30\x55\xd5\x5e\xe8\xdb\x20\x15\x8a\xdd\x38\x7b\xef\xa8\xf2\x06\x98\x45\xcc\x60\x22\x15\xfd\x84\xc0\xc9\x5c\x45\xf8\x29\xe0\x31\x4e\x4c\x7f\x22\x62\x8b\x2f\x84\xac\x2c\x5d\x55\x73\xfa\x06\xf3\xfa\x66\x3d\xa7\xeb\xe6\xa6\x30\xc4\xb8\xb5\xfb\x38\x60\x27\xb1\x0e\x3c\x6f\xb9\x84\xba\x90\xbf\x50\x69\x7b\x71\x85\x26\x57\x42\x83\xd9\x23\x44\xb9\x52\x28\x0c\x3c\xbb\x6f\x32\xae\x8f\xd3\xba\x70\x2f\xce\x45\x34\xf0\xf5\x03\x88\xb9\x64\xe6\xf7\xdf\xa0\x74\x71\x3a\x35\xba\xbe\x5f\xaf\x45\x2c\xc3\x36\x4d\x55\xd9\x56\x9b\x22\x43\x37\x28\x31\xb1\x27\x8e\xb7\x2c\x32\x52\x15\xa0\x8d\xca\x23\x53\x56\x5e\x9d\xc5\x8f\xa7\x8c\x02\x68\x71\xb8\x55\x32\xb5\xd0\xf8\xc2\xe2\xd3\x50\x25\x80\x49\xae\xd4\x17\x77\xb5\x8d\x55\x6d\x6b\xdd\x1f\xbd\x0b\x92\xde\x34\xb3\x5d\xf8\x6e\xc8\x8b\xcb\x93\x0f\xe4\x61\xdb\xc6\xb9\xa8\x8a\x9a\x42\x7e\x43\x98\xb3\xf3\x1f\x26\x6c\x5b\xff\x7b\x59\x46\x51\xf8\x4c\x1c\x6e\xeb\xd3\x44\xc2\xa0\xcd\x48\x31\x10\xac\x20\x0e\x5f\xf5\x8b\x89\x22\xf8\x03\x7e\xa1\x70\xad\xbf\xa7\x99\x29\xfc\xa0\x37\xce\x2d\x50\x03\x51\x9d\x0a\xd5\x75\xe1\xe2\x70\xee\x6c\x18\xce\xc1\x29\x8a\xe0\x3c\x48\x9a\x43\xbf\x6d\xe8\x24\x12\xe3\xb3\xc6\xa7\x05\x48\xbf\x90\x89\xf6\x1d\x25\x5c\xd1\x9d\x6e\x1c\x51\xb3\x77\x2b\x59\xc4\x74\xb3\x2a\xbc\x92\xae\x83\xbc\x5c\x8d\xb1\xbb\xc3\x97\x37\x5c\xfc\xc0\x9b\x10\x98\xd1\xcf\x1d\xc6\x2c\xe7\xe6\x55\x58\x41\xdc\x35\xe5\x02\xbc\x6d\xa7\x1e\x6a\xfa\x0d\x98\x78\x1c\xe8\x51\xbf\x5d\xe7\x46\xed\x6f\xb8\x1d\x9c\xec\xfe\x8d\x14\x06\x85\x79\x0f\x09\x26\x5c\x3f\x8c\x0b\x42\x9a\xb6\x45\x6b\xbd\x91\xd2\xfc\x3f\x5c\xf9\x75\xca\xe1\x9e\xe7\x8a\x71\xa8\xaa\x3f\x49\x5b\x75\x3f\xcd\x99\x8f\xa6\xd0\x04\xf2\x27\x99\x74\xa4\x5b\x1d\xa1\xa6\xc8\xf0\x4e\x5e\x6d\x90\x37\x8d\xdd\x53\xa6\xfd\x7e\x21\x83\x2f\x4d\x07\xd5\x58\xd7\xd5\x94\x8d\x4d\x69\xb7\x79\x8a\xdd\xfb\xba\x9a\x4a\x6d\x9f\xdb\xe5\x12\xda\x37\xb8\xbf\x09\xd4\x4f\xfe\x1b\xc3\xe5\xec\xdd\x42\xe0\x7e\xf9\x6f\x8d\x62\x1b\xbd\xdb\x10\xba\x9a\x2a\xaf\x4e\x7f\xcd\xb9\x83\x8d\x50\x77\x45\x30\xce\x01\xff\x21\x6d\x48\x24\xed\x70\x10\x6a\x97\x74\xe0\xe3\x07\xb0\x9d\x5e\x14\xfb\x38\x4d\x99\x7c\xfd\x2a\x79\x74\xa6\x2e\xdb\x9b\x2b\x87\x5e\xfd\xac\xdf\x4a\xd5\xdd\xbb\x0f\xa1\xed\xa5\x7b\xf9\x21\x96\xaa\xfe\x9d\xd0\x33\x1e\xf6\x91\x0e\xd1\x71\x9c\xe3\xda\xb6\x7d\x6c\xe6\xe6\x94\x8c\x9d\x01\xeb\xd7\xea\x54\x5b\x78\xf9\x5f\xf7\xf4\xa3\xdd\x3a\x47\xdb\x86\x3b\x89\x55\xb4\xca\xfb\x37\x00\x00\xff\xff\xc9\xa3\x7f\xed\xd5\x0f\x00\x00")

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

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 4053, mode: os.FileMode(420), modTime: time.Unix(1528338490, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x6f\xdb\x38\x12\x7f\xf7\xa7\xe0\x19\xd9\x5d\xfb\xe0\x28\x7d\xf6\x5e\x16\xe8\xb5\xd9\x45\x70\xed\x6e\xd0\x14\xdd\x87\xa2\xd8\x30\xd2\xd8\x61\x23\x91\x2a\x45\xa5\xc9\x19\xfa\xee\x07\xfe\x93\x48\x4a\xb2\xe9\x6c\x72\xc8\x83\xf3\x10\xd8\xc3\xe1\x6f\xfe\x91\xc3\x19\xd2\x25\x4e\x6f\xf1\x1a\xd0\x66\x83\x92\x4b\x10\xc9\x1b\x46\x57\x64\x5d\x73\x2c\x08\xa3\xc9\xef\xb8\x00\xd4\x34\x93\x09\x29\x4a\xc6\x05\x9a\x4d\x10\x42\x68\xba\x2a\xc4\x54\x7f\xaa\x1e\x68\x3a\x9d\xe8\xcf\x6b\x22\x6e\xea\xeb\x24\x65\xc5\x09\x2e\x19\x07\xc1\x8e\x09\x4d\x4f\x20\x87\x02\xa8\xc0\xb9\x9e\xb2\xd9\x20\x8e\xe9\x1a\x50\x72\x59\x42\x9a\x7c\x7c\x28\xe1\x82\xb3\x3b\x92\x01\xaf\xd0\x71\xd3\x68\x2c\xa9\x0e\x6a\x9a\x76\x0a\xd0\x4c\x0d\xce\x27\x93\xcd\x06\x1d\xe5\x58\x40\x25\x3e\x01\xaf\x08\xa3\x68\x79\x6a\xc0\xde\x29\xf2\x6b\x21\x38\xb9\xae\x05\x54\x96\x41\x5a\xd0\xca\x3d\x22\x0b\x74\x04\xb4\x2e\xe4\xbc\xeb\x9a\xe4\xd9\x19\xad\x8b\x4a\x43\x84\xd0\x4d\x33\x39\x39\x91\x0a\xa8\x19\x4a\x5b\xd4\x34\x88\x43\xc9\xa1\x02\x2a\x2a\x24\x6e\x00\x95\xac\xaa\xc8\x75\x0e\xe8\x0e\xe7\x35\x54\x68\xc5\x38\xc2\x56\x0b\x65\x8c\x9e\xde\x6a\x66\xfc\x3a\x4d\x26\x42\x22\xf6\xf0\x2b\xc1\x09\x5d\x4f\x26\x29\xa3\x95\xf5\xfa\x66\x73\x6c\x2d\xa0\xb8\x80\x05\x3a\x52\xd2\xa4\x15\x7a\xf2\x27\x2d\xdc\xb8\xd0\xa8\x4d\xb5\xa4\x50\x63\x3d\x55\x32\xe8\x4f\x4d\x93\x58\x57\xb7\x53\x7a\x5a\x9d\x6a\x53\xec\x0c\x2f\x38\x2a\x36\xdd\x67\xe3\x35\x1d\x96\xf7\x2c\x83\x3c\x39\xa3\x82\x88\x07\x63\xf9\x79\x06\xea\x6b\xa8\x57\x4b\x67\x2b\xf5\x9d\x5d\x7f\x85\x54\x24\x93\x3b\xcc\xe3\xf0\x4e\x51\xbb\xde\x92\x96\xb8\x51\x9a\x4a\xd6\x25\x6a\x97\x97\x03\xf5\x01\x2a\x21\x47\x9b\x66\xba\x50\xac\x6f\xb0\x80\x35\xe3\x0f\xcb\x21\x56\x56\xf3\xb4\x8d\xa0\xe6\xbf\xe0\xe4\x0e\x0b\x89\x1e\xb0\x9b\x81\xa6\x59\x4c\x9a\x89\x5a\x84\x64\x85\x28\x13\x1e\xd3\x79\xf5\x81\x31\xd1\xad\xb5\x61\x2b\x2f\xf2\x9a\xe3\x1c\x35\xcd\x3b\x52\x09\xd7\x6f\x18\xe5\x92\xc2\x56\x11\x73\xdb\xf5\x16\x23\xe3\xf3\x97\x7f\x8e\x72\x9a\x18\xbf\x61\x54\x00\x15\x4e\x38\x45\xcd\xa9\x8e\x25\x19\x8c\x65\x85\x08\x55\x5f\xa5\xd2\xc9\x64\x55\xd3\x14\xcd\x58\xa4\x4a\xf3\x50\xe0\x6c\x3e\x1c\x6f\x15\x15\xad\xcc\x38\x74\xb7\x6c\x26\xd6\x9a\xb2\x33\x01\xa3\x92\x11\x2a\x80\x23\xc1\x10\x46\xa9\x1c\x93\x7a\xc7\x69\xfa\x18\xcb\x4a\xdf\x1c\xcf\xd4\x15\xc1\x32\xc1\x18\xcb\x94\x32\xcb\x53\x84\xcb\x12\x68\x36\x8b\x13\xb1\x69\x16\x88\x25\x49\x32\x77\x9d\xf3\xa3\x84\x32\xe6\xbf\x56\x68\x06\xb4\xf2\x62\x26\x98\xfa\x8a\x11\x85\xef\x5a\xba\x09\xea\x73\x79\x43\xeb\x32\xb3\xf2\x93\x24\x09\xe3\xac\x3d\x12\xe9\x30\x56\x8b\xbf\xe9\x2f\x99\xd3\xff\x5a\x48\x87\x48\x20\x9d\x88\xad\x76\x3a\xbf\x58\x39\xad\x18\x56\x0b\x35\x21\x99\x6d\xdb\x47\x73\x8d\xdf\x78\x6b\x96\xd5\xc2\x04\x45\xed\xc4\x94\xd1\x3b\xe0\xc2\x8d\x89\x5a\x95\xb4\xb7\xfa\xb5\xd9\xd5\xe3\x9c\x2e\xff\x0f\xec\x28\x07\x33\xf0\xe7\x16\xce\x4d\xe3\xba\x8d\x08\x28\x1c\xbf\x6d\xf5\x98\xe4\xdd\xee\x93\xb7\xb0\xc2\x75\x2e\xfe\xe0\x19\x70\x2f\xe5\x64\x7a\x00\x31\x39\x42\xe8\x1a\xad\x08\xe4\x59\x65\x17\x6b\xaa\x17\xc8\xfe\x8e\x71\x05\xce\xe6\xe8\xf3\x17\x7d\x40\x07\x89\xc6\x92\x3b\xe3\x82\x62\xe7\x0f\xa3\x56\x57\xa1\x84\xf5\x86\x2d\x80\xda\x53\xca\x3f\x67\xfc\x03\x57\xbb\x48\xbb\xc4\x02\xb8\xde\xb8\x33\xb4\xbf\x6b\xbd\xc1\x9e\xcd\x11\xa1\x62\x28\xb9\x82\x48\x5e\x5f\x9c\x9f\xd3\x15\x4b\x9c\xc2\x49\x17\x5d\xb6\x20\xd8\x55\x11\x84\x95\x40\x21\x59\xa4\xea\x38\x9c\x66\x4f\xeb\xdd\xc7\x99\x29\xa5\xea\x54\x6a\x1d\xc4\x62\x7b\x0c\x36\x1b\x54\xe0\x5b\x90\x4c\x28\x41\x47\xc9\x45\x7d\x9d\x93\x54\xe2\x23\x35\x7a\xdc\xda\x25\x7d\xa1\xe4\xda\xe9\xd2\x47\x57\x5f\x2b\x46\x97\xd3\xe3\x29\xba\x56\x1f\xfe\x52\xd6\x98\x78\x4c\xaf\xf4\x2c\x59\x3b\x27\xef\x6b\x01\xf7\x26\x86\xbf\xc3\xf7\x1d\x1e\xb2\x27\x93\xcc\xc2\xe3\x39\x45\xea\xa5\x62\xbc\x03\x70\x36\xdf\x0e\x12\x84\xfa\xc7\x6d\xbc\xdd\x9a\x77\x9d\xb1\xdc\xb2\x3e\x16\xfd\x5d\x62\x0b\x5b\x55\xf6\xca\xb0\x31\xde\x15\xf7\x6d\xc0\xce\x29\x11\x04\xe7\xe4\xbf\xb2\x63\xd8\xb2\x7f\x74\xb1\x55\xdd\xb0\x3a\xcf\xfe\xe4\xc4\x9b\x88\x8e\x4c\xb9\xaf\x8a\xdd\x00\x24\x08\xb7\x53\x13\x2b\x7b\x3c\xf5\x9a\x66\xe1\x6f\xc8\xf1\x2d\x3a\x58\x24\x8d\x15\xbc\x76\x8f\x6e\x0d\xd0\x1c\x3d\x49\x1d\x84\x7a\x85\x90\xcd\xe6\x41\x7a\xd5\x95\xbf\xa7\xe9\x4f\x15\xaa\x29\xf9\x56\xdb\x72\x4f\xce\xd9\x53\x7b\x39\x65\x36\x47\x7e\x4a\xd5\xc1\xd3\x73\x1d\x6d\x6c\x7c\xed\x91\x90\xb4\x02\x3a\x26\x59\x07\xc8\xd3\x12\x32\xbb\x19\x6c\xfb\x04\x79\x05\x21\xc4\x74\xda\x0d\xeb\x98\x69\x17\x5c\x82\x70\xe4\x56\x20\x9e\xc7\x05\x9e\x98\x19\xc9\x8c\x1b\xe6\x91\x7e\x88\x73\x00\x3a\x45\x24\x1b\x36\x73\xe8\xdc\xb8\xc1\x3c\x4b\x59\x06\x59\x78\x82\xa8\x2c\x16\x6d\xda\xe3\x8f\x8d\x6d\x47\xbc\xed\x75\x46\x8e\xfa\x28\xed\x94\x7a\x2f\xfc\x4c\x77\x03\xaf\x0d\x78\x0b\x55\xca\x49\x29\xba\x7b\x89\xb7\x2c\xf5\xcb\x1f\x96\xd6\x2a\x07\x28\x1e\x59\x7a\x75\xab\x34\x36\x6a\x6f\x59\x3a\xb0\x19\x8f\xa5\x32\xf0\x6d\x54\x9f\xe9\x67\xca\x32\x96\x7e\x99\x86\xdb\x4b\x91\x2f\xf5\x75\x86\xb7\x0d\x7d\xb6\x2b\xe9\x81\x2a\xfd\x37\x4e\x6f\x05\x49\x6f\xab\x2d\x86\x5f\xf5\x17\xb2\x53\x66\x28\x2b\x63\xf6\x9d\x52\xa9\x67\xa9\x51\x67\x55\x88\xe4\xb2\xe4\x84\x8a\xd5\x6c\xfa\xaf\x1f\xaa\xe5\x0f\xd5\x2f\x53\xd9\x0d\x74\xf9\x56\x45\xb1\x23\xe9\x24\x36\x9f\xb8\x37\x4d\xbb\x0a\x0d\x5d\x67\xb4\x67\xd4\x6f\x40\x81\x63\x01\xbf\x81\x90\x5d\x67\xbf\xec\x38\x39\x41\xbf\x81\x90\x76\xf5\x36\xb8\xbb\x0a\x06\x19\xcc\x0e\xe6\x90\x02\xb9\x0b\xf3\xd3\xd1\x16\x47\x8d\x48\x9c\xcd\x7d\x39\xf6\x8e\xc8\xf7\xa3\x4e\x4f\xbd\x7c\xec\x55\x86\x43\x2e\xb8\xdc\xe2\x82\xcb\x11\x17\xb4\x29\x7a\x4d\xee\x80\x3e\xb1\x17\x46\x84\xce\x9c\xed\x3c\xe8\x8f\x36\x8f\x0f\x7b\x02\x9d\x22\x07\xc1\x5b\xc9\xfe\x5d\xda\x27\x9c\x93\x0c\x0b\x75\xfe\x90\x0c\xb4\xa1\x69\xcd\x39\x50\x81\x08\x5d\x31\x5e\xe8\x4d\x5f\x09\xc6\x21\x93\x19\x57\x77\xec\xba\x50\xa9\x39\xc4\xa7\x6d\x23\x4a\x96\x14\x9c\x33\x6e\x2d\x50\x5f\x2a\xbf\xeb\x3b\x53\xb4\x8d\xdd\xca\xdf\x6a\xc2\x21\x3b\xdb\xc6\x38\x74\x07\xbc\x75\x87\x74\x87\xa0\xba\x36\xfb\xc8\x31\xad\x88\xb4\xda\x1b\x4b\xce\xee\x4b\x56\x41\x57\x92\x1b\xf2\x07\xa3\x93\xcf\x2d\x13\x99\x0a\xd0\x54\xef\xfe\xa9\x1d\x96\x63\x9c\xfb\xaa\x5b\x7f\x58\x28\x93\x39\xfc\x4c\x3e\x12\xdf\xf9\xcf\x0a\xef\x1f\xa7\x88\x92\xdc\x69\x7a\x03\x57\xb5\xfd\xaf\x4f\x5f\xc8\xc9\xb6\x13\xf6\x8f\x89\x41\x63\x04\x29\x60\x2f\x53\x3e\x92\x02\x5e\xa2\x21\x70\x2f\x80\x53\x9c\xef\x65\xcc\x99\x99\xf4\x12\x0d\x22\x54\xc0\x1a\xf8\x5e\xf6\x9c\x53\xf1\x12\x4d\x59\xe5\x0c\x8b\xbd\x0c\xf9\x55\xce\x78\x21\xa6\xf4\x2d\x4b\x5e\xe7\x39\xfb\x0e\xd9\x9b\x1b\x46\xd2\xee\xf9\x64\x9b\x61\x3a\x03\x9c\x53\x75\x5d\x16\xd8\xa5\x13\xca\x6c\xc4\xbc\x45\x57\x4f\xca\x79\x5f\x19\xa1\x3d\x05\xae\xa6\x0b\x34\xbd\x92\x68\xcd\x42\x9d\x0e\xaf\x6b\xc1\xd6\xe6\x60\xcc\xb6\x38\x09\x02\xe7\xc0\x1e\xf1\xed\x94\xc0\x3c\xca\x07\x17\x58\x9e\xd0\x34\x2e\xac\x0b\x55\xdf\x85\x32\xae\xb4\x79\x4e\x8a\x7e\x0e\xcb\xd6\x02\x25\xef\xf1\xfd\x3b\xa0\x6b\x71\x83\x5e\xc5\xd8\xf6\x1e\xdf\x93\xa2\x2e\xf4\x94\x58\x0b\x25\xb5\x93\x23\x29\x2b\x9c\x57\xf0\x6c\x26\x11\xba\x97\x49\x84\x3e\xd2\xa4\x56\xce\xf3\x9b\x84\xef\xd5\x13\x26\x7a\x95\xbc\x1a\x3b\xaf\xe3\xb3\x8f\x09\xe2\x1e\xc9\xa7\x8d\xe1\x27\xf3\xc0\xf9\x74\xf6\x9a\x86\x27\x56\xe9\xe8\xd4\xbf\x90\x85\xde\x2c\x50\x7b\xfe\xd4\x71\xda\xb5\x10\x9f\x32\x6a\x7a\x9d\xee\x1f\x35\xab\xc5\x33\x44\x2d\x52\xe7\xc7\x04\xad\xd3\xfa\x99\x83\x16\xff\xcd\x9a\x9b\x03\x0d\x8e\xd7\x39\xfa\x05\xbd\x6a\x55\x32\x1d\x9e\xcf\xe2\xbe\xda\x18\x0c\x18\x98\xdb\xce\x86\xde\x2c\x7b\x6d\x40\x72\x7b\x0f\x57\x42\x4a\x56\x24\x55\xed\xcd\xaf\x8c\xb7\xcd\x82\xd7\xf3\xb6\x54\x8f\xbd\xbd\x02\xd1\x2d\x61\xf7\x93\x0c\x75\x9d\x7b\x0b\x0f\xb6\x29\xda\x75\x59\x30\xa6\xc3\x4c\x01\xd9\x2b\xbb\x6e\x69\x8c\xa8\xb3\x69\x3d\x73\xb7\x40\xec\x56\x2e\xa7\x6d\x82\xbb\xbe\xe8\x3d\x2e\x3f\x4b\x51\x5f\x7e\x96\xd3\x7a\x6e\xbc\x73\x3d\x78\x72\x82\xfe\x04\x94\xca\x76\x5a\xb5\x4c\x2b\x42\x33\x44\xc4\x02\x55\x0c\xe5\x20\x7e\xaa\x50\x7a\x03\xe9\x2d\x62\xe6\x19\x9e\x7d\x07\x8e\x52\x5c\x01\x22\x34\x83\x7b\xc8\x50\x55\x42\x8a\x0a\x5c\xc6\x5e\x1e\xbf\x93\x10\x6f\x70\x05\x03\x0a\xdb\xa7\xe5\x41\x87\x54\x5e\x0c\x57\x75\x9e\x3b\x31\xaa\x7c\xce\x02\x97\x91\xd1\x1a\x91\x35\x9b\x4b\x8c\xcf\x3a\x58\x5f\x62\x63\x15\x61\xbe\x67\xf5\x64\xf7\x6f\x5f\x3c\xfe\xde\xb3\x17\x2e\xd5\xa3\x57\xeb\x06\xb9\x84\xb7\xa1\xed\xfe\x6d\x8c\x2f\xef\x74\x1f\x2f\x38\xf7\x7f\x8f\x68\xd8\x9d\xa7\x97\x76\x86\x54\x21\xb8\xd6\xf1\x7e\xf4\x15\x26\xcb\xe9\x72\xe7\xb6\x72\xaf\x66\x8f\x47\x8b\x59\xf9\xe7\xd2\x97\xfd\x9a\x54\x96\xa4\x1e\x56\xf0\xac\xe3\x17\xea\xcb\x91\x06\xe2\xb8\x69\xf6\x2a\xf2\xbb\x93\xa6\x9d\xd6\xb4\xe9\x78\xd1\xb7\x2d\xe8\x06\x3a\xed\xdc\x81\xe5\x60\xe7\xb0\xd5\x3a\x2b\x60\xf0\xf5\x44\xfe\x79\x03\xcb\x91\x70\xc5\x89\xe0\xa0\x22\xf7\x07\xcd\x1f\x3c\x09\x0e\x5d\x5b\x10\x70\x46\xa1\x7b\x4f\x07\xce\xb8\x4b\xd7\xe8\x01\xe7\x3e\xe8\xf6\xdc\x76\xaf\xf0\x47\xee\x30\x8e\xfb\x3a\xa8\xd9\xed\x02\x0a\x20\xa5\x66\xee\x53\x65\xe3\x2c\x11\x4a\xf2\x91\xb5\xa1\x18\x7c\x0d\x68\x5d\x3c\x4a\xfa\xf8\x45\xe8\x50\xc8\x1d\x47\x3a\x18\x8f\xd0\xd7\xde\xc5\xed\xa7\xf1\x71\xd3\x4c\x07\x54\x98\x46\xeb\xb0\xb7\xb8\x5d\x06\xa3\x31\x69\xfd\x05\x35\xba\xc6\x4a\x0e\x69\xb8\xc1\x3b\xaa\x5d\xbd\x0e\x57\xe4\xda\xf5\x5e\x90\x3a\xe0\x96\xbc\x1c\x78\x89\x09\x9e\x5f\xa2\x24\x39\xd7\xb1\x76\xc8\x90\xb4\xee\xdd\x78\x14\xdc\xaf\x24\x17\xc0\xd5\x0f\xba\x9c\xd1\x8e\xaa\x41\x3d\xae\x38\x5c\xc6\x81\xac\xe9\x7f\xc0\x4b\x43\x1d\xd5\xe0\xba\x5c\xb1\xb8\x05\x16\x01\x66\x81\x85\xc9\x9a\xed\x70\x5c\xba\x34\x0f\x42\xce\x88\xa6\x68\xed\xda\xd1\x28\x2c\xe7\x0d\xd9\x19\xed\xa8\x1a\xd3\xe3\x8a\xc3\x95\xb5\xa3\x07\x29\x09\x06\xcd\x8c\x45\x01\xb9\x77\x28\xed\x60\x4b\x5c\xf6\xef\x59\x22\x41\x7b\x09\xdb\xd2\x96\xbd\xb6\x3f\x0a\xd1\xb9\x17\xe9\x20\x2d\x71\xd9\xbf\x3b\x89\x04\xed\xab\x69\x68\xcb\x5e\x9f\x1b\x83\x18\x9e\xe1\xce\xd1\xbd\xd7\x89\xad\x8e\xc8\x70\xff\xb5\x44\xad\x9b\xcb\x13\x05\x7a\xc1\x49\x81\xf9\x43\xb0\xfb\x3a\xaa\x86\xf5\xb8\xa2\x70\x3f\x00\xce\xc2\xd2\xc2\xd2\x96\xe6\xce\xb1\xe5\x88\x44\xf4\x9f\x91\x34\xa2\xa6\x2d\xc3\x5b\xcc\x28\xc4\xcb\xde\x9e\xbe\x74\xf6\xf4\xe5\x5e\x7b\xfa\x52\xbf\xf9\xb9\x58\x8a\x62\xb0\xec\x68\x1c\x56\x7d\x6d\x9e\x2e\x3b\x30\x4d\xb2\xbf\x85\x6f\x19\xe2\x56\x4e\xef\xc1\x4e\xfe\xb5\x44\xad\xa2\xcb\x13\x07\x1a\xa8\xe8\xe8\xb7\x53\x39\x23\x21\xa0\xf6\x7e\x18\xb3\x7f\xc3\xfb\x7f\x68\xe5\x46\x04\xbf\xec\x9e\x4e\xdf\x30\x24\x87\x86\xee\xd0\xd0\x1d\x1a\xba\x43\x43\x77\x68\xe8\x3a\xe5\x0e\x0d\xdd\xa1\xa1\x53\x7f\x87\x86\xee\xd0\xd0\x1d\x1a\xba\x43\x43\x77\x68\xe8\x0c\xde\xd3\x36\x74\xff\x0b\x00\x00\xff\xff\xc6\x45\xc4\x2c\xe6\x40\x00\x00")

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

	info := bindataFileInfo{name: "templates/model.gotpl", size: 16614, mode: os.FileMode(420), modTime: time.Unix(1525201046, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\xc1\x6a\xe3\x30\x10\x86\xef\x7a\x8a\xc1\x84\xa5\x85\xd4\xd9\x73\x20\x87\xa5\x0b\x25\x87\x2e\x25\x65\x4f\x21\x07\x45\x9a\x38\xa2\xf2\xc8\xc8\xe3\x2c\x41\xe8\xdd\x17\x5b\xc9\xc6\xed\x3a\x2d\x94\x5c\x4a\x9b\x9b\xf4\x6b\xbe\xd1\x7c\xd1\xc1\x95\x54\x4f\xb2\x40\x08\x01\xf2\x47\xe4\xfc\xd6\xd1\xc6\x14\x8d\x97\x6c\x1c\xe5\xbf\x64\x89\x10\xa3\x10\xa6\xac\x9c\x67\xc8\x0a\xc3\xdb\x66\x9d\x2b\x57\x4e\x64\xe5\x3c\xb2\xbb\x31\xa4\x26\x68\xb1\x44\x62\x69\x33\x21\x94\xa3\x9a\x81\x9c\x76\xea\x91\xbd\xa1\x02\x66\x90\x2d\xbb\xf5\x2a\x83\xc9\x04\xc8\x59\x43\x3c\x85\x9d\xf4\x6a\x8b\xea\x69\xac\x51\x6a\xe5\x34\x0a\xb1\x93\x1e\x3c\xda\xae\x79\xbd\x35\x55\xbd\xc0\xc2\xd4\xec\xf7\xf0\xaf\x43\xbe\x18\xca\x85\xd8\x34\xa4\xc0\x90\xe1\xab\x6b\x08\x42\x00\xc0\x19\xd2\xec\x2d\x56\x88\xa9\x3c\x04\xf0\x92\x0a\x84\x11\x12\x1b\xde\xb7\x32\xc6\x30\x3a\x52\x61\x3a\x4b\xca\x9e\x41\x5a\x5b\xa9\xf8\x06\xcc\x06\xea\xad\x6b\xac\x4e\x64\xf4\xfd\x93\x30\x6a\x8b\xfb\x6c\x18\xe5\x0f\xcd\xda\x1a\x75\xef\x34\x1e\x30\x83\x23\x2c\x43\x78\x56\x17\xe3\x5c\xa7\xe5\x0a\x66\xf0\x6d\x78\xbc\xd0\xf1\x7a\x57\x2b\x18\xae\x2c\xd2\x69\xa0\xfc\x87\xb5\xee\x4f\x7d\xeb\x51\x32\x5e\xc3\xf7\xe3\x28\xed\xaf\x1f\x4d\xa1\x94\xd5\xb2\xee\xfe\xdb\xd5\xda\x39\x0b\x27\xf6\x91\x7f\x10\x57\x49\x8f\xc4\xad\xa9\x53\x97\x3b\x64\xc8\x54\x47\xca\xfa\x2d\xce\x5a\x9b\x13\x0d\xab\x3b\xd0\xfb\xda\x5e\x02\xb3\x56\xd5\xe1\x5c\x8c\x19\x4c\x81\x7d\x83\xe3\xff\x9a\x22\xe9\xa1\xbb\xbc\xd8\x8e\x63\xf1\x4a\xfa\x86\xd7\xdf\x95\x3e\xe7\x35\x45\x97\xf0\xda\x74\xa4\x0f\xec\x35\x09\x79\x90\xac\xb6\x5f\x3e\xde\xf5\xce\x7e\xa2\xc5\x33\xef\x2c\x45\x97\xf0\xaa\x3b\xd2\xa7\xf2\x7a\x87\x3c\x28\x75\x81\xec\x0d\xee\x2e\xa2\xb5\x40\xfe\x6c\x4e\xef\x25\xed\x5f\xf5\xda\x1e\xb8\x90\xdb\x52\xd2\xfe\x03\xfb\x4d\x5e\xe6\xb4\x71\x5f\x3e\x06\xd2\xd3\x87\x57\x6f\x33\x84\xe3\x2a\x8a\xbf\x01\x00\x00\xff\xff\x45\xe2\xe7\x2b\xec\x0a\x00\x00")

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

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 2796, mode: os.FileMode(420), modTime: time.Unix(1528338461, 0)}
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
	"templates": &bintree{nil, map[string]*bintree{
		"README.md": &bintree{templatesReadmeMd, map[string]*bintree{}},
		"identities_registry.gotpl": &bintree{templatesIdentities_registryGotpl, map[string]*bintree{}},
		"model.gotpl": &bintree{templatesModelGotpl, map[string]*bintree{}},
		"relationships_registry.gotpl": &bintree{templatesRelationships_registryGotpl, map[string]*bintree{}},
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

