// package models

// import (
// 	"time"

// 	"example.com/rest-api/db"
// )

// type Event struct {
// 	ID          int64
// 	Name        string    `binding:"required"`
// 	Description string    `binding:"required"`
// 	Location    string    `binding:"required"`
// 	DateTime    time.Time `binding:"required"`
// 	UserID      int64
// }

// var events = []Event{}

// func (e *Event) Save() error {
// 	query := `
// 	INSERT INTO events(name, description, location, dateTime, user_id)
// 	VALUES (?, ?, ?, ?, ?)`
// 	stmt, err := db.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
// 	if err != nil {
// 		return err
// 	}
// 	id, err := result.LastInsertId()
// 	e.ID = id
// 	return err
// }

// func GetAllEvents() ([]Event, error) {
// 	query := "SELECT * FROM events"
// 	rows, err := db.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var events []Event

// 	for rows.Next() {
// 		var event Event
// 		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

// 		if err != nil {
// 			return nil, err
// 		}

// 		events = append(events, event)
// 	}

// 	return events, nil
// }

// func GetEventByID(id int64) (*Event, error) {
// 	query := "SELECT * FROM events WHERE id = ?"
// 	row := db.DB.QueryRow(query, id)

// 	var event Event
// 	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &event, nil
// }

// func (event Event) Update() error {
// 	query := `
// 	UPDATE events
// 	SET name = ?, description = ?, location = ?, dateTime = ?
// 	WHERE id = ?
// 	`
// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
// 	return err
// }

// func (event Event) Delete() error {
// 	query := "DELETE FROM events WHERE id = ?"
// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(event.ID)
// 	return err
// }

// func (e Event) Register(userId int64) error {
// 	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(e.ID, userId)

// 	return err
// }

// func (e Event) CancelRegistration(userId int64) error {
// 	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(e.ID, userId)

// 	return err
// }

package models

import (
	"time"

	"example.com/rest-api/db"
)

// Event represents the structure of an event in the system.
// @Description Event struct for the event management system
// @Property ID          int64     `json:"id"`
// @Property Name        string    `json:"name" binding:"required"`
// @Property Description string    `json:"description" binding:"required"`
// @Property Location    string    `json:"location" binding:"required"`
// @Property DateTime    time.Time `json:"dateTime" binding:"required"`
// @Property UserID      int64     `json:"userId"`
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var id int64
	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&id)
	if err != nil {
		return err
	}

	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = $1, description = $2, location = $3, dateTime = $4
	WHERE id = $5
	`
	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = $1"
	_, err := db.DB.Exec(query, event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = $1 AND user_id = $2"
	_, err := db.DB.Exec(query, e.ID, userId)
	return err
}
