package service

import (
	db "github.com/SJ22032003/go-ems/db"
	model "github.com/SJ22032003/go-ems/models"
)


func CreateEventService(event *model.Event) error {
	query := `INSERT INTO events (title, description, location, date, user_id) VALUES (?, ?, ?, ?, ?)`
	result, err := db.Execute(query, event.Title, event.Description, event.Location, event.Date, event.UserId)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err == nil {
		event.ID = id
	}
	return err
}

func GetAllEventsService() ([]model.Event, error) {
	var events = []model.Event{}
	
	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event model.Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.Date, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil

}

func GetEventByIdService(id int64, user_id int64) (*model.Event, error) {
	var event model.Event

	row := db.DB.QueryRow("SELECT * FROM events WHERE id = ? AND user_id LIMIT 1", id, user_id)

	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.Date, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil

}

func UpdateEventService(event *model.Event) error {
	query := `UPDATE events SET title = ?, description = ?, location = ?, date = ? WHERE id = ?`
	_, err := db.Execute(query, event.Title, event.Description, event.Location, event.Date, event.ID)
	return err
}

func DeleteEventByIdService(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := db.Execute(query, id)
	return err
}
