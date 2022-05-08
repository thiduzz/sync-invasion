package nodes

import (
	"fmt"
	"github.com/thiduzz/code-kata-invasion/tools"
	"reflect"
	"testing"
)

func TestAttackerCollection_Sort(t *testing.T) {
	type fields struct {
		Collection func(t *testing.T) *AttackerCollection
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
			name: "should return random slice of attacker identifiers",
			fields: fields{
				Collection: func(t *testing.T) *AttackerCollection {
					collection := NewAttackerCollection()
					for i := 1; i <= 10; i++ {
						factory := NewAttackerFactory(func() string {
							return fmt.Sprintf("test-%d", i)
						})
						attacker, err := factory.Generate(Attacker{}, uint(i))
						if err != nil {
							t.Fatal("fail generating attacker")
						}
						collection.Add(attacker)
					}
					return collection
				},
			},
			args: args{
				randomizerFunc: func() *tools.Randomizer {
					return tools.NewRandomizer(10)
				},
			},
			want: []uint{1, 9, 10, 2, 5, 7, 3, 8, 4, 6},
		},
		{
			name: "should return a different random slice of attacker if randomizer seed is different",
			fields: fields{
				Collection: func(t *testing.T) *AttackerCollection {
					collection := NewAttackerCollection()
					for i := 1; i <= 10; i++ {
						factory := NewAttackerFactory(func() string {
							return fmt.Sprintf("test-%d", i)
						})
						attacker, err := factory.Generate(Attacker{}, uint(i))
						if err != nil {
							t.Fatal("fail generating attacker")
						}
						collection.Add(attacker)
					}
					return collection
				},
			},
			args: args{
				randomizerFunc: func() *tools.Randomizer {
					return tools.NewRandomizer(11)
				},
			},
			want: []uint{2, 6, 3, 9, 10, 5, 4, 7, 8, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := tt.fields.Collection(t)
			if got := ac.Sort(tt.args.randomizerFunc()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
