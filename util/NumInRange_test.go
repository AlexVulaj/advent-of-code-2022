package util

import "testing"

func TestNumInRange(t *testing.T) {
	type args struct {
		num int
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Number in range",
			args: args{
				num: 2,
				min: 1,
				max: 3,
			},
			want: true,
		},
		{
			name: "Number matches bottom of range",
			args: args{
				num: 1,
				min: 1,
				max: 3,
			},
			want: true,
		},
		{
			name: "Number matches top of range",
			args: args{
				num: 3,
				min: 1,
				max: 3,
			},
			want: true,
		},
		{
			name: "Number is below range",
			args: args{
				num: 0,
				min: 1,
				max: 3,
			},
			want: false,
		},
		{
			name: "Number is above range",
			args: args{
				num: 4,
				min: 1,
				max: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumInRange(tt.args.num, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("NumInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
