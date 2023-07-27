package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "OK")
}

// func main() {
// 	server := http.Server{
// 		Addr:    "127.0.0.1:8080",
// 	}

// 	http.HandleFunc("/", index)

// 	server.ListenAndServe()
// }
