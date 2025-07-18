package database

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"mindenairport/models"
)

func (db Database) GetAirlineByID(id string) models.Airline {
	query := `BEGIN MindenAirport.GetAirlineByID(:1, :2); END;`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return models.Airline{}
	}

	var cursor driver.Rows
	_, err = stmt.Exec(sql.Out{Dest: &cursor})
	if err != nil {
		log.Printf("Failed to execute prepared statement: %v", err)
		return models.Airline{}
	}

	r := make([]driver.Value, len(cursor.Columns()))
	err = cursor.Next(r)
	if err != nil {
		log.Println(err)
	}

	defer cursor.Close()

	var airline models.Airline

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var airline models.Airline
		airline.ID = r[0].(string)
		airline.Name = r[1].(string)
		airline.Country = r[2].(string)
		airline.Logo = r[3].(string)
		airline.Active = r[4].(int64) == 1
	}

	return airline
}

func (db Database) GetAirlines() []models.Airline {
	query := `BEGIN MindenAirport.GetAllAirlines(:1); END;`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return []models.Airline{}
	}

	var cursor driver.Rows
	_, err = stmt.Exec(sql.Out{Dest: &cursor})
	if err != nil {
		log.Printf("Failed to execute prepared statement: %v", err)
		return []models.Airline{}
	}

	r := make([]driver.Value, len(cursor.Columns()))
	err = cursor.Next(r)
	if err != nil {
		log.Println(err)
	}

	defer cursor.Close()

	var airlines []models.Airline

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var airline models.Airline
		airline.ID = r[0].(string)
		airline.Name = r[1].(string)
		airline.Country = r[2].(string)
		airline.Logo = r[3].(string)
		airline.Active = r[4].(int64) == 1
		airlines = append(airlines, airline)
	}

	return airlines
}
