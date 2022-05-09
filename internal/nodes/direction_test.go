package nodes

import (
	"github.com/thiduzz/code-kata-invasion/tools"
	"reflect"
	"testing"
)

func TestDirections_GetRandomizedRoads(t *testing.T) {
	type fields struct {
		Roads map[string]map[uint]bool
	}
	type args struct {
		randomizerFunc func() *tools.Randomizer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []uint
	}{
		{
			name: "should return nil when no links are found",
			fields: fields{
				Roads: map[string]map[uint]bool{
					"north": {},
					"south": {},
				},
			},
			args: args{
				randomizerFunc: func() *tools.Randomizer {
					return nil
				},
			},
			want: nil,
		},
		{
			name: "should return value from when no links are found",
			fields: fields{
				Roads: map[string]map[uint]bool{
					"north": {1: true, 2: true, 3: true},
					"south": {},
				},
			},
			args: args{
				randomizerFunc: func() *tools.Randomizer {
					return tools.NewRandomizer(10)
				},
			},
			want: []uint{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Directions{
				Roads: tt.fields.Roads,
			}
			if got := d.GetRandomizedRoads(tt.args.randomizerFunc()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomizedRoads() = %v, want %v", got, tt.want)
			}
		})
	}
}
