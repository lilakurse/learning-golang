package main

import (
	"encoding/json"
	"flag"
	u "github.com/lilakurse/learning-golang/model/user"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	whereToConnect := flag.String("connection", "mongodb://localhost", "used in mgo.Dial()")
	dbToUse := flag.String("db", "myusers", "used in session.DB(...).C(...)")
	collectionToUse := flag.String("collection", "users", "used in session.DB(...).C(...)")
	pathToJSON := flag.String("path", "tools/mongo/users.json", "used in path, _ := filepath.Abs(...)")

	flag.Parse()

	users := []*u.User{}
	path, _ := filepath.Abs(*pathToJSON)
	file, err := os.Open(path)
	check(err)

	err = json.NewDecoder(file).Decode(&users)

	session, err := mgo.Dial(*whereToConnect)
	check(err)
	defer session.Close()

	c := session.DB(*dbToUse).C(*collectionToUse)
	for _, user := range users {
		err = c.Insert(user)
		if err != nil {
			log.Fatal(err)
		}
	}
}
