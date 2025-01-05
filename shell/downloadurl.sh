#!/bin/bash

echo URL?
read url
rm -rf test
echo $url
oj download $url
cp test/sample-1.in input
