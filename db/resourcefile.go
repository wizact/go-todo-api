// Code generated for package db by go-bindata DO NOT EDIT. (@generated)
// sources:
// db/migrations/1_create_table.down.sql
// db/migrations/1_create_table.up.sql
// db/migrations/2_alter_table.down.sql
// db/migrations/2_alter_table.up.sql
// db/migrations/3_create_table_users_aggregate.down.sql
// db/migrations/3_create_table_users_aggregate.up.sql
// db/migrations/4_alter_table_users_aggregate.down.sql
// db/migrations/4_alter_table_users_aggregate.up.sql
// db/migrations/5_create_table_users_aggregate.down.sql
// db/migrations/5_create_table_users_aggregate.up.sql
package db

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

var __1_create_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x48\x2d\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\xa3\x51\x3a\x6a\x1b\x00\x00\x00")

func _1_create_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_create_tableDownSql,
		"1_create_table.down.sql",
	)
}

func _1_create_tableDownSql() (*asset, error) {
	bytes, err := _1_create_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_create_table.down.sql", size: 27, mode: os.FileMode(420), modTime: time.Unix(1698133702, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_create_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x28\x48\x2d\x29\x56\xd0\xe0\x52\x50\xc8\x4b\xcc\x4d\x55\x28\x2e\x29\xca\xcc\x4b\xe7\xd2\xb4\x06\x04\x00\x00\xff\xff\x5a\x46\x3d\xd6\x24\x00\x00\x00")

func _1_create_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_create_tableUpSql,
		"1_create_table.up.sql",
	)
}

func _1_create_tableUpSql() (*asset, error) {
	bytes, err := _1_create_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_create_table.up.sql", size: 36, mode: os.FileMode(420), modTime: time.Unix(1698133705, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_alter_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x48\x2d\x29\xb6\x06\x04\x00\x00\xff\xff\x3d\x21\x1b\x85\x1a\x00\x00\x00")

func _2_alter_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_alter_tableDownSql,
		"2_alter_table.down.sql",
	)
}

func _2_alter_tableDownSql() (*asset, error) {
	bytes, err := _2_alter_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_alter_table.down.sql", size: 26, mode: os.FileMode(420), modTime: time.Unix(1698133724, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_alter_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x48\x2d\x29\x56\x70\x74\x71\x51\x28\x28\x4a\x4d\x49\x2c\xc9\x2f\x52\x48\xca\xcf\xcf\xb1\x06\x04\x00\x00\xff\xff\x3a\xd3\x70\xb1\x23\x00\x00\x00")

func _2_alter_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_alter_tableUpSql,
		"2_alter_table.up.sql",
	)
}

func _2_alter_tableUpSql() (*asset, error) {
	bytes, err := _2_alter_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_alter_table.up.sql", size: 35, mode: os.FileMode(420), modTime: time.Unix(1698133738, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_create_table_users_aggregateDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\x8e\x4f\x4c\x4f\x2f\x4a\x4d\x4f\x2c\x49\xb5\x06\x04\x00\x00\xff\xff\xc8\x0a\x5d\x57\x25\x00\x00\x00")

func _3_create_table_users_aggregateDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_create_table_users_aggregateDownSql,
		"3_create_table_users_aggregate.down.sql",
	)
}

func _3_create_table_users_aggregateDownSql() (*asset, error) {
	bytes, err := _3_create_table_users_aggregateDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_create_table_users_aggregate.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1704933804, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_create_table_users_aggregateUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8d\xcf\xca\x82\x40\x14\x47\xf7\x3e\xc5\x6f\xf9\x7d\xd0\x1b\xb4\xd2\xba\xd1\x90\x7f\x62\xbc\xa2\xae\x86\xa1\xb9\x89\x10\x1a\xa3\xf6\xfc\x31\x19\x41\xd0\xf2\x72\xce\x3d\xbf\x9d\xa6\x98\x09\x1c\x27\x29\x41\x1d\x90\x17\x0c\x6a\x54\xc9\x25\x96\x49\xfc\x64\x6c\xd7\x79\xe9\xec\x2c\xf8\x8b\x00\xe0\x73\x9b\xde\x81\xa9\x61\x9c\xb5\xca\x62\xdd\xe2\x44\xed\xe6\xa5\x3c\xec\x6d\x11\xe3\xec\x6c\x57\x21\x34\xf3\x2a\x4d\xdf\x54\xfc\xd4\x8f\x83\xe9\x87\xeb\x88\xbc\xca\x12\xd2\x2b\xb8\x78\x09\x59\x17\xb6\x7e\xfc\x2d\x77\xf7\x85\xa3\x7f\xd4\x8a\x8f\x45\xc5\xd0\x45\xad\xf6\xdb\x67\x00\x00\x00\xff\xff\x31\x1c\x0d\xb5\xcb\x00\x00\x00")

func _3_create_table_users_aggregateUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_create_table_users_aggregateUpSql,
		"3_create_table_users_aggregate.up.sql",
	)
}

func _3_create_table_users_aggregateUpSql() (*asset, error) {
	bytes, err := _3_create_table_users_aggregateUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_create_table_users_aggregate.up.sql", size: 203, mode: os.FileMode(420), modTime: time.Unix(1704935096, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_alter_table_users_aggregateDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xf0\x73\x04\x04\x00\x00\xff\xff\x26\x3e\x47\xfa\x05\x00\x00\x00")

