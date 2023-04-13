package calculate

import (
	"testing"
)

func TestAddition(t *testing.T) {
	type args struct {
		number_a int
		number_b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Addition positif number",
			args: args{20, 2},
			want: 22,
		},
		{
			name: "Addition negative number",
			args: args{-20, -2},
			want: -22,
		},
		{
			name: "Addition negative and positive number",
			args: args{-2, 20},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.args.number_a, tt.args.number_b); got != tt.want {
				t.Errorf("Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstrack(t *testing.T) {
	type args struct {
		number_a int
		number_b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "substrack positif number",
			args: args{20, 2},
			want: 18,
		},
		{
			name: "substrack negative number",
			args: args{-20, -2},
			want: -18,
		},
		{
			name: "substrack negative and positive number",
			args: args{-2, 20},
			want: -22,
		},

		{
			name: "substrack negative result",
			args: args{1, 10},
			want: -9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substrack(tt.args.number_a, tt.args.number_b); got != tt.want {
				t.Errorf("Substrack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiple(t *testing.T) {
	type args struct {
		number_a int
		number_b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Multiple positif number",
			args: args{20, 2},
			want: 40,
		},
		{
			name: "Multiple negative number",
			args: args{-20, -2},
			want: 40,
		},
		{
			name: "Multiple negative and positive number",
			args: args{-2, 20},
			want: -40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiple(tt.args.number_a, tt.args.number_b); got != tt.want {
				t.Errorf("Multiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		number_a float64
		number_b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Divide positif number",
			args: args{20, 2},
			want: 10,
		},
		{
			name: "Divide negative number",
			args: args{-20, -2},
			want: 10,
		},
		{
			name: "Divide negative and positive number",
			args: args{20, -2},
			want: -10,
		},
		{
			name: "Divide float result",
			args: args{2, 4},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.number_a, tt.args.number_b); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
