#!/bin/bash

cd "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" || exit 1

elegen github -r https://github.com/aporeto-inc/regolithe -p spec/tests -t "$GITHUB_TOKEN"

mkdir -p model
mv codegen/elemental/* ./model
rm -rf codegen

cat <<EOF >../data_test.go
package elemental

import (
	"fmt"
	"sync"
	"time"
)
EOF
{
    tail -n +10 model/list.go;
    tail -n +10 model/task.go;
    tail -n +10 model/unmarshalable.go;
    tail -n +10 model/user.go;
    tail -n +16 model/root.go;
    tail -n +5 model/identities_registry.go;
}>> ../data_test.go

sed -i '' 's/elemental\.//g' ../data_test.go
