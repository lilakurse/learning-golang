package user

import (
	"errors"
	"time"
)

//test user type
type User struct {
	Name         string `json:"username"`
	Registration *time.Time
	Age          int
	Birthday     *time.Time
	About        string
	Contacts     *Contacts
}

const ErrorBad = "invalid name or age"

func (u *User) AddIms(ims string) (*Contact, error) {
	return u.Contacts.AddIms(ims)
}

func (u *User) AddPhone(phone string) (*Contact, error) {
	return u.Contacts.AddPhone(phone)
}

func (u *User) AddEmail(email string) (*Contact, error) {
	return u.Contacts.AddEmail(email)
}

func (u *User) AddAddress(address string) (*Contact, error) {
	return u.Contacts.AddAddress(address)
}

func (u *User) Validate() error {
	if len(u.Name) <= 2 || u.Age < 18 {
		return errors.New(ErrorBad)
	}
	return nil
}
func New() *User {
	u := new(User)
	u.Name = "Lila"
	u.Age = 20
	u.Contacts = new(Contacts)
	return u
}
