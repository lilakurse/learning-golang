package user
import "errors"
//test user type
type User struct {
	Name string `json:"username"`
}

const ErrorBadName  = "invalid name"

func (u *User) Validate () error {
	if len(u.Name) <= 2 {
		return errors.New(ErrorBadName)
	}
	return nil
}

func New () *User {
	u := new(User)
	u.Name = "Lola"
	return u
}

