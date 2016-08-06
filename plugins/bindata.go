// Code generated by go-bindata.
// sources:
// plugins/plugins.toml
// DO NOT EDIT!

package plugins

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

var _pluginsPluginsToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x96\x4f\x6f\xea\x38\x14\xc5\xf7\x7c\x8a\xab\x74\x33\x33\x7a\x94\x79\x1d\x75\x76\xb3\xc8\x50\xfa\x6f\x68\x41\x09\x45\xaa\xaa\xaa\xbd\x24\x4e\xf0\xd4\xb1\x23\xdb\xa1\xe5\xdb\x3f\xdb\xfc\x6b\x21\x08\x07\xb5\x1b\x9a\xeb\xe3\x7b\x7e\x3e\xb1\x0d\x27\xd0\x15\xe5\x5c\xd2\x7c\xaa\xe1\xb7\xe4\x77\x38\xfb\xf3\xe7\x5f\xd0\xb6\x1f\xe7\x30\x61\x98\xbc\x69\x51\xc2\xad\x50\xd3\x0a\xe1\x0e\x29\x27\x3f\x20\x64\x0c\x22\x3b\x41\x41\x44\x14\x91\x33\x92\x9e\xb6\x4e\x20\x26\x04\xfa\x37\xdd\xde\x7d\xdc\x83\x4c\x48\x60\x34\x21\x5c\x11\xa0\xdc\x3c\x15\xa8\xa9\xe0\xa7\xad\xd6\xc9\xf7\xfc\x19\xbf\x61\xff\xe1\xea\xe6\xde\xe0\xf3\x8c\xe6\x95\x74\x06\xd0\xbc\xcf\x37\xf1\xb4\x9e\x9e\x4a\x56\xe5\x94\x3f\x3f\xb7\x00\x38\x16\x04\xfe\x81\x80\x2b\xc9\x02\xf3\x4c\x38\x4e\x18\x49\x4d\x29\x43\xa6\x88\xa9\xa4\x44\x25\x92\x96\x0e\xda\x08\xef\xe3\xa8\x0f\x17\xa8\x71\x82\x26\xb2\x6b\x54\x53\x93\x27\xca\x64\x6a\x67\x27\xa8\x49\x2e\xe4\xdc\x0a\x29\xd7\xc4\xb5\xa4\x05\xe6\xce\xa3\x40\x9b\x74\x67\x65\x55\xd0\x85\xf5\xd4\xf4\xb0\xcf\xf6\x53\xcf\x4b\xa2\x4c\xf1\xc9\xa8\xd3\xf3\x00\x9e\xf7\xe0\xce\xa8\xac\x94\x16\x1a\xb7\xa0\xb5\xac\x6a\x98\xc7\x56\x3d\xb2\x6a\xb3\x61\x32\xca\x8c\x85\x4a\x90\x03\xf2\xd4\xb9\x02\x13\xe2\xad\x2a\xbd\x57\xf0\xd5\xbd\x2a\x99\xc0\x75\x62\x60\xf7\x97\x5e\x15\xb5\x80\x57\xcb\xf4\x0a\x34\x83\xb9\xa8\xe0\x1d\xb9\xb6\xd5\xe5\xb8\xc2\xa2\x64\xc4\x16\x36\x3d\x4f\x13\x51\x98\xb6\x58\xd2\x37\xe2\x38\x1c\x57\x61\x1d\x82\x0d\xa8\x47\x7a\x3f\x20\x50\x53\xfc\xb9\xfc\x3c\x3b\xff\xdb\xe6\x69\xd3\x9a\x59\x4d\x70\x17\xda\x43\xf0\x32\x1e\xbd\x84\xc3\x9b\x60\x5f\xd2\x66\x66\x2a\xde\xdb\xee\xf4\x48\x9f\xb0\x63\x37\x21\x76\x7a\x13\xf7\x31\xf9\xee\x78\x36\x58\xec\xde\x2d\xa3\x09\x16\xed\x64\x5e\xc8\xca\x6b\x9f\x8f\x8c\xbc\x6b\xd5\x47\xae\xe1\xab\x5d\xb3\x05\xec\xbc\xa3\xd1\xf5\xcb\x43\xdc\x8b\xac\x66\x53\xf9\xaf\xf7\xb8\xf7\xad\xd9\x4d\x6e\xaf\x33\x9f\x17\xa6\x54\x4a\x48\xd9\x19\x45\xf4\xa2\x43\x3e\x68\xa6\x85\x60\xdb\xcb\x2c\x88\xc6\xd4\x1c\xfa\xba\x95\x7e\xf6\x5a\xad\xf3\x8f\x60\x0f\xd8\x1c\x25\xfa\x40\x3d\x86\x51\x08\xb1\x39\xa4\xdb\x24\x38\xab\x63\x58\xb5\x3d\xe8\x8f\x33\x54\xda\x07\x20\xb4\x42\x08\xb9\xa6\xee\xf6\xf0\xc3\x58\x77\xf7\xe0\xc8\xbd\x28\xc6\x57\xcd\x19\x72\x2f\x82\x09\xd5\x29\xc9\x08\x4f\xfd\x8e\xf5\xbf\x1b\x79\x53\xa2\x2d\xa7\x83\x64\x09\xc3\x62\xd1\xe9\x10\x54\xd7\x28\xc3\xb1\x1f\xc5\xa6\xeb\x61\x00\x51\x88\xd4\xeb\xec\x74\x9d\xb2\x69\x20\x9b\xfe\x07\x51\xb2\x52\x0a\xaf\xfd\x7a\xd9\x1e\x46\x83\x51\x53\x92\x75\xfb\x83\x20\x25\xf1\xba\x36\x87\x3d\x73\x5f\xda\x1b\xc4\x7e\xa7\x69\x49\xad\x5d\x29\xa4\xb6\xd3\x80\x7c\x90\xa4\x72\xff\xee\xe0\x99\xa1\x3a\xbe\x85\xeb\x0a\x0e\xcb\xd2\x54\xdd\xef\xa7\xce\x47\x3b\x15\xca\x36\xdc\x9b\x1c\x13\x4a\x79\x25\x47\x25\xe9\xcd\xcd\x4f\x42\x9c\x28\x18\x4c\xb2\x4a\x59\xb0\x14\x62\x83\xcf\x73\x88\x05\x5b\x7e\x0b\x79\xf0\xae\x4d\x8f\x42\x16\x59\x66\x9a\x78\xe5\x3c\x70\xd2\xdd\xac\x07\xfd\x5e\x27\x1a\x5d\x42\x2a\x92\xaa\x20\x5c\xef\x04\xbd\x1a\xa8\xa3\xdf\xf8\x1f\xde\x0e\x69\xe6\xb7\x1f\x2e\x2e\x77\x21\x6d\xf1\x28\xc0\xa5\x6b\x5d\xb8\x6e\xa8\x9e\xf5\x7f\x7b\x2d\x3b\x28\x2f\xe4\xdb\xb5\x7c\x97\xfc\x36\x86\xc5\x50\x23\xec\xaf\x00\x75\xf4\x9f\x15\xbf\x02\x00\x00\xff\xff\xd4\x69\x9a\x4c\xe9\x0c\x00\x00")

func pluginsPluginsTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsPluginsToml,
		"plugins/plugins.toml",
	)
}

func pluginsPluginsToml() (*asset, error) {
	bytes, err := pluginsPluginsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/plugins.toml", size: 3305, mode: os.FileMode(420), modTime: time.Unix(1470523620, 0)}
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
	"plugins/plugins.toml": pluginsPluginsToml,
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
	"plugins": &bintree{nil, map[string]*bintree{
		"plugins.toml": &bintree{pluginsPluginsToml, map[string]*bintree{}},
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

