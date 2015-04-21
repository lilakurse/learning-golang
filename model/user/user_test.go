package user
import (
	"testing"
	"fmt"
)

func TestUser(t *testing.T) {
 	u := New()
	fmt.Print("User Name: ",u.Name,"\n")
	if u.Name != "Lstatus 1ila" {
		t.Error("bad user name")
	}
}

func TestUseTwo(t *testing.T) {
	u := New()
	fmt.Print("User Name: ",u.Name,"\n")
	if u.Name != "Lstatus 1ila" {
		t.Error("bad user name")
	}
}