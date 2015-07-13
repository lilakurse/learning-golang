package mongostore

import (
	"github.com/lilakurse/learning-golang/model/user"
)

var userColl = "users"

type userStore struct {
	*mongoStore
}

func newUserStore(base *mongoStore) *userStore {
	base.collName = userColl
	return &userStore{base}
}

func (u *userStore) GetUsers(filter *user.UserFilter) (*[]*user.User, error) {
	users := &[]*user.User{}
	//use filter to bson.M
	err := u.Collection().Find(nil).All(users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
