#!/bin/bash

SITE=$1
FILE=$2

# Generate screenshot.
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
    --headless --disable-gpu --hide-scrollbars --screenshot --window-size=1280,5596 $SITE

# Rename file.
mv screenshot.png $FILE
