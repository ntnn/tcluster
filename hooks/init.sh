#!/usr/bin/env sh

cd $(git rev-parse --show-toplevel)

rm -r .git/hooks
ln -s ../hooks .git/hooks
