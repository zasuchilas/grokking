package recursion

import "testing"

func TestFact(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				x: 5,
			},
			want: 120, // 5*4*3*2*1
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fact(tt.args.x); got != tt.want {
				t.Errorf("Fact() = %v, want %v", got, tt.want)
			}
		})
	}
}
