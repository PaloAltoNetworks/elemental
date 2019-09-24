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
	"time"

    "github.com/globalsign/mgo/bson"
    "github.com/mitchellh/copystructure"
)

//lint:file-ignore U1000 auto generated code.
EOF
{
    tail -n +12 model/list.go;
    tail -n +11 model/task.go;
    tail -n +19 model/unmarshalable.go;
    tail -n +11 model/user.go;
    tail -n +19 model/root.go;
    tail -n +5 model/identities_registry.go;
    tail -n +4 model/relationships_registry.go;
}>> ../data_test.go

sed -i '' 's/elemental\.//g' ../data_test.go
