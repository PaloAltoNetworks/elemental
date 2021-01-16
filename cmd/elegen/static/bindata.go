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

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1578956520, 0)}
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

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 5396, mode: os.FileMode(420), modTime: time.Unix(1608240120, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x69\x73\xdb\x38\x96\x9f\xa3\x5f\x81\x56\xb9\x13\x2a\x23\xd3\xf9\xac\x1e\x4f\x55\xce\x5e\xd7\xe6\x70\xc5\x99\x9e\xaa\x4d\xa5\x26\x30\x09\xca\xe8\x50\x00\x1b\x00\xed\x78\x55\xfc\xef\x5b\x38\x48\x02\x3c\x41\x59\x4a\xbb\x77\xac\x0f\x89\x84\xe3\xe1\x5d\x00\xde\x7b\x00\x9e\x33\x18\x7d\x83\x6b\x04\xb6\x5b\x10\x5e\x20\x11\xbe\xa4\x24\xc1\xeb\x9c\x41\x81\x29\x09\xdf\xc3\x0d\x02\x45\x31\x9b\xe1\x4d\x46\x99\x00\xc1\x0c\x00\x00\xe6\xc9\x46\xcc\xf5\xb7\x35\x0d\x61\x46\x19\x12\x34\xc4\xf4\x04\xa5\x68\x83\x88\x80\x69\x59\x8b\xc5\x55\x7e\x19\x46\x74\x73\xb2\x4e\xe9\x25\x4c\x39\x5e\x93\x93\xcd\x9a\x9e\x5c\x72\x4a\xda\x8d\x36\x58\x44\x57\x28\x4d\xaf\x4e\x22\x9a\xdd\x72\xc1\xf2\x48\xe4\x0c\xe9\x86\xdb\xed\x31\x60\x90\xac\x11\x08\x2f\x32\x14\x85\x9f\x6e\x33\x74\xce\xe8\x35\x8e\x11\xe3\x12\x49\x05\x4d\xd2\x01\x8a\xa2\xee\x82\x48\x0c\x8e\x4d\x6d\x13\xc4\x6f\x30\xc5\xb1\xa2\xd4\x13\x50\x51\xcc\x16\xb3\xd9\x76\x0b\x8e\x52\x28\x10\x17\xbf\x21\xc6\x31\x25\x60\x75\x6a\x20\xbe\x55\xc5\xcf\x85\x60\xf8\x32\x17\x88\x97\x0d\x24\x0f\xb7\x5b\x33\xf8\x11\x5e\x82\x23\x44\xf2\x8d\xec\x77\x99\xe3\x34\x7e\x4d\xf2\x0d\xd7\x20\x9a\xa0\x8b\x62\x76\x72\x22\xc5\xa3\x7a\x28\xaa\x41\x51\x00\x86\x32\x86\x38\x22\x82\x03\x71\x85\x40\x46\x39\xc7\x97\x29\x02\xd7\x30\xcd\x11\x07\x09\x65\x00\x96\x58\x28\x62\x74\xf7\x0a\x33\x23\xd9\x79\x38\x13\x12\x62\x0b\x3e\x17\x0c\x93\xf5\x6c\x16\x51\xc2\x4b\xb9\xd7\xec\x3b\x22\x70\x83\x96\xe0\x48\x8d\x26\xa9\xd0\x9d\x7f\xd3\x83\x1b\x16\x1a\xb4\x89\x1e\xa9\x89\xb1\xee\x2a\x1b\xe8\x6f\x45\x11\x9a\x41\xea\x2e\x2d\xac\x4e\x35\x29\x65\x8f\x52\x38\xb5\x6c\xea\xef\x33\x89\x2d\x4e\x00\xa1\xc2\xc8\xe6\x1d\x8d\x51\x1a\xbe\x42\x02\x46\x57\x28\xae\x19\x6b\xd7\xbe\x26\x02\x8b\x5b\xc3\x9c\xb3\x18\xa9\x9f\x4d\xd4\xab\x72\x9a\xa8\xdf\xf4\xf2\x77\x14\x89\x70\x76\x0d\x99\x1f\xbc\x53\x50\xcd\x94\xb0\x2a\xdc\x2a\x62\x64\xd3\x15\xa8\x34\xd0\x02\xf5\x11\x71\x21\x6b\x8b\x62\xbe\x54\x4d\x5f\x42\x81\xd6\x94\xdd\xae\xba\x9a\xd2\x9c\x45\x95\x90\x75\xfb\x73\x3d\xd5\x57\x6d\xd0\xa6\xa6\x6e\xc9\xf0\x35\x14\xb2\x65\xb3\xa1\xae\x28\x8a\xe5\xac\x98\xc8\xeb\xed\xb6\xab\xc5\x19\xff\x48\xa9\x18\x93\xc5\x79\x9a\x33\x98\x82\xa2\x78\x8b\xb9\xb0\xa5\x01\x41\x2a\x4b\x68\xe2\xd1\xb7\x52\x74\x9f\x31\x3e\x7f\x79\xda\xdb\x52\x12\x7c\x72\x02\x2c\xed\x10\x39\x23\x5a\x35\x70\xa7\x6a\x70\x80\x89\xfa\x29\xb1\x0d\x67\x49\x4e\x22\x10\x50\x4f\x5c\x16\xd5\x48\xc1\xa2\x5b\x6f\x94\xcc\x34\x16\xfd\x30\x6b\xf5\x9b\x69\xfc\x5f\xd2\xac\xc6\x1d\x82\x8c\x62\x22\x10\x03\x82\x02\x08\xe4\xf2\xab\x10\xf6\x43\x71\x3a\x49\x72\xf0\x0e\x72\x12\x0c\x2f\x53\xc4\x4b\x9a\x14\x1a\xab\x53\x00\xb3\x0c\x91\x38\xf0\x03\xbe\x2d\x96\x80\x86\x61\xb8\xb0\xd9\xf2\x58\x82\x32\x84\x3f\x57\xd0\x0c\x50\xee\x88\x49\x50\xf5\x13\x02\x82\x6e\xf4\xe8\x46\x8e\x87\xe2\x83\xc6\x25\x28\xc7\x0f\xc3\xb0\x9b\x25\xa3\xac\xa2\xb9\xb8\x23\xa7\xe4\x96\xf1\xef\xa5\x64\x85\x04\xa4\xd7\xf9\x12\x2f\xbd\x36\x95\xe3\x54\xc3\xd0\x5c\xa8\x0e\x61\x30\x34\x5b\x16\x1a\x7e\xe1\xe8\x29\xcd\x85\x11\x87\x9a\x6f\x11\x25\xd7\x88\x09\x5b\x1a\x4a\x13\x49\x1f\xdd\xbb\xb1\x5b\xfe\xdb\xaf\x76\x0a\x13\x97\x9f\x1b\xf8\x0d\x05\x03\xcd\x97\x20\x45\x24\xa0\x8b\x9a\x85\x58\x76\x7b\xf6\x0b\xc0\xe0\xef\xa6\xee\x17\x80\xff\xf6\x37\x97\x85\x9f\xf1\x17\x70\x0a\xe8\x67\xfc\x65\x90\x35\xaf\x50\x02\xf3\x54\x7c\x60\x31\x62\xce\x32\x13\xeb\x0a\x40\x65\x0d\x26\x6b\x90\x60\x94\xc6\xbc\xd4\xd6\x88\x12\x81\xc8\x0e\xfc\xb1\x07\x0c\x16\xe0\xf3\x17\x6d\x06\x34\xd6\x98\xb2\xb8\x26\xa9\x32\x6d\xf4\x28\x0e\xde\xa5\xf1\xe5\x58\x55\x4b\xbb\xab\xd9\x45\x34\x27\x34\xe5\x9f\xe8\x45\x06\x19\x47\x0e\xd5\x9e\x6b\xb7\xd1\x25\x14\x4b\x0d\xd2\x60\x7c\xa7\xef\xc9\x09\xf8\xd0\x5e\xb1\xc1\x0d\x4e\x53\x40\x49\x7a\xab\x38\x0b\x4d\xd5\x1a\x5f\x23\x62\x38\x1f\x82\xf7\x54\x7f\x05\x1b\x04\x09\x07\x52\x4f\x18\x32\x45\x1c\xed\x20\x8b\x92\x05\x81\x91\x6d\x18\x86\x9a\xed\xbe\x6b\x81\xd2\xdd\x29\xf4\xdf\x59\x99\xc3\x06\xce\x72\x6d\x09\x83\xa7\x23\x38\x80\xa2\x18\x5e\x21\x4a\x53\xd8\xd6\x85\x6b\x53\x76\x57\x8d\x37\xb0\x83\x05\xc0\x44\x74\xed\xa5\x48\x84\xcf\xcf\xcf\xce\x48\x42\x43\xcb\x24\xd7\xe6\xbc\x51\x5c\xcb\x3b\x50\xdf\x8f\x30\x7f\x4d\x22\x76\x9b\x09\x29\x16\xc9\xc2\x04\xa6\x1c\x79\x58\x9c\x4d\x4b\x73\x23\x9b\x48\x1a\x61\xb3\x5b\x69\x0d\x8e\x1b\x36\xc6\x9a\xcf\x23\x49\x5e\xd3\x03\xaa\x3d\x95\xa6\xdf\x71\x5c\x14\xc6\xaa\x0b\x0d\x31\xca\x8e\xeb\xa0\xef\x14\x08\xa6\xcc\x71\x8b\x0f\xdb\xad\x72\x3f\x3e\xd1\x37\x6a\x02\x1c\x49\x3e\x1a\x2e\x84\x4d\x96\x49\x8e\x2b\xa4\xcb\xa1\xa5\x24\xbe\xfe\xce\x29\x59\xcd\x8f\xe7\x60\xc3\xd7\xd2\x49\x55\xdf\x2f\x55\xe1\xbf\x15\x5b\x8c\x06\xcc\xbf\x1a\x2d\x79\x8f\x6e\x46\x58\x5b\x9a\x3a\x72\x73\xef\xdf\xb0\x24\x4e\x4a\x8b\x46\x00\x06\x8b\x61\x20\x0d\x65\x7a\x3c\xd4\xb6\x9e\x4f\x36\x23\x56\x03\x1a\xb8\x9c\x59\x2b\xe8\xb1\xed\x58\x4a\xbe\x4b\x9d\xe3\x94\x59\x8e\x28\x08\x46\x04\xbe\x70\x16\x6a\x23\x7a\x7e\x45\xf3\x34\xfe\x17\xc3\x02\x9d\x11\x2c\x30\x4c\xf1\xff\x22\x26\xc5\xa9\x3c\x55\x39\x94\x8e\x11\x34\x94\xe7\x28\x3c\xcf\x2f\x53\x1c\x49\x6a\x40\x03\xec\x11\x26\x58\xad\x4f\x37\x1d\x60\x91\x70\x80\x37\xfb\xe2\xc4\x74\x77\xca\xbb\xca\x8e\xed\x5d\xc5\xaf\xc8\xcc\x6a\x1f\xaf\xb1\xd3\xf6\xef\x73\x0b\xcb\x25\x69\x50\x5b\xf6\x64\xe5\x03\xc7\xcc\x6f\x2e\x4c\xbe\x84\x25\xb8\x61\x6f\x68\x87\xdb\xa1\xeb\x09\x07\x39\xc1\x7f\xe4\xa5\xcf\x23\xfb\x4c\xa4\x55\x76\x09\x16\xc0\xb5\x31\xb4\x9f\xa8\xfb\x5a\xd8\x94\xca\x59\x6e\x0e\x61\x35\x40\xdd\x28\x7c\x59\xee\xfc\xe5\x3c\xae\xa4\x2c\x17\x9e\x06\x88\x79\x2b\xb4\xb3\x0b\xc3\x2e\x90\xb0\xb0\xe4\x48\x1c\x86\x61\xce\x30\x01\x8e\x41\x69\x0a\xf8\x71\xcd\x8f\x5d\xe0\x14\xe0\x78\x98\x29\x27\x27\xe0\x57\x24\x5e\x5c\x7c\x78\x0f\xf0\x26\xd3\x6a\xaa\x29\x96\x4b\x33\xd8\x40\xc6\xaf\x60\x2a\xc5\xa9\xbc\xc9\x04\x46\x48\x59\x55\x9f\xae\x30\x07\x98\x83\x9c\x6b\xb3\x4c\x30\x48\x78\x06\x19\x22\x42\x5b\x55\x12\x11\x70\xf6\x4a\xd6\xbd\xa3\x64\x4d\x5f\xbd\x38\x7b\x05\x20\x07\x1f\x2e\x51\x24\xce\x5e\x79\x33\xca\x60\x17\x2c\x40\x50\x61\x20\xfd\x1c\xc4\x18\x65\x15\xbb\x70\x02\x28\x38\x3d\x05\x04\xa7\x96\x2d\x63\x14\x83\xe0\x74\x29\xff\xb1\x6d\x12\x2e\x17\xac\xc7\x1b\x89\x59\xbd\x82\x0e\xae\xe8\xa5\xf2\xf9\x6f\xb7\x33\x6b\x95\x0b\x2f\x04\x65\x28\x06\x8d\x52\x4b\xb4\xa6\x46\x52\xa2\x84\xdb\x12\xe6\x4f\xa7\x60\x3e\xb7\xa8\xe3\xdd\xcd\x4e\x95\xe4\x42\x6d\xf6\x9e\xc5\xff\x85\xbe\x07\xdd\x00\x4b\x1b\xcd\x99\x53\x06\x8b\x5e\xd8\xdd\xa0\x9a\x3a\x36\xfc\xd3\x9e\xb4\x5c\x4b\x46\x6b\xe2\xc5\xbd\xd6\x44\x83\x5d\xc0\xe0\x8d\x66\xf1\x47\x78\xb3\xd0\x7a\xe8\xab\x86\xfb\xd0\x40\x9c\xc8\x31\xb5\x4b\x7f\x13\xfe\x93\x18\xc6\x04\x7c\xf1\x8b\xaa\xf8\xa9\x67\x78\xc4\x98\x23\xf0\x03\xeb\x71\x8f\x12\x9f\xf6\xa8\x56\x28\xf5\x74\xd1\xa9\x8b\x13\x21\xed\xae\x8b\x46\x11\xfd\xf6\x89\x2e\x1f\xe6\x0a\xb2\x38\xa2\x31\x8a\x9b\xde\x8c\xb2\x6f\xbd\x15\x6d\x67\x17\xa6\xb1\xb0\xbf\x48\xd1\x35\x52\x91\xf6\xe6\x84\x92\x15\xe1\xcb\x14\x72\xae\x65\x76\x56\xcf\x28\x4f\x1c\x2b\xd8\xad\xfd\xbe\xdc\x8d\xfb\xfd\x9b\xb9\x3f\x97\x7b\x03\x26\x65\x98\xb8\x27\x70\xe2\x45\x87\x22\xe4\x5e\x44\x48\x26\x9a\x29\xb6\x71\x50\xd6\xf2\x88\xe1\x4c\xd4\x07\x4c\xaf\x68\xe4\x46\x98\x68\x94\x2b\x1b\x54\xb5\x49\x28\xb3\x2c\x19\x5f\xa1\xbf\xa2\x51\x9f\xb8\xbf\x4a\xa2\x78\xf4\x02\x46\xdf\x04\x8e\xbe\xf1\x01\xec\xbe\x3a\x47\x0d\x13\x49\xf7\x5d\xab\x15\x8e\x7d\xc8\x26\x1b\x11\x5e\x64\x0c\x13\x91\x04\xf3\xbf\xff\xcc\x57\x3f\xf3\x7f\xcc\x97\x80\x86\xb5\xc9\xae\xbc\xa0\xba\x48\x5b\xb6\x8b\x96\xa8\x3c\x17\xd1\x4a\x66\xda\xff\xfa\x15\x11\xc4\xa0\x40\xbf\x22\x21\x10\x03\x61\xcb\xbd\xd2\x56\x59\xe7\xb2\xd7\x8c\x9f\xb5\x1a\x98\x25\x87\xa1\x08\xe1\xeb\xa6\x45\x7a\x34\x6c\x69\x75\x01\x0c\x16\xee\x38\xe5\xd1\x9d\xcb\xd2\x1e\xbb\xa0\x11\x56\x69\xb3\xe0\x62\x80\x05\x17\x3d\x2c\xa8\x8c\xf2\x8c\xd1\x0c\x31\xe9\x4c\x79\x30\x02\xe4\x5c\x6a\x42\x1d\xe8\x53\x26\xbd\x3f\x7b\x7a\xb0\x51\xa1\xf9\xf7\xf5\xf9\x66\x8b\x51\x95\x8d\xda\xbb\x8f\x59\x10\x1a\x53\xa3\x39\x33\x20\x89\x41\xd0\x37\x3d\x16\xed\x2a\x7d\x1a\xb7\x30\xfc\xec\x8c\xc1\x72\x5d\xd4\xbd\x61\x29\xf3\xaa\x6c\x8f\xe2\x32\x90\xbf\xd7\xf0\xe9\xc8\x4c\xf6\x8a\x9a\xea\x26\x76\xec\xd4\xb2\xc8\x52\x44\x4c\xe7\x85\xb4\xcd\x9e\x59\xa6\xd1\xc9\x09\x20\x34\xc5\x44\xac\xc0\x9a\xea\x2b\x11\xbc\x69\x37\x3d\x1e\x8f\x76\x5a\x10\x41\xc7\xad\x84\xc1\x75\xa1\xd9\x11\x27\xe0\x0a\xf2\x73\x86\x12\xfc\x1d\x04\x3a\xe6\xa6\x34\xc9\x09\xb9\x2d\xc0\xfc\xe9\xbc\xdd\xbd\xad\x5e\xab\x1e\xb5\x5b\xb6\x06\xb6\x4d\xae\x61\x88\x8f\xbd\x41\xba\xe1\x99\x9e\xe2\xc2\xb1\x8a\x33\x65\x16\xfb\xf0\xbc\xb0\x4f\xb9\x92\xfa\x8c\xcb\x28\x8a\xe5\x28\xdd\x60\x11\x5d\x81\x64\x5f\x62\x8a\x20\xd7\x57\x30\xca\x59\x3b\x5f\x39\xf5\x7b\x10\xa5\x66\x45\x9b\xcd\x63\x3e\x98\x8f\x50\x07\x61\x3f\x1e\x74\x16\xf7\x20\xe0\xd2\xef\xcb\x2c\x23\xb0\x11\x7f\xd6\xab\x95\x29\xb1\xa4\x82\x74\x89\x5e\xb7\x60\x5d\xbe\x81\xec\x1b\x8a\xa5\x4b\xf7\x15\x95\x91\xed\xaf\xad\xe5\xbe\xac\xf2\x8f\xd1\xb4\x30\x08\x2a\x18\xd6\xda\x53\x55\x97\x51\x75\xb6\x00\x81\x74\xc4\xaa\x08\x05\x98\xe2\x6f\x35\x1c\x2b\x3b\x54\x3f\x10\x1f\x28\x74\x44\x04\x9c\x5a\x64\x9a\xae\xc6\x14\xea\xec\x34\xe2\x32\x4a\x3b\xe9\xb5\xa4\x22\x09\xe6\x39\x51\xb2\x11\xb4\x1c\xc1\xba\x8e\xf4\xa4\x03\xf4\x13\x35\x33\x9f\x0c\x6e\xaa\x4f\x40\xf0\x33\x5f\xac\xc0\xcf\x7c\xde\x34\xb5\x14\x39\xad\x08\x85\xaf\x0b\xa7\x3c\x87\xa6\xfa\xc4\xe8\x2e\xea\x63\x7a\x4f\x50\x9f\x16\x06\x7f\x2d\xf5\x31\xe8\xef\x5d\x7d\x0c\x23\xef\xb3\xfa\xd8\xb5\x52\x97\xce\xa1\xdc\x3f\x60\x96\xa5\xfa\x12\x0d\xa1\xaa\x69\x1d\x14\x86\xc0\xe3\x4c\xb4\xbc\x8c\x32\xf1\x18\x41\x0d\x1e\x18\x33\x6d\xc8\xe4\xb1\x62\xc7\xc7\xfa\xb4\x85\xd7\xf7\x18\x7d\x35\xc6\xf4\xab\xb5\xe5\x27\x3d\xb2\xed\x1a\x9d\xf1\xd7\x7f\xe4\x30\x0d\x6c\x7f\x69\x61\x89\x3f\x83\x04\x47\xc1\x3c\x82\x44\xda\xa3\x99\x62\x5e\xc2\xe8\x06\x40\xa0\xa9\xb8\xc1\xe2\x0a\xc4\x38\x49\x10\x43\x44\x54\x77\xac\xe6\xce\xa9\x31\xa7\xea\xd0\x4b\x8f\xee\x77\xe6\x3c\x73\xb7\xf5\x16\x2d\xbc\x3f\xb2\xea\x2a\xf0\x1d\x76\xef\xfe\x68\xd5\xc8\xb6\xdd\xb5\x5d\xf7\x02\x7b\xea\x05\xcd\x0e\x32\xb4\xca\x86\x8e\x04\x5e\x21\x94\x35\xae\x93\xc5\x08\x65\xfa\x06\x15\x1e\xb9\x41\xa5\x6e\x7e\x7a\xaf\x91\x7a\x20\xdf\xb3\x57\x37\xc0\xfa\xe8\x91\x35\x6f\x1f\x15\xb3\xd9\x23\x73\x53\x62\xf8\x6c\xb6\x98\x3d\xa2\x61\x39\xf2\x19\x11\x34\xa0\xb9\x58\xcc\x66\x8f\x3a\xee\xeb\xd4\x8d\x24\xf1\x18\x71\xd7\xa7\xc4\xc4\x4c\x6a\xbd\x49\x0c\xd2\x30\x99\x29\x25\x6a\x63\xed\xb7\xb3\xd9\x23\x01\xd9\x1a\x89\x65\x19\x1a\x76\xae\x5b\x87\x8a\xc3\x74\x31\x7b\x64\x62\xc7\x3f\xd5\x0c\xd4\x73\xd5\x09\x88\xfc\xd3\x5a\xaa\x51\xa6\x44\x3e\x34\xbe\x59\x7f\xe5\x7a\xbb\xd0\x42\x78\xaa\xef\x94\x3d\xd5\x38\x0d\xdd\x25\x53\xb3\xd6\xdc\x09\xd1\x57\xb7\xd5\x49\x1b\x8e\x0d\x9f\xa3\x9c\xe9\x15\x82\x24\x94\x6d\x74\xe8\x8a\xeb\x00\x74\xc5\xf9\x9a\x4c\xef\xf8\xaa\x19\x2a\x68\x44\xef\xd5\x0f\xb5\x66\xd6\xcb\xac\xda\xbf\xf8\xb6\x3c\x68\xfc\x23\xc7\x0c\xc5\xaf\x87\x1a\xee\xb8\x5f\x57\x81\xaf\x4f\x0c\x12\x8e\x25\xd5\x4e\x5d\xf8\xfa\x7b\x46\x39\xaa\xb7\x2c\x53\xfc\xd1\xe0\xe4\xb6\x46\x7f\x00\x7d\xc7\x7a\xae\x9d\xe5\xb9\xb5\x0a\x1a\x15\xa9\x51\x2f\xf9\x51\x82\x32\x5b\xbe\xe3\xe1\x2c\x7b\xd6\xa2\x7e\x13\xc0\x61\xd5\x69\xa3\x20\x34\xf7\x24\x5b\xbb\xb4\xb3\x2b\x6b\x5a\x28\x03\x41\x4d\x0f\x22\xf9\x66\xbe\xb8\x3b\x39\xbc\xd7\xae\x91\x54\xfd\x00\xb2\x6a\x92\x04\xde\xa0\x49\x02\xfa\x84\x37\xe8\xbe\x8a\xe7\xbb\x40\x8c\xc0\x74\xbe\xb0\x4b\x53\xcc\xc5\x7c\x31\x81\xc2\xd7\x06\xcc\xbd\xa1\xb2\xa6\x05\x13\x81\xd6\x88\x4d\x12\xd8\x19\x11\xf7\x90\x92\x24\xa5\x50\x4c\xa2\xe3\x8d\xec\x71\x3f\x28\x19\x22\x8c\xa1\x64\xee\x75\x9e\xee\xa2\x56\xd3\xfd\x11\x71\x24\xcc\x91\xce\x1b\xca\xfe\x07\x31\xaa\x9f\xc2\x8c\x46\x47\x6a\x2e\x76\xb7\x0c\xeb\xcd\xa7\x87\x41\xd6\x4e\x74\x6a\xbe\xb4\x98\x02\xac\xa8\x8a\xe7\xc4\x64\x28\x79\xab\x66\x61\xa3\xf0\x1d\xcc\xea\xe5\xd4\x04\xd3\x78\x7e\x69\x5d\x19\xef\xe6\xde\xd6\x26\x59\x76\x68\x1d\x7b\x03\x75\xe1\x9f\x08\x4c\x72\xd4\xc0\xda\x97\xdb\x3c\xbf\xec\x62\x2d\xcf\x2f\x7f\x20\x1f\xc3\xe7\x69\x4a\x6f\x50\xfc\xf2\x8a\xe2\x08\x71\x9f\xf9\xa2\xb7\x9c\x33\xa2\xae\xa7\x4f\xda\x78\x96\xf5\x51\xa3\xec\xf7\x3b\xc5\xa4\x85\xc0\xd7\xf9\x12\xcc\xbf\x4a\x68\xc5\x52\x99\x66\xcf\x73\x41\xd7\xe6\x40\x25\x1e\x98\x7b\x63\xec\xf0\x62\x02\x64\x5e\x2c\x38\x87\x42\x2e\xe1\x7e\x8b\xc5\x52\x9d\x1f\x36\xc7\xf8\xda\x51\xfc\x0e\x71\x0e\xd7\x48\xd7\xca\x4a\xcb\xfe\x39\x00\xd9\x6b\x01\xc2\x77\xf0\xfb\x5b\x44\xd6\xe2\x0a\x3c\xf3\x21\xfc\x1d\xfc\x8e\x37\xf9\x46\x77\xf1\x25\x5f\x96\xd6\xe3\xc8\x12\xe5\x5f\x1e\x8a\x22\x4c\x26\x51\x84\xc9\x8e\x14\x55\xe3\x1c\x9c\x22\xf8\x5d\x2d\x19\xe0\x59\xf8\xac\xcf\x12\xf6\xdf\xee\x8c\x08\x27\xec\x76\x95\x04\x7f\x33\x2f\x19\xf7\x46\xae\x89\x08\xf8\xe2\xec\x6d\x69\x2c\xa5\x07\x15\x34\xb0\x5e\xec\x59\x4a\x63\x5a\xb8\x4f\x99\x69\x25\x9d\x2e\xb3\x12\x8b\xfd\xcb\xcc\x13\xe5\x5d\x44\x56\x23\xfd\x63\x44\x56\x5d\x42\x0f\x41\xeb\xe9\xb5\xba\xa4\xae\x9e\x3e\x0f\xbd\xbf\xae\x19\x21\xc1\x6d\x4a\x62\x15\xe5\xd6\xbd\xf3\x9a\x7c\x5d\xe8\x6b\x56\xfa\x92\x79\xdc\x63\x47\x7a\x30\xc1\xa1\xf7\xba\xa2\x54\x21\x56\xc5\x59\x75\xc0\xa1\xe6\x83\xfd\x78\xfa\x65\xce\x05\xdd\x94\x87\xe8\x35\x84\xb0\x8e\xda\x6e\x60\x96\x61\xb2\x56\x2f\xb0\xd5\x3d\xaf\x1a\xd2\x3b\x5d\x15\x9a\xff\xc1\xbc\x7e\x9c\xdf\x42\xa7\x11\xd3\x2d\xa1\x76\x8b\xc2\xc0\x2d\x05\x42\xf7\xc6\xe2\x2e\x8e\x9b\xf3\x78\xd7\xe8\x5f\x80\x7f\x58\xc7\xf2\x26\x0a\xe7\x36\x31\x23\xd8\x30\x50\x47\x5f\xfb\xb6\x63\xa3\xd7\x2e\x97\xfc\x64\x0d\x4e\x70\xa4\x38\xfb\x86\xb2\x2a\x8e\xe3\x5c\xa1\xa8\x4a\x9d\xe6\xd5\x1d\x2b\x1d\x1a\xac\x8f\x3b\xd4\x63\xf8\x6f\xe8\xb6\x8c\x57\x8d\x5d\x65\xea\xc3\x21\x50\x80\xda\x97\x21\x7a\xd0\xa9\x23\xa8\xd7\x4b\x40\xbf\x19\xf1\xf7\x0e\x5c\x87\xac\xde\xc1\xec\xb3\x1c\xea\xcb\x2f\xb2\x5b\x8b\xd3\xd7\x36\x93\x4f\x4e\xc0\xbf\x10\x88\x68\x9e\xc6\x8a\xb7\x09\x26\x31\xc0\x62\x09\x38\x05\x29\x12\x4f\x38\x88\xae\x50\xf4\x0d\x50\xf3\x18\x8f\xde\x20\xa6\xcf\xd3\x31\x89\xd1\x77\x14\x03\x9e\xa1\x08\x6c\x60\x66\xcb\x6c\x08\xcf\xb7\x12\xc4\x4b\xc8\x51\x07\xc2\xe5\xfb\xe0\x4e\x86\x70\x47\x86\x49\x9e\xa6\x96\x8c\xb8\xdb\x72\x03\x33\x4f\x69\xf5\x8c\x15\x2c\x24\x8c\xcf\x5a\x58\x5f\x7c\x65\xe5\x41\xbe\x43\x75\x1d\x4a\xcd\x51\xaf\xb6\xea\x43\xab\x1e\xe5\x74\x6e\x54\x43\x70\x8d\xd8\x2d\x80\xf1\x35\x24\x11\x8a\x81\x64\x80\x42\x4f\x5c\x41\x01\x6e\x69\x6e\xee\x72\x29\x49\x13\x84\x62\x70\x99\x0b\x80\x09\xe0\x74\x83\x24\x20\xd5\xbd\x64\x25\xc8\x39\x52\xa2\xf6\xbb\x9c\x69\x02\xb5\x2e\x21\xae\xca\x5b\xef\x01\x4a\x8e\x99\xab\x1e\xaa\xd9\xb6\xb1\x6e\x7b\x86\x62\x87\x6e\x77\x0c\x5f\x76\xeb\x58\xfe\xfa\x4e\xa7\xbd\x45\xda\x7a\x40\x08\x33\x75\xe0\x58\x89\x56\x0a\x72\xf8\xd4\x61\x2c\x8b\x85\x3b\xde\xe9\x14\x45\xdd\x36\xf7\xc6\x49\xe1\x6e\xeb\x31\x5a\xd5\x43\xa2\xd0\xb8\x0c\x78\x6c\x67\x70\x69\x32\x7d\xbe\x6a\x1f\xdb\x75\xfa\xaa\xf2\x63\x97\xaf\xda\xbe\xa5\xf4\x29\x1d\x58\x8d\x7b\x2c\xae\x1b\xbe\xea\x09\x0f\x1c\x17\xc5\x24\x17\xbe\x36\x18\xab\x6e\x45\x65\x7b\x2c\xdb\xb4\x35\x7c\xfd\x1a\x3b\xbb\x62\xd5\x19\x17\x18\xa4\xae\xfb\x82\xbf\xfc\xbc\xb8\xf8\xf0\x5e\x3d\xfb\x7c\xaf\x72\x99\xcc\xcd\x5b\x50\xa7\xb8\x7d\xc7\xb9\x77\x80\xde\x33\x4b\xa7\x62\xd5\x23\x6f\xbf\x21\x18\x52\xfa\xf9\x81\xa4\xb7\xce\x08\x56\xb9\x66\x51\xa3\xa5\x17\x74\x13\xaa\x2a\x2d\xf0\xaa\xde\x2e\x57\xd0\xd5\x83\x48\xa7\xb5\xf3\x22\x32\xf4\x1f\x30\x63\x28\x6a\x0a\xbc\x2e\xd5\xa4\x38\xad\x3c\xe1\x3a\xb7\xc6\x6b\xc0\x55\xf1\xaa\xe3\x62\x77\xe3\x36\xb7\xd7\x48\xad\x7b\x28\xf2\x53\x15\x6a\xfc\xed\x36\x7e\x40\xeb\x13\xb1\x0a\xa4\x2e\x32\x00\xab\x7a\x2f\x70\x6f\x70\x2a\x10\x2b\x6f\xa0\x95\xb5\x75\xa9\x06\xea\xb4\xf2\x83\x4b\x19\xc2\x6b\xf2\xdf\xc8\x51\xc5\xba\xd4\xc0\xb5\x5b\x79\xc1\x35\xb7\xc7\xad\x1a\x5d\xa2\xe1\x55\xb5\x5e\xb0\xda\xef\x77\xe4\xa7\x2e\xd5\x30\x9d\x56\x5e\x70\xed\x98\x56\x55\x59\x15\xae\xda\x71\x2f\x4f\xa0\xad\xb9\x57\x96\xad\x5a\x81\x18\x2f\x88\x56\xa0\xaa\x06\x59\x16\xae\xda\xc1\x2c\x4f\xa0\x6d\x34\x4d\xd9\xaa\x15\x7b\xf0\x81\xd8\x5c\x30\xad\x75\x72\xd2\xf2\xa8\x5e\xa9\x34\x15\xbd\x2a\xd4\xb8\xd9\x6d\xbc\x80\x9e\x33\xbc\x81\xec\xb6\xa1\xe6\x75\xa9\x06\xeb\xb4\xf2\x82\xfb\x11\xc1\xb8\xb9\x8e\x97\x65\x2b\x13\x02\xae\x5a\x78\x42\x74\x8f\xcc\x35\x44\x5d\xb6\x6a\x06\x95\xfd\xf6\x4c\x14\x31\xe4\x3c\x58\xd7\x25\xe5\x83\x7f\x53\xeb\x09\xab\x39\xad\x2f\xac\x69\x7d\x31\x69\x5a\x5f\xe0\x35\x71\xe9\xd4\x25\x06\x56\x59\xbb\xa3\x5d\xa0\x4b\x0c\xac\xb2\xd6\x0f\x56\x7e\x69\x5e\x47\xd4\xc0\x74\x51\x99\xfb\xac\x6a\xe0\xa7\xd1\xad\x4b\x13\xf2\x53\x15\x6a\x14\xed\x36\x7e\x40\x1b\x28\x5a\xf8\x8d\x22\x67\x46\xe8\x8f\x86\x8c\xfb\x01\xdd\x9e\xed\x0f\x70\x08\x7a\x06\xbe\xdf\x9e\x81\x0e\x25\x84\x0f\x6e\xc1\x83\x5b\xf0\xe0\x16\x3c\xb8\x05\x0f\x6e\xc1\x83\x5b\xf0\xe0\x16\x3c\xb8\x05\x0f\x6e\xc1\x83\x5b\xf0\xd7\x72\x0b\xb6\xed\x84\x08\x77\x78\xf6\xad\x4f\x31\xfd\x53\x46\x76\x27\x63\xf6\x85\xa0\x33\x17\x4e\x1a\xef\xf3\x97\xb1\xa7\x3d\x7b\x4b\xcf\x3c\x05\xaf\x3f\x35\x49\xf3\xb4\x1c\xa7\xbb\x91\xb7\x53\xc2\xe6\x29\x43\x1c\x24\x6d\xf3\x8f\xe0\xcc\x81\x52\x38\xef\xce\xbb\xbb\x25\x72\x1e\x9d\x5d\x3f\x20\x9d\xf3\x34\x01\xfc\xa7\x26\x75\x9e\xc6\xa5\x7b\x91\xb8\x48\xa7\x15\x39\x4f\x21\x76\x33\x61\x4d\xda\x02\x9c\xfc\xce\x87\x9d\xdb\x06\xd7\xfb\xa5\x5d\x61\x85\xd5\xa0\x9e\xed\x92\x35\x79\x1a\x77\x76\xcf\x9d\x3c\x6e\x65\x74\x24\x43\x6e\xa7\x9e\x19\xca\x8a\x1c\x7a\x19\x17\x4e\x72\xe4\x0e\x75\xf7\x49\x74\xd7\x91\xe9\x58\x25\x44\x0e\x3b\xec\x35\xf5\x7b\x3f\xc9\x8e\xbd\x18\x68\xa7\x3c\xf6\xe0\x45\x58\x65\x3e\x1e\x6f\x1b\x2c\xbc\x5e\xb1\x6f\x9d\xfd\x7c\xbc\xc3\xb6\x54\x10\xaf\x1c\xbb\x46\x27\x5a\x6f\xe4\x3d\x9e\x7f\x1f\x22\xe1\xae\x6f\x0a\x5d\x07\xed\x91\xc4\xb0\xfe\xa4\xec\x92\x4f\xb7\x7a\xef\x34\x9e\x1c\xb6\x27\x6d\x83\xc9\xa3\x5b\xd8\xec\x7a\x7a\xa0\xfc\xbc\xbe\x19\x77\xf7\xcd\xdf\x3b\xa6\xdf\xc5\x09\xc0\x71\x2b\x21\xab\x77\x52\xde\xc7\x26\x2b\x6f\xa1\xf9\xb4\x03\x84\x3a\x97\x68\x27\x4b\xff\xec\x94\xbe\x3e\x22\xf8\x91\x89\x7d\xbd\x56\xa9\x1d\x77\x8b\x3d\xa6\xf7\x75\x89\xf4\xcf\xef\xfb\x74\x7a\x82\xdf\xbe\x3c\x30\x2e\x1a\xbb\xa4\x01\x6e\xa9\xe4\xf0\x4f\x7b\x85\xb8\x57\xc9\x80\x3d\x97\x91\x83\xa7\x04\xf6\xd6\xdd\xbf\x48\x62\x60\x1c\xab\x27\x9c\x63\x59\x80\x7b\x53\x9f\x3c\xb6\x53\x9a\xbb\x1a\xdd\x01\xb4\xa9\xd1\xbb\x24\x13\xbe\x83\x46\xd7\xea\x7c\xa7\x7c\xc1\x3e\xba\x78\x27\xe3\xbd\xcb\x87\xcb\x54\x49\x03\x33\xb3\x11\xef\x82\x60\x97\xef\xa5\x4a\xba\xf2\x36\x1a\xdf\x6b\xf4\x4f\x76\x54\xcc\x9f\x74\xcb\x60\x6c\xf5\xdb\x53\x2a\xa0\x5c\xec\xb0\x7e\x76\xe9\xf6\x08\xb4\x9e\xf5\xbf\x37\x08\x3d\xa4\xb1\xc6\xdf\xdc\xfe\xc9\x99\xf9\x7c\x14\xea\xbe\xe6\xe7\xeb\x12\xc7\x78\x82\xbe\xee\x5e\x87\xcf\xd0\x57\x71\xfa\xff\x6f\x9e\x3e\x1f\x65\xba\xaf\xd9\xfa\x7c\x95\xc9\x4d\xd7\xb7\x47\x65\x9a\x94\xaf\xef\x4f\x54\xa6\xed\x5f\x37\x31\xb7\x07\xd7\x86\xd2\x73\xab\x1c\x61\xdb\xa1\xd4\xd3\xfd\x5e\x47\x8f\x65\xda\x11\x88\xec\xdb\x64\xee\x67\x8a\x6f\x18\xc7\x0c\xf1\x2a\xd6\xde\x9d\xf1\xdb\x8b\xef\x87\xcb\xfb\xfd\x78\x3b\x9e\xf8\xdb\x33\x19\x9f\x77\x24\xd0\x7f\x39\xac\x13\xf3\x79\x05\x05\x3b\x9c\x9d\xfe\xf4\x7c\x5e\x3e\xcd\xc1\x92\xf4\x1d\x8c\x59\x75\xc2\x3e\x9f\x5e\x87\x4f\xdb\x37\x8e\x85\x47\xf2\x3e\x9f\xc4\x9b\x8e\xca\xaa\x30\xfd\x84\xbf\x6b\xd3\x8c\xd7\x4f\xf8\x83\x86\x83\x0e\xa8\x09\xe3\x2b\x37\xbf\xe7\xaf\x16\x0e\xef\x31\x7b\xb9\x80\xd1\xc5\x8d\xe9\x67\x18\x87\xe6\x49\xef\xf9\x46\x8b\x25\xd6\x8f\xff\x0b\x00\x00\xff\xff\x94\x4b\x5f\xfc\x6b\x80\x00\x00")

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

	info := bindataFileInfo{name: "templates/model.gotpl", size: 32875, mode: os.FileMode(420), modTime: time.Unix(1605051107, 0)}
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

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 17509, mode: os.FileMode(420), modTime: time.Unix(1595349174, 0)}
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

