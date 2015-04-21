package user
//test user type
type User struct {
	Name string `json:"username"`
}

func New () *User {
	u := new(User)
	u.Name = "Lola"
	return u
}