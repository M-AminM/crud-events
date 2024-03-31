package models

import (
	"crud-events/db"

	_ "github.com/go-sql-driver/mysql"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
}

func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(eventId int64) (*Event, error) {
	row := db.DB.QueryRow("SELECT * from events where id=?", eventId)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description)
	if err != nil {
		return nil, err
	}

	return &event, err
}

func CreateNewEvent(event Event) error {
	insert, err := db.DB.Query(
		"INSERT INTO events (name,description) VALUES (?,?)",
		event.Name, event.Description)

	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

func DeleteEvent(id int64) error {
	_, err := db.DB.Query("DELETE FROM events where id=?", id)
	if err != nil {
		return err
	}

	return nil
}
