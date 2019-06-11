package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)
var pathstring string = ""
func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	pathstring = exPath
	fmt.Println(exPath)
	fmt.Println("jacobgasser's webserver is now online!")
	http.Handle("/", http.FileServer(http.Dir(pathstring + "/src")))
	http.HandleFunc("/help", index)
	http.ListenAndServe(":80", nil)

}



func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(pathstring)
	http.ServeFile(w, r, pathstring + "/src/index.html")
}


