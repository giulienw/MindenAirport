package database

import (
	"log"
	"mindenairport/models"
)

func (db Database) GetFlightByID(id string) models.Flight {
	var flight models.Flight

	err := db.QueryRow("SELECT ID, \"FROM\", \"TO\", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM FROM FLIGHT WHERE ID = :1", id).Scan(&flight.ID, &flight.From, &flight.To, &flight.PilotID, &flight.PlaneID, &flight.TerminalID, &flight.StatusID, &flight.ScheduledDeparture, &flight.ActualDeparture, &flight.ScheduledArrival, &flight.ActualArrival, &flight.Gate, &flight.BaggageClaim)
	if err != nil {
		log.Fatal("Error querying the database:", err)
	}

	return flight
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
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get flights with pagination
	query := `SELECT ID, "FROM", "TO", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM FROM FLIGHT ORDER BY ID DESC OFFSET :1 ROWS FETCH NEXT :2 ROWS ONLY`

	rows, err := db.Query(query, offset, limit)
	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}
		flightList = append(flightList, flight)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return flightList, total, nil
}

func (db Database) CreateFlight(flight models.Flight) {
	_, err := db.Exec("INSERT INTO FLIGHT (ID, \"FROM\", \"TO\", PILOT, PLANE, TERMINAL, STATUS, SCHEDULED_DEPARTURE, ACTUAL_DEPARTURE, SCHEDULED_ARRIVAL, ACTUAL_ARRIVAL, GATE, BAGGAGE_CLAIM) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13)", flight.ID, flight.From, flight.To, flight.PilotID, flight.PlaneID, flight.TerminalID, flight.StatusID, flight.ScheduledDeparture, flight.ActualDeparture, flight.ScheduledArrival, flight.ActualArrival, flight.Gate, flight.BaggageClaim)
	if err != nil {
		log.Fatal("Error inserting into the database:", err)
	}
}

func (db Database) UpdateFlight(flight models.Flight) {
	query := `UPDATE FLIGHT SET 
	"FROM" = :1, 
	"TO" = :2, 
	PILOT = :3, 
	PLANE = :4, 
	TERMINAL = :5, 
	STATUS = :6, 
	SCHEDULED_DEPARTURE = :7, 
	ACTUAL_DEPARTURE = :8, 
	SCHEDULED_ARRIVAL = :9, 
	ACTUAL_ARRIVAL = :10, 
	GATE = :11, 
	BAGGAGE_CLAIM = :12 
	WHERE ID = :13`
	_, err := db.Exec(query, flight.From, flight.To, flight.PilotID, flight.PlaneID, flight.TerminalID, flight.StatusID, flight.ScheduledDeparture, flight.ActualDeparture, flight.ScheduledArrival, flight.ActualArrival, flight.Gate, flight.BaggageClaim, flight.ID)
	if err != nil {
		log.Fatal("Error updating the database:", err)
	}
}

func (db Database) DeleteFlight(id string) {
	_, err := db.Exec("DELETE FROM FLIGHT WHERE ID = :1", id)
	if err != nil {
		log.Fatal("Error deleting from the database:", err)
	}
}
