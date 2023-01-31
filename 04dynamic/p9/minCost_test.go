package p9

import "testing"

func TestMinCost1(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		ic int
		dc int
		rc int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				"abc", "adc", 5, 3, 2,
			},
			want: 2,
		},
		{
			name: "02",
			args: args{
				"abc", "adc", 5, 3, 100,
			},
			want: 8,
		},
		{
			name: "02",
			args: args{
				"abc", "abc", 5, 3, 100,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCost1(tt.args.s1, tt.args.s2, tt.args.ic, tt.args.dc, tt.args.rc); got != tt.want {
				t.Errorf("MinCost1() = %v, want %v", got, tt.want)
			}
		})
	}
}
