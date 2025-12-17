package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"structs-code/user"
	"time"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	u := user.User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}

	// u.ClearDetails()
	// user2 := User{
	// 	firstName,
	// 	lastName,
	// 	birthdate,
	// 	time.Now(),
	// }
	// fmt.Println(user)
	// fmt.Println(user2)
	u.OutputUserDetails()
	adult := user.Adult{
		u,
		18,
	}
	fmt.Println(adult.Age)
	jsonData, _ := json.Marshal(adult)
	fmt.Println(string(jsonData))
	err := os.WriteFile("hello.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
	err = readFromJsonFile()

}
func readFromJsonFile() error {
	bytearr, err := os.ReadFile("hello.json")
	if err != nil {
		return errors.New("error reading file")
	}
	var a user.Adult
	err = json.Unmarshal(bytearr, &a)
	if err != nil {
		return errors.New("error while unmarshalling file")
	}
	fmt.Println(a)
	return nil

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
