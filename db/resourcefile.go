// Code generated for package db by go-bindata DO NOT EDIT. (@generated)
// sources:
// db/migrations/10_create_table_user_registration_verifications.down.sql
// db/migrations/10_create_table_user_registration_verifications.up.sql
// db/migrations/11_drop_users_token_view.down.sql
// db/migrations/11_drop_users_token_view.up.sql
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

var __10_create_table_user_registration_verificationsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x36\x00\xc9\xff\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x75\x73\x65\x72\x5f\x72\x65\x67\x69\x73\x74\x72\x61\x74\x69\x6f\x6e\x5f\x76\x65\x72\x69\x66\x69\x63\x61\x74\x69\x6f\x6e\x73\x3b\x0a\x03\x00\x3c\x6b\x21\x19\x36\x00\x00\x00")

func _10_create_table_user_registration_verificationsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__10_create_table_user_registration_verificationsDownSql,
		"10_create_table_user_registration_verifications.down.sql",
	)
}

func _10_create_table_user_registration_verificationsDownSql() (*asset, error) {
	bytes, err := _10_create_table_user_registration_verificationsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10_create_table_user_registration_verifications.down.sql", size: 54, mode: os.FileMode(420), modTime: time.Unix(1790534939, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __10_create_table_user_registration_verificationsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcb\xc1\xaa\x82\x40\x14\x87\xf1\xbd\x4f\xf1\x5f\xde\x0b\xbd\x41\x2b\xab\x53\x0d\x99\xc6\x78\x44\x5d\x0d\xe2\x9c\x64\x36\x25\x33\x63\xf4\xf8\x81\xae\x0a\x5c\x7e\xf0\xfb\xf6\x9a\x52\x26\x70\xba\xcb\x08\xea\x88\xbc\x60\x50\xa3\x4a\x2e\x31\x05\xf1\xc6\xcb\xe0\x42\xf4\x5d\x74\xcf\x87\x79\x89\x77\x77\xd7\xcf\x11\xf0\x97\x00\x58\x94\xb3\x60\x6a\x78\xbe\xf3\x2a\xcb\x70\xd3\xea\x9a\xea\x16\x17\x6a\x37\x33\x0b\xd2\x7b\x89\xc6\xba\x41\x42\xfc\xc6\x0b\x90\xf7\xe8\xbc\x04\xd3\x45\xa8\x9c\xe9\x44\xfa\x07\xf4\x5e\xba\x28\x76\x1d\x4c\xa3\x5d\x03\xc9\x3f\x6a\xc5\xe7\xa2\x62\xe8\xa2\x56\x87\x6d\xf2\x19\x00\xd5\xb7\x28\xc9\xf8\x00\x00\x00")

func _10_create_table_user_registration_verificationsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__10_create_table_user_registration_verificationsUpSql,
		"10_create_table_user_registration_verifications.up.sql",
	)
}

func _10_create_table_user_registration_verificationsUpSql() (*asset, error) {
	bytes, err := _10_create_table_user_registration_verificationsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10_create_table_user_registration_verifications.up.sql", size: 248, mode: os.FileMode(420), modTime: time.Unix(1790535070, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __11_drop_users_token_viewDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xc1\x4a\xc4\x30\x14\x45\xf7\xf9\x8a\xbb\x74\xc0\x3f\x98\x55\x9c\x89\x18\xac\x89\xa4\x6f\x98\x76\x15\x42\xf3\x84\x60\x69\xa5\x49\xeb\xef\x0b\xd5\x85\xda\xea\x36\xe7\xdc\xc0\x3b\x27\xa7\x24\x29\x90\xbc\xab\x14\xf4\x3d\x8c\x25\xa8\x46\xd7\x54\x63\xce\x3c\x65\x5f\xc6\x57\x1e\xfc\x92\xf8\x1d\x37\x02\xc0\xfa\xec\x53\x04\xa9\x86\x56\xdd\x5c\xaa\x0a\xcf\x4e\x3f\x49\xd7\xe2\x51\xb5\xb7\xab\xb6\xf0\x94\x5e\x52\x17\x4a\x1a\x87\xcf\x4f\x7e\x2e\x76\xac\x1c\xfa\xb2\x27\x75\x13\x87\xc2\xd1\x87\x5d\x3a\xbf\xc5\x7f\x68\xe4\x9e\xbf\x51\x71\xc0\x55\xd3\x83\xbd\x10\x9c\xbd\xea\xf3\x51\x88\xaf\x02\xda\x9c\x55\xf3\xab\xc0\xf6\x06\x1f\x72\x07\x6b\xfe\x6a\xb3\x1d\x40\xd6\xa7\xc3\x51\x7c\x0c\x00\xd0\x58\xee\xb0\x67\x01\x00\x00")