func _4_alter_table_users_aggregateDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_alter_table_users_aggregateDownSql,
		"4_alter_table_users_aggregate.down.sql",
	)
}

func _4_alter_table_users_aggregateDownSql() (*asset, error) {
	bytes, err := _4_alter_table_users_aggregateDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_alter_table_users_aggregate.down.sql", size: 5, mode: os.FileMode(420), modTime: time.Unix(1704946160, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_alter_table_users_aggregateUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\x8e\x4f\x4c\x4f\x2f\x4a\x4d\x4f\x2c\x49\xb5\x06\x04\x00\x00\xff\xff\xc8\x0a\x5d\x57\x25\x00\x00\x00")

func _4_alter_table_users_aggregateUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_alter_table_users_aggregateUpSql,
		"4_alter_table_users_aggregate.up.sql",
	)
}

func _4_alter_table_users_aggregateUpSql() (*asset, error) {
	bytes, err := _4_alter_table_users_aggregateUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_alter_table_users_aggregate.up.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1704946545, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_create_table_users_aggregateDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\x8e\x4f\x4c\x4f\x2f\x4a\x4d\x4f\x2c\x49\xb5\x06\x04\x00\x00\xff\xff\xc8\x0a\x5d\x57\x25\x00\x00\x00")

func _5_create_table_users_aggregateDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__5_create_table_users_aggregateDownSql,
		"5_create_table_users_aggregate.down.sql",
	)
}

func _5_create_table_users_aggregateDownSql() (*asset, error) {
	bytes, err := _5_create_table_users_aggregateDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "5_create_table_users_aggregate.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1704946524, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_create_table_users_aggregateUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xca\xb1\x0a\xc2\x30\x10\x87\xf1\xbd\x4f\xf1\x1f\x15\x7c\x03\xa7\xaa\x27\x06\x6b\x2b\xe9\x95\xb6\x53\x38\xcc\x51\x84\x0e\x92\x26\x3e\xbf\x10\x17\x05\x71\xfd\x7e\xdf\xde\x52\xc9\x04\x2e\x77\x15\xc1\x1c\x51\x37\x0c\x1a\x4c\xcb\x2d\xd2\xa2\x61\x71\x32\x4d\x41\x27\x89\x8a\x55\x01\x20\x57\x77\xf7\x60\x1a\x38\xdf\x75\x57\x55\xb8\x5a\x73\x29\xed\x88\x33\x8d\x9b\xbc\x3d\x65\x4e\xea\xbc\x44\xf9\x3e\xdf\x7a\x0b\x2a\x51\xbd\x93\xf8\x4b\xd3\xc3\xff\x51\xaf\xb3\x7e\x68\xb1\x46\x6f\xf8\xd4\x74\x0c\xdb\xf4\xe6\xb0\x7d\x05\x00\x00\xff\xff\xa4\x4e\x0c\x33\xd2\x00\x00\x00")

func _5_create_table_users_aggregateUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__5_create_table_users_aggregateUpSql,
		"5_create_table_users_aggregate.up.sql",
	)
}

func _5_create_table_users_aggregateUpSql() (*asset, error) {
	bytes, err := _5_create_table_users_aggregateUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "5_create_table_users_aggregate.up.sql", size: 210, mode: os.FileMode(420), modTime: time.Unix(1704946962, 0)}
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
	"1_create_table.down.sql":                 _1_create_tableDownSql,
	"1_create_table.up.sql":                   _1_create_tableUpSql,
	"2_alter_table.down.sql":                  _2_alter_tableDownSql,
	"2_alter_table.up.sql":                    _2_alter_tableUpSql,
	"3_create_table_users_aggregate.down.sql": _3_create_table_users_aggregateDownSql,
	"3_create_table_users_aggregate.up.sql":   _3_create_table_users_aggregateUpSql,
	"4_alter_table_users_aggregate.down.sql":  _4_alter_table_users_aggregateDownSql,
	"4_alter_table_users_aggregate.up.sql":    _4_alter_table_users_aggregateUpSql,
	"5_create_table_users_aggregate.down.sql": _5_create_table_users_aggregateDownSql,
	"5_create_table_users_aggregate.up.sql":   _5_create_table_users_aggregateUpSql,
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
	"1_create_table.down.sql":                 &bintree{_1_create_tableDownSql, map[string]*bintree{}},
	"1_create_table.up.sql":                   &bintree{_1_create_tableUpSql, map[string]*bintree{}},
	"2_alter_table.down.sql":                  &bintree{_2_alter_tableDownSql, map[string]*bintree{}},
	"2_alter_table.up.sql":                    &bintree{_2_alter_tableUpSql, map[string]*bintree{}},
	"3_create_table_users_aggregate.down.sql": &bintree{_3_create_table_users_aggregateDownSql, map[string]*bintree{}},
	"3_create_table_users_aggregate.up.sql":   &bintree{_3_create_table_users_aggregateUpSql, map[string]*bintree{}},
	"4_alter_table_users_aggregate.down.sql":  &bintree{_4_alter_table_users_aggregateDownSql, map[string]*bintree{}},
	"4_alter_table_users_aggregate.up.sql":    &bintree{_4_alter_table_users_aggregateUpSql, map[string]*bintree{}},
	"5_create_table_users_aggregate.down.sql": &bintree{_5_create_table_users_aggregateDownSql, map[string]*bintree{}},
	"5_create_table_users_aggregate.up.sql":   &bintree{_5_create_table_users_aggregateUpSql, map[string]*bintree{}},
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
