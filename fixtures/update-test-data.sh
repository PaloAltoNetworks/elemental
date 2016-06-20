#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR

monogen -f ./specs -L elemental

echo -n > ./data_test.go
cat ./codegen/elemental/1.0/list.go >> ./data_test.go
cat ./codegen/elemental/1.0/task.go | (read; read; cat) >> ./data_test.go
cat >> ./data_test.go << EOF

var UnmarshalableListIdentity = Identity{Name: "list", Category: "lists"}

type UnmarshalableList struct {
	List
}

func NewUnmarshalableList() *UnmarshalableList {
	return &UnmarshalableList{List: List{}}
}

func (o *UnmarshalableList) Identity() Identity { return UnmarshalableListIdentity }

func (o *UnmarshalableList) UnmarshalJSON([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

func (o *UnmarshalableList) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

func (o *UnmarshalableList) Validate() Errors { return nil }
EOF

mv ./data_test.go ../

rm -rf codegen
rm -f ./data_test.go-e
cd -
