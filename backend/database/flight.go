package database

import (
	"database/sql"
	"fmt"
	"log"
	"mindenairport/models"
)

func (db Database) GetFlightByID(id string) (models.Flight, error) {
	query := `SELECT ID, "FROM", "TO", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM FROM FLIGHT WHERE ID = :1`

	var flight models.Flight
	row := db.QueryRow(query, id)
	
	err := row.Scan(
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
		if err == sql.ErrNoRows {
			return models.Flight{}, fmt.Errorf("flight with ID %s not found", id)
		}
		return models.Flight{}, fmt.Errorf("error scanning flight data: %w", err)
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

	// First get the total count
	countQuery := `SELECT COUNT(*) FROM FLIGHT`
	err := db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("error getting flight count: %w", err)
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get flights with pagination
	query := `SELECT ID, "FROM", "TO", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM 
	          FROM (
	              SELECT ROW_NUMBER() OVER (ORDER BY ID) as rn, ID, "FROM", "TO", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM 
	              FROM FLIGHT
	          ) 
	          WHERE rn > :1 AND rn <= :2`
	
	rows, err := db.Query(query, offset, offset+limit)
	if err != nil {
		return nil, 0, fmt.Errorf("error querying flights: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var flight models.Flight
		err := rows.Scan(
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
			return nil, 0, fmt.Errorf("error scanning flight: %w", err)
		}
		flightList = append(flightList, flight)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating flight rows: %w", err)
	}

	return flightList, total, nil
}
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
