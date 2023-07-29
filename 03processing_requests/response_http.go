package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func cookie(w http.ResponseWriter, r *http.Request){
	c1 := &http.Cookie{
		Name: "fcookie",
		Value: "322",
		HttpOnly: true,
	}

	c2 := &http.Cookie{
		Name: "2cookie",
		Value: "3223",
		HttpOnly: true,
		Secure: true,

	}

	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())

	fmt.Println(r.Header["Cookie"])
	fmt.Println(r.Cookies())
}

//flash messages in the cookies
func setMessage(w http.ResponseWriter, r *http.Request){
	c := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString([]byte("Hello, world!")),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request){
	c, _ := r.Cookie("flash")
	val, _ := base64.URLEncoding.DecodeString(c.Value)
	rc := http.Cookie{
		Name: "flash",
		MaxAge: -1, //delete cookie
		Expires: time.Unix(1, 0),
	}

	http.SetCookie(w, &rc)
	fmt.Fprint(w, string(val))
}


func main(){
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", root)
	http.HandleFunc("/error", wExample)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/json", jsonResponse)
	http.HandleFunc("/cookie", cookie)
	http.HandleFunc("/flash", showMessage)
	http.HandleFunc("/set_flash", setMessage)
	server.ListenAndServe()
}
