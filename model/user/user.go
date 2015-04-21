package user
import (
	"errors"
	"time"
)
//test user type

type ContactType int
const(
ContactTypeIms ContactType = iota
ContactTypePhone
ContactTypeEmail
ContactTypeAddress
)

type Contact struct{
	Value string
	Type ContactType
}

type User struct {
	Name string `json:"username"`
	Registration *time.Time
	Age int
	Birthday *time.Time
	About string
	Contacts *[]*Contact
}

const ErrorBadName  = "invalid name"
const ErrorAge= "invalid value"

func (u *User) Validate () error {
	if len(u.Name) <= 2 {
		return errors.New(ErrorBadName)
	}
	if u.Age <18  {
		return errors.New(ErrorAge)
	}
	return nil
}

func New () *User {
	u := new(User)
	u.Name = "Lola"
	return u
}

