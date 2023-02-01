package p14

import "testing"

func TestWin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "01", args: args{[]int{1, 2, 100, 4}}, want: 101},
		{name: "02", args: args{[]int{1, 100, 2}}, want: 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Win(tt.args.nums); got != tt.want {
				t.Errorf("Win() = %v, want %v", got, tt.want)
			}
		})
	}
}
