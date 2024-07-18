package models

import "time"

type Event struct {
	Id          int
	Name        string    `binding:"required"` // We use the "binding" keyword to specify the data is required else an error ocurrs
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
