package main

import (
	//"fmt"
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"time"
	"strings"
	"github.com/lilakurse/learning-golang/model/mongo"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "WELCOME") // write data to response
}

var mongoconn *mongo.MongoDB

func registration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("C:/Users/Samsung/dev/go/src/github.com/lilakurse/learning-golang/model/blog/tmpl/registration" + ".html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of registration
		name := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		//birthday := r.FormValue("birthday")
		fmt.Println("username:", r.Form["name"])
		fmt.Println("password:", r.Form["password"])

		mongoconn.Add_user(name, password, email)

	}
}
func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/registration", registration)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	mongoconn=mongo.Mongoconn()
	mongoconn.Connection("127.0.0.1")
	defer mongoconn.Stop()

}
