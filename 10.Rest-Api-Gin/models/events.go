package models

import (
	"fmt"
	"rest-api/db"
	"strconv"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"datetime"`
	UserId      int       `json:"userId"`
}

var events []Event

// func (E *Event) Save() error {
// 	query := `INSERT INTO  Event(Name,Description,Location,DateTime,UserId)
// 	 VALUES(?, ?, ?, ?, ?),`
// 	stmt, err := db.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	result, err := stmt.Exec(E.Name, E.Description, E.Location, E.DateTime, E.UserId)
// 	if err != nil {
// 		return err
// 	}
// 	id, err := result.LastInsertId()
// 	E.ID = id

// 	return err

// }
func (E *Event) Save() error {
	res, err := db.DB.Exec(`INSERT INTO  Event(Name,Description,Location,DateTime,UserId) 
	 VALUES(?, ?, ?, ?, ?)`, E.Name, E.Description, E.Location, E.DateTime, E.UserId)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	E.ID = id
	return err

}
func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM Event`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
		)
		if err != nil {
			return nil, err

		}
		events = append(events, event)

	}

	return events, nil
}

func (e Event) GetEventById(id string) (Event, error) {
	idval, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return Event{}, err
	}
	var event Event
	err = db.DB.QueryRow(`SELECT * FROM Event where ID=?`, idval).Scan(&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId)
	fmt.Println(event)
	if err != nil {
		return Event{}, err
	}

	return event, nil

}
