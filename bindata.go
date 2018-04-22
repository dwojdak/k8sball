// Code generated by go-bindata. DO NOT EDIT.
// sources:
// site/README.txt
// site/index.html
// site/index.js
// site/license.txt
// site/style.css

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _readmeTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcd\x31\x4e\xc6\x30\x0c\x47\xf1\xbd\xa7\xf8\x9f\x20\x59\x11\x1b\x65\x43\x20\x15\x75\x81\xd1\x72\xdc\xd6\x10\xd9\x51\xe2\xf6\xfc\xa8\x7c\xdb\x1b\x7e\xd2\x7b\xc1\x22\x06\xee\x42\x21\x05\x14\x78\xf5\x22\x8b\x58\x52\x4f\xf8\xf6\x13\x4c\x86\x4d\xad\x20\x0e\x1d\x70\x93\x1b\x1d\x11\x6d\x3c\xe7\xcc\x5e\xa4\xfd\xe3\x3b\x75\x53\x29\xec\xc6\xd2\x62\xe4\x26\x96\xaf\xaf\xfd\xfa\x7c\x4f\xd3\x84\x37\xba\x68\xe5\xae\x2d\xd0\x3a\x71\x28\xcb\xe3\xaa\xb6\x83\xf0\x41\xbb\x32\x9e\x30\x53\xad\x09\x4b\xf7\x1f\xe1\xc0\xe6\x1d\x84\xf5\x57\x6b\xe5\x7e\x8e\x03\x73\x3d\xa5\x75\xb5\x48\x7f\x01\x00\x00\xff\xff\xab\x18\xc8\x13\xb8\x00\x00\x00")

func readmeTxtBytes() ([]byte, error) {
	return bindataRead(
		_readmeTxt,
		"README.txt",
	)
}

