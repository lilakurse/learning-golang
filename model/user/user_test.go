package user

import (
	"fmt"
	"testing"
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
	fmt.Print("User Age: ", u.Age, "\n")
	err := u.Validate()
	if err != nil {
		t.Error(err)
	}
}

type testpair struct {
	values string
	wanted string
	isError bool
	contactType ContactType
}
var tests = []testpair{
	{"gmail.com","error my",true,ContactTypeEmail},
	{"123@gmail.com","123@gmail.com",false,ContactTypeEmail},
	{"789-99-89","789-90-89",true,ContactTypePhone},
	{"Obrucheva13","Obrucheva13",true,ContactTypeAddress},
	{"hjj","hjj",false,ContactTypeIms},
}

func TestContact(t *testing.T) {
  var  (
	  v *Contact
	  err error
  )
	u:=New()
	for _, pair :=range tests{
		switch pair.contactType {
			case ContactTypeEmail:
				v,err =u.AddEmail(pair.values)
			case ContactTypePhone:
				v,err =u.AddPhone(pair.values)
			case ContactTypeAddress:
				v,err=u.AddAddress(pair.values)
			case ContactTypeIms:
				v,err =u.AddIms(pair.values)
		}
		if pair.isError {
			if v != nil && pair.wanted != err.Error() {
				t.Errorf("Incorrect %s for %s expected %s got %s",pair.contactType, pair.values, pair.wanted, err)
			}
		} else {
			if (v == nil && err != nil) || (v != nil && v.Value != pair.wanted && u.Contacts.IsExist(v)) {
				t.Errorf("Incorrect %s for %s expected %s got %s", pair.values, pair.wanted,err)
			}
		}
	}

}
