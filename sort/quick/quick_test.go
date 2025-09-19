package quick

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				list: []int{10, 5, 2, 3},
			},
			want: []int{2, 3, 5, 10},
		},
		{
			name: "success",
			args: args{
				list: []int{5, 3, 1, 7, 9},
			},
			want: []int{1, 3, 5, 7, 9},
		},
		{
			name: "success",
			args: args{
				list: []int{5},
			},
			want: []int{5},
		},
		{
			name: "success",
			args: args{
				list: []int{5, -3, 0, 12, 4, -6, 0, 10},
			},
			want: []int{-6, -3, 0, 0, 4, 5, 10, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
