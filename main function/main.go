package main

import (
	"bufio"
	"fmt"
	"guess"
	"os"
	"strconv"
	"math"
)

func main() {
	file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)
	var data []float64
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("unable to convert data set to float values")
			os.Exit(0)
		}
		data = append(data, number)
	}
	lower, upper := NextValueRange(data)
	cleandata := RemoveOutliers(data)
	Prediction(cleandata, upper, lower)
}

func RemoveOutliers(d []float64) []float64 {
	clean := []float64{}
	// make a copy of the slice so that the slice outside function is not affected
	copyofd := make([]float64, len(d)) 
	copy(copyofd, d)
	//use q1 and q3 to determine the possible range of the dataset and remove outliers
	q1 := math.Round(guess.Median(copyofd[:len(copyofd)/2]))
	q2 := math.Round(guess.Median(copyofd[len(copyofd)/2:]))
	iqr := q2 - q1
	upperBoundary := q2 + 1.5*iqr
	lowerBoundary := q1 - 1.5*iqr

	for _, num := range d {
		if lowerBoundary > num || num > upperBoundary {
			num = math.Round(guess.Median(copyofd))
		}
	clean = append(clean, num)
	}
	return clean
}

func NextValueRange(d []float64) (float64, float64) {
	clean := RemoveOutliers(d)
	mean := guess.Average(clean)
	stddev := guess.StandardDev(clean)
	lowerBoundary := mean - stddev
	upperBoundary := mean + stddev
	return lowerBoundary, upperBoundary
}

func Prediction(d []float64, u float64, l float64) {
	count := 0
	for i := range d {
		if d[i] == d[len(d)-1] {
			continue
		}
		if d[i] > l && d[i] < u {
			count++
		}
		fmt.Printf("%d --> the standard input\n", int(math.Round(d[i])))
		fmt.Printf("%.0f - %.0f --> range for the next input. In this case for the number %.0f\n", l, u, d[i+1] )
	}
	accuracy := 100 * count/len(d)
	fmt.Printf("Prediction  accuracy: %d%%\n", accuracy)
}