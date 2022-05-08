package engine

import (
	"errors"
	"github.com/golang/mock/gomock"
	mock_nodes "github.com/thiduzz/code-kata-invasion/internal/mock/nodes"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"testing"
)

func TestEngine_PrepareAttackers(t *testing.T) {
	type fields struct {
		AttackersQty uint
	}
	type args struct {
		factoryMockFunc func(ctrl *gomock.Controller) nodes.AttackerFactoryInterface
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *nodes.AttackerCollection
		wantErr bool
	}{
		{
			name: "should set 3 attackers in the engine",
			fields: fields{
				AttackersQty: 3,
			},
			args: args{
				factoryMockFunc: func(ctrl *gomock.Controller) nodes.AttackerFactoryInterface {
					mock := mock_nodes.NewMockAttackerFactoryInterface(ctrl)
					mock.EXPECT().Generate(nodes.Attacker{}, uint(3)).Return(&nodes.Attacker{
						Id:   uint(3),
						Name: "Test3",
					}, nil)
					mock.EXPECT().Generate(nodes.Attacker{}, uint(2)).Return(&nodes.Attacker{
						Id:   uint(2),
						Name: "Test2",
					}, nil)
					mock.EXPECT().Generate(nodes.Attacker{}, uint(1)).Return(&nodes.Attacker{
						Id:   uint(1),
						Name: "Test1",
					}, nil)
					return mock
				},
			},
			want: &nodes.AttackerCollection{
				Collection: map[uint]*nodes.Attacker{
					1: {
						Id:   1,
						Name: "Test1",
					},
					2: {
						Id:   2,
						Name: "Test2",
					},
					3: {
						Id:   3,
						Name: "Test3",
					},
				},
				ReferenceMap: map[string]uint{"Test1": 1, "Test2": 2, "Test3": 3},
			},
			wantErr: false,
		},
		{
			name: "should throw an error",
			fields: fields{
				AttackersQty: 3,
			},
			args: args{
				factoryMockFunc: func(ctrl *gomock.Controller) nodes.AttackerFactoryInterface {
					mock := mock_nodes.NewMockAttackerFactoryInterface(ctrl)
					mock.EXPECT().Generate(nodes.Attacker{}, gomock.Any()).Return(nil, errors.New("any-error"))
					return mock
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			en := &Engine{
				Attackers:    nodes.NewAttackerCollection(),
				AttackersQty: tt.fields.AttackersQty,
			}
			if err := en.PrepareAttackers(tt.args.factoryMockFunc(ctrl)); (err != nil) != tt.wantErr {
				t.Errorf("PrepareAttackers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