func readmeTxt() (*asset, error) {
	bytes, err := readmeTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.txt", size: 184, mode: os.FileMode(420), modTime: time.Unix(1524444438, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5b, 0x74, 0xbe, 0x64, 0xe6, 0x40, 0x48, 0xa3, 0xc8, 0x4e, 0xd3, 0x92, 0x54, 0x58, 0xbe, 0x1a, 0xa5, 0xb, 0x5, 0x5a, 0xa6, 0xa4, 0xab, 0x74, 0x13, 0x5f, 0x28, 0xa1, 0x57, 0xc4, 0x34, 0x3e}}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4d\x73\x9b\x30\x10\x3d\xc3\xaf\xd8\xea\x50\x0e\x4d\x10\x4e\x32\xad\xeb\x0a\x3a\x76\x42\x13\x4f\x3e\xdc\x18\x7b\x32\x39\xca\xa0\x18\xc5\x42\x10\x56\x4e\xe2\x4e\x7f\x7c\x07\x81\xeb\x7c\xf4\x52\x8f\x67\xc4\x5b\xed\xee\x7b\xbb\xbc\x81\x7d\x38\x99\x1c\xcf\x6e\x7f\xc6\x90\x9b\x42\x45\x2e\x6b\x0e\x50\x5c\x2f\x43\x22\x34\x81\xc8\x75\x59\x2e\x78\x16\xb9\x60\x7f\xac\x10\x86\x43\x9a\xf3\x1a\x85\x09\xc9\x7c\xf6\x63\xbf\x4f\xfe\x5e\x1a\x69\x94\x88\x4e\x78\x51\x6a\x0f\xe1\x92\x2f\x65\x0a\xab\x3e\xc2\x88\x2b\xc5\x68\x7b\xeb\x3a\x00\x4c\x49\xbd\x82\xbc\x16\x77\xa1\x97\x1b\x53\xe1\x80\xd2\xbb\x52\x1b\xf4\x97\x65\xb9\x54\x82\x57\x12\xfd\xb4\x2c\x68\x8a\xf8\xfd\x8e\x17\x52\x6d\xc2\x49\x25\xf4\xa7\x84\x6b\x1c\x1c\x05\x81\x34\x5c\xc9\x74\xef\x28\x08\xf6\xbe\x04\xc1\xef\xcb\xa6\x56\xd4\x35\x37\x83\x2e\xe4\x41\x2d\x54\xe8\xa1\xd9\x28\x81\xb9\x10\xc6\x03\xb3\xa9\x44\xe8\x19\xf1\x6c\x9a\xb6\x5e\x2b\x04\xd3\x5a\x56\x06\xb0\x4e\x77\x4a\xf8\x3d\x7f\x7e\x2b\xa4\x89\x51\x25\x17\x48\xef\x1f\xd6\xa2\xde\xd0\x9e\xdf\x3b\xf0\x0f\x3a\xe4\x17\x52\xfb\xf7\xe8\x45\x8c\xb6\x0d\xdf\x35\x27\xff\xd3\x7c\x2d\x9b\xf6\x3d\xff\xb0\xc3\xfb\x6b\xd9\x31\x90\x17\x0c\xdd\xce\xed\x2a\x9b\x61\xc9\x6e\x58\xd2\x0e\x4b\xb6\xc3\x92\x76\xd7\x56\xc4\x80\x52\x55\xa6\x5c\xe5\x25\x9a\xc1\x61\x10\x04\xd4\xd6\xf9\x4d\x5a\xe4\x3a\x8c\xb6\xaf\xdb\x75\xd8\xa2\xcc\x36\x91\xeb\x38\x2c\x93\x8f\x90\x2a\x8e\x18\x92\x8a\x2f\x05\x69\x6e\x1d\xc7\xfa\x42\xd4\x4d\x46\x03\x7a\xd1\xe5\xf0\x74\x7c\xfc\x51\x2f\xb0\xfa\xb6\x62\x58\x71\x1d\xf5\x19\xb5\x27\xb6\xd1\xd1\xf0\xe2\x82\xd1\xbc\x67\x4b\x5a\xa2\xa6\xde\xa2\x17\x1c\x0b\xae\x14\xe9\xda\xca\x62\x09\x32\x0b\x49\xdf\x06\x5f\xaf\x12\x0f\x7d\x5e\xf0\x5f\xa5\xe6\x4f\xed\x1e\x0b\x91\x49\xee\xe3\x4a\x2a\x95\xd6\x6b\xcc\x6d\x70\x07\xe9\x53\xb5\x9f\x96\xda\x08\x6d\xe8\xba\x52\x25\xcf\x90\x1e\x04\xbd\xcf\x34\xf8\x4a\x8b\xc6\xab\x96\xe4\x7a\x2d\xd0\xc8\x52\xfb\x95\x5e\x6e\x55\x54\x56\x03\xd7\xf8\x24\x6a\xb2\x95\xd9\xc1\x68\x76\x36\x4e\x60\x9c\xc0\xcd\x59\x3c\x8d\x61\x76\x16\xc3\xf0\x2a\xb9\x89\xa7\x70\x3a\x89\x13\x46\xab\x6e\xda\x4c\x3e\xfe\x63\xd4\xb5\x31\xa5\xde\xd2\xb4\xc8\x72\x3d\x74\x2a\x46\x5d\xc2\x30\x39\x87\xdb\xc9\x7c\x0a\xe7\xf3\x51\x3c\xbd\x8a\x67\x71\x02\xd7\xf3\x38\x99\x8d\x27\x57\x8c\xb6\x85\xaf\x89\x00\xba\x47\xeb\x92\xce\x8a\xbb\x05\xbe\xb7\x81\xd4\x99\x78\x7e\xe3\x31\x87\xd1\xd6\x05\x8c\xb6\xdf\x07\x80\xe6\xef\xba\x7f\x02\x00\x00\xff\xff\x9a\x9d\xcd\xda\x38\x04\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 1080, mode: os.FileMode(420), modTime: time.Unix(1524430873, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf5, 0x1e, 0x57, 0x12, 0xd2, 0x57, 0x37, 0x51, 0x3d, 0x2c, 0x74, 0x7e, 0xd7, 0x1d, 0x97, 0xea, 0x6f, 0xbb, 0xab, 0x2d, 0x6d, 0xc5, 0x69, 0xf4, 0xc, 0x19, 0xe5, 0x24, 0xe0, 0x50, 0xe2, 0xf3}}
	return a, nil
}

