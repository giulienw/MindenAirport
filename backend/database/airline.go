package database

import (
	"log"
	"mindenairport/models"
)

func (db Database) GetAirlineByID(id string) models.Airline {
	var airline models.Airline

	err := db.QueryRow("SELECT * FROM airline WHERE id = :1", id).Scan(&airline.ID, &airline.Name, &airline.Country, &airline.Logo, &airline.Active)
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}

	return airline
}

func (db Database) GetAirlines() []models.Airline {
	var airlines []models.Airline

	rows, err := db.Query("SELECT ID,NAME,COUNTRY,LOGO_URL,ACTIVE FROM airline")
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var airline models.Airline
		err := rows.Scan(&airline.ID, &airline.Name, &airline.Country, &airline.Logo, &airline.Active)
		if err != nil {
			log.Fatal("Error scanning the database:", err)
		}
		airlines = append(airlines, airline)
	}

	return airlines
}
