#!/bin/sh

# Save the current working directory
original_dir=$(pwd)

# Change the working directory to the directory the script is running from
cd "$(dirname "$0")" || exit

go run `find  -maxdepth 1 -type f -name '*.go' | grep -v '_test.go$'` $*

# Restore the original working directory
cd "$original_dir" || exit