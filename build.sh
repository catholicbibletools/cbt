#!/bin/bash
echo "CATHOLiC BiBLE TOOLS"
cd src
export GOPATH=$PWD
function build_app {
	echo "Building app for '$3-$2' ..."
	rm -f ../dist/$3-$2/$4
	mkdir ../dist/$3-$2
	env GOOS=$1 GOARCH=$2 go build -o ../dist/$3-$2/$4 app/*.go
	cp app/cbt.cfg ../dist/$3-$2
	cp -rf templates ../dist/$3-$2
	cp -rf client ../dist/$3-$2
	cp ../README.md ../dist/$3-$2
}
function clear_dist {
	rm -rf ../dist
	mkdir ../dist
}
clear_dist
build_app "linux" "386" "linux" "cbt"
build_app "linux" "amd64" "linux" "cbt"
build_app "darwin" "386" "macos" "cbt.app"
build_app "darwin" "amd64" "macos" "cbt.app"
build_app "windows" "386" "windows" "cbt.exe"
build_app "windows" "amd64" "windows" "cbt.exe"
