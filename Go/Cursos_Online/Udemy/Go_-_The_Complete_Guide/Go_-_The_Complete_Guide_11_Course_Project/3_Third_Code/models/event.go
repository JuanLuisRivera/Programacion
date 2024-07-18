package models

import (
	"time"

	"example.com/Course_Project/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"` // We use the "binding" keyword to specify the data is required else an error ocurrs
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	query := `
	INSERT INTO event(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	` // By adding the "?" symbol we create nameholders that prevent sql injection atacks
	stm, err := db.DB.Prepare(query) // We use the function "Prepare" so the query is stored in memory and reused for different values
	if err != nil {
		return err
	}
	defer stm.Close() // As we are closing the sql statement the is no potential advantage of using the prepare function, but will be keep there as a way to use it

	result, err := stm.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // The function "Exec" is used as there is a need to change the data in a database

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM event"
	rows, err := db.DB.Query(query) // There is no need to use the function "Prepare" as the instruction is easy and we do not use "Exec" as we do not want to change data on the database

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
