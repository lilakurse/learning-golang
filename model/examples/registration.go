package main

import (
//"fmt"
	"html/template"
	"log"
	"net/http"
	"fmt"
)

func registration(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("C:/Users/Samsung/dev/go/src/github.com/lilakurse/learning-golang/model/examples/tmpl/registration"+ ".html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of registration
		fmt.Println("username:", r.Form["name"])
		fmt.Println("email:", r.Form["email"])
	}
}
func main() {
	http.HandleFunc("/registration", registration)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}