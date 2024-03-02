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
// db/migrations/6_create_table_users_email_view.down.sql
// db/migrations/6_create_table_users_email_view.up.sql
// db/migrations/7_alter_table_users_aggregate_and_email_view.down.sql
// db/migrations/7_alter_table_users_aggregate_and_email_view.up.sql
// db/migrations/8_create_table_users_token_view.down.sql
// db/migrations/8_create_table_users_token_view.up.sql
// db/migrations/9_create_index_users_token_view_verification_token.down.sql
// db/migrations/9_create_index_users_token_view_verification_token.up.sql
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

	info := bindataFileInfo{name: "1_create_table.down.sql", size: 27, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
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

	info := bindataFileInfo{name: "1_create_table.up.sql", size: 36, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
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

	info := bindataFileInfo{name: "2_alter_table.down.sql", size: 26, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
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

	info := bindataFileInfo{name: "2_alter_table.up.sql", size: 35, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
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

	info := bindataFileInfo{name: "3_create_table_users_aggregate.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
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

	info := bindataFileInfo{name: "3_create_table_users_aggregate.up.sql", size: 203, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
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

	info := bindataFileInfo{name: "4_alter_table_users_aggregate.down.sql", size: 5, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
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

	info := bindataFileInfo{name: "4_alter_table_users_aggregate.up.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
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

	info := bindataFileInfo{name: "5_create_table_users_aggregate.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
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

	info := bindataFileInfo{name: "5_create_table_users_aggregate.up.sql", size: 210, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __6_create_table_users_email_viewDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\x8e\x4f\xcd\x4d\xcc\xcc\x89\x2f\xcb\x4c\x2d\xb7\x06\x04\x00\x00\xff\xff\x57\xa2\x1f\x52\x26\x00\x00\x00")

func _6_create_table_users_email_viewDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__6_create_table_users_email_viewDownSql,
		"6_create_table_users_email_view.down.sql",
	)
}

func _6_create_table_users_email_viewDownSql() (*asset, error) {
	bytes, err := _6_create_table_users_email_viewDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "6_create_table_users_email_view.down.sql", size: 38, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __6_create_table_users_email_viewUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcc\xc1\xaa\xc2\x30\x10\x85\xe1\x7d\x9f\xe2\x2c\xef\x05\xdf\xc0\x55\xd5\x51\x83\xb5\x95\x74\x4a\xdb\x55\x08\x66\xc4\x40\x05\x69\xda\xfa\xfa\x42\xba\x51\x11\xb7\xf3\x7f\x73\xd6\x9a\x52\x26\x70\xba\xca\x08\x6a\x8b\xbc\x60\x50\xa3\x4a\x2e\x31\x06\xe9\x83\x91\x9b\xf5\x9d\x99\xbc\x3c\xf0\x97\x00\x88\x67\xe3\x1d\x98\x1a\x8e\x3c\xaf\xb2\x0c\x27\xad\x8e\xa9\x6e\x71\xa0\x76\x11\x59\xfc\x7b\x47\x73\xb8\xda\x60\x26\xe9\xfd\xc5\x8b\x9b\xd7\xa1\x72\xa6\x1d\xe9\x0f\x78\xee\xc5\x0e\xe2\x8c\x1d\xbe\xcd\x8c\x77\xf7\xa3\x3a\xe9\xe4\xa5\x26\xff\xa8\x15\xef\x8b\x8a\xa1\x8b\x5a\x6d\x96\xcf\x00\x00\x00\xff\xff\xd1\x30\x2e\x31\xf7\x00\x00\x00")

func _6_create_table_users_email_viewUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__6_create_table_users_email_viewUpSql,
		"6_create_table_users_email_view.up.sql",
	)
}

func _6_create_table_users_email_viewUpSql() (*asset, error) {
	bytes, err := _6_create_table_users_email_viewUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "6_create_table_users_email_view.up.sql", size: 247, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_alter_table_users_aggregate_and_email_viewDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xcd\x4e\xc3\x30\x10\x84\xef\x7e\x8a\x39\x82\xc4\x1b\xf4\x94\xd2\x05\x2c\xd2\xa4\x72\x37\x6a\x7b\xb2\x2c\xbc\x0d\x96\x42\x8b\x9c\x1f\x5e\x1f\xa5\x29\x2d\xd0\x06\x71\xb4\xe7\xdb\xd9\x9d\x99\x99\x7c\x01\x4e\xa6\x29\x41\x3f\x80\xd6\x7a\xc9\x4b\xb4\xb5\xc4\xda\xba\xb2\x8c\x52\xba\x46\x26\xea\xde\x50\xc2\x74\xe6\xb2\x9c\x47\x58\xdc\x28\x00\x38\xbd\x6d\xf0\x60\x5a\x33\x16\x46\xcf\x13\xb3\xc1\x33\x6d\xee\x0e\x48\xe7\xaa\x56\xac\x77\x8d\x1b\x80\xde\x33\x2b\xd2\xf4\xa8\x4a\xac\xc3\x7e\x67\xc3\x6e\xbb\x47\x56\xcc\xa7\x64\x06\xe1\x25\x4a\x6f\xeb\xfb\x5d\x57\xe6\xda\x77\xff\x43\x56\xb7\x58\x69\x7e\xca\x0b\x86\xc9\x57\x7a\x36\x51\xea\x8f\xc8\xf2\xe6\x42\x65\xbb\x20\x1f\xff\xc8\x7c\x86\x8f\xa1\xfb\xef\x53\xde\xaf\xb3\x2e\x83\x1f\xe6\xae\xdd\xfe\xea\x6a\xdb\x49\x0c\xdb\x20\x7e\x70\x87\xce\x98\x1e\xc9\xfc\x02\x87\x0e\xbc\x75\xcd\x78\x05\x63\xaa\x97\x4a\xbe\xa9\x17\xfd\x7c\x06\x00\x00\xff\xff\xa1\xaa\x43\x50\x11\x02\x00\x00")

func _7_alter_table_users_aggregate_and_email_viewDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__7_alter_table_users_aggregate_and_email_viewDownSql,
		"7_alter_table_users_aggregate_and_email_view.down.sql",
	)
}

func _7_alter_table_users_aggregate_and_email_viewDownSql() (*asset, error) {
	bytes, err := _7_alter_table_users_aggregate_and_email_viewDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "7_alter_table_users_aggregate_and_email_view.down.sql", size: 529, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_alter_table_users_aggregate_and_email_viewUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x90\x41\x4e\xc3\x30\x10\x45\xf7\x3e\xc5\x2c\x41\xe2\x06\x5d\xa5\x74\x00\x8b\x34\xa9\xdc\x89\xda\xae\xac\x11\x9e\x06\x4b\xa1\x45\x8e\x13\xae\x8f\x88\x5b\x55\x2a\x09\x62\xd1\xed\xff\x6f\xbe\xec\xb7\x30\xe5\x0a\x28\x9b\xe7\x08\xfa\x09\x70\xab\xd7\xb4\x86\xae\x95\xd0\x5a\xae\xeb\x20\x35\x47\x99\xa9\x47\x83\x19\xe1\x85\x2b\x4a\x9a\x60\xe1\x4e\x01\xc0\x90\x5a\xef\x80\x70\x4b\xb0\x32\x7a\x99\x99\x1d\xbc\xe2\xee\x61\x68\x7b\x6e\x3a\xb1\x8e\x23\x27\xe0\x67\xae\xa8\xf2\xfc\xd4\x4a\x68\xfd\xf1\x60\xfd\x61\x7f\x84\xa2\x5a\xce\xd1\xa4\xe2\x2d\x08\x47\x71\x96\x23\xe8\x82\xf0\x19\xcd\xd5\x65\xf7\xe9\xae\x80\x94\x3b\x69\xe4\x94\x9f\x0b\x75\x0f\x1b\x4d\x2f\x65\x45\x60\xca\x8d\x5e\xcc\x94\xfa\x43\x85\x7c\xb0\x6f\x6c\xef\xe5\xeb\x1f\x2e\x2e\xf0\x98\x8c\xf3\x8b\x7f\x5b\x19\xee\xc6\x84\xbc\x73\x6b\x7b\x09\x7e\xef\xc5\xa5\xf5\x89\xff\xdf\x42\xd0\x94\x9f\xef\x00\x00\x00\xff\xff\xeb\xfe\xd4\xf8\x29\x02\x00\x00")

func _7_alter_table_users_aggregate_and_email_viewUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__7_alter_table_users_aggregate_and_email_viewUpSql,
		"7_alter_table_users_aggregate_and_email_view.up.sql",
	)
}

func _7_alter_table_users_aggregate_and_email_viewUpSql() (*asset, error) {
	bytes, err := _7_alter_table_users_aggregate_and_email_viewUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "7_alter_table_users_aggregate_and_email_view.up.sql", size: 553, mode: os.FileMode(420), modTime: time.Unix(1706693682, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __8_create_table_users_token_viewDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\x8e\x2f\xc9\xcf\x4e\xcd\x8b\x2f\xcb\x4c\x2d\xb7\xe6\x02\x04\x00\x00\xff\xff\x94\xa1\x68\xbb\x27\x00\x00\x00")

func _8_create_table_users_token_viewDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__8_create_table_users_token_viewDownSql,
		"8_create_table_users_token_view.down.sql",
	)
}

func _8_create_table_users_token_viewDownSql() (*asset, error) {
	bytes, err := _8_create_table_users_token_viewDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "8_create_table_users_token_view.down.sql", size: 39, mode: os.FileMode(420), modTime: time.Unix(1709280262, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __8_create_table_users_token_viewUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcc\x41\xcb\x82\x40\x10\xc6\xf1\xbb\x9f\xe2\x39\xbe\x2f\xf4\x0d\x3a\x59\x4d\xb4\x64\x1a\xeb\x88\x7a\x5a\x16\x77\x82\x25\xd1\xd0\xd5\xbe\x7e\x60\x97\x02\xe9\x3a\xff\xdf\x3c\x7b\x4d\x31\x13\x38\xde\x25\x04\x75\x44\x9a\x31\xa8\x52\x39\xe7\x98\x46\x19\x46\x13\xfa\xbb\x74\x66\xf6\xf2\xc4\x5f\x04\x60\x39\x1b\xef\xc0\x54\xf1\xc2\xd3\x22\x49\x70\xd5\xea\x12\xeb\x1a\x67\xaa\x37\x0b\x9b\x65\xf0\x37\xdf\xd8\xe0\xfb\xee\x3d\xf2\xfd\xb1\xa2\x46\xdb\x86\x35\xd4\x0c\x62\x83\x38\x63\x57\xeb\xf4\x70\x3f\xaa\x93\x56\x3e\x6a\xf4\x8f\x52\xf1\x29\x2b\x18\x3a\x2b\xd5\x61\x1b\xbd\x02\x00\x00\xff\xff\x1e\x6c\xf4\xa2\x01\x01\x00\x00")

func _8_create_table_users_token_viewUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__8_create_table_users_token_viewUpSql,
		"8_create_table_users_token_view.up.sql",
	)
}

func _8_create_table_users_token_viewUpSql() (*asset, error) {
	bytes, err := _8_create_table_users_token_viewUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "8_create_table_users_token_view.up.sql", size: 257, mode: os.FileMode(420), modTime: time.Unix(1709280219, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __9_create_index_users_token_view_verification_tokenDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x4b\x2d\xca\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x2f\xc9\xcf\x4e\xcd\x8b\x4f\x2c\x4e\xb6\xe6\x02\x04\x00\x00\xff\xff\xc0\x06\xc0\xb2\x2d\x00\x00\x00")

func _9_create_index_users_token_view_verification_tokenDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__9_create_index_users_token_view_verification_tokenDownSql,
		"9_create_index_users_token_view_verification_token.down.sql",
	)
}

func _9_create_index_users_token_view_verification_tokenDownSql() (*asset, error) {
	bytes, err := _9_create_index_users_token_view_verification_tokenDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "9_create_index_users_token_view_verification_token.down.sql", size: 45, mode: os.FileMode(420), modTime: time.Unix(1709281532, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __9_create_index_users_token_view_verification_tokenUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x4b\x2d\xca\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x2f\xc9\xcf\x4e\xcd\x8b\x4f\x2c\x4e\x56\xf0\xf7\x53\x28\x2d\x4e\x2d\x2a\x86\x0a\x95\x65\xa6\x96\x2b\x68\x70\x29\x28\x28\x60\xd1\xa0\xe0\x18\xec\xac\x69\xcd\x05\x08\x00\x00\xff\xff\xc2\x76\x55\xef\x65\x00\x00\x00")

func _9_create_index_users_token_view_verification_tokenUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__9_create_index_users_token_view_verification_tokenUpSql,
		"9_create_index_users_token_view_verification_token.up.sql",
	)
}

func _9_create_index_users_token_view_verification_tokenUpSql() (*asset, error) {
	bytes, err := _9_create_index_users_token_view_verification_tokenUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "9_create_index_users_token_view_verification_token.up.sql", size: 101, mode: os.FileMode(420), modTime: time.Unix(1709281172, 0)}
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
	"1_create_table.down.sql":                                     _1_create_tableDownSql,
	"1_create_table.up.sql":                                       _1_create_tableUpSql,
	"2_alter_table.down.sql":                                      _2_alter_tableDownSql,
	"2_alter_table.up.sql":                                        _2_alter_tableUpSql,
	"3_create_table_users_aggregate.down.sql":                     _3_create_table_users_aggregateDownSql,
	"3_create_table_users_aggregate.up.sql":                       _3_create_table_users_aggregateUpSql,
	"4_alter_table_users_aggregate.down.sql":                      _4_alter_table_users_aggregateDownSql,
	"4_alter_table_users_aggregate.up.sql":                        _4_alter_table_users_aggregateUpSql,
	"5_create_table_users_aggregate.down.sql":                     _5_create_table_users_aggregateDownSql,
	"5_create_table_users_aggregate.up.sql":                       _5_create_table_users_aggregateUpSql,
	"6_create_table_users_email_view.down.sql":                    _6_create_table_users_email_viewDownSql,
	"6_create_table_users_email_view.up.sql":                      _6_create_table_users_email_viewUpSql,
	"7_alter_table_users_aggregate_and_email_view.down.sql":       _7_alter_table_users_aggregate_and_email_viewDownSql,
	"7_alter_table_users_aggregate_and_email_view.up.sql":         _7_alter_table_users_aggregate_and_email_viewUpSql,
	"8_create_table_users_token_view.down.sql":                    _8_create_table_users_token_viewDownSql,
	"8_create_table_users_token_view.up.sql":                      _8_create_table_users_token_viewUpSql,
	"9_create_index_users_token_view_verification_token.down.sql": _9_create_index_users_token_view_verification_tokenDownSql,
	"9_create_index_users_token_view_verification_token.up.sql":   _9_create_index_users_token_view_verification_tokenUpSql,
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
	"1_create_table.down.sql":                                     &bintree{_1_create_tableDownSql, map[string]*bintree{}},
	"1_create_table.up.sql":                                       &bintree{_1_create_tableUpSql, map[string]*bintree{}},
	"2_alter_table.down.sql":                                      &bintree{_2_alter_tableDownSql, map[string]*bintree{}},
	"2_alter_table.up.sql":                                        &bintree{_2_alter_tableUpSql, map[string]*bintree{}},
	"3_create_table_users_aggregate.down.sql":                     &bintree{_3_create_table_users_aggregateDownSql, map[string]*bintree{}},
	"3_create_table_users_aggregate.up.sql":                       &bintree{_3_create_table_users_aggregateUpSql, map[string]*bintree{}},
	"4_alter_table_users_aggregate.down.sql":                      &bintree{_4_alter_table_users_aggregateDownSql, map[string]*bintree{}},
	"4_alter_table_users_aggregate.up.sql":                        &bintree{_4_alter_table_users_aggregateUpSql, map[string]*bintree{}},
	"5_create_table_users_aggregate.down.sql":                     &bintree{_5_create_table_users_aggregateDownSql, map[string]*bintree{}},
	"5_create_table_users_aggregate.up.sql":                       &bintree{_5_create_table_users_aggregateUpSql, map[string]*bintree{}},
	"6_create_table_users_email_view.down.sql":                    &bintree{_6_create_table_users_email_viewDownSql, map[string]*bintree{}},
	"6_create_table_users_email_view.up.sql":                      &bintree{_6_create_table_users_email_viewUpSql, map[string]*bintree{}},
	"7_alter_table_users_aggregate_and_email_view.down.sql":       &bintree{_7_alter_table_users_aggregate_and_email_viewDownSql, map[string]*bintree{}},
	"7_alter_table_users_aggregate_and_email_view.up.sql":         &bintree{_7_alter_table_users_aggregate_and_email_viewUpSql, map[string]*bintree{}},
	"8_create_table_users_token_view.down.sql":                    &bintree{_8_create_table_users_token_viewDownSql, map[string]*bintree{}},
	"8_create_table_users_token_view.up.sql":                      &bintree{_8_create_table_users_token_viewUpSql, map[string]*bintree{}},
	"9_create_index_users_token_view_verification_token.down.sql": &bintree{_9_create_index_users_token_view_verification_tokenDownSql, map[string]*bintree{}},
	"9_create_index_users_token_view_verification_token.up.sql":   &bintree{_9_create_index_users_token_view_verification_tokenUpSql, map[string]*bintree{}},
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
