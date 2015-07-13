package mongostore

import (
	"gopkg.in/mgo.v2"
	//	"github.com/lilakurse/learning-golang/model/user"
)

type mongoStore struct {
	*mgo.Session
	collName string
}

func NewMongoStore(uri string) interface{} {
	if len(uri) == 0 {
		uri = "mongodb://localhost/myusers"
	}

	sess, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}

	base := &mongoStore{sess, ""}

	return &struct {
		*userStore
	}{
		newUserStore(base),
	}
}

func (m *mongoStore) Collection() *mgo.Collection {
	return m.DB("").C(m.collName)
}
