#!/bin/bash

cd "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" || exit 1

elegen folder -d ../../regolithe/spec/tests || exit 1

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
    tail -n +11 model/list.go;
    tail -n +9 model/task.go;
    tail -n +9 model/unmarshalable.go;
    tail -n +9 model/user.go;
    tail -n +17 model/root.go;
    tail -n +5 model/identities_registry.go;
    tail -n +4 model/relationships_registry.go;
}>> ../data_test.go

sed -i '' 's/elemental\.//g' ../data_test.go
