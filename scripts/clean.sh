#!/bin/bash
## DeGOps: 0.0.4
set -o errexit
set -o nounset

rm -rf vendor
rm -rf tmp
rm -f coverage.out
rm -f coverage.html
