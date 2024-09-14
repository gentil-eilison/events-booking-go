package models

import (
	"time"

	"github.com/gentil-eilison/events-booking-go/db"
)

type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int
}

var events = []Event{}

func (event *Event) Save() error {
	createEventQuery := `
	INSERT INTO event(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?);
	`
	stmt, err := db.DB.Prepare(createEventQuery)
	
	if err != nil {
		return err
	}

	/*
		Must come after the err check because there's no point in closing
		A connection which hasn't been opened
	*/
	defer stmt.Close()

	result, err := stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserID,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id
	events = append(events, *event)
	return err
}

func GetAllEvents() ([]Event, error) {
	var events []Event
	
	getEventsQuery := "SELECT * FROM event;"
	rows, err := db.DB.Query(getEventsQuery)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	

	return events, nil
}