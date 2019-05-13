#!/usr/bin/env bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )

go build -o $DIR/bin/zmake $DIR/main.go