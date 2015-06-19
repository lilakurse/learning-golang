package mongo

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//"time"
	//"fmt"
)

type MongoDB struct {
	session *mgo.Session
}

type User struct {
	username     string
	email    string
	//birthday *time.Time
	password string
}
func Mongoconn() * MongoDB{
	return &MongoDB{}
}

func(m *MongoDB) Connection(url string) *mgo.Session{
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	m.session = session
	return m.session
}

func (m *MongoDB) Stop(){
	m.session.Close()
}

func (m *MongoDB) Add_user(username, password, email string)(err error){
	m.session.SetMode(mgo.Monotonic, true)
	c :=m.session.DB("blog").C("blog")
	//insert
	err = c.Insert(&User{username,email,password})
	if err!=nil{
		panic(err)
	}
	//Query one
	/*result := User{}
	err = c.Find(bson.M{"name": "Andrew"}).Select(bson.M{"email": 0}).One(&result)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Email", result)*/
	return  nil
}