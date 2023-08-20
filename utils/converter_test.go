package utils

import (
	"reflect"
	"testing"
)

func TestIntegerToHex(t *testing.T) {
	type args struct {
		input int64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Success Convert Integer To Hex",
			args: args{
				input: 1,
			},
			want: []byte{0, 0, 0, 0, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntegerToHex(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegerToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