var _indexJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x54\x61\x6f\xe3\x44\x10\xfd\x9e\x5f\x31\x32\x95\xce\x41\xc5\x09\x9c\x84\x0e\xa2\x4a\x5c\xaf\x80\x2a\x71\x57\xa0\x07\x12\x3a\xf1\x61\xb2\x1e\x7b\x57\x5e\xef\x98\x9d\xd9\x06\xb7\xea\x7f\x47\x5e\x27\x6d\x02\xfd\xcc\x27\x6b\x67\xdf\xbe\x79\x6f\xfc\x34\x67\x65\xcd\x26\xf5\x14\x74\x59\x45\xc2\x7a\x2c\x9b\x14\x8c\x3a\x0e\xe5\xf2\x61\xb1\x00\xb8\xc3\x08\x3d\xb6\xce\xbc\xb9\x44\xef\xe1\x02\x1e\x1e\x37\x0b\x38\x2a\x55\xde\x89\x72\x83\x41\x76\x14\x05\x2e\xe0\x53\xd1\xa5\x2d\x19\xf5\x50\x93\x27\x25\x08\xd8\x93\x0c\x68\xa8\x38\x87\xe2\xe3\x8e\x48\xe1\xbb\x8e\xbc\xd0\x68\x5d\x6b\x95\x77\x14\xa1\xe1\x08\x96\xfc\x50\x4d\x98\x77\x96\x4c\x07\xad\x53\x9b\xb6\xe0\x44\x12\x49\xae\x5f\xeb\x2b\x81\x2b\x36\x1d\xc5\x7c\x3e\xf4\x31\x1c\x6b\x0e\x70\xf6\xf0\xe1\xe6\xea\xfb\xc7\xcc\x10\x51\xec\x4f\xcc\xc3\x25\x9a\x8e\x9b\xe6\x1c\x50\x3a\xf8\x2b\x91\x4c\xd6\x00\x5b\x74\x01\x3c\xea\x9e\xe8\x07\x17\x69\x87\xde\x53\x9d\x8f\x1f\x6d\xe4\x1d\x8c\x9c\xa2\x90\x6f\x00\x15\xd4\x12\xf4\x14\xcd\x08\xdc\xc0\xed\xf5\x8f\xc0\x01\x6e\x3d\x9a\x2e\xe3\xff\xd8\xeb\xbb\x75\x6d\x10\x18\xd8\x05\x05\x65\x38\xc8\x6b\x49\x61\xe0\x8c\xf8\x95\x06\x3f\x82\xc5\xfb\xf1\x1c\x34\x8e\xf0\x3c\x29\x31\xd1\x6d\x09\xa6\x81\x4e\xc0\xb7\xd2\x1d\xab\x3c\x87\xeb\x1e\x5b\xfa\x39\x79\x3f\x39\xba\x69\x9a\xec\x92\x83\xa1\xa0\x11\x95\x00\x43\x7d\xd2\x90\xee\x28\xe8\x2c\xeb\x8a\xc3\x2b\x05\xc3\x29\xe8\xa4\xdb\x69\xae\xbe\x1f\x21\x66\x35\x4e\x20\x70\x4d\xf3\x57\x41\x8c\xa5\x3a\x79\xdc\x7a\x3a\xe0\x84\x53\x34\x24\x20\x38\x3e\x41\x15\x5d\xd0\xfd\xbc\x6e\x92\x7a\xe6\x6e\x7e\xce\x60\x09\xbd\xda\xfb\x7c\xf5\x3b\xc5\x11\x6a\x4e\x5b\x6d\x92\xaf\xe0\xb7\xa1\x8d\x58\x13\x18\x9f\x44\x29\x4e\x53\x9a\xfc\x89\xc2\x1d\x45\x71\x1c\xaa\xe2\xcf\xcd\xe2\x34\x5f\x2d\xe9\xdb\x1c\x2e\xb8\x80\xe7\x70\x2e\x00\x1e\x16\x00\x73\x40\x23\x86\x9a\xfb\x0f\xa9\xdf\x66\xd4\x7b\x54\x5b\xcd\xb5\x72\xb9\xf9\x17\xea\x89\x2b\xa3\x1a\xcf\x1c\xcb\x93\xf7\x9f\x83\x5a\x27\xa7\xb1\xae\x3c\x85\x56\xed\x11\x19\x1e\x68\xfe\x0b\xfe\x74\xdc\x68\xb6\x03\x70\x56\x16\x9f\xbd\xd9\xa2\xf7\xc5\xb2\xa2\xa6\x21\xa3\x25\x14\x62\xb1\xa3\x02\xf6\xb4\x13\x64\xa6\x28\x96\x95\xd2\xdf\x5a\x1e\xba\xbc\x04\x68\xb0\xa6\xeb\x50\xbe\x5e\xaf\xd7\x47\xd7\x87\x16\xa8\x1a\xcb\x42\xa2\x99\x7e\x82\x55\x1d\xe4\xdb\xd5\x4a\x5e\x57\xd8\xe3\x3d\x07\xdc\x49\x65\xb8\x5f\xf5\x54\x3b\xac\xa4\x73\xde\x9b\x98\xc4\xe6\xe2\xf3\x71\xb5\x1b\xbe\x30\x1c\x94\x82\xae\xd2\xe0\x19\x6b\x59\x7d\xb5\xfe\xf2\xeb\xd5\xfa\x9b\xd5\xde\xab\xab\xa9\x1a\x42\x5b\x2c\xb3\xcd\xbc\x1a\x4e\x54\x5a\x57\x53\x39\x5f\x4e\x53\xe3\xf0\xce\x3b\xd3\xbd\xfc\x27\x5f\x7c\xf8\xbf\x1b\x9b\x93\x37\x75\xfb\x65\xbf\x2e\x0e\x06\x27\x29\x2f\xe5\x72\x96\xf9\x98\x4d\x4e\x52\x0f\x6b\xe6\x32\xa9\x72\x28\x96\x95\x99\x3c\x97\x4f\xe6\x97\x9b\xc5\xe3\x72\xb3\xf8\x27\x00\x00\xff\xff\x46\xaa\x3f\xa8\x7b\x05\x00\x00")

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 1403, mode: os.FileMode(420), modTime: time.Unix(1524435067, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa2, 0xa, 0x5, 0x37, 0x5f, 0xb5, 0xc3, 0x51, 0xcd, 0xda, 0x62, 0xb4, 0x93, 0x2e, 0x59, 0xb4, 0xbd, 0xe7, 0xa5, 0x88, 0x33, 0x99, 0x9a, 0xc1, 0x50, 0x5d, 0x1b, 0x86, 0x44, 0xfe, 0xb, 0xbc}}
	return a, nil
}

