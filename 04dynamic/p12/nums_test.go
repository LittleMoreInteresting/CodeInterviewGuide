package p12

import "testing"

func Test_num2(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := num2(tt.args.str); got != tt.want {
				t.Errorf("num2() = %v, want %v", got, tt.want)
			}
		})
	}
}
