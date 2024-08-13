## Guess-it-1
### scope
* Guess-it-1 is an implementation of math-skills' functions(median, mean, variance and standard deviation) that given a number as standard input, prints out a range in which the next number provided should be.
* Guess-it-1 uses the follwing range to manage overflows within input dataset and output math.MinInt64 > num > math.MaxInt64


*Example of a data set* 


    189


    113


    121


    114


    145


    110


    ...

### Packages
`strconv` for string manipulation

`os` to read from the terminal and to open the file

`bufio` to read the file line by line

`fmt` to print to the terminal

`math` to allow usage of functions such as `math.Sqrt`

### Functions
`func Median(x []float64) float64` determines the median of the data set


`func Average(x []float64) float64` determines the average/ mean of the data set


`func StandardDev(x []float64) float64` determines the standard deviation od the data set

`func Variance(x []float64) float64` determines the variance of the data set


`func Sort(x []float64) []float64` sorts the values of the data set in ascending order. This helps in determining median

`func NextRangeValue(x []float64) (float64, float64)` sorts the values of the data set in ascending order. This helps in determining median

### Test case I
#### Data set
    6
    7
    3
    8
    10
    2  

#### Output  
    6
    7
    6 8
    3
    2 9
    8
    2 10
    2

### Run program Locally

Clone the project

```bash
  git clone https://learn.zone01kisumu.ke/git/pochieng/guess-it-1.git
```

Go to the project directory

```bash
  cd guess-it-1
  cd student
```

Run the program

```bash
  usage: Usage: go run [program name]
  Example1: go run main.go
  Example2: go run .
```


## Running Tests

To run tests, run the following command

```bash
  go test
```