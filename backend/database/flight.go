package database

import (
	"database/sql"
	"fmt"
	"log"
	"mindenairport/models"
)

func (db Database) GetFlightByID(id string) (models.Flight, error) {
	query := `BEGIN GetFlightByID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, id, sql.Out{Dest: &cursor})
	if err != nil {
		log.Fatal("Error calling stored procedure:", err)
	}
	defer cursor.Close()

	var flight models.Flight

	if cursor.Next() {
		err := cursor.Scan(&flight.ID, &flight.From, &flight.To, &flight.PilotID, &flight.PlaneID, &flight.TerminalID, &flight.StatusID, &flight.ScheduledDeparture, &flight.ActualDeparture, &flight.ScheduledArrival, &flight.ActualArrival, &flight.Gate, &flight.BaggageClaim)
		if err != nil {
			return models.Flight{}, fmt.Errorf("error scanning flight data: %w", err)
		}
	}

	return flight, nil
}

// Deprecated: Because missing error handling and not using context. Use GetAllFlights instead
func (db Database) GetFlights() []models.Flight {
	var flights []models.Flight

	rows, err := db.Query("SELECT ID, \"FROM\", \"TO\", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM FROM FLIGHT")
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var flight models.Flight
		err := rows.Scan(&flight.ID, &flight.From, &flight.To, &flight.PilotID, &flight.PlaneID, &flight.TerminalID, &flight.StatusID, &flight.ScheduledDeparture, &flight.ActualDeparture, &flight.ScheduledArrival, &flight.ActualArrival, &flight.Gate, &flight.BaggageClaim)
		if err != nil {
			log.Fatal("Error scanning the database:", err)
		}
		flights = append(flights, flight)
	}
	return flights
}

func (db Database) GetAllFlights(page, limit int) ([]models.Flight, int, error) {
	var flightList []models.Flight
	var total int

	// First get the total count using stored procedure
	countQuery := `BEGIN GetFlightCount(:1); END;`
	_, err := db.Exec(countQuery, sql.Out{Dest: &total})
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get flights with pagination using stored procedure
	query := `BEGIN GetAllFlights(:1, :2, :3); END;`
	var cursor *sql.Rows
	_, err = db.Exec(query, offset, limit, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close()

	for cursor.Next() {
		var flight models.Flight
		err := cursor.Scan(
			&flight.ID,
			&flight.From,
			&flight.To,
			&flight.PilotID,
			&flight.PlaneID,
			&flight.TerminalID,
			&flight.StatusID,
			&flight.ScheduledDeparture,
			&flight.ActualDeparture,
			&flight.ScheduledArrival,
			&flight.ActualArrival,
			&flight.Gate,
			&flight.BaggageClaim,
		)
		if err != nil {
			return nil, 0, err
		}
		flightList = append(flightList, flight)
	}

	if err = cursor.Err(); err != nil {
		return nil, 0, err
	}

	return flightList, total, nil
}

func (db Database) CreateFlight(flight models.Flight) {
	query := `BEGIN CreateFlight(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13); END;`
	_, err := db.Exec(query, flight.ID, flight.From, flight.To, flight.PilotID, flight.PlaneID, flight.TerminalID, flight.StatusID, flight.ScheduledDeparture, flight.ActualDeparture, flight.ScheduledArrival, flight.ActualArrival, flight.Gate, flight.BaggageClaim)
	if err != nil {
		log.Fatal("Error calling stored procedure:", err)
	}
}

func (db Database) UpdateFlight(flight models.Flight) {
	query := `BEGIN UpdateFlight(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13); END;`
	_, err := db.Exec(query, flight.ID, flight.From, flight.To, flight.PilotID, flight.PlaneID, flight.TerminalID, flight.StatusID, flight.ScheduledDeparture, flight.ActualDeparture, flight.ScheduledArrival, flight.ActualArrival, flight.Gate, flight.BaggageClaim)
	if err != nil {
		log.Fatal("Error calling stored procedure:", err)
	}
}

func (db Database) DeleteFlight(id string) {
	query := `BEGIN DeleteFlight(:1); END;`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("Error calling stored procedure:", err)
	}
}
