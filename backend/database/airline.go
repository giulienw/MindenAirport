package database

import (
	"database/sql"
	"log"
	"mindenairport/models"
)

func (db Database) GetAirlineByID(id string) models.Airline {
	query := `BEGIN GetAirlineByID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, id, sql.Out{Dest: &cursor})
	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return models.Airline{}
	}

	if cursor == nil {
		log.Printf("Cursor is nil after stored procedure call")
		return models.Airline{}
	}
	defer cursor.Close()

	var airline models.Airline

	if cursor.Next() {
		err := cursor.Scan(&airline.ID, &airline.Name, &airline.Country, &airline.Logo, &airline.Active)
		if err != nil {
			log.Printf("Error scanning airline data: %v", err)
			return models.Airline{}
		}
	}

	return airline
}

func (db Database) GetAirlines() []models.Airline {
	query := `BEGIN GetAllAirlines(:1); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, sql.Out{Dest: &cursor})
	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return []models.Airline{}
	}

	if cursor == nil {
		log.Printf("Cursor is nil after stored procedure call")
		return []models.Airline{}
	}
	defer cursor.Close()

	var airlines []models.Airline

	for cursor.Next() {
		var airline models.Airline
		err := cursor.Scan(&airline.ID, &airline.Name, &airline.Country, &airline.Logo, &airline.Active)
		if err != nil {
			log.Printf("Error scanning airline data: %v", err)
			continue
		}
		airlines = append(airlines, airline)
	}

	if err = cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
	}

	return airlines
}
