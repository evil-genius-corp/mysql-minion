#!/bin/bash
set -e

if [ ! -f setenv.sh ]; then
    echo "please create setenv.sh"
    exit 
fi

source setenv.sh
./main