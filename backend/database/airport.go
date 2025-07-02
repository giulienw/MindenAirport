package database

import (
	"database/sql"
	"log"
	"mindenairport/models"
)

func (db Database) GetAirports() []models.Airport {
	query := `BEGIN GetAllAirports(:1); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, sql.Out{Dest: &cursor})
	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return []models.Airport{}
	}

	if cursor == nil {
		log.Printf("Cursor is nil after stored procedure call")
		return []models.Airport{}
	}
	defer cursor.Close()

	var airports []models.Airport

	for cursor.Next() {
		var airport models.Airport
		err := cursor.Scan(&airport.ID, &airport.Name, &airport.Country, &airport.City, &airport.Timezone, &airport.Elevation, &airport.NumberOfTerminal, &airport.Latitude, &airport.Longitude)
		if err != nil {
			log.Printf("Error scanning airport data: %v", err)
			continue
		}
		airports = append(airports, airport)
	}

	if err = cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
	}

	return airports
}

func (db Database) GetAirportByID(id string) models.Airport {
	query := `BEGIN GetAirportByID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, id, sql.Out{Dest: &cursor})
	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return models.Airport{}
	}

	if cursor == nil {
		log.Printf("Cursor is nil after stored procedure call")
		return models.Airport{}
	}
	defer cursor.Close()

	var airport models.Airport

	if cursor.Next() {
		err := cursor.Scan(&airport.ID, &airport.Name, &airport.Country, &airport.City, &airport.Timezone, &airport.Elevation, &airport.NumberOfTerminal, &airport.Latitude, &airport.Longitude)
		if err != nil {
			log.Printf("Error scanning airport data: %v", err)
			return models.Airport{}
		}
	}

	return airport
}
