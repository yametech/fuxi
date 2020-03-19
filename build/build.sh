#!/bin/bash

dir=`pwd`

build() {
	for d in $(ls ./$1); do
		echo "building $1/$d"
		pushd $dir/$1/$d >/dev/null
		GOOS=linux CGO_ENABLED=0 go build -ldflags '-w' -o $dir/bin/$1-$d
		popd >/dev/null
	done
}

build srv && build cmd
