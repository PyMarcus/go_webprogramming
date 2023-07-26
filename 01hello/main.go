package main 

import (
	"fmt"
	"net/http"
)

func handle(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Hello, %s", request.URL.Path[1:])
	fmt.Println("Receiving connection from ", request.RemoteAddr)
	fmt.Println("Connection method: ", request.Method)
}

func main(){
	//routes
	http.HandleFunc("/", handle)
	fmt.Println("Running at 8080")
	http.ListenAndServe(":8080", nil)
}