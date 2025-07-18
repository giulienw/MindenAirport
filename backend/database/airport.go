package database

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"mindenairport/models"
	"strconv"

	"github.com/godror/godror"
)

func (db Database) GetAirports() []models.Airport {
	var query = `
	BEGIN 
		GetAllAirports(:1); 
	END;
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return []models.Airport{}
	}

	var cursor driver.Rows
	_, err = stmt.Exec(sql.Out{Dest: &cursor})
	if err != nil {
		log.Printf("Failed to execute prepared statement: %v", err)
		return []models.Airport{}
	}

	r := make([]driver.Value, len(cursor.Columns()))
	err = cursor.Next(r)
	if err != nil {
		log.Println(err) // column count mismatch: we have 10 columns, but given 0 destination
	}

	defer cursor.Close()

	var airports []models.Airport

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var airport models.Airport
		airport.ID = r[0].(string)
		airport.Name = r[1].(string)
		airport.Country = r[2].(string)
		airport.City = r[3].(string)
		airport.Timezone = r[4].(string)
		airport.Elevation, _ = strconv.ParseFloat(r[5].(godror.Number).String(), 64)
		airport.Latitude, _ = strconv.ParseFloat(r[7].(godror.Number).String(), 64)
		airport.Longitude, _ = strconv.ParseFloat(r[8].(godror.Number).String(), 64)
		airports = append(airports, airport)
	}

	return airports
}

func (db Database) GetAirportByID(id string) models.Airport {
	stmt, err := db.Prepare(`
	BEGIN 
	GetAirportByID(:1, :2); END;
	`)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return models.Airport{}
	}

	var cursor driver.Rows
	_, err = stmt.Exec(id, sql.Out{Dest: &cursor})
	if err != nil {
		log.Printf("Failed to execute prepared statement: %v", err)
		return models.Airport{}
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var airport models.Airport

	err = cursor.Next(r)
	if err == nil {
		airport.ID = r[0].(string)
		airport.Name = r[1].(string)
		airport.Country = r[2].(string)
		airport.City = r[3].(string)
		airport.Timezone = r[4].(string)
		airport.Elevation, _ = strconv.ParseFloat(r[5].(godror.Number).String(), 64)
		airport.NumberOfTerminal, _ = strconv.Atoi(r[6].(godror.Number).String())
		airport.Latitude, _ = strconv.ParseFloat(r[7].(godror.Number).String(), 64)
		airport.Longitude, _ = strconv.ParseFloat(r[8].(godror.Number).String(), 64)
	}

	return airport
}
