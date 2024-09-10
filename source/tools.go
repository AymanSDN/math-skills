package source

import (
	"log"
)

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatalln(msg)
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

// func DownloadHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/download" {
// 		http.Error(w, "404 Not Found", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	w.Header().Set("Content-Disposition", "attachment; filename=filexd.txt")

// 	fileContent := "This is the content of filexd.txt"

// 	_, err := w.Write([]byte(fileContent))
// 	if err != nil {
// 		http.Error(w, "Error generating file", http.StatusInternalServerError)
// 		return
// 	}
// }
