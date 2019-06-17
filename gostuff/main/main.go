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
	port := "80"
	if os.Args[0] != "" {
		port = os.Args[1]
	}
	exPath := filepath.Dir(ex)
	pathstring = exPath
	fmt.Println(exPath)
	fmt.Println("jacobgasser's webserver is now online!")
	http.Handle("/", http.FileServer(http.Dir(pathstring+"/src")))
	http.HandleFunc("/help", index)
	fmt.Println("Starting on port " + port)
	errr := http.ListenAndServe(":"+port, nil)
	if errr != nil {
		panic(errr)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(pathstring)
	http.ServeFile(w, r, pathstring+"/src/index.html")
}
