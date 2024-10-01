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
	UserID int64
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

func GetEventById(eventId int64) (*Event, error) {
	query := "SELECT * FROM event WHERE id = ?"
	row := db.DB.QueryRow(query, eventId)

	var event Event
	err := row.Scan(
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

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE event
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := `
	DELETE FROM event WHERE id = ?
	`
	_, err := db.DB.Exec(query, event.ID)

	return err
}

func (event Event) Register(userId int64) error {
	query := "INSERT INTO registration(user_id, event_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, event.ID)
	return err
}

func (event Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registration WHERE user_id = ? AND event_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, event.ID)

	return err
}