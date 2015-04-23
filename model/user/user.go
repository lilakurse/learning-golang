package user

import (
	"errors"
	"time"
)

//test user type
type User struct {
	Name         string `json:"name"`
	Registration *time.Time `json:"registration"`
	Birthday     *time.Time `json:"birthday,omitempty"`
	About        string `json:"about,omitempty"`
	Contacts     *Contacts `json:"contacts"`
}

const ErrorBad = "invalid name"

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


func (u *User)HowAreYouOld() int{
	if u.Birthday == nil {
		return 0
	}
	now := time.Now()
	years := now.Year() - u.Birthday.Year()
	if now.YearDay() < u.Birthday.YearDay(){
		years--
	}
	return years
}

func (u *User) Validate() error {
	if len(u.Name) <= 2  {
		return errors.New(ErrorBad)
	}

	return u.Contacts.Validate()
}

func New() *User {
	u := new(User)
	u.Name = "Lila"
	u.Contacts = new(Contacts)
	return u
}
