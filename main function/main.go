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
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error Opening file! Check file's name or path and try again!")
		os.Exit(0)
	}
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
	counter := 0
	if len(data) == 0 { // check if file is empty
		fmt.Println("There is no data to work with!")
		return
	}
	for range data {
		counter++
	}
	lower, upper := NextValueRange(data)
	cleandata := RemoveOutliers(data)
	Prediction(cleandata, upper, lower)
	fmt.Println("Data set count: ", counter)
}

func RemoveOutliers(d []float64) []float64 {
	clean := []float64{}
	// make a copy of the slice so that the slice outside function is not affected
	copyofd := make([]float64, len(d)) 
	copy(copyofd, d)
	sorted := guess.Sort(copyofd)
	//use q1 and q3 to determine the possible range of the dataset and remove outliers
	q1 := math.Round(guess.Median(sorted[:len(sorted)/2]))
	q2 := math.Round(guess.Median(sorted[len(sorted)/2:]))
	iqr := q2 - q1
	upperBoundary := q2 + 3*iqr
	lowerBoundary := q1 - 3*iqr

	for _, num := range d {
		if lowerBoundary > num || num > upperBoundary {
			num = math.Round(guess.Median(sorted))
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
	for i := range d[:len(d)-1] {
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