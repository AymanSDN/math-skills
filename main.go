package main

import (
	"fmt"
	"math"
	"os"

	"math-skills/source"
)

func main() {
	// source.GenerateData(999)
	if len(os.Args[1:]) == 1 {
		filePath := os.Args[1]
		buffer, err := os.ReadFile(filePath)
		source.CheckError("error: cannot read Input file", err)

		numsArr, size := source.CalcAverage(string(buffer))
		source.CalcMedian(numsArr, size)
		source.CalcVarianceAndStdDiv(numsArr, size)

		fmt.Printf("Average: %v\nMedian: %v\nVariance: %v\nStandard Deviation: %v\n", math.Round(source.Average), math.Round(source.Median), int(math.Round(source.Variance)), math.Round(source.StdDeviation))

	} else {
		fmt.Println("wrong usage")
	}
}
ahabchi
yhamdoun
abouchik
hhouda
aelhadda
yhayyani