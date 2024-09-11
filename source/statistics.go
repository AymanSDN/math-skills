package source

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

var Average, Median, Variance, StdDeviation float64

func CalcAverage(buffer string) ([]float64, int) {
	var numsArray []float64
	data := RemoveEmptyStr(strings.Split(string(buffer), "\n"))

	for i := 0; i < len(data); i++ {
		tmpNum, err := strconv.Atoi(data[i])
		CheckError("error: strconv.Atoi failed to convert ["+data[i]+"]", err)
		numsArray = append(numsArray, float64(tmpNum))
		Average += float64(tmpNum)
	}

	size := len(numsArray)
	Average /= float64(size)

	return numsArray, size
}

func CalcMedian(numsArray []float64, size int) {
	sort.Float64s(numsArray)

	if size%2 == 0 {
		Median = (numsArray[size/2] + numsArray[size/2-1]) / 2
	} else {
		Median = numsArray[int(size/2)]
	}
}

func CalcVarianceAndStdDiv(numsArray []float64, size int) {
	for i := 0; i < size; i++ {
		Variance += math.Pow(numsArray[i]-Average, 2)
	}
	
	Variance /= float64(size)
	StdDeviation = math.Sqrt(Variance)
}
