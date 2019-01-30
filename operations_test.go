// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import "testing"

func TestParseOperation(t *testing.T) {
	type args struct {
		op string
	}
	tests := []struct {
		name    string
		args    args
		want    Operation
		wantErr bool
	}{
		{
			"create",
			args{
				"create",
			},
			OperationCreate,
			false,
		},
		{
			"delete",
			args{
				"delete",
			},
			OperationDelete,
			false,
		},
		{
			"info",
			args{
				"info",
			},
			OperationInfo,
			false,
		},
		{
			"patch",
			args{
				"patch",
			},
			OperationPatch,
			false,
		},
		{
			"retrieve",
			args{
				"retrieve",
			},
			OperationRetrieve,
			false,
		},
		{
			"retrieve-many",
			args{
				"retrieve-many",
			},
			OperationRetrieveMany,
			false,
		},
		{
			"update",
			args{
				"update",
			},
			OperationUpdate,
			false,
		},
		{
			"invalid",
			args{
				"invalid",
			},
			OperationEmpty,
			true,
		},
		{
			"CREATE",
			args{
				"CREATE",
			},
			OperationEmpty,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOperation(tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseOperation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
