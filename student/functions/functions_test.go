package guess

import (
	"reflect"
	"testing"
)

func TestVariance(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.x); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardDev(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDev(tt.args.x); got != tt.want {
				t.Errorf("StandardDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.x); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	type args struct {
		x []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.x); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}
