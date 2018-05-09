#!/bin/bash
## DeGOps: 0.0.4
set -o errexit
set -o nounset

mkdir -p tmp
BIN=tmp/cropper_tmp

go build -o $BIN \
    ./cmd/cropper/...

$BIN -port=9900 \
    -screenshot-dir=$PWD/tmp \
    -exec="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
