package s2

import "testing"

func TestMultiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{10, 4}, 40},
		{"test2", args{-10, 4}, -40},
		{"test3", args{10, -4}, -40},
		{"test4", args{0, 4}, 0},
		{"test5", args{10, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
