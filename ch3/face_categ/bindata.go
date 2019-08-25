// Code generated by go-bindata. DO NOT EDIT.
// sources:
// faces.dat
// partial_faces.dat

package main


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataFacesdat = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x97\xd1\x6e\x23\x35\x18\x85\xaf\xbb\xcf\x01\x77\x73\x31\x67\x66\xe2" +
	"\x19\x57\xa8\x12\x08\x04\x48\xc0\x2d\x12\xa8\xaa\x82\x1a\x89\x0a\x36\xad\xba\xe1\xa2\x6f\x8f\x36\xbb\x9b\x9d\xef" +
	"\xd8\x9e\x49\x96\xb6\xa4\x01\x59\x5a\xf5\xd8\x59\xfb\x1c\xfb\xcf\xaf\x2f\x57\xdf\x9d\x9f\x7d\xf6\xd3\xf2\xf5\xea" +
	"\xec\xf3\xef\xd7\x77\x7f\x6d\x7e\x6d\xce\xeb\xaa\xbe\xfc\xa2\x39\x57\xa8\x14\x2e\xc6\xd3\xba\x1c\xab\x06\xaa\x85" +
	"\xea\xa0\x16\x50\x01\xaa\x87\x1a\xa0\x22\x94\x6a\x4a\xba\x11\xed\x88\x7e\x44\x43\x1a\x3b\x52\x55\x43\x09\xaa\x81" +
	"\x6a\xa1\x3a\x28\xee\x19\xa0\x7a\xa8\x01\x2a\xf2\x74\x33\x43\x37\xa2\x1d\xd1\x8f\x68\x08\x29\x1b\xa4\x6c\x90\xb2" +
	"\x41\xca\x06\x29\x1b\xa4\x6c\x2a\xee\x19\xa0\x7a\xa8\x01\x2a\xf2\x74\x33\x43\x37\xa2\x1d\xd1\x8f\x68\x08\x29\x5b" +
	"\xa4\x6c\x91\xb2\x45\xca\x16\x29\x5b\xa4\x6c\x2b\xee\x19\xa0\x7a\xa8\x01\x2a\xf2\x74\x33\x43\x37\xa2\x1d\xd1\x8f" +
	"\x68\x08\x29\x3b\xa4\xec\x90\xb2\x43\xca\x0e\x29\x3b\xa4\xec\x2a\xee\x19\xa0\x7a\xa8\x01\x2a\xf2\x74\x33\x43\x37" +
	"\xa2\x1d\xd1\x8f\x68\x08\x29\x17\x48\xb9\x40\xca\x05\x52\x2e\x90\x72\x81\x94\x8b\x8a\x7b\x06\xa8\x1e\x6a\x80\x8a" +
	"\x3c\xdd\xcc\xd0\x8d\x68\x47\xf4\x23\x1a\x42\xca\x80\x94\x01\x29\x03\x52\x06\xa4\x0c\x48\x19\x2a\xee\x19\xa0\x7a" +
	"\xa8\x01\x2a\xf2\x74\x33\x43\x37\xa2\x1d\xd1\x8f\x68\x08\x29\x7b\xa4\xec\x91\xb2\x47\xca\x1e\x29\x7b\xa4\xec\x2b" +
	"\xee\x19\xa0\x7a\xa8\x01\x2a\xf2\x74\x33\x43\x37\xa2\x1d\xd1\x8f\x68\x08\x29\x07\xa4\x1c\x90\x72\x40\xca\x01\x29" +
	"\x07\xa4\x1c\x2a\xee\x19\xa0\x7a\xa8\x01\x2a\xf2\x74\x33\x43\x37\xa2\x1d\xd1\x8f\x68\x08\x29\x23\x52\x46\xa4\x8c" +
	"\x48\x19\x91\x32\x22\x65\xac\xb8\x67\x80\xea\xa1\x06\xa8\xc8\xd3\xcd\x0c\xdd\x88\x76\x44\x3f\xa2\x21\x52\x41\x4d" +
	"\x2c\x20\xfd\x88\xf8\x23\xf2\x8f\x08\x40\x22\x01\x89\x08\x24\x32\x90\x08\x41\x22\x05\xc9\x30\x48\xc6\x41\x32\x10" +
	"\x92\x91\x90\x0c\x85\xe4\x2c\x64\x30\x64\x34\x64\x38\x64\x3c\x64\x40\x64\x44\x64\x48\x64\x4c\x64\x50\x64\x54\xe4" +
	"\x58\xe4\x5c\xe4\x60\xe4\x64\xe4\x68\x64\x6c\x24\xc2\x91\x48\x47\x22\x1e\x89\x7c\x24\x02\x92\x48\x48\x22\x22\x89" +
	"\x8c\x24\x42\x92\x48\x49\x32\x4c\x92\x71\x92\x0c\x94\x64\xa4\x24\x43\x25\x19\x2b\x89\xb0\x24\xd2\x92\x88\x4b\x22" +
	"\x2f\x89\xc0\x24\x12\x93\x88\x4c\x22\x33\x89\xd0\x24\x52\x93\x0c\x9b\x64\xdc\x24\x03\x27\x19\x39\xc9\xd0\x49\xc6" +
	"\x4e\x22\x3c\x89\xf4\x24\xe2\x93\xc8\x4f\x22\x40\x89\x04\x25\x22\x94\xc8\x50\x22\x44\x89\x14\x25\xc3\x28\x19\x47" +
	"\xc9\x40\x4a\x46\x52\x32\x94\x92\xb1\x94\x08\x53\x22\x4d\x89\x38\x25\xf2\x94\x08\x54\x22\x51\x89\x48\x25\x32\x95" +
	"\x08\x55\x22\x55\xc9\xb0\x4a\xc6\x55\x32\xb0\x92\x91\x95\x0c\xad\xf4\x9e\xad\xbe\x79\x7d\xbb\xb9\xb9\x5d\x8f\x7e" +
	"\xca\x56\xcd\x05\xa7\xdf\x1e\xf3\xed\x6a\x7d\xbd\xba\xf7\x4f\x8d\x66\xb7\x5e\xae\x57\xeb\xcd\xcd\xe6\x61\xfc\x31" +
	"\xd5\x17\xb6\x90\x7c\xb2\xf1\x89\xd6\x27\x3a\x9f\x58\xf8\x44\xf0\x89\xde\x27\x06\x9f\x88\x97\xaf\xae\xbe\x3e\x3f" +
	"\xfb\xf2\xcf\xdf\x56\xf7\x9b\xdb\xab\xdf\x97\x77\x77\x0f\x67\xf5\x6e\x68\x34\xea\x64\x68\x42\x7d\xd4\x1a\xfd\x95" +
	"\x5b\x4d\x55\x6e\xb7\xd2\x9e\xf9\xbf\x55\xd8\x69\xdf\xf5\x77\xb3\x82\x3f\x65\xfc\xcd\xe9\xf2\x8d\xcd\xad\xe7\x3e" +
	"\x51\xba\xe1\xf2\x8e\xd3\xef\x57\x1a\xe9\x1d\xed\x06\xaa\xe5\xcd\xf2\x7a\xef\xb3\xf6\xa9\x95\xf4\xdf\xf4\xd6\x72" +
	"\x37\x78\x0c\xb5\x92\x3a\xe7\xff\x4a\x4f\x3e\x8d\x5a\x49\xd3\xed\xc6\xb6\x56\xbe\x5a\x6d\x36\x0f\xef\xfb\x8a\xec" +
	"\x55\xf2\xdf\xb6\xf4\x6e\x64\xb7\x2a\x7b\x9f\xfc\x6a\x4e\xe7\x9d\x6a\xa2\x33\xa5\x37\x9d\x5b\x3f\xe4\x2d\x34\x91" +
	"\x36\x97\x67\xfe\xed\x0e\x7d\xeb\x7d\x6a\x61\x6a\x35\x7d\xbf\xfd\x46\xfa\x1a\xbb\x31\xaa\x95\xb7\x5d\xe5\x31\x2b" +
	"\x25\xdf\xc1\xfd\x84\xf4\xb6\x8f\xa1\x52\x52\xa7\xe9\x1b\x8c\x77\x3e\x95\x4a\xf1\x37\x19\x8d\x6d\xa5\xfc\x70\xf3" +
	"\x66\xf9\x1f\x6a\x2a\x73\xfb\x2b\xdb\x54\xfc\xfc\x7f\xb7\xa9\x3c\x55\x29\x95\xbf\x1e\xa3\x52\xf9\xbf\xa7\x50\x1d" +
	"\x73\x4f\x79\xba\x9e\x93\xff\x25\xf2\xa1\x50\x7e\x5c\xde\xff\xf1\x62\x7e\x00\x95\x3b\x45\x79\xae\x46\x59\xee\xb3" +
	"\x9e\x2b\xf5\x52\x17\xc9\xeb\x43\x9f\xee\xf0\x52\xde\xe7\x05\xca\x27\xe6\xdf\xb7\xf4\x26\x28\x95\x97\xf1\xeb\xe7" +
	"\x79\x0a\x25\x77\x69\xec\x8a\xee\xe2\x34\x0a\x65\xa2\x93\x6e\x0b\xe5\xe7\xd5\xfa\xfa\x61\xd7\x54\x84\x5b\xe3\xdf" +
	"\xb9\xfd\x5c\xf1\x0b\x99\x76\x72\x7f\x85\x5c\xf7\xce\xdd\x45\xb9\x2d\xa5\x37\x9d\x5b\xcf\xff\x5d\x5a\x17\xda\x49" +
	"\xee\xfc\x52\x8b\x79\x9c\xb7\xfc\xa7\xeb\x73\xdf\xf5\xd2\xc8\xbf\xd4\x76\x8c\x6a\xe5\x5d\x57\xf9\x94\x4a\xc9\xdf" +
	"\xa3\xb7\xef\xf4\xfc\x3c\x69\x1f\x43\xa5\x78\x4a\x3f\xff\x25\x90\xca\xa7\x54\x4a\xe9\xa6\x3f\x54\xca\x2f\xcb\xf5" +
	"\xea\x19\x49\x25\x6d\x2a\x73\xed\x37\xbf\x67\xe9\x81\xa7\x0b\x68\x6e\x7d\xbe\xd4\xdd\x71\x5e\x4f\xdd\xd8\xe3\x97" +
	"\xc6\xe1\xa5\xaa\xec\x0b\x4f\xbc\xc6\xc7\x52\x79\x2e\x52\x99\x2b\xa4\x63\x28\x94\xd4\xb9\x77\x54\x3f\xf9\x34\x0a" +
	"\xa5\xdc\xff\xeb\x57\x7f\x07\x00\x00\xff\xff\x6c\x38\x01\xf7\xd3\x29\x00\x00")

