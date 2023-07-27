package main

// chaining two functions

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "hello, world")
}

func log(f http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request){
		functionCalledNamed := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Println("Function called: ", functionCalledNamed)
		f(writer, request)
	}
}

// func main(){
// 	server := http.Server{
// 		Addr: "localhost:8080",
// 	}

// 	http.HandleFunc("/", log(hello))

// 	server.ListenAndServe()
// }