package mock

import (
	"github.com/lilakurse/learning-golang/model/user"
)

var (
	EmptyUser = &user.User{}
	GoodUser  = user.NewUserWithAllFields()
	BadUser   = getBadUserEmail()
	BadUserID = getBadUserId()
)

func getBadUserEmail() *user.User {
	u := user.New()
	u.Contacts.AddEmail("lila.gmail.com", user.ContactTypeEmail)
	return u
}

func getBadUserId() *user.User {
	u := user.NewUserWithAllFields()
	u.ID = "53700d9cd83e146623e6bfb4"
	return u
}
