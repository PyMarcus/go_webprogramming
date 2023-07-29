package main

import (
	"encoding/json"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request){
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello World</h1></body>
	</html>`
	w.Write([]byte(str))
}

func wExample(w http.ResponseWriter, r *http.Request){
	
	w.WriteHeader(666)
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>ERROR</h1></body>
	</html>`
	w.Write([]byte(str))
}

func redirect(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("location", "www.google.com")
	w.WriteHeader(302)
}

type jsonStruct struct{
	User uint64 
	Name string
}

func jsonResponse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	js := &jsonStruct{
		User: 31223,
		Name: "fulano",
	}
	j, _ := json.Marshal(js)
	w.Write(j)
}


func main(){
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", root)
	http.HandleFunc("/error", wExample)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/json", jsonResponse)
	server.ListenAndServe()
}
