#!/bin/bash

set -x

BIN=`which $1`

LIB_PATH=/tmp/host/lib
BIN_PATH=/tmp/host/bin

mkdir -p $LIB_PATH $BIN_PATH

cp $BIN $BIN_PATH

libs=`ldd $BIN | grep -v "linux-vdso.so" | grep -v "ld-linux-x86-64.so" | awk '{print $3}'`

for lib in $libs
do
	cp $lib $LIB_PATH
done
