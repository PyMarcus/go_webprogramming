package main

import (
	"fmt"
	"time"
	"html/template"
	"math/rand"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Println("Index")
	t, _ := template.ParseFiles("templates/index.html", "templates/to_include.html")
	// t.Execute(w, "Hi")
	rand.Seed(time.Now().Unix())
	// t.Execute(w, rand.Intn(10) > 5)
	fmt.Println("Processing...")
	v := t.Execute(w, "OK")
	fmt.Println(v)
}

func main(){
	server := &http.Server{
		Addr: ":8080",
	}	
	fmt.Println("Running")
	http.HandleFunc("/", index)
	server.ListenAndServe()
}