func _11_drop_users_token_viewDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__11_drop_users_token_viewDownSql,
		"11_drop_users_token_view.down.sql",
	)
}

func _11_drop_users_token_viewDownSql() (*asset, error) {
	bytes, err := _11_drop_users_token_viewDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "11_drop_users_token_view.down.sql", size: 359, mode: os.FileMode(420), modTime: time.Unix(1791104997, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __11_drop_users_token_viewUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x54\x00\xab\xff\x44\x52\x4f\x50\x20\x49\x4e\x44\x45\x58\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x76\x65\x72\x69\x66\x69\x63\x61\x74\x69\x6f\x6e\x5f\x74\x6f\x6b\x65\x6e\x5f\x61\x73\x63\x3b\x0a\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x75\x73\x65\x72\x73\x5f\x74\x6f\x6b\x65\x6e\x5f\x76\x69\x65\x77\x3b\x0a\x03\x00\x06\xfa\xcb\x32\x54\x00\x00\x00")

func _11_drop_users_token_viewUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__11_drop_users_token_viewUpSql,
		"11_drop_users_token_view.up.sql",
	)
}

func _11_drop_users_token_viewUpSql() (*asset, error) {
	bytes, err := _11_drop_users_token_viewUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "11_drop_users_token_view.up.sql", size: 84, mode: os.FileMode(420), modTime: time.Unix(1791104997, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_create_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x1b\x00\xe4\xff\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x70\x65\x74\x73\x3b\x0a\x03\x00\xa3\x51\x3a\x6a\x1b\x00\x00\x00")

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

	info := bindataFileInfo{name: "1_create_table.down.sql", size: 27, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_create_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x24\x00\xdb\xff\x43\x52\x45\x41\x54\x45\x20\x54\x41\x42\x4c\x45\x20\x70\x65\x74\x73\x20\x28\x0a\x20\x20\x6e\x61\x6d\x65\x20\x73\x74\x72\x69\x6e\x67\x0a\x29\x3b\x03\x00\x5a\x46\x3d\xd6\x24\x00\x00\x00")

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

	info := bindataFileInfo{name: "1_create_table.up.sql", size: 36, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_alter_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x1a\x00\xe5\xff\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x70\x65\x74\x73\x3b\x03\x00\x3d\x21\x1b\x85\x1a\x00\x00\x00")

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

	info := bindataFileInfo{name: "2_alter_table.down.sql", size: 26, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_alter_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x23\x00\xdc\xff\x41\x4c\x54\x45\x52\x20\x54\x41\x42\x4c\x45\x20\x70\x65\x74\x73\x20\x41\x44\x44\x20\x70\x72\x65\x64\x61\x74\x6f\x72\x20\x62\x6f\x6f\x6c\x3b\x03\x00\x3a\xd3\x70\xb1\x23\x00\x00\x00")

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

	info := bindataFileInfo{name: "2_alter_table.up.sql", size: 35, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_create_table_users_aggregateDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x25\x00\xda\xff\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x75\x73\x65\x72\x73\x5f\x61\x67\x67\x72\x65\x67\x61\x74\x65\x3b\x03\x00\xc8\x0a\x5d\x57\x25\x00\x00\x00")

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

	info := bindataFileInfo{name: "3_create_table_users_aggregate.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_create_table_users_aggregateUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8d\xcf\xca\x82\x40\x14\x47\xf7\x3e\xc5\x6f\xf9\x7d\xd0\x1b\xb4\xd2\xba\xd1\x90\x7f\x62\xbc\xa2\xae\x86\xa1\xb9\x89\x10\x1a\xa3\xf6\xfc\x31\x19\x41\xd0\xf2\x72\xce\x3d\xbf\x9d\xa6\x98\x09\x1c\x27\x29\x41\x1d\x90\x17\x0c\x6a\x54\xc9\x25\x96\x49\xfc\x64\x6c\xd7\x79\xe9\xec\x2c\xf8\x8b\x00\xe0\x73\x9b\xde\x81\xa9\x61\x9c\xb5\xca\x62\xdd\xe2\x44\xed\xe6\xa5\x3c\xec\x6d\x11\xe3\xec\x6c\x57\x21\x34\xf3\x2a\x4d\xdf\x54\xfc\xd4\x8f\x83\xe9\x87\xeb\x88\xbc\xca\x12\xd2\x2b\xb8\x78\x09\x59\x17\xb6\x7e\xfc\x2d\x77\xf7\x85\xa3\x7f\xd4\x8a\x8f\x45\xc5\xd0\x45\xad\xf6\xdb\xe7\x00\x31\x1c\x0d\xb5\xcb\x00\x00\x00")

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

	info := bindataFileInfo{name: "3_create_table_users_aggregate.up.sql", size: 203, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_alter_table_users_aggregateDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x05\x00\xfa\xff\x2d\x2d\x20\x4e\x41\x03\x00\x26\x3e\x47\xfa\x05\x00\x00\x00")

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

	info := bindataFileInfo{name: "4_alter_table_users_aggregate.down.sql", size: 5, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_alter_table_users_aggregateUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x25\x00\xda\xff\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x75\x73\x65\x72\x73\x5f\x61\x67\x67\x72\x65\x67\x61\x74\x65\x3b\x03\x00\xc8\x0a\x5d\x57\x25\x00\x00\x00")

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

	info := bindataFileInfo{name: "4_alter_table_users_aggregate.up.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_create_table_users_aggregateDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x25\x00\xda\xff\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x75\x73\x65\x72\x73\x5f\x61\x67\x67\x72\x65\x67\x61\x74\x65\x3b\x03\x00\xc8\x0a\x5d\x57\x25\x00\x00\x00")

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

	info := bindataFileInfo{name: "5_create_table_users_aggregate.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_create_table_users_aggregateUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xca\xb1\x0a\xc2\x30\x10\x87\xf1\xbd\x4f\xf1\x1f\x15\x7c\x03\xa7\xaa\x27\x06\x6b\x2b\xe9\x95\xb6\x53\x38\xcc\x51\x84\x0e\x92\x26\x3e\xbf\x10\x17\x05\x71\xfd\x7e\xdf\xde\x52\xc9\x04\x2e\x77\x15\xc1\x1c\x51\x37\x0c\x1a\x4c\xcb\x2d\xd2\xa2\x61\x71\x32\x4d\x41\x27\x89\x8a\x55\x01\x20\x57\x77\xf7\x60\x1a\x38\xdf\x75\x57\x55\xb8\x5a\x73\x29\xed\x88\x33\x8d\x9b\xbc\x3d\x65\x4e\xea\xbc\x44\xf9\x3e\xdf\x7a\x0b\x2a\x51\xbd\x93\xf8\x4b\xd3\xc3\xff\x51\xaf\xb3\x7e\x68\xb1\x46\x6f\xf8\xd4\x74\x0c\xdb\xf4\xe6\xb0\x7d\x0d\x00\xa4\x4e\x0c\x33\xd2\x00\x00\x00")

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

	info := bindataFileInfo{name: "5_create_table_users_aggregate.up.sql", size: 210, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __6_create_table_users_email_viewDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x26\x00\xd9\xff\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x75\x73\x65\x72\x73\x5f\x65\x6d\x61\x69\x6c\x5f\x76\x69\x65\x77\x3b\x03\x00\x57\xa2\x1f\x52\x26\x00\x00\x00")

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

	info := bindataFileInfo{name: "6_create_table_users_email_view.down.sql", size: 38, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __6_create_table_users_email_viewUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcc\xc1\xaa\xc2\x30\x10\x85\xe1\x7d\x9f\xe2\x2c\xef\x05\xdf\xc0\x55\xd5\x51\x83\xb5\x95\x74\x4a\xdb\x55\x08\x66\xc4\x40\x05\x69\xda\xfa\xfa\x42\xba\x51\x11\xb7\xf3\x7f\x73\xd6\x9a\x52\x26\x70\xba\xca\x08\x6a\x8b\xbc\x60\x50\xa3\x4a\x2e\x31\x06\xe9\x83\x91\x9b\xf5\x9d\x99\xbc\x3c\xf0\x97\x00\x88\x67\xe3\x1d\x98\x1a\x8e\x3c\xaf\xb2\x0c\x27\xad\x8e\xa9\x6e\x71\xa0\x76\x11\x59\xfc\x7b\x47\x73\xb8\xda\x60\x26\xe9\xfd\xc5\x8b\x9b\xd7\xa1\x72\xa6\x1d\xe9\x0f\x78\xee\xc5\x0e\xe2\x8c\x1d\xbe\xcd\x8c\x77\xf7\xa3\x3a\xe9\xe4\xa5\x26\xff\xa8\x15\xef\x8b\x8a\xa1\x8b\x5a\x6d\x96\xcf\x01\x00\xd1\x30\x2e\x31\xf7\x00\x00\x00")

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

	info := bindataFileInfo{name: "6_create_table_users_email_view.up.sql", size: 247, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_alter_table_users_aggregate_and_email_viewDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xcf\x4e\x02\x31\x10\xc6\xef\x7d\x8a\x39\x6a\xe2\x1b\x70\x5a\x64\xd4\xc6\x65\x97\x94\xd9\x00\xa7\x66\x62\x87\xb5\xc9\x0a\xa6\xfb\xc7\xd7\x37\xa5\x08\x2a\x8b\xf1\xd8\x7c\xbf\x99\xce\xef\x9b\x99\x72\x01\x94\x4d\x73\x04\xfd\x00\xb8\xd6\x4b\x5a\x42\xdf\x4a\x68\x2d\xd7\x75\x90\x9a\x3b\x99\xa8\x7b\x83\x19\xe1\x99\x2b\x4a\xba\xc2\xc2\x8d\x02\x00\x38\xbd\xad\x77\x40\xb8\x26\x58\x18\x3d\xcf\xcc\x06\x9e\x71\x73\x77\x40\x06\x6e\x7a\xb1\x8e\x3b\x4e\x40\xdc\x59\x54\x79\x7e\x4c\x25\xb4\x7e\xbf\xb3\x7e\xb7\xdd\x43\x51\xcd\xa7\x68\x52\xf0\x12\x24\xae\x75\xf1\xaf\x91\xb9\xfe\xdd\xfd\x88\xd5\x2d\xac\x34\x3d\x95\x15\x81\x29\x57\x7a\x36\x51\xea\x0f\x65\x79\x63\xdf\xd8\xc1\xcb\xc7\x3f\x9c\xcf\xf0\x51\x3a\x56\x71\xf2\xfd\x3a\xeb\x52\xfc\x30\x37\x76\xfb\x2b\xb7\x76\x90\xe0\xb7\x5e\x9c\x4d\x94\x2e\x08\x1f\xd1\xfc\x02\x53\x07\xce\x72\x37\xb6\x26\x55\x70\x2d\x75\xd2\xc8\xb7\xf4\xa2\x9f\xcf\x01\x00\xa1\xaa\x43\x50\x11\x02\x00\x00")

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

	info := bindataFileInfo{name: "7_alter_table_users_aggregate_and_email_view.down.sql", size: 529, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_alter_table_users_aggregate_and_email_viewUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x90\xcf\x4e\xc2\x40\x10\xc6\xef\xfb\x14\x73\xd4\xc4\x37\xe0\x54\x64\xd4\x8d\xa5\x25\xcb\x34\xc0\x69\x32\x71\x87\xba\x49\x05\xb3\xfd\xe3\xeb\x1b\xdb\x12\x92\x0a\xc6\x03\xd7\xfd\x7e\xdf\x97\x9d\xdf\xc2\xe5\x2b\xa0\x64\x9e\x22\xd8\x27\xc0\xad\x5d\xd3\x1a\xda\x5a\x63\xcd\x52\x96\x51\x4b\x69\x74\x66\x1e\x1d\x26\x84\x67\x2e\xcb\xe9\x0a\x0b\x77\x06\x00\xfa\x05\x0e\x1e\x08\xb7\x04\x2b\x67\x97\x89\xdb\xc1\x2b\xee\x1e\xfa\xb4\x93\xaa\x55\xf6\xd2\xc8\x00\xfc\xcc\x65\x45\x9a\x8e\xa9\xc6\x3a\x1c\x0f\x1c\x0e\xfb\x23\x64\xc5\x72\x8e\x6e\x08\xde\xa2\x4a\xa3\x9e\xa5\x01\x9b\x11\x3e\xa3\x9b\x34\xdb\x4f\x3f\x01\x86\xa2\xd7\x4a\xc7\xf7\x53\x60\xee\x61\x63\xe9\x25\x2f\x08\x5c\xbe\xb1\x8b\x99\x31\x7f\xa8\xd0\x0f\x09\x15\x77\x41\xbf\xfe\xe1\xe2\x0c\x5f\x92\x71\xfa\xf1\x6f\x2b\x7d\xef\x92\x90\x77\xa9\xb9\xd3\x18\xf6\x41\x3d\x0f\xd4\x78\xc5\x04\xbc\x85\xa0\x6b\x7e\xbe\x07\x00\xeb\xfe\xd4\xf8\x29\x02\x00\x00")

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

	info := bindataFileInfo{name: "7_alter_table_users_aggregate_and_email_view.up.sql", size: 553, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __8_create_table_users_token_viewDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x27\x00\xd8\xff\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x75\x73\x65\x72\x73\x5f\x74\x6f\x6b\x65\x6e\x5f\x76\x69\x65\x77\x3b\x0a\x03\x00\x94\xa1\x68\xbb\x27\x00\x00\x00")

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

	info := bindataFileInfo{name: "8_create_table_users_token_view.down.sql", size: 39, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __8_create_table_users_token_viewUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcc\x41\xcb\x82\x40\x10\xc6\xf1\xbb\x9f\xe2\x39\xbe\x2f\xf4\x0d\x3a\x59\x4d\xb4\x64\x1a\xeb\x88\x7a\x5a\x16\x77\x82\x25\xd1\xd0\xd5\xbe\x7e\x60\x97\x02\xe9\x3a\xff\xdf\x3c\x7b\x4d\x31\x13\x38\xde\x25\x04\x75\x44\x9a\x31\xa8\x52\x39\xe7\x98\x46\x19\x46\x13\xfa\xbb\x74\x66\xf6\xf2\xc4\x5f\x04\x60\x39\x1b\xef\xc0\x54\xf1\xc2\xd3\x22\x49\x70\xd5\xea\x12\xeb\x1a\x67\xaa\x37\x0b\x9b\x65\xf0\x37\xdf\xd8\xe0\xfb\xee\x3d\xf2\xfd\xb1\xa2\x46\xdb\x86\x35\xd4\x0c\x62\x83\x38\x63\x57\xeb\xf4\x70\x3f\xaa\x93\x56\x3e\x6a\xf4\x8f\x52\xf1\x29\x2b\x18\x3a\x2b\xd5\x61\x1b\xbd\x06\x00\x1e\x6c\xf4\xa2\x01\x01\x00\x00")

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

	info := bindataFileInfo{name: "8_create_table_users_token_view.up.sql", size: 257, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __9_create_index_users_token_view_verification_tokenDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x2d\x00\xd2\xff\x44\x52\x4f\x50\x20\x49\x4e\x44\x45\x58\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x76\x65\x72\x69\x66\x69\x63\x61\x74\x69\x6f\x6e\x5f\x74\x6f\x6b\x65\x6e\x5f\x61\x73\x63\x3b\x0a\x03\x00\xc0\x06\xc0\xb2\x2d\x00\x00\x00")

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

	info := bindataFileInfo{name: "9_create_index_users_token_view_verification_token.down.sql", size: 45, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __9_create_index_users_token_view_verification_tokenUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x65\x00\x9a\xff\x43\x52\x45\x41\x54\x45\x20\x49\x4e\x44\x45\x58\x20\x49\x46\x20\x4e\x4f\x54\x20\x45\x58\x49\x53\x54\x53\x20\x76\x65\x72\x69\x66\x69\x63\x61\x74\x69\x6f\x6e\x5f\x74\x6f\x6b\x65\x6e\x5f\x61\x73\x63\x20\x4f\x4e\x20\x75\x73\x65\x72\x73\x5f\x74\x6f\x6b\x65\x6e\x5f\x76\x69\x65\x77\x20\x28\x0a\x20\x20\x20\x20\x76\x65\x72\x69\x66\x69\x63\x61\x74\x69\x6f\x6e\x5f\x74\x6f\x6b\x65\x6e\x20\x41\x53\x43\x29\x3b\x0a\x03\x00\xc2\x76\x55\xef\x65\x00\x00\x00")

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

	info := bindataFileInfo{name: "9_create_index_users_token_view_verification_token.up.sql", size: 101, mode: os.FileMode(420), modTime: time.Unix(1767324376, 0)}
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
	"10_create_table_user_registration_verifications.down.sql":    _10_create_table_user_registration_verificationsDownSql,
	"10_create_table_user_registration_verifications.up.sql":      _10_create_table_user_registration_verificationsUpSql,
	"11_drop_users_token_view.down.sql":                           _11_drop_users_token_viewDownSql,
	"11_drop_users_token_view.up.sql":                             _11_drop_users_token_viewUpSql,
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
	"10_create_table_user_registration_verifications.down.sql":    &bintree{_10_create_table_user_registration_verificationsDownSql, map[string]*bintree{}},
	"10_create_table_user_registration_verifications.up.sql":      &bintree{_10_create_table_user_registration_verificationsUpSql, map[string]*bintree{}},
	"11_drop_users_token_view.down.sql":                           &bintree{_11_drop_users_token_viewDownSql, map[string]*bintree{}},
	"11_drop_users_token_view.up.sql":                             &bintree{_11_drop_users_token_viewUpSql, map[string]*bintree{}},
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
