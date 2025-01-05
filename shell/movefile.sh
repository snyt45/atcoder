#!/bin/bash

contest=$(cat ./contest)
echo CONTEST: $contest
echo QUESTION?
read q
move_dir=result/$contest/$q
mkdir -p $move_dir
mv -i main.go $move_dir
