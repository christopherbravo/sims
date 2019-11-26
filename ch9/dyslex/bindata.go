// Code generated by go-bindata. DO NOT EDIT.
// sources:
// train_pats.tsv
// semantics.tsv

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

var _bindataTrainpatstsv = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5d\x5d\x8f\xe3\x36\xb2\x7d\x9e\xfb\x2b\xf4\x90\xfb\x36\x03\xb4\x3e" +
	"\xbb\x7b\x70\x11\xcc\x00\x99\x8b\x9b\x97\x3b\x0b\x2c\xf6\x29\x08\x08\xb6\x44\xdb\x8a\x25\xd2\x21\x69\x7b\x3d\xbf" +
	"\x7e\x21\xc9\x5f\x52\x15\x29\x52\x56\x77\xdc\x81\xb0\xdb\x3d\x24\x2d\xd5\x39\xae\xc3\xaa\x12\xc9\x0e\xf9\xbf\xcf" +
	"\x1f\x7e\xfa\x7f\x5a\xb3\x0f\xff\xfd\x5d\xea\x95\x58\x4a\xba\x59\x1d\x7e\x8b\x3e\x3f\x7c\x7c\xf8\xfd\x7f\xa2\xcf" +
	"\xd9\xc7\xa7\x9f\x61\x57\xf8\x3b\x6c\x8b\x90\xb6\x18\x69\x4b\x90\xb6\x14\x69\xcb\x90\xb6\x47\xd0\x16\x7e\x7c\x40" +
	"\xda\xa0\x7d\x21\x62\x5f\x88\xd8\x17\x22\xf6\x85\x88\x7d\x21\x62\x5f\x88\xd8\x17\x21\xf6\x45\x88\x7d\x11\x62\x5f" +
	"\x84\xd8\x17\x21\xf6\x45\x88\x7d\x11\x62\x5f\x84\xd8\x17\x23\xf6\xc5\x88\x7d\x31\x62\x5f\x8c\xd8\x17\x23\xf6\xc5" +
	"\x88\x7d\x31\x62\x5f\x8c\xd8\x97\x20\xf6\x25\x88\x7d\x09\x62\x5f\x82\xd8\x97\x20\xf6\x25\x88\x7d\x09\x62\x5f\x82" +
	"\xd8\x97\x22\xf6\xa5\x88\x7d\x29\x62\x5f\x8a\xd8\x97\x22\xf6\xa5\x88\x7d\x29\x62\x5f\xda\xda\xf7\x4f\x56\x53\xae" +
	"\xcb\x5c\x5d\x44\x1c\x3e\x7c\x0c\xa3\x9f\x87\x5d\x21\x18\x1c\x81\x96\x18\xb4\x24\xa0\x25\x05\x2d\x19\x68\x81\x86" +
	"\x3d\x81\x96\x67\xd0\x12\x3e\xc0\xa6\xa1\xd5\xdd\x04\xd0\x6f\x81\x63\x86\xc8\x42\x80\x2c\x04\xc8\x42\x80\x2c\x04" +
	"\xc8\x42\x80\x2c\x04\xc8\x42\x80\x2c\x84\xc8\x42\x88\x2c\x02\xc8\x22\x80\x2c\x02\xc8\x22\x80\x2c\x02\xc8\x22\x80" +
	"\x2c\x02\xc8\x22\x80\x2c\x02\xc8\x22\x80\x2c\x82\xc8\x22\x88\x2c\x06\xc8\x62\x80\x2c\x06\xc8\x62\x80\x2c\x06\xc8" +
	"\x62\x80\x2c\x06\xc8\x62\x80\x2c\x06\xc8\x62\x80\x2c\x86\xc8\x62\x88\x2c\x01\xc8\x12\x80\x2c\x01\xc8\x12\x80\x2c" +
	"\x01\xc8\x12\x80\x2c\x01\xc8\x12\x80\x2c\x01\xc8\x12\x80\x2c\x81\xc8\x12\x88\x2c\x05\xc8\x52\x80\x2c\x05\xc8\x52" +
	"\x80\x2c\x05\xc8\x52\x80\x2c\x05\xc8\xe0\xd4\x96\x02\x64\x29\x40\x96\x42\x64\x29\x44\x96\x01\x64\x19\x40\x96\x01" +
	"\x64\x19\x40\x96\x01\x64\x19\x40\x96\x01\x64\x19\x40\x96\x01\x64\x19\x40\x96\x41\x64\x19\x44\xf6\x08\x90\x3d\x02" +
	"\x64\x8f\x00\xd9\x23\x40\xf6\x08\x90\x3d\x02\x64\x8f\x00\xd9\x23\x40\xf6\x08\x90\x3d\x02\x64\x8f\x10\xd9\x23\x44" +
	"\xf6\x04\x90\x3d\x01\x64\x4f\x00\xd9\x13\x40\xf6\x04\x90\x3d\x01\x64\x4f\x00\xd9\x13\x40\xf6\x04\x90\x3d\x01\x64" +
	"\x4f\x10\xd9\x13\x44\xf6\x0c\x90\x3d\x03\x64\xcf\x00\xd9\x33\x40\xf6\x0c\x90\x3d\x03\x64\xcf\x00\xd9\x33\x40\xf6" +
	"\x0c\x90\x3d\x03\x64\xcf\x10\xd9\x73\x87\xec\x1f\x2b\xc1\x45\x25\x96\x87\xdf\x92\x26\x19\xf9\xd8\x26\x24\x49\x13" +
	"\x36\x3f\x3e\x7e\x6c\x52\x12\x38\x00\xfb\x58\x17\xe4\x61\x2b\x36\x36\x42\xc7\x46\xe8\xd8\x18\x1d\x1b\xa3\x63\x13" +
	"\x74\x6c\x82\x8e\x4d\xd1\xb1\x29\x3a\x36\x43\xc7\x66\xc8\xd8\xb0\xa5\x0f\x6b\xc5\xc6\x62\x9c\x85\x28\x67\x21\xca" +
	"\x59\x88\x72\x16\xa2\x9c\x85\x28\x67\x21\xca\x59\x88\x72\x16\xa2\x9c\x85\x28\x67\x21\xca\x59\x88\x72\x16\xa1\x9c" +
	"\x45\x28\x67\x11\xca\x59\x84\x72\x16\xa1\x9c\x45\x28\x67\x11\xca\x59\x84\x72\x16\xa1\x9c\x45\x28\x67\x11\xca\x59" +
	"\x84\x72\x16\xa1\x9c\x45\x28\x67\x31\xca\x59\x8c\x72\x16\xa3\x9c\xc5\x28\x67\x31\xca\x59\x8c\x72\x16\xa3\x9c\xc5" +
	"\x28\x67\x31\xca\x59\x8c\x72\x16\xa3\x9c\xc5\x28\x67\x31\xca\x59\x8c\x72\x96\xa0\x9c\x25\x28\x67\x09\xca\x59\x82" +
	"\x72\x96\xa0\x9c\x25\x28\x67\x09\xca\x59\x82\x72\x96\xa0\x9c\x25\x28\x67\x09\xca\x59\x82\x72\x96\xa0\x9c\x25\x28" +
	"\x67\x29\xca\x59\x8a\x72\x96\xa2\x9c\xa5\x28\x67\x29\xca\x59\x8a\x72\x96\xa2\x9c\xa5\x28\x67\x29\xca\x59\x8a\x72" +
	"\x96\xa2\x9c\xa5\x28\x67\x29\xca\x59\x8a\x72\x96\xa1\x9c\x65\x28\x67\x19\xca\x59\x86\x72\x96\xa1\x9c\x65\x28\x67" +
	"\x19\xca\x59\x86\x72\x96\xa1\x9c\x65\x28\x67\x19\xca\x59\x86\x72\x96\xa1\x9c\x65\x1d\x67\x3f\xfd\x2f\xa3\x7a\x2b" +
	"\x99\xfa\x2f\xf2\xcb\xe7\x0f\x9a\x4a\x4d\xb4\x6e\x7e\xe8\x0f\x0f\xe8\x15\x1e\xff\x67\xbb\xf0\xfe\x61\x6b\x38\xf8" +
	"\x2d\x6c\x2f\x38\x22\xec\x3d\x37\xbc\xc1\x02\xd3\x38\xf3\x53\x7d\xef\xe7\x7f\x99\xef\x6c\x7b\x26\xd6\x77\x2b\x27" +
	"\x2e\xcf\xab\xe9\xbf\x89\x2a\x7f\x30\x52\x31\xa5\xc8\x42\x08\x1d\xd4\xb4\xe4\x44\xad\xe8\x86\x91\xe8\x97\x20\x97" +
	"\x42\x29\xa2\x58\xae\x4b\xc1\x89\x64\xb9\xa6\x7c\xb9\xad\xa8\x1c\xf4\xe4\xa5\xcc\x91\x66\xa1\x57\x4c\x06\x2f\x52" +
	"\xec\x79\x50\x50\xb9\x0e\x94\x58\xe8\x40\xed\x19\xd3\x41\xc9\x0b\x21\xa4\x0a\x4a\x4e\xd6\xa5\xce\x57\x8c\x07\x82" +
	"\x13\xb5\x95\x0b\x9a\xb3\x60\x29\x34\x59\x48\x51\x93\x4d\x45\xb9\x56\x01\x2f\x73\x16\xd4\x94\x93\x9a\x16\x2c\x58" +
	"\x08\x49\x18\xd5\x25\x5f\x92\x42\x96\x7c\x5d\xf2\x65\xdb\x56\x6d\x79\xbe\x22\x45\xc9\x39\x93\x41\x45\xe5\x92\x75" +
	"\x3a\x60\xbc\xd5\x01\xe3\x56\x1d\x98\x39\x9b\xfa\xdd\x0d\x7f\x9a\x47\xf4\xdb\x2e\xca\x70\xbf\xa0\x87\x87\x88\x2a" +
	"\x4d\xcf\x1c\x63\xc5\x57\x31\xf3\xe9\x60\xea\xbd\x7c\x9e\x77\xd6\xc1\x52\x32\xaa\x99\x24\x7a\x2f\xc8\x81\xca\x42" +
	"\x5d\xeb\x21\x1e\xea\xa1\x73\xef\x1d\x95\x25\x2b\x48\x2e\xaa\xc6\x9d\x5b\x0f\x17\x9c\x2c\xa5\xd8\xf2\xa2\xf1\xee" +
	"\x5c\x6c\xb9\x96\x87\x60\xd1\x34\x90\xbd\x10\x85\x3a\xfe\xce\x19\x95\x44\x69\xc9\x68\x7d\x6a\xaa\x9b\xc1\xb4\xe4" +
	"\xea\xa2\x00\xca\xcb\x9a\x56\x47\x09\x14\x94\x2f\x99\x14\x5b\x85\x6a\x60\x43\xa5\x2e\x5b\x21\x56\x07\x42\x95\x12" +
	"\x39\xc9\x57\x65\x55\x04\x0b\x56\x97\xbc\xe4\x2c\xa8\x19\x55\x5b\xc9\x6a\xc6\x75\xb0\xd8\x36\xa1\x82\x6c\x44\x23" +
	"\x8c\x92\x56\x81\x64\x15\xd5\x4c\x91\x5a\x70\x76\x38\xff\x6b\x2f\xe4\xba\x95\x51\x23\x4c\xb2\x58\x2c\xbe\x2a\xa5" +
	"\x2c\xe4\x9a\x03\xc3\x34\x57\x30\x4b\x13\x0b\x39\x0f\xe7\xdf\x30\xd1\xf5\x65\x81\x8f\xba\xf5\xf2\x0b\x56\xfe\x7c" +
	"\xf8\xf4\xf9\xa0\xbb\xfd\xe9\xb7\x84\x93\x46\x3d\x9d\xa0\x1a\x45\x08\xbe\xec\x94\x54\x8b\x1d\x53\xe7\x58\xd1\xf6" +
	"\xef\x4b\xc5\x88\xda\x6e\x36\x42\x6a\x56\x04\x62\xab\xdb\x4e\xd2\x48\xad\xd4\x87\x6b\xc9\x1d\xa3\x09\x11\x0b\xf2" +
	"\x22\x8a\x43\x40\x5f\xc4\x8e\x91\x3d\x2d\x95\x0e\x78\x93\x28\xd1\x2a\xa8\x69\x5d\xd3\x2a\x78\x29\x65\x11\x54\xe5" +
	"\xae\x91\xd1\x6a\x5b\x53\x1e\xd0\xd6\xca\xe0\xcf\x2d\xad\x4a\x7d\x20\x45\xb9\x58\x34\xda\xd2\x87\x56\x0c\x05\x63" +
	"\x92\x14\x45\xf1\x4d\x4a\xe9\xf5\x15\x4c\x8d\xdb\xfd\x88\x31\x9c\x99\xed\xff\x6f\x8e\x29\xf8\x1d\xcc\x76\xb8\xb4" +
	"\xce\x21\xa9\xf9\xa5\xe0\x63\xd9\x1c\x4f\x9f\x18\x53\xce\x99\xd4\x8a\x2a\x52\xb1\xa5\x6a\x7f\xe1\x2c\x5f\x13\x21" +
	"\x9b\x20\xd3\xf4\x0d\x33\xaa\x4e\x25\x73\x44\x9d\x81\x2a\xf6\x4d\xf0\x28\x04\x53\x44\x6e\xf9\x49\x1e\x6d\x20\xda" +
	"\x2a\x56\x90\x26\x06\x2d\x69\xcd\x54\x63\x9b\x64\x79\x83\xb4\xd1\xcc\x55\x94\x69\xc5\x92\x0b\xaa\xc9\x7a\xbd\xfe" +
	"\xae\xb5\x76\xf8\x12\xec\xc9\xd8\xf8\x17\x82\x49\xe5\x61\xe0\xfe\xa1\xe1\x29\xb6\xc8\xe1\x63\x13\x96\x34\xbd\x5e" +
	"\xd9\x31\x7e\xff\xd7\x8d\x1f\xb7\x3e\xf9\x2c\x96\x26\x68\x10\x2d\x70\xb1\x58\x23\x48\x23\x13\x2a\x6b\x54\x2f\xfd" +
	"\x14\xed\xa2\x9b\x39\xe3\xc9\xb0\x62\xc1\xf3\xb7\x5e\x09\xb3\x67\x54\x36\x82\x6a\x35\xb2\x94\x65\xa3\x5f\x59\x72" +
	"\xce\xbd\x79\xf5\xfd\x06\xc7\xca\x02\xac\xc7\x2d\xbb\xb2\xdf\xd3\xdf\xc3\x6e\xed\x9f\x32\x76\x5e\xad\xcc\xa5\x4a" +
	"\xcf\xec\xaa\x2b\x4e\xde\x3e\xaf\xba\x0e\x12\x5d\x22\xa5\x6a\x5a\x55\x81\x2e\x6b\x46\x5e\xd8\x42\xc8\xae\x28\xaf" +
	"\x44\xbe\x26\x55\x55\xd1\xf5\x7a\xed\xc0\x89\xb9\x94\xf6\xfd\x0e\x4c\xe5\xac\xd9\x7b\xe1\x8b\xac\x70\x46\x1f\x77" +
	"\x43\xe6\x7b\x57\xb7\x71\x73\xf7\xcd\x79\x3f\xbb\xbf\x83\xc4\x69\xa4\x9a\x58\x51\x59\xf8\xf9\x7b\x97\x1a\x95\x9c" +
	"\x6c\xb6\x2f\x55\x99\x93\x97\x6d\x59\x15\x25\x5f\xaa\x4b\x8f\x96\x94\xab\xe6\xf3\x97\xa6\x05\xcd\xb5\x90\x25\x6b" +
	"\x02\x57\xd1\x2a\xa4\x66\xba\x4d\xa8\xae\x26\xfe\x4e\x9b\x48\xb5\x4e\x8b\x6d\xa5\xcf\x95\xf7\x46\x28\xc5\x94\x2a" +
	"\x05\x6f\x15\x23\xc5\x86\x11\x29\xe5\xf7\xcd\x66\xe3\xc0\xe1\x78\xaa\xee\x32\xb7\x9a\xfc\xdc\xae\x80\x10\x51\x9c" +
	"\x8f\x62\xe0\x3d\x30\xbb\xa6\xe4\x53\xee\x9f\x99\xcf\x97\x6f\x9b\x21\xdc\xfb\xdc\x4a\x8e\xd0\x58\x72\xec\x57\xa5" +
	"\x66\xc7\xe2\xa2\xd5\x8b\x87\x4e\x44\xe3\xfa\xb2\xf6\x15\x47\xa7\x84\x61\xee\xe4\x23\x16\x6b\x21\x72\x51\xd2\x9e" +
	"\xc9\x56\x44\x2b\x2a\x19\x59\xad\x56\x5f\xbb\xba\xdd\x44\xb2\x7b\xda\xe1\x1a\x90\x6c\x55\xb6\xad\x6e\x3f\xfd\x84" +
	"\x35\xf6\xf4\xba\x7d\x6a\xc8\x71\xbd\xe6\x0d\x2b\x6f\x2b\x59\xb7\x52\xc4\xa1\x6e\xf7\x2e\xd2\x2f\x55\xf8\x40\x52" +
	"\xce\x45\x39\xf6\x4e\xd8\xaa\x8f\x2e\x13\xa3\x4a\x35\x99\xd8\x97\xcb\x7b\x5d\x8c\xb2\x29\xb3\x98\x2d\xcf\xea\xcf" +
	"\xe1\x78\x5c\xba\xae\x3c\x86\x55\x08\x76\xdf\x10\xdc\xf7\xc1\xb8\xc4\x38\x25\x87\xb9\xed\x9a\x4b\x16\x7e\x16\xde" +
	"\xfe\xd4\x89\x92\x18\x56\xe7\xad\x2a\x1c\x0b\x92\xb3\x4a\x6c\x65\xc8\x40\x17\xad\x24\xd4\xbe\xac\x71\x71\x18\x17" +
	"\x47\xba\x2a\xe5\xb4\x44\xd2\x2d\x75\x54\x94\x93\x45\x55\x89\x4b\x31\x3e\x5e\xf8\x4d\x0d\xf3\xe3\x05\xc7\xf0\x29" +
	"\xa1\xc5\xad\xa7\xb9\xca\x30\x8a\xd8\xdb\x6e\xbf\x5c\xee\x33\xd5\x71\x6d\x23\xe6\x92\xa0\x67\x31\x6e\x5d\x1f\x47" +
	"\x16\x10\xdf\x7c\x89\x7c\x55\xf2\xa2\x49\x8b\x7e\xe5\x45\x31\x42\x89\x4f\x82\x31\x4e\xe7\x30\x89\xb1\x27\x44\x7d" +
	"\x81\xf8\x2e\x67\xf8\x89\xc1\x15\x8b\xfb\x35\x6f\x5a\xe4\x66\x9b\x5f\x35\x62\xee\x9b\x2d\x2d\xba\xbb\xe5\x0c\x6b" +
	"\x82\x84\x2d\x9c\xb7\x82\xd9\xd3\x1d\x23\xfb\xfd\xfe\xeb\x6e\xb7\xb3\xd0\xe9\x3e\xcf\xb9\x4a\xc6\x14\x6d\xc6\x5d" +
	"\xdf\x45\xcc\xd3\xc4\xed\x2e\x98\xdb\xc4\x34\x4d\x08\x53\xab\x9d\xdb\xe4\x7a\xd3\x9e\x12\x24\x4b\x02\xf2\x40\x6a" +
	"\x73\x20\x94\x4e\x1c\x8c\x9e\x85\xd0\x2a\xe0\x54\x7c\x57\xe5\x9f\xdb\xb2\xb8\xda\x53\xe2\xba\xa8\x07\x5e\xe6\x2e" +
	"\x2a\x46\x9b\x7c\xe9\xdb\xa7\x4f\x9f\xbc\xbe\x2b\xf7\x00\x63\xff\xbc\x79\x7a\xef\x97\x10\xc3\x71\xe1\x60\xd4\xb0" +
	"\x77\x2c\x6c\xe0\x0a\x87\x19\xdd\xad\xd7\x6b\xe4\x4b\x53\x0a\xa0\xa9\x7d\x9e\x2f\x73\x3b\x15\x9c\x63\x86\xbd\x42" +
	"\xb8\xa4\x44\x63\xef\xa6\x86\xeb\x16\x3d\x55\xb4\x71\x61\x51\x1d\x4e\x71\x21\xa7\x92\x97\x3b\x21\x59\xb0\xe5\x9b" +
	"\x8a\x51\x45\xb9\xbe\xd2\x8a\xb1\x9c\xb8\x24\x55\x4a\xb7\xa1\x49\xd3\xfe\xbb\x26\x1b\x63\x3e\x11\xc3\xd4\x0a\xe3" +
	"\x09\x3e\xcb\x43\x4f\xf7\x7b\xb2\x39\x76\x8c\x47\x34\xff\xec\xdc\xd7\xba\x5b\xfb\x6e\xa9\xac\x5c\xfb\x66\x5a\xf6" +
	"\x1e\xa4\x53\xed\xab\xdb\xae\x82\x70\x7a\x81\xdb\x6e\x7d\xba\x14\xcf\xfd\x57\xb3\x57\xbe\xcf\xd9\x92\xea\x72\xd7" +
	"\xdf\x44\x58\xf2\x9d\xa8\x76\x4c\x91\x7c\xd5\x48\xa3\x5b\xa8\x60\xac\x20\x52\xca\x6f\x45\xaf\x98\x80\x0c\x4c\xc9" +
	"\x3e\x4c\x23\x6c\x51\xde\x5c\x6e\x5c\x3e\x8d\x67\x35\x63\x77\x9e\xee\x23\xf3\x5d\xf3\x78\xfd\x6d\xb9\x95\x5f\xdf" +
	"\xc4\x85\x0a\xf7\x0d\x83\x96\xda\xe1\x9c\x12\x21\xe5\x43\x2f\x1e\x1c\xc3\x40\x5b\x5d\x5f\x2d\x44\x1c\x17\xaf\x05" +
	"\x27\x55\x55\xfd\xeb\x7a\xb3\x86\x19\xb3\xeb\xfc\x32\x1e\x91\xb1\xdc\xfe\xc1\xb1\xf0\xb5\x95\xcb\x0f\xbd\xfc\x68" +
	"\x4a\x66\xf0\x3a\xd7\xfc\xbe\x67\xc7\x30\xd7\xf3\xe6\x28\x97\x07\xd3\xfa\xbc\x35\xc1\x48\xc1\xdc\xc6\x84\x7e\x5a" +
	"\x74\x89\x10\x20\x41\x6a\x5f\x3d\x75\x3b\xfd\xa8\x62\x64\xbd\x5e\xb7\x7b\xc4\xa7\x7e\x09\x0f\xce\x93\x2e\xe6\xe6" +
	"\xf8\x18\x5c\x10\xfd\x71\x63\x85\x89\x1b\x0e\x5c\x86\x3e\x97\x3b\x3b\x73\xf4\xf9\xd8\x78\xfb\x73\x67\x58\x4b\xe8" +
	"\xbf\x2f\xed\x6f\xec\x30\xaf\x1b\x78\x6d\xe7\xe8\xe7\x40\xe7\x97\xaa\xb9\x68\x65\xc2\xa4\x7d\x61\xba\xab\x03\xe0" +
	"\xda\xf3\xa2\xa2\xcb\xa6\x4a\xfe\xb2\x5c\x2e\x47\x49\xf3\x15\xc0\xd8\x1d\xfc\x6a\x06\x7b\x3d\xe0\x2a\x0e\xd3\x93" +
	"\xcd\x62\x75\xbf\xfb\xd4\x4f\xfa\xe0\x7a\x8b\x1a\xe0\xfa\x9a\xa1\x1e\x40\x16\x13\xcc\xb1\x01\x88\xc4\xbe\x2d\x03" +
	"\xae\x30\x1c\x6b\x04\xdb\x6b\xa3\xee\x18\x45\xeb\xff\x1b\xa1\x34\xd9\x6c\x36\xdf\x15\x72\x0e\x0f\xf7\x12\x77\xfe" +
	"\x4c\x9f\xbf\x78\xf8\xd8\x64\x6d\x6e\x33\x65\x5d\xf6\x6f\xd8\x65\xa9\xfa\xf5\xae\xb7\x0f\x11\xb7\xcf\x54\x6e\xfe" +
	"\x6f\xaf\x0b\xf0\x7c\x0a\x89\x19\xa3\xb1\x02\x6e\x69\x32\x6f\xee\xf3\xdc\xcd\x34\xbe\x29\x43\xd3\xbc\x3d\xb3\xfa" +
	"\x65\xed\x75\x76\xfb\x96\x6f\x6a\x18\x71\xb0\x08\xe4\x97\xd6\x9c\x54\x80\x57\x9a\x36\x15\x8c\x6b\xd5\xe5\x2e\x0f" +
	"\x83\xda\x67\xaa\x87\xba\x47\xd9\x71\x3e\x5c\x79\x1b\xeb\xe3\xa2\x95\x4a\xd0\xfc\x6c\x64\xd1\xfc\xd2\x73\xfe\xa6" +
	"\xa1\x49\x4e\x54\xdb\xd3\x38\x7e\xf3\xcb\x82\xb1\xaa\xf9\x59\x89\xfc\xf8\x5a\xff\xe4\x9d\xbc\x29\x69\x8f\xfb\x25" +
	"\x04\xc9\x2b\xaa\xda\x4f\xd2\x5c\x07\x9c\x6d\x35\x6b\x3f\x2e\x45\xc5\x2e\x2f\x82\xda\x57\x9d\x81\xd2\x54\xb7\x8f" +
	"\x2f\xb6\xf2\xb8\x6a\x56\xe6\xbd\xed\x77\x92\xfd\xb9\x65\x4a\x23\x67\xe9\xce\x4d\x42\x2e\x29\x2f\x7f\xb0\xe2\xdc" +
	"\xa2\x18\x57\xa5\x2e\x77\xe5\xf1\xc4\x9d\x64\x5c\x13\x29\xa5\xe1\x14\xf7\xf8\xec\xec\x3f\x23\xde\x93\x22\xf0\x08" +
	"\x67\x8f\x5c\x78\x81\xe3\x66\xf1\xdc\xb9\xcd\x6d\xf9\xde\xbd\x2b\x42\x10\xba\x58\xb0\xa6\x43\x90\x9a\x2e\x79\xa9" +
	"\xb7\xc5\x49\x19\x25\x5f\x08\x59\x77\xf7\xd6\xac\xde\x08\x49\xe5\x01\xd9\xe4\xdd\x3b\x71\xed\xa0\x9d\xd3\xa1\x6c" +
	"\x4d\x16\x8b\xc5\x55\x9c\x70\x89\xe0\x7f\x97\x38\x31\x56\x5d\xd8\x6a\x20\x37\x2f\x7e\x3b\x1d\xf8\x8d\x71\x7d\xfe" +
	"\x5d\xa8\xe2\x78\x04\x48\x6e\xd9\xe9\x58\xf5\x75\xb8\x38\xb9\x3a\xdb\x31\x7e\x71\xf4\x36\x53\x3a\x9d\xb5\x2e\xda" +
	"\xb3\xd6\xdd\x7a\x82\xbb\x2f\xd8\xfd\xd7\x85\xd9\x7b\xf0\x71\x93\x9e\x43\xf0\x39\x38\xdf\xbb\x30\x30\xb7\x87\x8f" +
	"\xf5\xce\xff\xbc\x7b\xc8\x85\x3a\x17\xbf\x5e\x1a\xeb\x66\xff\xd6\xe9\xaf\xf6\x4b\x98\xff\xec\xc6\xc0\xfd\x4f\x09" +
	"\xd1\xf1\x04\xb5\x6a\x4f\x50\x0b\xe5\x71\x82\xda\xb5\xe7\xbd\x4e\xf3\x98\xdc\x4d\x77\xc1\xd3\x74\x9f\xb0\x38\x5d" +
	"\x0a\xb7\x7e\xfe\xdd\x4c\xf3\x48\xf2\xc3\x05\xd1\x72\xab\x57\xe8\x84\x8f\xe4\x3e\xc3\x74\xe7\x12\x05\x96\xb4\xe4" +
	"\x64\xb9\x5c\x7e\xb5\x1f\x90\xf6\x9d\xf3\x5d\xc8\xbd\x07\x09\xe0\x4e\x6d\x92\x00\xd6\x32\xce\xcc\x5f\x91\xeb\xcc" +
	"\xf7\xd4\x7b\x88\x02\x36\x05\x5c\x07\x81\xfe\x9c\x8f\xe8\x01\xfa\x7f\x45\xbb\x03\xd3\x5f\xec\x07\xa6\xfd\x32\x7d" +
	"\xb7\xde\x7b\xf0\xff\x5b\x43\x80\xbb\x8f\xce\xe5\x8f\xaf\x19\x14\xb0\xbe\x77\x19\x02\xba\xbf\x35\x06\xbd\xbd\xb9" +
	"\x6d\x7b\xd8\xb9\xaa\xaa\x09\x5c\xfb\x7d\xab\xfd\x11\xf7\xeb\xed\x26\x5b\xdc\x63\xde\x94\x99\xda\xaf\xef\xf6\x54" +
	"\xc6\xad\xef\x2e\xbc\xbd\x4b\xfa\xe1\x6b\x1e\xcc\xd3\xdb\xb7\x3a\xd0\xd1\x57\x65\x77\x20\xf9\x57\xfb\x81\x64\x77" +
	"\x82\xfe\x7e\x8e\x8e\x27\x31\x97\xbb\x87\x57\xcf\xf6\x4d\xff\xa6\xf6\xbd\x5d\x6a\x73\x77\x8e\xbe\xe5\xed\x5e\xcf" +
	"\xe3\xe7\xbb\xa9\xdd\x58\xd4\x8e\xbd\xde\x6c\x77\x47\x5c\x65\xff\x79\xb9\x91\x27\x8b\x1d\x5f\x81\x56\xa2\x3b\xbf" +
	"\x2c\xec\xe7\x97\xa7\x7c\xa7\x63\xd3\xe5\xfd\x2a\x68\xde\xc2\x60\x8c\xb9\xe9\x7d\xaf\x77\xe7\x4b\xdf\x5d\x28\xe8" +
	"\x15\x2b\x83\x4d\x45\x39\xd9\x54\xd5\x97\xe1\x6e\x54\x8c\x22\xcc\x55\xa7\x93\x7b\xbf\x02\xc0\x6d\xc1\xee\xee\x17" +
	"\x6a\xe7\xc8\xa0\x7c\x38\x99\xa3\xef\x2e\x04\x60\xca\x95\x56\x54\x5d\xad\x1c\xdb\xff\x26\xf1\xc9\xf7\x5f\x24\xdd" +
	"\x31\x79\x38\x1d\x5d\xd6\x4d\x02\x55\xf6\xd7\x85\xdd\xbf\x55\x7f\x96\xdf\x83\xf7\x0f\x53\x25\x3c\x24\xb8\xfa\xe6" +
	"\xeb\x4d\xfd\xd3\x14\xe5\xf3\xec\xbb\xf0\xfe\x6e\xab\xc4\x71\xf1\xeb\xda\xfd\x3b\x41\x38\xfa\xfd\x20\xf3\xd9\xd3" +
	"\x65\x77\x08\xf9\x8f\x3f\xfe\xb0\xd0\x30\x8f\x14\xfa\x23\xee\xd7\xf5\x4d\x53\xbc\x89\x05\x3c\x33\x1a\x67\xc0\x8f" +
	"\x3d\xd7\xb0\x3b\xaf\xd0\xee\xc2\xf5\x47\xb6\x44\x38\x2e\x0c\x0c\x0b\x05\xd3\x32\xc1\xa6\x62\xb4\x49\x86\xae\x8e" +
	"\x22\x9b\x59\x82\x5e\xeb\xce\xb8\xf9\xbb\xbd\x57\x4d\x84\x46\xeb\x1e\xce\xcf\x1b\x3e\x63\x9a\x16\x7c\x3d\x75\xae" +
	"\xfb\xb8\xf4\xdd\x9f\x26\xe0\xa2\xf1\x69\x03\xdd\xc8\x5a\xf1\xe9\x5f\xd5\xb0\x62\xee\xff\x17\x1c\xf0\xaa\xba\xe4" +
	"\x9a\xc9\x0d\x93\x4a\x70\x5a\x99\x22\x8c\xd2\xf4\x40\x94\xd6\x5f\x31\x35\xd9\xbd\xcd\xed\x32\x25\x2f\xf7\xac\x26" +
	"\x2c\x5e\x60\x05\xc6\x30\x01\x73\xf1\xcf\xdb\x3d\xfe\x35\xef\x0d\xaf\xbb\x50\xd3\x75\x69\x71\x2e\xaa\xaf\x5e\x53" +
	"\x0d\x0e\x29\x9b\x75\xe3\xa3\x94\x4b\xc8\xe1\x8c\x15\x84\x73\x8e\x9c\x77\xb6\x27\x16\x53\x93\xb0\xf7\x20\x12\xcc" +
	"\x21\xf1\xbb\xdc\x36\x5d\x4c\xed\x7b\x4b\x99\xdc\x85\x48\x7a\x21\xa7\xdd\x95\x04\x37\x6e\xf7\xe2\xc7\xf1\xe5\x2a" +
	"\x6d\x4f\x3a\x7f\xf7\x39\xe9\xec\x5e\x46\xd8\xfa\xde\x83\x6f\x9b\x6d\xb9\x04\x04\x37\x5b\xe7\xf6\xba\xb7\x7a\xde" +
	"\xfd\xf9\xb6\x77\x75\xed\xb9\x15\xbb\xd5\x05\xa3\x8a\x91\x4f\x9f\x3e\x7d\xfb\xf1\xe3\x87\x85\x33\x77\x25\xf8\x8f" +
	"\xfc\x2b\x75\x61\x4e\xe2\x87\xd6\x8c\x17\x11\xf0\xd3\xaf\x91\xa8\xf8\xf0\x72\xeb\xb3\xc2\x3b\xd9\x8d\x34\x38\x9f" +
	"\x03\x77\x61\x98\x4f\xdf\x5c\xb2\x9b\x45\x45\xf7\xed\xdf\x42\xf5\xf9\xdb\x5e\xee\x4c\xdb\xca\xc1\x7b\xf0\x74\x9f" +
	"\x08\x60\xc3\x62\x6e\xf3\xc9\xf7\x7c\xfb\x86\x23\xfe\xb6\x25\xc0\xc8\x4b\xa6\xde\x22\xc3\xf8\x9a\xc2\x86\x76\x27" +
	"\x95\xbf\xd8\x4f\x2a\x4f\xcb\x73\xec\x91\xe2\xbd\xf9\xbc\xc9\x56\x77\xbe\xe6\xf0\x49\xbf\x99\xdd\x3e\xda\xe5\x99" +
	"\xef\xc1\xe7\xaf\x0a\x60\x6c\x81\xb9\x3b\x89\xdc\x9f\xef\xff\x13\x00\x00\xff\xff\xe1\x06\x51\xf3\x15\x8f\x00\x00" +
	"")

