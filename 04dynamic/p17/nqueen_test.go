package p17

import "testing"

func TestNqueen(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "01", args: args{n: 1}, want: 1},
		{name: "02", args: args{n: 2}, want: 0},
		{name: "03", args: args{n: 3}, want: 0},
		{name: "08", args: args{n: 8}, want: 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Nqueen(tt.args.n); got != tt.want {
				t.Errorf("Nqueen() = %v, want %v", got, tt.want)
			}
		})
	}
}
