#!/bin/sh 
echo "generating the go-bin data"

path_to_go_bindata=$(which go-bindata)
if [ -x "$path_to_go_bindata" ]; then 
    echo "found go-bindata: $path_to_go_bindata"
else
    echo "could not find go-bindata. installing go-bindata"
    go install github.com/go-bindata/go-bindata/go-bindata
fi

go-bindata -prefix "migrations/" -o resourcefile.go -pkg dbmigration migrations/**.sql