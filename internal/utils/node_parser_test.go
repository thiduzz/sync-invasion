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
		// TODO: Add test cases.
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
