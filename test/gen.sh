#!/bin/bash

cd "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" || exit 1

elegen github -r https://github.com/aporeto-inc/regolithe -p spec/tests -t "$GITHUB_TOKEN"

mkdir -p model
mv codegen/elemental/* ./model
rm -rf codegen
