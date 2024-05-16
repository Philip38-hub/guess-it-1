package guess

import "math"

func Median(x []float64) float64 {
	data := Sort(x)
	out := 0.0
	if len(data)%2 == 1 {
		out = data[(len(x) / 2)]
	} else {
		out = (data[len(data)/2] + data[(len(data)/2)-1]) / 2
	}
	return out
}

func Sort(x []float64) []float64 {
	data := x
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}

func Average(x []float64) float64 {
	sum := 0.0
	counter := 0.0
	for _, num := range x {
		sum += num
		counter++
	}
	return sum / counter
}

func StandardDev(x []float64) float64 {
	out := Variance(x)
	return math.Sqrt(out)
}

func Variance(x []float64) float64 {
	mean := Average(x)
	out := []float64{}

	for _, num := range x {
		dev := num - mean
		out = append(out, dev*dev)
	}
	return Average(out)
}
