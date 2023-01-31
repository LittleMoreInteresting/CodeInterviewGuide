package p8

import "testing"

func TestLcst2(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "t1", args: args{s1: "1AB2345CD", s2: "12345EF"}, want: "2345"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcst2(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Lcst2() = %v, want %v", got, tt.want)
			}
		})
	}
}
