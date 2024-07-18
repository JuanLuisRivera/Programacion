package models

import (
	"time"

	"example.com/Course_Project/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e Event) Cancel(userID int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(e.Id, userID)
	return err
}

func (e Event) Register(userID int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(e.Id, userID)
	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM event WHERE id = ?"
	stm, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(e.Id)
	return err
}

func (e Event) Update() error {
	query := `
	UPDATE event
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stm, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)
	return err
}

func (e *Event) Save() error {
	query := `
	INSERT INTO event(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()

	result, err := stm.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM event"
	rows, err := db.DB.Query(query)

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

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM event WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
