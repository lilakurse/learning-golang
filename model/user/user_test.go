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
/*
func TestContact(t *testing.T) {
	u := New()
	fmt.Print("User Data",u.Contacts,"\n")
	err :=u.Validate()
	if err !=nil{
		t.Error(err)}

}
*/