func bindataTrainpatstsvBytes() ([]byte, error) {
	return bindataRead(
		_bindataTrainpatstsv,
		"train_pats.tsv",
	)
}



func bindataTrainpatstsv() (*asset, error) {
	bytes, err := bindataTrainpatstsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "train_pats.tsv",
		size: 36629,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1574762607, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataSemanticstsv = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x96\xcf\xb2\xdb\xba\x0d\x87\xd7\xea\x53\x78\x91\x17\x68\xbb\xcb\xae" +
	"\x93\x2c\xba\xea\x2b\x60\x60\x12\x92\x31\x87\x04\x14\x80\xb4\xeb\x3c\xfd\x1d\xca\xe7\xfa\x9a\x94\x9c\xdc\xcd\x99" +
	"\x83\xef\xf7\xf1\x8f\x49\x91\x12\xfc\xf7\xeb\xf4\xe5\x7f\x98\x69\xfa\xf2\x0d\x0b\x2d\xd3\x97\x6f\x2a\xf0\x9f\xb3" +
	"\xff\x03\xbe\x7f\x9d\x32\xfe\x1f\x9c\x7f\x12\x24\x72\x87\x59\xb5\x9c\xa6\x56\x4f\x41\xa5\x17\x5a\x06\x45\xa1\xdc" +
	"\x14\xee\x68\xd1\xdf\x89\x8b\x11\x16\xb2\xb7\xa2\xe8\xe6\x8d\xad\x59\xc0\x2f\xb8\x12\xfc\xf3\xfb\x69\xda\xfe\x3b" +
	"\x0a\xff\xf5\xab\xf0\xdf\xbb\xb0\x8d\xb5\x81\x1e\x07\x53\x77\x70\x0a\x85\x55\xc0\x28\x14\x94\xa5\x26\xb4\x53\x1f" +
	"\xbd\x69\x10\xd8\xc2\xdf\xb7\xb5\x5c\xc8\xde\x98\xa2\x30\x04\x87\xda\x05\x1d\x12\x2d\x3e\xad\x68\xc5\x3b\x8c\x96" +
	"\xfd\x74\xc0\x85\xc2\x07\xa8\x41\xd0\xb4\xcd\xb4\x37\x44\xe1\x01\x7a\x7c\xbb\x70\xa1\xd3\x14\x34\xa9\x3d\xe1\xd9" +
	"\xf4\x26\x23\xdc\xaa\xc7\x4f\x03\x2f\xa6\xb2\x8c\xc6\x15\x8d\x29\xc2\x06\x7d\x0c\x23\xda\xc7\xc8\xda\x5a\x7c\x82" +
	"\x57\x7c\x41\x8b\xd3\x4c\x94\x9e\xc4\x75\x2e\xa7\x1e\x89\xc2\x56\xf7\xde\x8d\xa8\x9c\xa6\x82\x5e\x5e\x1e\x17\xbd" +
	"\x92\x9f\xa6\xac\xdd\x02\xb3\x44\x55\xf3\x29\x69\xc0\x21\x80\x0f\x2e\xe1\x42\x72\xda\x87\x2a\xb0\x98\x56\x89\x87" +
	"\x91\x57\x9b\x31\xd0\x51\xbb\xb6\x6c\x37\x76\x02\xaf\xeb\xaa\x56\x28\x1e\x59\xb5\x6c\x93\x02\x6e\x4f\x5c\xb9\x1f" +
	"\x28\x2d\xd1\x2a\xc5\x8e\xc2\xb9\xcd\x0c\x6e\xaa\xed\x04\xbe\x49\x85\xd0\xc0\x09\x7f\x23\x14\x23\xcc\xef\x7b\xc9" +
	"\x6d\x0e\xc8\xf2\xde\x50\x81\xf9\xf1\xa4\xbe\x11\x58\x60\xad\xe7\xc4\x01\xce\x95\x53\x64\x59\x0e\x36\xe3\xa9\x16" +
	"\x43\xf1\xb6\x70\xbf\xea\x6f\xc6\x50\xd4\x98\x8e\x06\xfd\xdc\x1b\xd0\x19\xce\x1a\x8f\x56\x0f\xcf\x7a\x25\xb8\x21" +
	"\xfb\xd1\x20\xa2\xf0\x64\xfb\x10\x4b\x35\x4c\xa7\xc7\x3f\xaf\x37\x95\x40\xc6\x48\xbb\x20\xf1\x95\xdb\x01\xfa\x60" +
	"\x59\xa2\xe6\x97\x51\x04\x1e\xd9\x2e\x5a\x13\x4a\xd9\xb7\xc8\x98\x73\x1b\x39\x24\xf4\xbf\x8e\xf5\x99\x2d\x8e\xec" +
	"\x52\x33\xca\xc0\x6e\x9c\xe2\x80\x02\x9a\xf0\x55\x8d\xc6\xf6\xed\xb4\x3e\x40\x87\xa3\x92\xc3\x9c\xee\xa7\x09\xfb" +
	"\x2b\x6c\x0b\xfc\xc6\xf9\x38\xb1\x2a\xbb\x40\x14\x30\x94\x91\xce\x94\x59\x58\x68\x5a\x48\x22\xbd\x5c\x1f\x54\x0b" +
	"\xd9\x48\x33\xa6\x9d\xd9\xf6\xa0\xed\x7c\xa6\xd2\xd6\xca\xeb\xd9\x0b\x4a\xa0\x9d\x90\xf8\x47\xe5\xf8\x2b\xe3\x71" +
	"\xb7\xef\xf3\x45\x0b\xcc\xa6\x19\xb6\x7d\xf2\xa3\x2e\x9e\x0a\x0a\x67\x4c\x87\x4e\x7b\x7d\x3d\xe1\x41\xcc\x81\x26" +
	"\x9c\x67\x0a\xe5\xc9\xaa\xac\x89\xd0\xb7\xa7\x63\x88\x22\xca\x42\xa6\xd5\x77\xc9\xaa\xce\x85\xaf\xbb\xce\x84\x16" +
	"\x3c\xe4\x0a\x9f\x64\x08\x82\x6e\x17\x01\xd9\x69\x32\x4d\xf4\x72\x28\x0d\x08\x0b\xcb\x02\xd1\x58\x3e\xb6\xc7\x7d" +
	"\x67\xdc\x08\xed\x38\xd9\xd6\xf9\x80\xa7\x2a\xe1\x02\x91\x65\x3f\x64\x75\x8a\xd0\x9c\x05\x33\x79\x7b\x15\x1a\x85" +
	"\xf6\x69\xc2\x2a\x83\xda\x5e\x82\xbc\xbd\xd1\xd3\x1d\xd0\x5d\x03\x84\x0b\xa7\xf8\x7b\x0d\x63\x4d\x65\xd0\x44\x61" +
	"\xab\x47\x98\x71\x11\x2e\x35\xd2\x94\x71\x99\xf0\xf3\x03\xcc\x33\xa6\xd4\x91\x84\xb6\xf4\x4e\x26\xf4\x6a\x94\x49" +
	"\x4a\xdf\xb6\xae\x64\x6a\x91\x05\x4b\xdf\xa2\x58\xa5\xf6\xa7\x5c\x9e\x68\xe6\xc7\x39\xea\x29\xcb\xac\x96\xf1\x20" +
	"\x11\x85\x07\xe8\xf1\xe7\x69\xf4\xd2\x86\x7c\x4e\x64\xab\x7a\xd6\x3e\x42\x62\xb5\xcf\xae\x39\xd3\x6b\xcf\xc7\x41" +
	"\xdb\x4c\x94\xa5\xdd\x77\x1d\x67\xb9\x6a\xba\x92\xc3\x16\x53\x1f\x16\xca\xab\x1a\xda\xbd\xc7\x2b\x59\x46\x69\x2b" +
	"\xd6\xdb\x9c\x09\xce\x34\xab\x0d\xdd\xcc\xb5\x5d\xc9\xb0\x6a\x21\x29\x8c\xa9\x4f\x8d\x12\x16\x72\xa0\x6b\xeb\xb1" +
	"\xaf\x46\xe7\xf9\x3e\xd8\x81\xd1\xcc\x2a\x74\x1f\xaa\xd1\x59\xd5\x9d\xdc\x5f\xfb\x7b\x41\xa3\x7d\x53\xfb\xe8\x8b" +
	"\x7d\x7f\x37\xb2\xa1\x1a\x1d\xa3\xc0\xab\x8d\x3f\xa3\xa7\xfb\x36\x3f\x2a\x79\xd9\xd5\xa3\xc7\x52\xc8\x56\x32\x57" +
	"\xc1\xf4\x86\x8e\x6d\x1e\xb7\x6c\x5f\xfd\xe9\xfc\xa8\x98\xb8\xdc\x21\xf2\x3c\xb7\xc3\x59\xee\x47\x68\xb4\xd5\x16" +
	"\x14\xfe\x49\xf1\x80\x8c\xee\xd9\xf0\x4a\x76\xdf\xd5\xa3\xe7\x24\xdb\x2d\xca\x2f\x53\x78\x65\xbb\x39\x6c\xbf\xa3" +
	"\xaf\x9a\xf3\x47\x00\x00\x00\xff\xff\xdb\xd6\xb3\xd7\xb1\x0d\x00\x00")

func bindataSemanticstsvBytes() ([]byte, error) {
	return bindataRead(
		_bindataSemanticstsv,
		"semantics.tsv",
	)
}



func bindataSemanticstsv() (*asset, error) {
	bytes, err := bindataSemanticstsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "semantics.tsv",
		size: 3505,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1574762607, 0),
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
	"train_pats.tsv": bindataTrainpatstsv,
	"semantics.tsv":  bindataSemanticstsv,
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
	"semantics.tsv": {Func: bindataSemanticstsv, Children: map[string]*bintree{}},
	"train_pats.tsv": {Func: bindataTrainpatstsv, Children: map[string]*bintree{}},
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