package main 

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func index(writer http.ResponseWriter, request *http.Request){
	// urlInfo := request.URL
	// header := request.Header
	fmt.Println(request.Body)
	fmt.Println(request.Method)
	x, _ := ioutil.ReadAll(request.Body)
	fmt.Println(string(x))
	fmt.Fprint(writer, string(x))
	writer.Write(x)
}

// func main(){
// 	server := http.Server{
// 		Addr: "localhost:8080",
// 	}

// 	http.HandleFunc("/", index)
// 	server.ListenAndServe()
// }