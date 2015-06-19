package main

import (
	//"fmt"
	"fmt"
//	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
	//"time"
	"strings"
	"github.com/lilakurse/learning-golang/model/mongo"
)
/*
var (
	IsDrop = true
)

type MongoDB struct {
	session *mgo.Session
}

type User struct {
	username string
	email    string
	password string
}

func Mongoconn() *MongoDB {
	return &MongoDB{}
}

func (m *MongoDB) Stop() {
	m.session.Close()
}
func (m *MongoDB) Connect() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	m.session = session
	return m.session
}
func (m *MongoDB) Add_user(username, password, email string) (err error) {
	if IsDrop {
		err = m.session.DB("data").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	//err = c.Insert(bson.M{"name":username,"password":password})  еще один способ сделать insert
	c:=m.session.DB("blog").C("blog")
	//c := m.session.DB("test").C("blog")
	err = c.Insert(&User{username, email, password})
	//err = c.Insert(bson.M{"name":username,"password":password,"email":email})
	if err != nil {
		panic(err)
	}
	return nil
}
*/
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
		t, _ := template.ParseFiles("C:/Users/Samsung/dev/go/src/github.com/lilakurse/learning-golang/model/blog/tmpl/login" + ".html")
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
		//mongoconn.Stop()
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
