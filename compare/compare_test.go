package compare

import (
	"reflect"
	"testing"
)

func TestTernary(t *testing.T) {
	type args struct {
		boolean bool
		ifTrue  string
		ifFalse string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "true",
		args: args{
			true,
			"true",
			"false",
		},
		want: "true",
	},
		{
			name: "false",
			args: args{
				false,
				"true",
				"false",
			},
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ternary(tt.args.boolean, tt.args.ifTrue, tt.args.ifFalse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}
