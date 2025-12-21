package models

import (
	"fmt"
	"rest-api/db"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

func (U *User) Save() error {
	res, err := db.DB.Exec(`INSERT INTO User (Email, Password)
	values (?,?)
	
	`, U.Email, U.Password)
	fmt.Println(res)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	U.ID = id
	return err

}

func (U *User) GetUserByEmail(email string) error {
	fmt.Println("calling")

	err := db.DB.QueryRow(`Select id, email, password from  User where email= ?
	
	`, email).Scan(&U.ID, &U.Email, &U.Password)
	fmt.Print(err)
	if err != nil {
		return err
	}

	return nil

}
