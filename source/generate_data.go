package source

import (
	"log"
	"math/rand"
	"os"
	"strconv"
)

func GenerateData(maxValue int) {
	var result string
	var size uint
	for size == 0 {
		size = uint(rand.Intn(2000))
	}
	for i := 0; i < int(size); i++ {
		result += strconv.Itoa(rand.Intn(maxValue))
		if i+1 != int(size) {
			result += "\n"
		}
	}
	if err := os.WriteFile("data.txt", []byte(result), 0o666); err != nil {
		log.Fatal(err)
	}
	// print(size)
}
