package user

import (
	"errors"
	"regexp"
	"time"
)

//test user type
type ContactType int

const (
	ContactTypeIms ContactType = iota
	ContactTypePhone
	ContactTypeEmail
	ContactTypeAddress
)

var contactTypes = []string{"ims", "phone", "email", "address"}

type Contact struct {
	Value string
	Type  ContactType
}

type Contacts []*Contact

type User struct {
	Name         string `json:"username"`
	Registration *time.Time
	Age          int
	Birthday     *time.Time
	About        string
	Contacts     *Contacts
}

const ErrorBadName = "invalid name"
const ErrorAge = "invalid age value"

func (c ContactType) String() string {
	return contactTypes[int(c)]
}

//метод проверяющий существование контакта в Contacts
func (c *Contacts) IsExist(contact *Contact) bool {
	for _, cc := range *c {
		if cc.Type == contact.Type && cc.Value == contact.Value {
			return true
		}
	}
	return false
}
/*
func(c *Contact) Add(t *Contacts) {
	//t :=new(Contact)
	*(c.Contacts)= append(*(u.Contacts),t)
}
*/
func (u *User) AddIms(ims string) (*Contact, error) {
	i := new(Contact)
	i.Type = ContactTypeEmail
	i.Value = ims
	err := i.Validate()
	if err != nil {
		return nil, err
	}
	if u.Contacts.IsExist(i) {
		return nil, errors.New("")
	}
	//u.Contacts.Add(i)
	//*(u.Contacts) = append(* (u.Contacts),i)
	return i, nil
}
func (u *User) AddPhone(phone string) (*Contact, error) {
	p := new(Contact)
	p.Type = ContactTypePhone
	p.Value = phone
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	if u.Contacts.IsExist(p) {
		return nil, errors.New("")
	}
	//u.Contacts.Add(p)
	//*(u.Contacts) = append(*(u.Contacts),p)
	return p, nil
}
func (u *User) AddEmail(email string) (*Contact, error) {
	e := new(Contact)
	e.Type = ContactTypeEmail
	e.Value = email
	err := e.Validate()
	if err != nil {
		return nil, err
	}
	if u.Contacts.IsExist(e) {
		return nil, errors.New("")
	}
	//u.Contacts.Add(e)
	//*(u.Contacts) = append(*(u.Contacts),e)
	return e, nil
}

func (u *User) AddAddress(address string) (*Contact, error) {
	a := new(Contact)
	a.Type = ContactTypeAddress
	a.Value = address
	err := a.Validate()
	if err != nil {
		return nil, err
	}
	if u.Contacts.IsExist(a) {
		return nil, errors.New("")
	}
	//u.Contacts.Add(a)
	//*(u.Contacts) = append(*(u.Contacts),a)
	return a, nil
}

var Ims = regexp.MustCompile(`^[a-z]`)
var Email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+$`)
var Phone = regexp.MustCompile(`^[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]$`)
var Address = regexp.MustCompile(`^[a-z]+\,+\s+[0-9][0-9][0-9]+.+[0-9]`)

func (c *Contact) Validate() error {
	switch c.Type {
	case  ContactTypeIms:
		if !Ims.MatchString(c.Value) {
			return errors.New("incorrect ims")
		}
	case ContactTypePhone:
		if !Phone.MatchString(c.Value) {
			return errors.New("incorrect phone")
		}
	case ContactTypeEmail:
		if !Email.MatchString(c.Value) {
			return errors.New("incorrect email")
		}
	case ContactTypeAddress:
		if !Address.MatchString(c.Value) {
			return errors.New("incorrect address")
		}
	}
	return nil
}

func (u *User) ValidateUser() error {
	if len(u.Name) <= 2 {
		return errors.New(ErrorBadName)
	}
	if u.Age < 18 {
		return errors.New(ErrorAge)
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
