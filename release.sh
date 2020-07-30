#!/bin/bash
# Builds our release artifacts and puts

cp -r static $TMPDIR

export GOOS GOARCH
for GOOS in linux darwin; do
    for GOARCH in amd64 arm; do
        echo "Building $GOOS-$GOARCH"
        TMPDIR="$(mktemp -d)"

        cp -r static $TMPDIR
        go build -o $TMPDIR/go-shorten github.com/thomasdesr/go-shorten

        tar -cf go-shorten-$GOOS-$GOARCH.tar \
            -C $TMPDIR \
            go-shorten static

        rm -r $TMPDIR
    done
done

echo "Release complete"
