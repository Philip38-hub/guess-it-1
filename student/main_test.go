package main

import "testing"

func TestNextValueRange(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NextValueRange(tt.args.data)
			if got != tt.want {
				t.Errorf("NextValueRange() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NextValueRange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
