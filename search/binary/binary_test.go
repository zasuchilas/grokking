package binary

import (
	"github.com/zasuchilas/grokking/utils/benchutil"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, item := range benchutil.Items {
			_, _ = Search(benchutil.List, item)
		}
	}
}

func TestSearch(t *testing.T) {
	type args struct {
		list []int
		item int
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
				list: []int{1, 3, 5, 7, 9},
				item: 7,
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Search(tt.args.list, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
