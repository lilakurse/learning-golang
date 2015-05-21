package main

import (
	//"fmt"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
	"time"
)

func registration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("C:/Users/Samsung/dev/go/src/github.com/lilakurse/learning-golang/model/examples/tmpl/registration" + ".html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of registration
		fmt.Println("username:", r.Form["name"])
		fmt.Println("email:", r.Form["email"])
	}
}

type Person struct {
	Name     string
	Phone    string
	email    string
	birthday *time.Time
}

func main() {
	http.HandleFunc("/registration", registration)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//work with mgo, open session
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("data").C("people")
	err = c.Insert()

}
