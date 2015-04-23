package user

import (
	"errors"
	"time"
)

//test user type
type User struct {
	Name         string `json:"username"`
	Registration *time.Time
	Birthday     *time.Time
	About        string
	Contacts     *Contacts
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

func (u *User)IsExist(b *time.Time) bool{
		if u.Birthday !=nil && u.Birthday==b{
			return true
		}
	return false
}

func (u *User)HowAreYouOld(b *time.Time) (int,error){
	if u.IsExist(b){
		return 0,errors.New("no birthday date")
	}
	now := time.Now()
	years := now.Year() - b.Year()
	if now.YearDay() < b.YearDay(){
		years--
	}
	return years,nil
}

func (u *User) Validate() error {
	if len(u.Name) <= 2  {
		return errors.New(ErrorBad)
	}
	return nil
}

func New() *User {
	u := new(User)
	u.Name = "Lila"
	u.Contacts = new(Contacts)
	return u
}
