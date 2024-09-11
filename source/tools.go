package source

import (
	"log"
)

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func RemoveEmptyStr(arr []string) []string {
	var result []string

	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			result = append(result, arr[i])
		}
	}
	return result
}