var _licenseTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x52\x4d\x8f\xe2\x36\x18\xbe\xe7\x57\x3c\x9d\xd3\x8c\x94\x19\xda\x9e\xaa\x55\x55\xc9\x43\x0c\x58\x0d\x36\x75\xcc\xb2\x1c\x4d\x62\x88\xab\x60\x47\xb1\x99\x11\xff\xbe\xb2\x61\x3b\x3b\x7b\x4a\xe2\xf7\x7d\x3e\x9d\xa2\xf8\xf3\x97\xe7\xe7\x62\xee\xc7\xeb\x64\x4f\x7d\xc4\x63\xfb\x84\xdf\x7f\xfd\xed\x0f\x1c\xae\x58\xeb\xc9\x38\x3c\xf6\x31\x8e\xe1\xcb\x6c\xd6\xfa\xce\x8c\xc6\xbd\x58\x9f\x5e\xed\xd1\x9a\xae\xf5\xae\x35\x63\x0c\xb3\xd1\xb8\xd9\xdb\xb7\xd3\xdb\x3f\xf5\x53\x51\x14\x1b\x33\x9d\x6d\x08\xd6\x3b\xd8\x80\xde\x4c\xe6\x70\xc5\x69\xd2\x2e\x9a\xae\xc4\x71\x32\x06\xfe\x88\xb6\xd7\xd3\xc9\x94\x88\x1e\xda\x5d\x31\x9a\x29\x78\x07\x7f\x88\xda\x3a\xeb\x4e\xd0\x68\xfd\x78\x4d\x9b\xb1\xb7\x01\xc1\x1f\xe3\xbb\x9e\x0c\xb4\xeb\xa0\x43\xf0\xad\xd5\xd1\x74\xe8\x7c\x7b\x39\x1b\x17\x75\x4c\x7a\x47\x3b\x98\x80\xc7\xd8\x1b\x3c\x34\x77\xc4\xc3\x53\x16\xe9\x8c\x1e\x60\x1d\xd2\xec\xfb\x08\xef\x36\xf6\xfe\x12\x31\x99\x10\x27\xdb\x26\x8e\x12\xd6\xb5\xc3\xa5\x4b\x1e\xbe\x8f\x07\x7b\xb6\x77\x85\x04\xcf\x5d\x85\x44\x7a\x09\xa6\xcc\x3e\x4b\x9c\x53\x29\xe9\x69\x72\xac\xf1\x72\x18\x6c\xe8\x4b\x74\x36\x51\x1f\x2e\xd1\x94\x08\xe9\xb0\x35\x2e\xa1\xb4\xeb\x66\x7e\x42\x30\xc3\x90\x18\xac\x09\xb7\xac\x1f\xee\xf2\x4e\x52\x19\x53\xa1\xf1\x5e\x51\xd6\x7d\xef\xfd\xf9\x73\x12\x1b\x70\xbc\x4c\xce\x86\xde\x64\x4c\xe7\x11\x7c\x56\xfc\xd7\xb4\x31\x9d\xa4\xf5\xa3\x1f\x06\xff\x9e\xa2\xb5\xde\x75\x36\x25\x0a\x5f\x8a\x42\xf5\x06\xfa\xe0\xdf\x4c\xce\x72\xfb\x15\x9c\x8f\xb6\xbd\xd5\x9d\x2f\x60\xfc\xb8\xd5\xfb\x28\xf4\x7a\x18\x70\x30\xf7\xc2\x4c\x97\xea\xd5\x3f\xc4\x99\x92\x7c\x88\xda\x45\xab\x07\x8c\x7e\xca\x7a\x3f\xc7\x7c\x29\x0a\xb5\xa2\x68\xc4\x42\xed\x88\xa4\x60\x0d\x36\x52\x7c\x65\x15\xad\xf0\x40\x1a\xb0\xe6\xa1\xc4\x8e\xa9\x95\xd8\x2a\xec\x88\x94\x84\xab\x3d\xc4\x02\x84\xef\xf1\x37\xe3\x55\x09\xfa\x6d\x23\x69\xd3\x40\x48\xb0\xf5\xa6\x66\xb4\x2a\xc1\xf8\xbc\xde\x56\x8c\x2f\xf1\xba\x55\xe0\x42\xa1\x66\x6b\xa6\x68\x05\x25\x90\x04\xef\x54\x8c\x36\x89\x6c\x4d\xe5\x7c\x45\xb8\x22\xaf\xac\x66\x6a\x5f\x62\xc1\x14\x4f\x9c\x0b\x21\x41\xb0\x21\x52\xb1\xf9\xb6\x26\x12\x9b\xad\xdc\x88\x86\x82\xf0\x0a\x5c\x70\xc6\x17\x92\xf1\x25\x5d\x53\xae\x5e\xc0\x38\xb8\x00\xfd\x4a\xb9\x42\xb3\x22\x75\x9d\xa5\xc8\x56\xad\x84\xcc\xfe\xe6\x62\xb3\x97\x6c\xb9\x52\x58\x89\xba\xa2\xb2\xc1\x2b\x45\xcd\xc8\x6b\x4d\x6f\x52\x7c\x8f\x79\x4d\xd8\xba\x44\x45\xd6\x64\x49\x33\x4a\xa8\x15\x95\x79\xed\xee\x6e\xb7\xa2\xf9\x88\x71\x10\x0e\x32\x57\x4c\xf0\x14\x63\x2e\xb8\x92\x64\xae\x4a\x28\x21\xd5\xff\xd0\x1d\x6b\x68\x09\x22\x59\x93\x0a\x59\x48\xb1\x2e\x91\xea\x14\x8b\xdc\x19\x4f\x38\x4e\x6f\x2c\xa9\x6a\x7c\xba\x11\x21\xf3\xf7\xb6\xa1\x1f\x5e\x2a\x4a\x6a\xc6\x97\x4d\x02\xff\xb8\xfc\x52\x3c\x3f\xff\x55\xfc\x17\x00\x00\xff\xff\x14\xf9\xdc\xd8\x59\x04\x00\x00")

