package user

import (
	"fmt"
	"testing"
	"time"
	"encoding/json"
)

func TestUser(t *testing.T) {
	u := New()
	fmt.Print("User Name: ", u.Name, "\n")
	if u.Name != "Lila" {
		t.Error("bad user name")
	}
}

func TestUseTwo(t *testing.T) {
	u := New()
	fmt.Print("User Name: ", u.Name, "\n")
	//fmt.Print("User Age: ", u.Age, "\n")
	err := u.Validate()
	if err != nil {
		t.Error(err)
	}
}

type testpair struct {
	values      string
	wanted      string
	isError     bool
	contactType ContactType
}

var goodUser = []byte(`{
	"name":"Lila","registration":"2015-04-23T11:58:34.841Z",
	"contacts":[{"v":"test@mail.ru","t":2},
				 {"v" :"555-99-69","t":1}]
	}`)

func TestUserJson (t *testing.T) {
	u,err := getUserFromJson()
	if err != nil {
		t.Error(err)
	}
	err = u.Validate()
	if err != nil {
		t.Error(err)
	}
	answer,err := json.Marshal(u)
	if err != nil {
		t.Error(err)
	}
	fmt.Print(string(answer))
	/*if string(answer) != string(goodUser) {
		t.Error("Must be equals")
	}*/
}

func getUserFromJson() (*User,error) {
	u := &User{}
	err := json.Unmarshal(goodUser,u)
	if err != nil {
		return nil,err
	}
	return u,nil
}

func getUserWithBirthday () *User {
	u := New()
	b := time.Now().AddDate(-10,0,0)
	u.Birthday = &b
	return u
}


var tests = []testpair{
	{"gmail.com", "error my", true, ContactTypeEmail},
	{"123@gmail.com", "123@gmail.com", false, ContactTypeEmail},
	{"789-90-89", "789-90-89", false, ContactTypePhone},
	{"998-12-12", "998-12-12", false, ContactTypePhone},
	{"Obrucheva13", "Obrucheva13", true, ContactTypeAddress},
	{"aab_c", "aab_c", false, ContactTypeIms},
}

func TestContact(t *testing.T) {
	var (
		v   *Contact
		err error
	)
	u := New()
	for _, pair := range tests {
		switch pair.contactType {

		case ContactTypeIms:
			v, err = u.AddIms(pair.values)
		case ContactTypePhone:
			v, err = u.AddPhone(pair.values)
		case ContactTypeEmail:
			v, err = u.AddEmail(pair.values)
		case ContactTypeAddress:
			v, err = u.AddAddress(pair.values)
		}
		if pair.isError {
			if v != nil && pair.wanted != err.Error() {
				t.Errorf("Incorrect %s for %s expected %s got %s", pair.contactType, pair.values, pair.wanted, err)
			}
		} else {
			if (v == nil && err != nil) || (v != nil && v.Value != pair.wanted && u.Contacts.IsExist(v)) {
				t.Errorf("Incorrect %s for %s expected %s got %s", pair.contactType, pair.values, pair.wanted, err)
			}
		}
	}

}
