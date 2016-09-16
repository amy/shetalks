package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// initializes MySQL schema
func InitDB(db *sql.DB) {

	db, err := sql.Open("mysql", "amy:password@/shetalks")

	if err != nil {
		fmt.Println("Database opening error")
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Unable to open connection: %v", err)
	}

	initSpeakerEventPivot(db)
	initOrganizationEventPivot(db)
	initOrganizationSpeakerPivot(db)
	initEventTable(db)
	initSpeakerTable(db)
	initOrganizationTable(db)

	db.Close()

}

func initSpeakerEventPivot(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE speaker_event (
			speaker INT,
			event INT 
		)`,
	)

	if err != nil {
		fmt.Println("Error in creating speaker_event table: %v", err)
		// @TODO Log error
	}
}

func initOrganizationEventPivot(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE organization_event (
			organization INT,
			event INT
		)`,
	)

	if err != nil {
		fmt.Println("Error in creating organization_event table: %v", err)
		// @TODO Log error
	}
}

func initOrganizationSpeakerPivot(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE organization_speaker (
			organization INT,
			speaker INT
		)`,
	)

	if err != nil {
		fmt.Println("Error in creating organization_speaker table: %v", err)
		// @TODO Log error
	}
}

func initEventTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE events (
			id INT AUTO_INCREMENT primary key, 
			name VARCHAR(100), 
			description VARCHAR(400)
		)`,
	)

	if err != nil {
		fmt.Println("Error in creating event table: %v", err)
		// @TODO Log error
	}

}

func initOrganizationTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE organizations (
			id INT AUTO_INCREMENT primary key, 
			name VARCHAR(100), 
			description VARCHAR(400)
		)`,
	)

	if err != nil {
		fmt.Println("Error in creating organization table: %v", err)
		// @TODO Log error
	}
}

func initSpeakerTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE speakers (
			id INT AUTO_INCREMENT primary key, 
			name VARCHAR(100), 
			description VARCHAR(400)
		)`,
	)

	if err != nil {
		fmt.Println("Error in creating speaker table: %v", err)
		// @TODO Log error
	}
}
