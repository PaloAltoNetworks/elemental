#!/bin/bash

cd "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" || exit 1

monogen -f ./specs -L elemental

mkdir -p model
mv codegen/elemental/* ./model
rm -rf codegen
