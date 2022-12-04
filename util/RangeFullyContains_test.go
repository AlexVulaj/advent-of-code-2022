package util

import "testing"

func TestRangeFullyContains(t *testing.T) {
	type args struct {
		firstMin  int
		firstMax  int
		secondMin int
		secondMax int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First contains second entirely",
			args: args{
				firstMin:  1,
				firstMax:  4,
				secondMin: 2,
				secondMax: 3,
			},
			want: true,
		},
		{
			name: "First contains second with same lower bound",
			args: args{
				firstMin:  1,
				firstMax:  3,
				secondMin: 1,
				secondMax: 2,
			},
			want: true,
		},
		{
			name: "First contains second with same upper bound",
			args: args{
				firstMin:  1,
				firstMax:  3,
				secondMin: 2,
				secondMax: 3,
			},
			want: true,
		},
		{
			name: "Second contains first entirely",
			args: args{
				firstMin:  2,
				firstMax:  3,
				secondMin: 1,
				secondMax: 4,
			},
			want: true,
		},
		{
			name: "Second contains first with same lower bound",
			args: args{
				firstMin:  1,
				firstMax:  2,
				secondMin: 1,
				secondMax: 3,
			},
			want: true,
		},
		{
			name: "Second contains first with same upper bound",
			args: args{
				firstMin:  2,
				firstMax:  3,
				secondMin: 1,
				secondMax: 3,
			},
			want: true,
		},
		{
			name: "First and second match exactly",
			args: args{
				firstMin:  1,
				firstMax:  3,
				secondMin: 1,
				secondMax: 3,
			},
			want: true,
		},
		{
			name: "Partial overlap with lower first",
			args: args{
				firstMin:  1,
				firstMax:  3,
				secondMin: 2,
				secondMax: 4,
			},
			want: false,
		},
		{
			name: "Partial overlap with higher first",
			args: args{
				firstMin:  2,
				firstMax:  4,
				secondMin: 1,
				secondMax: 3,
			},
			want: false,
		},
		{
			name: "Ranges don't overlap at all",
			args: args{
				firstMin:  1,
				firstMax:  2,
				secondMin: 3,
				secondMax: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeFullyContains(tt.args.firstMin, tt.args.firstMax, tt.args.secondMin, tt.args.secondMax); got != tt.want {
				t.Errorf("RangesOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
