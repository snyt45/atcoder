#!/bin/bash

if [ -z "$1" ]; then
	echo "nを指定してください（使用法：make sw n=2）"
	exit 1
fi

selected=./test/sample-$1.in
cp $selected input
echo Switched to test case $selected
