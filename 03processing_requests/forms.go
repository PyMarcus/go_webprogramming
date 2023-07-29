package main 

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func process(writer http.ResponseWriter, request *http.Request){
	fmt.Println("calling1")
	request.ParseForm()
	fmt.Fprintln(writer, "(1)", request.FormValue("name"))
	fmt.Fprintln(writer, "(2)", request.PostFormValue("name"))
	fmt.Fprintln(writer, "(3)", request.PostForm)
	fmt.Fprintln(writer, "(4)", request.Form)
	fmt.Fprint(writer, request.Form.Get("name"), request.Form.Get("password"))
}

func process2(writer http.ResponseWriter, request *http.Request){
	fmt.Println("calling2")
	request.ParseMultipartForm(2048)
	request.ParseForm()
	fmt.Fprintln(writer, "(1)", request.FormValue("name"))
	fmt.Fprintln(writer, "(2)", request.PostFormValue("name"))
	fmt.Fprintln(writer, "(3)", request.PostForm)
	fmt.Fprintln(writer, "(4)", request.Form)
	fmt.Fprint(writer, request.Form.Get("name"), request.Form.Get("password"), request.MultipartForm)
	fmt.Println("reading file", request.MultipartForm.File["upload"])
	file := request.MultipartForm.File["upload"][0]
	f, e := file.Open()
	if e == nil{
		data, er := ioutil.ReadAll(f)
		if er == nil{
			fmt.Println(string(data))
		}
	}
}


func main(){
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	server.ListenAndServe()
}