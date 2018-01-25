#!/bin/bash

VERSION="v0.2"
PROJECT="autossh"

function build() {
    os=$1
    arch=$2
    alias_name=$3
    package="${PROJECT}-${alias_name}-${arch}_${VERSION}"

    echo "build ${package} ..."
    mkdir -p "./releases/${package}"
    CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -o "./releases/${package}/autossh" main.go
    cp ./servers.example.json "./releases/${package}/servers.json"
    cd ./releases/
    zip -r "./${package}.zip" "./${package}"
    echo "clean ${package}"
    rm -rf "./${package}"
    cd ../
}

# OS X Mac
build darwin amd64 macOS

# Linux
build linux amd64 linux
build linux 386 linux
build linux arm linux
