package mock

import (
	"github.com/lilakurse/learning-golang/model/user"
	//"gopkg.in/mgo.v2/bson"
)

var (
	EmptyUser = &user.User{}
	GoodUser  = user.NewUserWithAllFields()
	BadUser   = getBadUserEmail()
	BadUserID = getBadUserId()
)

func getBadUserEmail() *user.User {
	u := user.New()
	u.Contacts.AddEmail("lila.gmail.com")
	return u
}

func getBadUserId() *user.User {
	u := user.NewUserWithAllFields()
	u.ID = "53700d9cd83e146623e6bfb4"
	return u
}

/*UserExistID    = bson.ObjectIdHex(ExistingID)
UserNonExistID = bson.ObjectIdHex(NonexistentID)

ExistingID    = "5571545e1f72901f6b000003"
NonexistentID = "5571545e1f72901f6b000004"
*/
