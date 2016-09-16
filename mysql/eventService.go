package mysql

import (
	"database/sql"
	"strconv"

	"github.com/amy/shetalks"
)

// EventService is a mysql implementation of shetalks.EventService
type EventService struct {
	DB *sql.DB
}

func (es *EventService) Create(e shetalks.Event) (shetalks.Event, error) {

	stmt, err := es.DB.Prepare("INSERT events (name, description) VALUES (?, ?)")
	if err != nil {
		return e, err
	}

	res, err := stmt.Exec(e.Name, e.Description)
	if err != nil {
		return e, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return e, err
	}

	stmt, err = es.DB.Prepare("INSERT INTO speaker_event (speaker, event) VALUES (?, ?)")
	if err != nil {
		return e, err
	}

	for _, speaker := range e.Speakers {
		_, err := stmt.Exec(strconv.Itoa(speaker), strconv.FormatInt(id, 10))
		if err != nil {
			return e, err
		}
	}

	/* Edge Case: insert into events table success. insert into speaker_event failure.
	Need to delete from events table. Look into how to execute query all at once? */

	return e, nil

}

func (es *EventService) Read(id int) (shetalks.Event, error) {

	var eventID int
	var name string
	var description string
	var speakers []int

	err := es.DB.QueryRow("SELECT * FROM events WHERE id=?", id).Scan(&eventID, &name, &description)
	if err != nil {
		return shetalks.Event{}, err
	}

	rows, err := es.DB.Query("SELECT * FROM speaker_event WHERE event=?", id)

	var speaker int
	for rows.Next() {
		err = rows.Scan(&speaker, &eventID)
		speakers = append(speakers, speaker)
	}

	return shetalks.Event{
		Name:        name,
		Description: description,
		Speakers:    speakers,
	}, nil

}

func (es *EventService) Update(id int, name string, description string, speakers []int) (shetalks.Event, error) {

	/* Flaw: Currently must always update name and description and speakers. Need to include
	   switch/if logic to use correct query. */

	stmt, err := es.DB.Prepare("UPDATE events SET name=?, description=? WHERE id=?")
	if err != nil {
		return shetalks.Event{}, err
	}

	_, err = stmt.Exec(name, description, id)
	if err != nil {
		return shetalks.Event{}, err
	}

	stmt, err = es.DB.Prepare("DELETE FROM speaker_event WHERE event=?")
	if err != nil {
		return shetalks.Event{}, err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return shetalks.Event{}, err
	}

	stmt, err = es.DB.Prepare("INSERT INTO speaker_event (speaker, event) VALUES (?, ?)")
	if err != nil {
		return shetalks.Event{}, err
	}

	for _, speaker := range speakers {
		_, err := stmt.Exec(speaker, id)
		if err != nil {
			return shetalks.Event{}, err
		}
	}

	return shetalks.Event{
		Name:        name,
		Description: description,
		Speakers:    speakers,
	}, nil

}

func (es *EventService) Delete(id int) error {

	stmt, err := es.DB.Prepare("DELETE FROM events WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	stmt, err = es.DB.Prepare("DELETE FROM speaker_event WHERE event=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
