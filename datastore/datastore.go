package datastore

import (
	"github.com/lilakurse/learning-golang/model/user"
)

type DataStore interface{
	CreateUser(u *user.User) error
	GetBadId(Id string)(*user.User,error)
}

func CreateUser(db DataStore, u *user.User) error {
	return db.CreateUser(user)
}

func GetBadId(db DataStore, ID string) (*user.User, error){
	return db.GetBadId(ID)
}