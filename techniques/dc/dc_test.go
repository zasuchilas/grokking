package dc

import "testing"

func TestLoopSum(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				list: []int{3, 5, 7, 9},
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoopSum(tt.args.list); got != tt.want {
				t.Errorf("LoopSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				list: []int{3, 5, 7, 9},
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.list); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				list: []int{3, 5, 7, 9},
			},
			want: 4,
		},
		{
			name: "success 0",
			args: args{
				list: []int{},
			},
			want: 0,
		},
		{
			name: "success 1",
			args: args{
				list: []int{3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.list); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				list: []int{3, 5, 7, 9},
			},
			want:    9,
			wantErr: false,
		},
		{
			name: "success unsorted",
			args: args{
				list: []int{9, 5, 3, 7},
			},
			want:    9,
			wantErr: false,
		},
		{
			name: "failure 0",
			args: args{
				list: []int{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failure 1",
			args: args{
				list: []int{3},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Max(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("Max() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Max() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFarmSquarePlots(t *testing.T) {
	type args struct {
		w int
		h int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				w: 1680,
				h: 640,
			},
			want: 80,
		},
		{
			name: "success w < h",
			args: args{
				w: 640,
				h: 1680,
			},
			want: 80,
		},
		{
			name: "success 0",
			args: args{
				w: 0,
				h: 0,
			},
			want: 0,
		},
		{
			name: "success 1",
			args: args{
				w: 7,
				h: 4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FarmSquarePlots(tt.args.w, tt.args.h); got != tt.want {
				t.Errorf("FarmSquarePlots() = %v, want %v", got, tt.want)
			}
		})
	}
}
