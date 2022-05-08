package main

import (
	"testing"
)

func Test_validateInput(t *testing.T) {
	type args struct {
		path *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should throw error when file name is empty",
			args: args{
				path: ptrStr(""),
			},
			wantErr: true,
		},
		{
			name: "should not throw error when file name is value",
			args: args{
				path: ptrStr("/test"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateInput(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("validateInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func ptrStr(value string) *string {
	return &value
}
