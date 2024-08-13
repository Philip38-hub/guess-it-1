package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	guess "guess/functions"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run [program name] [filename]")
		os.Exit(0)
	}
	var data []float64
	// Create a scanner to read from standard input
	input := os.Stdin
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Invalid input: %v\n", err)
			continue
		}
		if num < math.MinInt64 || num > math.MaxInt64 {
			fmt.Printf("cannot use an overflow value %.f\n", num)
			return
		}
		data = append(data, num)
		if len(data) > 1 {
			lower, upper := NextValueRange(data)
			if lower < math.MinInt64 || upper > math.MaxInt64 {
				fmt.Printf("cannot use an overflow value %.f\n", num)
				return
			}
			fmt.Printf("%.0f %.0f\n", lower, upper)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading standard input: %v", err)
	}
}

func NextValueRange(data []float64) (float64, float64) {
	if len(data) == 0 {
		return 0, 0
	}
	// cleandata:= RemoveOutliers(data)
	m := guess.Average(data)
	std := guess.StandardDev(data)
	lower := m - std*2
	upper := m + std*2
	return lower, upper
}
