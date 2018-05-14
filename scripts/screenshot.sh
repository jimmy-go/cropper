#!/bin/bash

SITE=$1
if [ -z "$SITE" ]; then
    echo "Site not set"
    exit
fi
FILE=$2
if [ -z "$FILE" ]; then
    echo "File not set"
    exit
fi

# Generate screenshot.
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
    --headless --disable-gpu --hide-scrollbars --screenshot \
    --window-size=1280,5596 \
    --default-background-color=00000000 \
    $SITE

# Rename file.
mv screenshot.png $FILE
