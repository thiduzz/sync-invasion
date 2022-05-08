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
		want    func() *nodes.LocationCollection
		wantErr bool
	}{
		{
			name: "should return error when provided path is empty",
			args: args{
				filePath: ptrStr(""),
			},
			want:    func() *nodes.LocationCollection { return nil },
			wantErr: true,
		},
		{
			name: "should return error when file does not exist",
			args: args{
				filePath: ptrStr("/doesnt-exist"),
			},
			want:    func() *nodes.LocationCollection { return nil },
			wantErr: true,
		},
		{
			name: "should return error when provided empty file",
			args: args{
				filePath: ptrStr("../../resources/empty-map.txt"),
			},
			want:    func() *nodes.LocationCollection { return nil },
			wantErr: true,
		},
		{
			name: "should return error when provided map with empty row",
			args: args{
				filePath: ptrStr("../../resources/map-with-empty-row.txt"),
			},
			want:    func() *nodes.LocationCollection { return nil },
			wantErr: true,
		},
		{
			name: "should be able to read existing file",
			args: args{
				filePath: ptrStr("../../resources/test-map.txt"),
			},
			want: func() *nodes.LocationCollection {
				return &nodes.LocationCollection{
					ReferenceMap: map[string]uint{
						"Hamburg": 1, "Beijing": 2, "Berlin": 3, "Bremen": 4, "Moscow": 5,
					},
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNodes(tt.args.filePath)
			want := tt.want()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if want != nil && !reflect.DeepEqual(got.ReferenceMap, want.ReferenceMap) {
				t.Errorf("ParseNodes() got = %v, want %v", got.ReferenceMap, want.ReferenceMap)
			}
		})
	}
}

func ptrStr(value string) *string {
	return &value
}
