package database

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"mindenairport/models"
	"strconv"

	"github.com/godror/godror"
)

func (db Database) GetFlightStatuses() []models.FlightStatus {
	stmt, err := db.Prepare(`
	BEGIN 
	GetFlightStatuses(:1); END;
	`)
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}

	var cursor driver.Rows
	_, err = stmt.Exec(sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Failed to execute prepared statement:", err)
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var flightStatuses []models.FlightStatus

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var flightStatus models.FlightStatus
		flightStatus.ID, _ = strconv.Atoi(r[0].(godror.Number).String())
		flightStatus.Name = r[1].(string)
		if r[2] != nil {
			flightStatus.Description = r[2].(string)
		}
		flightStatuses = append(flightStatuses, flightStatus)
	}

	return flightStatuses
}

func (db Database) GetFlightStatusByID(id int) models.FlightStatus {
	stmt, err := db.Prepare(`
	BEGIN 
	GetFlightStatusByID(:1, :2); END;
	`)
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}

	var cursor driver.Rows
	_, err = stmt.Exec(id, sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Failed to execute prepared statement:", err)
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var flightStatus models.FlightStatus

	err = cursor.Next(r)
	if err == nil {
		flightStatus.ID, _ = strconv.Atoi(r[0].(godror.Number).String())
		flightStatus.Name = r[1].(string)
		if r[2] != nil {
			flightStatus.Description = r[2].(string)
		}
	}

	return flightStatus
}
