package datastore

import (
	"github.com/lilakurse/learning-golang/model/user"
	"golang.org/x/net/context"
)

type DataStore interface {
	CreateUser(u *user.User) error
	GetUserByID(ID string) (*user.User, error)
}

func CreateUser(ctx context.Context, u *user.User) error {
	return FromContext(ctx).CreateUser(u)
}

func GetUserByID(ctx context.Context, ID string) (*user.User, error) {
	return FromContext(ctx).GetUserByID(ID)
}

/*func CreateUser(db DataStore, u *user.User) error {
	return db.CreateUser(user)
}

func GetUserByID(db DataStore, ID string) (*user.User, error) {
	return db.GetUserByID(ID)
}*/
