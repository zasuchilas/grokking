package listutil

import (
	"reflect"
	"testing"
)

func TestMakeOrderedIntList(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{length: 10},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeOrderedIntList(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeOrderedIntList() = %v, want %v", got, tt.want)
			}
		})
	}
}
