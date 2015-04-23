package user

import (
	"errors"
	"regexp"

	"fmt"

	"strings"
)

type ContactType int

const (
	ContactTypeUndefined ContactType = iota
	ContactTypeIms
	ContactTypePhone
	ContactTypeEmail
	ContactTypeAddress
)

type Contact struct {
	Value string      `json:"v"`
	Type  ContactType `json:"t"`
}

type Contacts []*Contact

var contactTypes = []string{"undefined", "ims", "phone", "email", "address"}

func (c ContactType) String() string {
	return contactTypes[int(c)]
}
//поиск строки и выставление ее индекса
func (c ContactType) lookFor(str string) ContactType {
	for index, ct := range contactTypes {
		if ct == str {
			fmt.Println(ct," = ",str)
			return ContactType(index)
		}
	}
	return ContactTypeUndefined
}
//вытаскиваем значение в string и ищем его в массиве. если нету ->ошибка
func (c *ContactType)UnmarshalJSON(data []byte) error{
	*c=c.lookFor(strings.Trim(string(data),`"`))
	if *c==ContactTypeUndefined{
		return errors.New("contact type not found")
	}
	return nil
}
func (c *ContactType)MarshalJSON() ([]byte,error){
	return []byte(`"`+c.String()+`"`),nil
}
//в массив контакты добавляем элменты-контакты
func (c *Contacts) Add(t *Contact) {
	*c = append(*c, t)
}

func (c *Contacts) AddIms(ims string) (*Contact, error) {
	return c.addType(ims, ContactTypeIms)
}
func (c *Contacts) AddPhone(phone string) (*Contact, error) {
	return c.addType(phone, ContactTypePhone)
}

func (c *Contacts) AddEmail(email string) (*Contact, error) {
	return c.addType(email, ContactTypeEmail)
}

func (c *Contacts) AddAddress(address string) (*Contact, error) {
	return c.addType(address, ContactTypeAddress)
}

func (c *Contacts) addType(val string, someType ContactType) (*Contact, error) {
	i := new(Contact)
	i.Type = someType
	i.Value = val
	err := i.Validate()
	if err != nil {
		return nil, err
	}
	if c.IsExist(i) {
		return nil, errors.New("incorrect value for " + someType.String())
	}
	c.Add(i)
	return i, nil
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

var Ims = regexp.MustCompile(`^\w+$`)
var Email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+$`)
var Phone = regexp.MustCompile(`^[\d\+]{1}[\d\-\(\)]+$`)
var Address = regexp.MustCompile(`^[a-z]+\,+\s+[\d\+]+.+[0-9]`)

var err []error

func (c *Contacts) Validate() error {
	var err error
	for _, cc := range *c {
		err = cc.Validate()
		if err != nil {
			return errors.New("Error in " + cc.Value + " (" + cc.Type.String() + ") : " + err.Error())
		}
	}
	return nil
}

func (c *Contact) Validate() error {
	switch c.Type {
	case ContactTypeIms:
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