func licenseTxtBytes() ([]byte, error) {
	return bindataRead(
		_licenseTxt,
		"license.txt",
	)
}

func licenseTxt() (*asset, error) {
	bytes, err := licenseTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "license.txt", size: 1113, mode: os.FileMode(420), modTime: time.Unix(1524444438, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf4, 0x39, 0xcf, 0xed, 0x4c, 0x31, 0xbe, 0x6f, 0x91, 0xb, 0xc9, 0xcc, 0x9e, 0x8f, 0x70, 0xd4, 0x25, 0x3a, 0xfd, 0xcc, 0xf8, 0xcb, 0x7d, 0xf, 0xc3, 0x4b, 0x54, 0x89, 0x3b, 0xba, 0xde, 0x8d}}
	return a, nil
}

var _styleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\xcd\x92\x9b\x30\x0c\xbe\xf3\x14\x9a\xdd\xc9\x0d\x18\x20\x21\xcd\x9a\x5b\xef\x7d\x08\x25\x08\xf0\x14\x6c\xd7\x16\xdd\xb4\x9d\xbc\x7b\xc7\x36\x64\x93\xcd\xf6\xd0\x1d\x4e\x96\x84\xf4\xfd\xe5\x06\x7b\x82\x3f\x09\xc0\x84\xe7\xec\x55\xb6\x3c\x08\x78\xd9\x17\xe6\xdc\xf8\x9a\x54\x6b\x6d\x5b\xac\x35\xb4\xbd\x54\x02\x0a\xc0\x99\x75\x93\x5c\x92\x64\xe0\x69\x0c\x3b\x3a\xad\x38\xeb\x70\x92\xe3\x2f\x01\x4f\xdf\xb4\x62\x47\xd6\x22\x3f\xa5\xe0\x50\xb9\xcc\x91\x95\x5d\xb3\x0e\x3a\xf9\x9b\x04\x94\xfb\xb8\xf7\xa4\x47\x6d\x05\x3c\xef\x6a\xff\xdd\x5d\x8a\x03\x06\xdb\x56\xaa\xfe\xfa\x3e\xe2\xe9\x7b\x6f\xf5\xac\xda\x6c\xfd\xb7\x3b\xf8\xcf\x37\x99\xce\x9c\xe1\x28\x7b\x25\xe0\x44\x8a\xc9\x46\xa4\x84\x2d\xd9\x80\xf5\xe3\x89\x65\xc0\x19\x54\x6f\x8c\x22\xd0\x55\x94\x9b\x52\x45\xd3\x2d\xf4\xae\xdc\x17\xf5\x4b\x58\x53\xa6\x30\x54\x29\x0c\xdb\x14\x86\x5d\x0a\x43\x9d\xc2\xb0\xff\x1f\x8d\x2e\x49\x92\x1f\x71\xf4\xba\x1a\xed\x24\x4b\xad\x04\x58\x1a\x91\xe5\x4f\xf2\xed\xd8\x95\x53\xef\x97\x2e\x26\x95\x45\xb1\x69\xee\xbd\xdc\x15\x1e\xe3\x25\xc9\x51\xb9\xd7\x85\xfa\x8a\xb7\xaa\xaa\xab\x56\x6c\x51\xb9\x4e\xdb\x49\xc0\x6c\x0c\xd9\x13\x3a\xba\x33\x21\x58\x70\x45\x82\x47\xa7\xc7\x99\xc3\x08\x6b\x23\x60\x7b\x08\x97\x47\xea\x58\xc0\x6e\x1f\x1e\x2b\xaa\x72\xd3\xbc\x97\x72\x13\x18\x3e\xff\x98\xc9\xf9\x85\x5f\x67\x66\xad\x3e\x9d\xa1\xed\xa3\x35\x65\x7e\xf8\x52\xbf\xb3\xe7\x9a\x8e\x8f\xa2\xb3\x78\x07\x70\xd4\xb6\x25\x9b\x59\x6c\xe5\xec\x04\xd4\x4b\xda\x42\x55\x80\xd2\x8a\xee\xd2\x98\xd7\x34\x41\x19\x2f\x2d\x84\x0f\x0f\x2e\x94\x0b\x94\xa8\x66\x16\x14\x2b\xa3\x31\x49\xa7\x35\x93\x05\xf3\x96\xca\x7f\x79\x71\x4b\xaf\x7a\x60\x9c\x47\xbe\x97\xe4\x6f\x00\x00\x00\xff\xff\xe4\x2f\x75\x5f\xd4\x03\x00\x00")

func styleCssBytes() ([]byte, error) {
	return bindataRead(
		_styleCss,
		"style.css",
	)
}

func styleCss() (*asset, error) {
	bytes, err := styleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "style.css", size: 980, mode: os.FileMode(420), modTime: time.Unix(1524422865, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x13, 0x89, 0x51, 0x96, 0xa9, 0x42, 0xca, 0x17, 0xc1, 0xdb, 0xae, 0x7, 0x24, 0x80, 0xf1, 0x2a, 0xd0, 0xad, 0xed, 0xb9, 0x91, 0xf2, 0x5a, 0x58, 0x14, 0xa4, 0x83, 0xdd, 0x1, 0x64, 0x1c, 0xe8}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"README.txt": readmeTxt,

	"index.html": indexHtml,

	"index.js": indexJs,

	"license.txt": licenseTxt,

	"style.css": styleCss,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"README.txt":  &bintree{readmeTxt, map[string]*bintree{}},
	"index.html":  &bintree{indexHtml, map[string]*bintree{}},
	"index.js":    &bintree{indexJs, map[string]*bintree{}},
	"license.txt": &bintree{licenseTxt, map[string]*bintree{}},
	"style.css":   &bintree{styleCss, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