func bindataFacesdatBytes() ([]byte, error) {
	return bindataRead(
		_bindataFacesdat,
		"faces.dat",
	)
}



func bindataFacesdat() (*asset, error) {
	bytes, err := bindataFacesdatBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "faces.dat",
		size: 10707,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1566723673, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPartialfacesdat = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x98\x5f\x6b\x24\x45\x14\xc5\x9f\xe3\xe7\xd0\xb7\x7e\xe8\xd3\x33\xd3" +
	"\x7f\x82\x2c\x28\x8a\x0a\xea\xab\xa0\x84\x30\x92\x01\x83\xee\x24\x64\xc7\x87\x7c\x7b\xd9\x90\x65\xfa\x77\xea\x56" +
	"\x75\x75\x66\xa5\x59\xd8\xdb\xdd\x55\x75\x4e\xf5\xe9\xdb\xbf\xc9\xed\x8f\xd7\x57\x5f\xfe\xba\x7f\x7f\xb8\xfa\xea" +
	"\xa7\xe3\xe3\xbf\xa7\x3f\xba\xeb\xb6\x69\x6f\xbe\xee\xae\xd5\x37\xea\xdf\xcd\x4f\xeb\x66\x5e\x75\xa8\x36\xa8\xb6" +
	"\xa8\x76\xa8\x7a\x54\x03\xaa\x11\xd5\x84\x4a\x2d\x4b\xaa\x11\xe5\x88\x7a\x44\x41\x9a\x2b\x52\xd3\xa2\x12\xaa\x0e" +
	"\xd5\x06\xd5\x16\x15\xe7\xec\x51\x0d\xa8\x46\x54\x13\x57\x37\x31\x54\x23\xca\x11\xf5\x88\x82\xe0\xb2\x83\xcb\x0e" +
	"\x2e\x3b\xb8\xec\xe0\xb2\x83\xcb\xae\xe1\x9c\x3d\xaa\x01\xd5\x88\x6a\xe2\xea\x26\x86\x6a\x44\x39\xa2\x1e\x51\x10" +
	"\x5c\x6e\xe0\x72\x03\x97\x1b\xb8\xdc\xc0\xe5\x06\x2e\x37\x0d\xe7\xec\x51\x0d\xa8\x46\x54\x13\x57\x37\x31\x54\x23" +
	"\xca\x11\xf5\x88\x82\xe0\x72\x0b\x97\x5b\xb8\xdc\xc2\xe5\x16\x2e\xb7\x70\xb9\x6d\x38\x67\x8f\x6a\x40\x35\xa2\x9a" +
	"\xb8\xba\x89\xa1\x1a\x51\x8e\xa8\x47\x14\x04\x97\x3b\xb8\xdc\xc1\xe5\x0e\x2e\x77\x70\xb9\x83\xcb\x5d\xc3\x39\x7b" +
	"\x54\x03\xaa\x11\xd5\xc4\xd5\x4d\x0c\xd5\x88\x72\x44\x3d\xa2\x20\xb8\xec\xe1\xb2\x87\xcb\x1e\x2e\x7b\xb8\xec\xe1" +
	"\xb2\x6f\x38\x67\x8f\x6a\x40\x35\xa2\x9a\xb8\xba\x89\xa1\x1a\x51\x8e\xa8\x47\x14\x04\x97\x03\x5c\x0e\x70\x39\xc0" +
	"\xe5\x00\x97\x03\x5c\x0e\x0d\xe7\xec\x51\x0d\xa8\x46\x54\x13\x57\x37\x31\x54\x23\xca\x11\xf5\x88\x82\xe0\x72\x84" +
	"\xcb\x11\x2e\x47\xb8\x1c\xe1\x72\x84\xcb\xb1\xe1\x9c\x3d\xaa\x01\xd5\x88\x6a\xe2\xea\x26\x86\x6a\x44\x39\xa2\x1e" +
	"\x51\x10\x5c\x4e\x70\x39\xc1\xe5\x04\x97\x13\x5c\x4e\x70\x39\x35\x9c\xb3\x47\x35\xa0\x1a\x51\x4d\x5c\xdd\xc4\x50" +
	"\x8d\x28\x47\xd4\x23\x0a\x22\x15\xb4\xc4\x02\xd2\x8f\x88\x3f\x22\xff\x88\x00\x24\x12\x90\x88\x40\x22\x03\x89\x10" +
	"\x24\x52\x90\x0c\x83\x64\x1c\x24\x03\x21\x19\x09\xc9\x50\x48\xce\x42\x06\x43\x46\x43\x86\x43\xc6\x43\x06\x44\x46" +
	"\x44\x86\x44\xc6\x44\x06\x45\x46\x45\x8e\x45\xce\x45\x0e\x46\x4e\x46\x8e\x46\xc6\x46\x22\x1c\x89\x74\x24\xe2\x91" +
	"\xc8\x47\x22\x20\x89\x84\x24\x22\x92\xc8\x48\x22\x24\x89\x94\x24\xc3\x24\x19\x27\xc9\x40\x49\x46\x4a\x32\x54\x92" +
	"\xb1\x92\x08\x4b\x22\x2d\x89\xb8\x24\xf2\x92\x08\x4c\x22\x31\x89\xc8\x24\x32\x93\x08\x4d\x22\x35\xc9\xb0\x49\xc6" +
	"\x4d\x32\x70\x92\x91\x93\x0c\x9d\x64\xec\x24\xc2\x93\x48\x4f\x22\x3e\x89\xfc\x24\x02\x94\x48\x50\x22\x42\x89\x0c" +
	"\x25\x42\x94\x48\x51\x32\x8c\x92\x71\x94\x0c\xa4\x64\x24\x25\x43\x29\x19\x4b\x89\x30\x25\xd2\x94\x88\x53\x22\x4f" +
	"\x89\x40\x25\x12\x95\x88\x54\x22\x53\x89\x50\x25\x52\x95\x0c\xab\x64\x5c\x25\x03\x2b\x19\x59\xc9\xd0\x4a\xaf\x6c" +
	"\xf5\xfd\xfb\x87\xd3\xfd\xc3\x71\xf6\x53\xb6\xe9\xde\xf1\xf4\xc7\x65\x7e\x38\x1c\xef\x0e\x4f\x7e\xd7\xec\xec\x8b" +
	"\x96\xbb\xc3\xf1\x74\x7f\x7a\x9e\xdf\xa6\xf6\x9d\x5d\x48\xee\xec\xfc\xc4\xc6\x4f\x6c\xfd\xc4\xce\x4f\xf4\x7e\x62" +
	"\xf0\x13\xa3\x9f\x98\x6e\xbe\xb8\xfd\xee\xfa\xea\x9b\x7f\xfe\x3c\x3c\x9d\x1e\x6e\xff\xda\x3f\x3e\x3e\x5f\xb5\xb3" +
	"\x43\xaf\x47\x1b\x1c\xf1\x59\xbf\x9a\x1b\x3d\xbf\x2b\xbd\x63\x69\xc5\xe8\x7f\xf1\x28\x99\x9b\xf2\x18\xbf\xa2\x42" +
	"\x55\xf2\x51\xaf\xbf\xb4\xfa\x9a\x23\xef\x3d\xff\x04\x4b\x33\x85\x63\x90\x96\x0f\xfb\xbb\xba\x51\x95\x59\xa9\x71" +
	"\xbf\x94\x80\xf4\x4e\xe1\xce\xf4\xff\xb9\x55\xe3\xf1\x6b\x3d\xc5\xf7\xa9\xe2\xf9\x2e\xf9\xab\x1d\x5d\xab\xfc\x53" +
	"\x4a\xd6\x65\xa5\xcd\x8f\x79\xc9\xca\xb7\x87\xd3\xe9\x79\xb1\xaf\x28\xf8\x97\x5e\x71\xf5\x4a\xf6\x52\x96\xa5\xba" +
	"\x3d\x54\x61\x76\x1f\xbb\xd4\x61\x6a\xf6\x2b\x55\x5b\xda\xdb\x35\x49\x89\xea\xa5\xde\x5a\x76\x94\x5f\x7b\x7d\x2a" +
	"\xfd\xe9\xcc\x8e\x59\x56\x3e\x75\x15\xbe\x7f\x51\x52\x72\xb3\x97\x7c\xe4\x92\x9f\x8e\x89\x32\x1d\xef\x5c\x2e\x29" +
	"\xe5\xeb\xb5\xcf\xd5\xdd\xc7\x8e\xdf\x96\x94\xa5\xab\xcb\xcf\xb8\x2e\x29\x6b\x7b\x4a\xdc\x9d\x5f\x8e\x97\xa4\xfc" +
	"\x7c\xff\x61\xff\xda\x54\xd2\x70\x2c\x47\x25\x9a\x98\xcb\xe5\x23\x16\xc5\x38\xb6\xb7\x66\xb3\x4b\xf1\xcc\x45\x31" +
	"\x6e\x92\xb1\xa6\xf5\x2d\xec\xb2\xa8\xfc\x3f\x4d\x27\x82\xd1\xfc\xeb\x31\x8b\xca\xc7\x9e\x92\xbe\x9b\x3e\x2c\x0e" +
	"\x4c\xbe\xfb\x94\xfb\x77\x7d\x4f\xa9\xdf\x8c\xb8\xa7\x94\x1f\xee\xba\x9e\x93\xce\xbc\x76\x7c\x69\xf5\x9a\x19\x3e" +
	"\x57\x50\x4a\x67\xa3\xa0\xfc\xb2\x7f\xfa\xfb\xa2\x1f\x40\xb9\x8f\x4f\x79\x6c\xda\x77\x6a\x46\xd5\x35\xde\xf4\xfa" +
	"\xf2\x47\x29\xde\x40\x25\xdd\x75\x69\x96\xcb\x3e\x45\x35\x63\x96\xa3\x55\x13\xdd\x08\x7a\x73\xcf\x04\x51\xe1\xaf" +
	"\x9f\x36\xb3\x41\xeb\x6d\xf9\x99\xf8\x4d\xaf\x0f\x4a\x7c\x36\xea\x48\x51\x50\x96\x3e\x22\x1e\xa8\x38\xf4\xf9\xde" +
	"\xfb\x96\xdd\x59\xb7\x7f\x9f\xe3\xe3\x76\x7e\xb6\x51\xc4\xf2\x41\xf9\xed\x70\xbc\x3b\xff\xfa\x99\x7f\x39\x94\x27" +
	"\x9c\xcc\xec\xad\xed\x73\x29\x01\xb9\x99\xeb\xd3\x99\xef\xac\x7e\x3d\xf7\xfb\x3a\x9e\x2b\xa2\x79\xcf\xd5\xd2\x8e" +
	"\x94\xcf\xac\x7b\xf6\x6f\xf9\x5e\x6a\xe1\xb3\x90\x57\x9e\x79\x9b\x66\x59\x39\xff\xfa\x51\xb2\x5f\x4a\xde\x60\xee" +
	"\x7c\x7e\x1f\xa3\x84\xfa\x1b\x5c\xf3\x9e\xc5\xe3\xa3\xbd\x8a\x3b\xd5\x9a\xcf\x5b\xea\x39\x7e\xfb\xe2\x2b\xb5\x4f" +
	"\xe4\x92\xa3\x6e\xc7\xd6\xf6\xbb\xd4\x8f\x27\xe5\xf7\xfd\xf1\x10\x92\x4a\x7e\x23\xca\xa1\x2e\x3f\x4a\xb6\x93\xf5" +
	"\x0d\xa4\xfc\x09\x59\xfa\xf3\x5a\x6e\xb4\x37\x91\xf8\x25\xc9\xeb\x5e\x8a\x62\xad\xbb\x78\x86\xa5\x39\xd7\x8e\xf7" +
	"\x4f\x47\xc5\x1e\x9d\xa3\x92\xfe\x9d\x36\xdf\xbd\xea\x82\x92\xdf\x56\xbf\x33\xff\xf8\x4a\x6b\xa7\x5f\x8c\x9a\x20" +
	"\xe4\x74\xba\x16\x65\x15\x46\x71\x5c\xf7\x28\xd7\x5d\xbd\x3c\x28\xe9\xfd\x9a\x85\xa4\xf4\x65\x3d\x07\xe5\xbf\x00" +
	"\x00\x00\xff\xff\x0f\x0d\x16\x5d\xd3\x29\x00\x00")

func bindataPartialfacesdatBytes() ([]byte, error) {
	return bindataRead(
		_bindataPartialfacesdat,
		"partial_faces.dat",
	)
}



func bindataPartialfacesdat() (*asset, error) {
	bytes, err := bindataPartialfacesdatBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "partial_faces.dat",
		size: 10707,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1566723766, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"faces.dat":         bindataFacesdat,
	"partial_faces.dat": bindataPartialfacesdat,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"faces.dat": {Func: bindataFacesdat, Children: map[string]*bintree{}},
	"partial_faces.dat": {Func: bindataPartialfacesdat, Children: map[string]*bintree{}},
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