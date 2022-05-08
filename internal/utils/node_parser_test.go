package utils

import (
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"reflect"
	"testing"
)

func TestParseNodes(t *testing.T) {
	type args struct {
		filePath *string
	}
	tests := []struct {
		name    string
		args    args
		want    *nodes.LocationCollection
		wantErr bool
	}{
		{
			name: "should return error when provided path is empty",
			args: args{
				filePath: ptrStr(""),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error when file does not exist",
			args: args{
				filePath: ptrStr("/doesnt-exist"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should be able to read existing file",
			args: args{
				filePath: ptrStr("../../resources/world-map.txt"),
			},
			want:    nodes.NewLocationCollection(),
			wantErr: false,
		},
		{
			name: "should return empty map when provided empty file",
			args: args{
				filePath: ptrStr("../../resources/empty-map.txt"),
			},
			want:    nodes.NewLocationCollection(),
			wantErr: false,
		},
		{
			name: "should return error when provided map with empty row",
			args: args{
				filePath: ptrStr("../../resources/map-with-empty-row.txt"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNodes(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func ptrStr(value string) *string {
	return &value
}
