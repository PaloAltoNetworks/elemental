#!/bin/bash

cd "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)" || exit 1

/Users/knainwal/apomux/workspace/code/go/src/go.aporeto.io/elemental/cmd/elegen/elegen folder -d ../../regolithe/spec/tests || exit 1

mkdir -p model
mv codegen/elemental/* ./model
rm -rf codegen

cat <<EOF >../data_test.go
package elemental

import (
	"fmt"
	"time"

    "go.mongodb.org/mongo-driver/bson"
    "github.com/mitchellh/copystructure"
)

//lint:file-ignore U1000 auto generated code.
EOF
{
	tail -n +14 model/list.go
	tail -n +13 model/task.go
	tail -n +21 model/unmarshalable.go
	tail -n +13 model/user.go
	tail -n +21 model/root.go
	tail -n +7 model/identities_registry.go
	tail -n +7 model/relationships_registry.go
} >>../data_test.go

sed 's/elemental\.//g' ../data_test.go >../data_test.go.new
mv ../data_test.go.new ../data_test.go
rm -f ../data_test.go.new
