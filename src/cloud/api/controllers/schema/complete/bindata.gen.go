// Code generated for package complete by go-bindata DO NOT EDIT. (@generated)
// sources:
// 01_base_schema.graphql
// 02_unauth_schema.graphql
// 03_auth_schema.graphql
package complete

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __01_base_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\x14\x54\x5c\x2a\x10\x4a\x49\x4f\x81\xe0\x07\x1c\x7b\x38\x47\x72\xbc\x3e\xef\x46\x47\x84\xf8\x77\x94\xcb\x5d\x77\xd5\x6c\x31\xf3\xb4\x4f\x43\xe2\xec\xf1\xeb\x80\xd3\xc2\xb6\x0e\xf8\xd8\xc2\x01\xf3\x62\xde\x26\x29\x03\xde\xaf\x97\xfb\x73\xae\xc3\x57\x22\xb4\x32\x20\x0a\xb5\x3c\x1a\x7c\xce\x72\x06\xe7\x6a\x2b\x6c\xad\xd4\x83\xeb\xf0\x29\x38\x13\xa1\xd1\x1b\x51\x7d\x0e\x4c\x92\x23\x9b\x22\xb1\x11\xbe\xc4\xeb\xce\x12\x95\xfb\x0e\x26\x18\xe9\x3a\xf0\xc7\x58\x22\x23\xc6\x15\x62\x89\x0d\xdf\x53\xde\xb9\xc9\xac\xea\xd0\xf7\xc7\xc9\xd2\x32\x1e\x82\xcc\xfd\xb1\xf9\x9a\x4e\xf9\x96\x4f\xdb\x73\xfd\xa4\xba\x50\xfb\xe7\x97\x57\xe7\x36\xf8\xae\x75\xf1\x2c\x22\x75\xc0\x9b\x48\xa6\x2f\x0f\x9b\xd4\xa5\x70\xb3\xbc\xdf\xf9\x0f\x00\x00\xff\xff\x6f\xc4\xb8\xef\x28\x01\x00\x00")

func _01_base_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__01_base_schemaGraphql,
		"01_base_schema.graphql",
	)
}

func _01_base_schemaGraphql() (*asset, error) {
	bytes, err := _01_base_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_base_schema.graphql", size: 296, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_unauth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\xb1\x0a\xc2\x40\x0c\x86\xf7\x7b\x8a\xdc\x56\x5f\xe1\x36\x1d\x84\x0e\x0a\xa2\x9b\x38\x84\x9a\xd6\x60\x2f\x29\x77\xb1\x58\xc4\x77\x17\x0b\xa5\x15\xb7\x9f\xe4\xfb\xbf\x9f\x9e\x46\x72\x05\x1b\x3a\x82\xc3\x83\xd2\x00\x2f\x07\x80\xc9\xb8\xc6\xca\x72\x31\xa5\x3d\x46\x0a\x70\xb4\xc4\xd2\xf8\x55\x80\xf5\x44\x94\x52\xab\x77\x00\x3d\x25\xae\x87\x52\x7a\x36\x3a\xe9\x9d\xa4\xe0\x39\x2f\x9b\x1b\xd5\x96\x50\xbc\x7b\x3b\x37\xce\xfe\xa8\xc6\x79\x36\x8a\x39\xc0\x79\xfa\xf8\xcb\x3f\x3d\x82\x3d\xa5\xcc\x3a\xeb\x1d\x40\x75\x43\x69\xa8\xd5\x66\x79\x34\x8e\x94\x0d\x63\xb7\xcb\x01\xb6\xad\xa2\x7d\x85\x9f\x00\x00\x00\xff\xff\xef\x77\x02\x34\xfc\x00\x00\x00")

func _02_unauth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__02_unauth_schemaGraphql,
		"02_unauth_schema.graphql",
	)
}

