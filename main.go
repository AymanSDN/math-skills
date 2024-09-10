package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"math-skills/source"
)


func calcAverage(buffer string) ([]float64, int){
	var numsArray []float64
	data := source.RemoveEmptyStr(strings.Split(string(buffer), "\n"))

	for i := 0; i < len(data); i++ {
		tmpNum, err := strconv.Atoi(data[i])
		source.CheckError("error: strconv.Atoi failed to convert ["+data[i]+"]", err)
		numsArray = append(numsArray, float64(tmpNum))
		Average += float64(tmpNum)
	}
	size := len(numsArray)
	Average /= float64(size)
	
	return numsArray, size
}

func calcMedian(numsArray []float64, size int)  {
	sort.Float64s(numsArray)
	if size % 2 == 0 {
		Median = (numsArray[size/2] + numsArray[size/2-1])/2
	} else {
		Median = numsArray[int(size/2)]
	}
	
}

func calcVarianceAndStdDiv(numsArray []float64, size int) {
	for i := 0; i < size ; i++ {
		Variance += math.Pow(numsArray[i]-Average, 2)
	}
	Variance /= float64(size)
	StdDeviation = math.Sqrt(Variance)
}

var Average, Median, Variance, StdDeviation float64

func main() {
	// source.GenerateData(999)
	if len(os.Args[1:]) == 1 {
		filePath := os.Args[1]
		buffer, err := os.ReadFile(filePath)
		source.CheckError("error: cannot read Input file", err)
		numsArr, size := calcAverage(string(buffer))
		calcMedian(numsArr, size)
		calcVarianceAndStdDiv(numsArr, size)
		// fmt.Println(math.Round(3.5))
		// fmt.Println(math.Round(4.49))
	
		fmt.Printf("Average: %v\nMedian: %v\nVariance: %v\nStandard Deviation: %v\n", math.Round(Average), math.Round(Median), int(math.Round(Variance)), math.Round(StdDeviation))

	} else {
		fmt.Println("wrong usage")
	}
}
