#!/bin/bash

contest=$(cat ./contest)
echo -e CONTEST: $contest
echo QUESTION?
read q
rm -rf test
url="https://atcoder.jp/contests/${contest}/tasks/${contest}_${q}"
echo $url
oj download $url
cp test/sample-1.in input