func _02_unauth_schemaGraphql() (*asset, error) {
	bytes, err := _02_unauth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_unauth_schema.graphql", size: 252, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_auth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x5a\xdd\x8f\xe3\xb6\x11\x7f\xf7\x5f\x31\x97\x7b\xc8\x2e\xb0\x3d\x04\x45\x13\x14\xfb\x54\xc5\xd6\xe5\xd4\xdd\xf5\xba\xb6\x37\x69\x10\x1c\x0e\xb4\x34\xb6\x88\x95\x48\x85\xa4\xbc\xeb\x16\xf9\xdf\x8b\x21\x29\x89\x94\xb5\xf7\xd5\xa2\x6f\x16\x3f\x66\x7e\x33\x1c\xce\x17\x8d\xcf\x06\x45\x01\xe6\xd4\x20\xfc\xa3\x45\x75\x82\x7f\xcf\x00\x5a\x8d\xea\x1a\x1e\x34\xaa\x4c\xec\xe5\xab\x19\x80\x54\x87\x6b\xb8\x57\x87\xee\x9b\x56\x6c\xd0\x18\x2e\x0e\xda\xad\xec\xbe\xba\xd9\xc4\x18\xc5\x77\xad\x41\x3f\x3f\x7c\x7b\x7a\x34\xa8\xaf\xe1\xb7\x9e\xcd\x7b\x9a\xc8\xab\x56\x1b\x54\x17\xbc\xb8\x86\x6c\xf1\xea\xf2\x1a\xe6\x6e\xa4\xe3\xec\x17\xfc\x78\x5a\xb2\x1a\x2f\x04\xab\xf1\x1a\x36\x46\x71\x71\x78\x79\x31\xb1\x09\x67\x3c\x27\x9e\x28\xc3\xf7\x2c\x37\x17\xcc\xff\xd8\x9e\x1a\xbc\x86\x24\xf8\xb2\x44\x6f\xb3\x6e\x88\x36\xb2\xd6\xc8\x5c\xd6\x4d\x85\x06\x2f\xb8\x68\x5a\xd3\x21\xb8\x82\xbc\x55\x5a\xaa\x95\xd4\xd7\x90\x09\x73\x05\x2c\x37\x5c\x8a\x6b\x48\x82\x3d\x89\x1d\x23\xe2\x57\x1d\xc0\x87\x6c\xd1\xd1\xb8\x8c\x17\xaf\x51\xb7\xd5\x19\xdb\xb7\x1c\xab\x62\xcc\x7b\x4f\x83\x5e\x82\x60\x6d\x2a\x0c\x37\xa7\x1b\x2e\x8a\xab\x19\x00\x80\xc2\xdf\x5b\xae\xb0\x48\xd4\x81\x16\x93\x6e\xa6\x97\xbf\xff\x0c\x78\x16\x48\x87\x71\x06\xf0\x1a\x36\xb9\xe2\x8d\xa9\x0f\x0a\x50\x14\x8d\xe4\xc2\xe8\x2b\x50\xb8\x47\x05\x46\x42\x21\x73\x0d\x5c\x40\x5e\xc9\xb6\x60\x0d\x7f\xd3\x28\x69\xe4\x0c\xa0\xe2\x47\xfc\x99\xe3\x13\xc1\xb9\xf5\xbf\xef\xd0\xb0\x82\x19\xe6\xce\xab\x5b\x31\x97\xc2\xa0\x30\x3a\x30\x91\xdb\xd1\x14\x2d\xd7\x16\x07\x91\x73\x88\x62\x62\x6e\x76\x82\xd4\x26\x9a\xf0\x32\x2d\xb0\xa9\xe4\x09\x1e\xf1\xa4\x67\x00\x85\xfd\xaa\x51\x98\x1b\x3c\x11\x83\x45\x38\x10\xf3\x89\xd6\x06\x6c\xa2\x2d\x9e\x4b\xb2\xca\x3a\x16\xac\xe1\x9e\x76\xb2\xca\xce\x88\xba\xd9\x80\x9a\x5b\xe4\xc9\xac\xaa\xf6\xc0\xc5\x0c\xa0\xb1\x3f\xf4\xc5\x23\x17\xc5\xb5\x1f\xa6\x73\xbd\xbc\x86\xdf\xdc\x97\x23\xa7\x90\x64\xe5\x52\xb8\x41\xba\x21\x96\xb6\xbf\x55\x57\x9e\xd0\xcf\xa8\xb4\xb5\xe5\xe1\xb6\x0d\x1b\x5e\x59\xd6\xdb\x92\x6b\x78\xe2\x55\x05\x3b\x04\x85\x4d\xc5\x72\x2c\x60\x77\x1a\xb3\x98\x4b\xb1\xe7\x07\x90\x22\x47\x60\x55\x05\xb9\x14\xba\xad\x51\x69\x28\xd9\x11\x61\x87\x28\xa0\x6d\x0a\x66\xb0\x78\xe3\x9c\xc5\x7a\x8a\x40\x88\x72\x10\xca\xcd\x4d\x8a\x36\xb9\x6d\x92\x74\xb4\x79\xd3\x9b\xd2\x3a\x1e\x1a\x31\x71\x83\x23\xf2\x0b\x34\x8c\x57\x58\x8c\xb7\xce\xfe\x98\xcd\x42\xef\x7b\xd7\x1a\x46\xd3\xd6\x01\xcf\x15\x32\x83\xde\x65\x45\x5e\x0d\xfe\x56\x60\xa3\x30\x27\xdd\x5c\x28\x64\x9a\x4e\xe4\x1b\xbf\x40\x03\x53\x08\x42\x3e\x41\x6e\x09\x14\x70\xe4\x0c\x9a\x67\x6f\x86\xdf\x5c\xf6\xa4\x23\xfb\x3b\x33\x47\x80\x05\xd2\xed\x5e\xbc\x60\xbd\x3f\x4a\x59\x21\x13\xaf\x7a\x72\xce\x00\x07\x43\xec\x08\xb8\xef\xe9\x9d\x0f\xf6\x80\xc3\xd0\x71\xa1\xfb\x88\x92\x16\xdc\xb0\x5d\x15\x4d\xd3\xfe\x71\xa4\xd9\xa0\x89\x83\xcb\x05\x0b\xe2\x4e\x48\x25\x88\x3f\x97\x53\x11\x29\x13\x47\xee\xe0\x5c\x60\xcd\x78\x15\xd8\xff\x9e\x2b\x6d\x96\x61\xa4\xb9\x82\x8a\x8d\x86\x2e\xbb\x80\x49\x64\x62\xf9\x56\xa8\x6a\xae\xe9\xf2\xe8\x0b\x0a\x8d\xe4\x4d\xb3\xc5\xab\x2b\x1b\x27\x83\xc9\x18\x70\x30\x31\x10\x77\x37\xcd\x29\xfd\x5e\x1d\x2e\xa4\x3a\x8c\x51\x64\x8b\x81\xfb\xbd\x3a\xf4\xca\x95\xea\xd0\x33\x96\xc3\xf8\xc0\x34\x58\x4c\x74\x82\x68\xef\xf8\x39\xd1\xb6\xf2\x11\x45\x40\xec\xb2\xe7\x3d\x03\x58\xe3\x51\x3e\x62\x52\x55\xc1\x5a\x1d\x2f\x0e\x2c\x60\x8d\xb5\x3c\x5a\x59\xdf\x2a\x59\x93\x38\x81\x76\xc2\xa5\xb1\x5f\x73\xa2\x7d\xd2\x23\x5c\x01\x0a\x12\xab\xe8\x09\xf5\x23\x23\x57\x76\x45\x0e\x68\xcf\x43\x5d\x84\x44\xf5\xa4\xe1\xae\x27\x6e\xbd\xd5\xad\x0b\x2d\x03\xa9\xd1\xc2\xa9\xeb\x33\xa6\xf5\x69\x12\xee\x90\xdd\x25\x7b\x01\x49\xc8\xe8\x8f\xd9\xcc\x3a\x99\xce\x88\xac\x93\xf1\xeb\x66\x00\x51\x16\x35\x03\x88\x2f\x00\x85\x12\x9e\x9b\x56\x45\x6b\xc6\x96\xe7\x86\x86\x44\x81\x06\xb8\x4e\x9a\x46\xc9\x63\x70\x06\x03\x96\x6c\x91\xae\x98\x29\x2d\x94\x6c\x91\x8e\x89\x35\xcc\x94\xc3\x77\xb7\xc9\x1b\xe5\x27\xf0\x17\xb2\x66\x5c\x8c\x29\xba\xc3\x77\x88\x58\xa5\xa3\x73\xe0\x05\x12\x18\xf2\xf1\x1e\x17\xf9\xf6\x50\x6d\xdd\xd5\xb0\xac\x99\x60\xd5\xc9\xf0\x5c\xdf\x37\x46\x52\x0e\x16\x91\x72\xb0\xc2\xcd\x83\xab\xb1\xdb\x8d\x6c\xd5\x06\x51\xbc\xb4\xcf\x26\x76\x2f\x78\xaf\x69\x02\xd3\xbb\x3e\x0b\x73\x0f\x34\x4e\x35\x46\x2a\xf6\x01\x25\x31\x77\xfa\x1a\xde\x56\x92\x19\x97\xde\xe8\xfc\xfc\x90\x1c\xa1\x11\x81\x47\x8a\x0f\xc3\x61\x7c\x09\xbd\xc9\xfc\xea\xbf\xc0\x17\xd1\xfb\x9f\xc0\x44\xd1\xd6\x13\x49\xf7\xc6\x30\x83\x96\x41\x92\x6e\x3e\x3c\x2c\x6f\x96\xf7\xbf\x2c\xfd\xd7\x2a\x5d\x2e\xb2\xe5\x4f\xfe\x6b\xfd\xb0\x5c\x0e\x5f\x6f\x93\xec\x36\x5d\xf8\x8f\x6d\xba\xbe\xcb\x96\xc9\x36\x5d\x4c\x72\x1a\xaa\x09\xc7\x28\xd9\x06\x8c\x5e\x43\x22\x00\x0b\x6e\x7c\x21\x02\x32\xa7\x0a\x05\xf8\x1e\x98\x8d\x3e\x50\x32\x0d\xb5\x2c\xf8\x9e\x63\x01\xa6\x44\x70\x56\x64\xf0\xd9\x50\xe2\xc6\x85\x46\x45\x36\x04\x52\x41\x41\xee\x86\x7e\xe7\x25\x53\x2c\xa7\x84\xe3\xcd\x90\xf7\x71\x91\x57\x6d\x81\x9a\xd2\x19\xbb\x41\x58\x7a\x8f\x78\xda\x49\xa6\x0a\x60\xa2\x80\x86\x69\x47\x40\xd6\x35\x13\x85\xdd\x4e\x88\xd3\x45\xb6\x75\x70\x41\x63\x85\xf9\x80\x57\x54\xa7\x69\xd0\x79\x29\x35\x0a\x60\x22\x2a\x8c\x40\xb7\x87\x03\x6a\xda\xfb\xa6\x83\x55\x70\xca\x96\x34\xd8\x3a\xe3\xb5\x05\x15\x6d\x31\x25\x33\xc0\x0d\xe8\x52\xb6\x55\x01\x14\x93\xec\x22\x62\xf5\xad\xf6\x25\x1d\x15\x2f\x34\x28\x48\x31\x8c\x7c\x48\xa3\x38\x9d\xae\x61\xbb\x4e\x8a\x4d\x7a\x9b\xce\xb7\x1f\xb1\x07\xca\xbe\xbd\x39\xdc\x44\xe6\x70\xf3\x61\x75\xbf\xf0\xbf\x36\x3f\xcf\xbb\x5f\xf3\x75\xb6\xda\xfa\x8f\x65\x72\x97\x6e\x56\xc9\x3c\xed\xbe\xef\x17\xe9\x70\xe3\x02\x56\x9b\x5e\x03\x96\x95\xcb\xfe\xa7\xb1\x8c\x5c\xa7\xb7\x6c\x0a\x22\x41\x74\x9c\x01\xd4\xcc\xe4\x25\x16\x99\x28\xf0\xd9\x16\x8c\x99\x30\xef\xa9\x8a\x22\xfb\x9e\x22\x6e\x0d\xbf\x47\xb7\x65\xbb\x11\x28\x32\x19\x32\xb5\x02\x9f\x41\xee\xad\x62\x0d\xdb\xb9\x93\x30\x25\xea\xf0\x1c\x5d\x52\xbb\x97\x8a\xd4\x6c\xd8\xce\xa2\xb0\xe5\xb5\x25\xf4\x4b\x89\xa6\x44\xe5\xed\x86\x8c\x8b\x05\x9b\x69\x1f\x18\xb2\x03\xa2\xef\x18\xda\x02\xa5\x66\x8f\xee\x94\xbd\x29\x02\x3e\x63\xde\x5a\xcf\x49\x7c\x86\xaf\x64\x6f\xc8\x91\x12\xf1\xc1\x65\x42\x88\x6f\x54\x40\x0f\xa2\xbe\x9f\x3c\x1f\x57\x2d\x07\x6a\xd8\x4b\x55\x33\x43\xd9\xba\xbb\x7b\x04\xb6\xbf\x88\xda\x67\x28\x4f\x25\xcf\x4b\x6b\xf8\xb6\x3a\x6a\x98\xd2\xae\xb4\x3a\x37\x67\xd9\xdb\xbc\xb3\x77\xb6\xdb\x18\xd9\x40\x23\x35\xb7\x78\x49\xbe\x9e\x67\x16\xf6\x10\x22\x85\x8e\x31\x10\x2e\x06\x47\x56\xf1\xe2\x2a\xd0\x4f\xa7\xc0\x37\x36\xde\xa7\xfd\x78\xa8\xac\xd7\x90\x54\x55\x74\xa4\x74\x2c\xc8\xf2\x32\x38\x7d\x02\xa9\xfd\x19\x6f\x22\xed\x46\xf6\x33\xad\xd4\xa0\x0f\x11\x68\xf6\x05\xcf\xa0\xbd\x55\x74\xf2\x51\x42\xc0\x0b\x2c\x3e\xf7\x58\x5f\x45\x7a\x92\x0a\x84\xb4\x66\x4b\x95\x60\xab\x04\x16\xa0\x2c\x12\x67\xb9\x0d\x53\x86\xb3\x0a\x2e\x8c\x6a\xf1\x92\x96\xf7\x90\x2e\xf6\xac\xd2\x48\x55\x59\xc9\x74\x52\x14\xf6\x7c\x58\x75\x67\xaf\x9b\x9e\xc8\x99\xe6\x52\x18\xc6\x05\x2a\xba\x60\xad\x8b\xeb\xe3\xe4\x67\x3a\x64\xf9\xab\x3a\x2c\xab\x51\x6b\x76\x88\x86\xba\x72\x32\x1c\xd1\x86\x29\x33\x97\xad\x30\xf6\xca\x0d\x50\x6e\xfe\xaa\xd3\x23\x0a\xa7\xee\x09\x62\xb6\x68\xda\xf2\x1a\x23\x18\x54\x36\x8d\x06\x3b\x82\x2b\x59\x7c\x95\x54\xad\xfe\x62\xb1\xf2\x4e\x8d\xb6\x31\x18\xeb\xd4\xd5\xf4\x48\xa2\xd1\x6c\x27\x66\x57\xea\x4f\xe9\xc3\x7a\x7b\x5f\x82\x07\x22\x38\x1b\x2c\x70\xcf\xc8\x2a\xed\x01\x50\x0c\x13\xd2\x94\xfe\x3a\x3d\x0a\xf9\x24\xc8\xe4\xe7\x9b\x28\x68\xd3\x3e\xbf\x5e\x43\x89\xac\x32\xe5\x89\xb6\x96\xc8\x94\xd9\x21\xf3\x96\xa5\x30\x47\x7e\xc4\x82\x42\xad\xc2\x43\x5b\x31\x05\x5c\x18\x54\x94\xde\xda\x78\x6b\x4a\xe7\x03\x7c\x1f\x8f\xc8\x29\xd4\x8d\x14\x05\x21\x30\xd2\x36\x03\x51\x1b\xed\x41\xbc\x4b\x93\xdb\xed\xbb\x5f\xcf\x41\xb4\x22\x80\x61\xdd\xe6\x40\x31\x97\x42\x60\x4e\xfe\xcb\x48\x58\xf1\x67\x8e\x30\xaf\x64\xeb\x22\x3e\xd7\xfe\x7a\x75\xee\x65\x90\xe1\x0a\x76\xd6\xdb\x89\x6f\x0d\xfc\xde\xa2\x3a\x59\x77\x42\x57\x53\xcb\x1a\xfd\xb1\xf9\x28\xae\x50\x63\xbd\xab\x50\xc3\xbb\xed\x76\xf5\xad\x86\xef\xbf\xfb\xce\x9f\x7e\xaf\xbf\x69\xf0\xd6\xdb\x1f\xa4\x6d\x3e\x72\x3d\x60\xf5\x72\xfc\xb4\x5e\xcd\x3b\x09\x28\x5e\xec\x14\xb2\x47\xfd\xc6\x12\x28\x65\x83\xce\x1b\x33\xd3\xa7\x0e\x9d\xe0\x96\x6e\x4e\x40\x77\x2c\x7f\xa4\x44\x85\x0b\xb4\x22\xd3\xe5\xaf\xc9\xb7\x80\x47\xe4\x90\x78\x9c\x8b\x6c\x33\xbf\x5f\x2e\xd3\xf9\xd6\x66\x78\x63\x3d\x53\x6d\x49\x67\xf3\x54\xa2\x18\x2b\x9a\xbb\x91\x46\xc9\x1c\xb5\x26\xd7\xd9\x2d\xef\x74\xb0\x5a\x24\x5b\x97\x46\x3a\xba\x47\xfe\x2f\xde\xe5\x4b\x9d\xe4\x4e\xed\x34\x44\x6e\x4b\xd3\x15\x66\xe2\x04\xd2\x3a\xb3\x7d\xab\x5c\x34\x75\x66\xec\x9a\x70\x1a\xd8\x4e\xb6\x4e\x05\x4f\xde\xeb\x71\x13\xda\xa6\x54\x63\x28\xe7\x32\x7a\x2c\x4f\x4c\x83\x51\x27\x6f\x7f\x8e\x81\x83\xb4\xb7\x7d\xb2\xce\x6a\x84\x7c\x22\x81\x19\xec\x58\x11\x29\xd0\x0a\x99\x0e\x39\xf2\x48\x83\x05\x1e\x14\x2b\x86\x03\x0e\xf4\x57\xf1\x47\xac\x4e\xc4\x76\x87\x81\xc5\x11\xef\x9a\x1f\x4a\x43\xc3\xb6\xe5\xe2\x4d\x95\xca\x8c\xee\xd4\xd2\x9f\xd6\xc9\xc2\xa5\xe0\xce\x13\x07\x2d\xb9\xb8\x82\xe8\x7c\x52\xe4\x10\x3a\xdf\xf7\xae\xb3\xfe\xc8\x8d\x39\xd5\x8c\x1b\xac\xc3\xfb\xc5\xcb\x33\x67\x55\xb3\x42\x63\x4e\xf3\xe9\xc9\xf3\x6e\x7e\xe7\x0c\x95\xac\x56\x15\x13\xd8\xfb\x60\x9b\xe5\xf5\x5f\xce\xf9\xf5\x3e\x60\xc1\x0c\xfb\xf4\x72\xd1\xd6\x4b\x59\xa0\xf6\x7e\xd2\x0e\x64\x42\x1b\xd5\x52\xe5\x85\x45\x3c\xe9\xb4\x76\x77\xee\xbd\x1b\x85\x47\x2e\x5b\xbd\x99\x52\xeb\xd9\x7c\x14\x5b\x46\x6d\x8f\x23\xf7\xc5\xd8\x59\x7b\x83\xdb\xb9\x5b\x2e\x1e\xa3\xaa\xee\x35\xac\x3f\xf1\x6e\x61\xa9\x8f\x9f\x2b\x3e\xd5\x9c\x18\x17\x8f\x5f\xc8\xa6\x7b\x9b\xf0\x81\xd7\xf1\xbc\x3e\x43\x61\x75\xf7\x5c\x75\xab\x43\x04\x47\xae\xff\xbe\xb9\x5f\x7e\x0d\x88\xf8\x2d\xe5\x8b\x24\xb5\x49\x4e\x87\x32\xce\x6d\xbe\x88\xf9\x0b\xf2\x8f\x5e\x79\xbc\x61\xc7\xa2\xf7\x95\x59\xf0\xc0\x67\xc9\x00\x44\x65\xb3\xfd\xbc\xcd\x96\x0f\xff\xfc\x90\xdc\x2d\x7e\xf8\x4b\x37\xb4\x48\xd6\xbf\x64\xcb\x78\x6c\x7e\xbf\xdc\x26\xd9\x32\x5d\x7f\xd8\xa4\xdb\x0f\xbf\x26\x77\xb7\x9b\xe9\xa9\x09\x7a\xf1\x82\x6d\x7a\xb7\xba\x25\xd7\xe6\x88\xf4\x9e\x66\x78\x7d\x74\x8f\xb3\x2a\xb2\x5d\x5d\xb2\x3f\x7f\xff\x43\x24\xe3\x79\x23\x28\xe8\x25\xbb\x33\x3b\x6f\xcc\x9d\x6f\x0c\xfa\xc1\xee\xda\xbc\xd0\x3f\x73\x27\xe8\x3a\xa6\x7f\x52\x58\xd9\x87\x07\x82\xae\xdf\x74\xa9\x9d\x9d\x9b\xcc\xeb\x82\x86\xed\x74\xf9\x69\x3d\xa7\x3c\xc8\xa0\x46\x21\x0e\xda\x4c\x38\x45\xdd\x36\x8d\x54\x46\xf7\x1d\xd1\xa8\xc1\xd6\xbf\xd3\xa4\xa3\xbe\xf0\xd0\x1c\x1c\x77\x86\x7b\x8b\x19\x5e\xcf\xac\x14\xab\xb0\x7c\x5f\xdd\x7c\x58\xa7\xdb\x74\xb9\xcd\xee\x97\x43\x36\x1b\x3e\x74\x4d\x09\x7e\x64\x55\x8b\xe7\xbd\xa9\xe1\x49\xcd\xee\xea\x1b\xd3\xd1\xe3\xd6\x26\x2f\xb1\xee\x1e\x03\xab\x4a\x3e\xcd\x5b\x6d\x64\x9d\x3e\x93\xf4\x0f\xeb\xdb\x48\x32\xbb\x20\x13\x1a\xf3\x56\xe1\xf6\x76\x13\x4d\xfa\xec\x34\xd8\x39\x0d\x28\xe4\x3b\x29\xce\xe4\xc9\x9d\x99\xd4\x57\x68\xe5\x65\x02\x7a\xa4\xa1\xa9\x35\xfe\xa9\x7f\xac\x9e\xde\x94\xf8\x84\x5e\x7a\xd1\x27\x9f\x18\x3e\x72\x2c\x5f\xcb\xec\x35\x50\x44\x1d\xec\x33\xbe\x3c\xa3\xf6\xfe\x67\xf8\xdc\x89\x1b\xb4\xb7\x89\xbe\xc8\x4f\x9b\x3e\xe2\x8e\x5f\x47\x46\xff\x97\xc8\x16\x4e\x1c\xf7\xf8\x3b\xee\xeb\xaf\x28\x21\x37\x13\x15\xea\x0b\xcf\x9d\xff\x6f\xd4\xe7\xbe\x7f\x5a\x92\x8f\x1c\xd6\xa4\x8c\xb1\x35\x4e\x09\xf9\x59\x0d\xb5\x91\x60\x13\x72\x9d\x8b\x35\x21\xd5\x84\x50\x1f\x91\xe9\x8f\xd9\x7f\x02\x00\x00\xff\xff\x18\xbb\x7f\x87\xee\x23\x00\x00")

func _03_auth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__03_auth_schemaGraphql,
		"03_auth_schema.graphql",
	)
}

func _03_auth_schemaGraphql() (*asset, error) {
	bytes, err := _03_auth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_auth_schema.graphql", size: 9198, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"01_base_schema.graphql":   _01_base_schemaGraphql,
	"02_unauth_schema.graphql": _02_unauth_schemaGraphql,
	"03_auth_schema.graphql":   _03_auth_schemaGraphql,
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
	"01_base_schema.graphql":   &bintree{_01_base_schemaGraphql, map[string]*bintree{}},
	"02_unauth_schema.graphql": &bintree{_02_unauth_schemaGraphql, map[string]*bintree{}},
	"03_auth_schema.graphql":   &bintree{_03_auth_schemaGraphql, map[string]*bintree{}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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
