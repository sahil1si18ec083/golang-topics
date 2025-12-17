package user

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string    `json:"fname"`
	LastName  string    `json:"lname"`
	Birthdate string    `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}
type Adult struct {
	User
	Age int
}

func (user User) OutputUserDetails() {
	fmt.Print(user.FirstName, user.LastName, user.Birthdate)
}
func (user *User) ClearDetails() {
	user.FirstName = ""
	user.LastName = ""
	user.Birthdate = ""
}
