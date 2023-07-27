package main 

import (
	"golang.org/x/net/http2"
	"net/http"
	"fmt"
)

func indexx(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Worked!")
}

func main(){
	fmt.Println("running...")

	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/", indexx)

	http2.ConfigureServer(&server, &http2.Server{})

	server.ListenAndServeTLS("cert.pem", "key.pem")
}