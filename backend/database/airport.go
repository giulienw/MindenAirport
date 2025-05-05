package database

import (
	"log"
	"mindenairport/models"
)

func (db Database) GetAirports() []models.Airport {
	var airports []models.Airport

	rows, err := db.Query("SELECT * FROM Airport")
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airport models.Airport
		err := rows.Scan(&airport.ID, &airport.Name, &airport.Country, &airport.City, &airport.Timezone, &airport.Elevation, &airport.NumberOfTerminal, &airport.Latitude, &airport.Longitude)
		if err != nil {
			log.Fatal("Error scanning the database:", err)
		}
		airports = append(airports, airport)
	}

	return airports
}

func (db Database) GetAirportByID(id string) models.Airport {
	var airport models.Airport

	err := db.QueryRow("SELECT * FROM AIRPORT WHERE ID = :1", id).Scan(&airport.ID, &airport.Name, &airport.Country, &airport.City, &airport.Timezone, &airport.Elevation, &airport.NumberOfTerminal, &airport.Latitude, &airport.Longitude)
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}

	return airport
}
