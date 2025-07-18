// Package database provides flight-related database operations for the MindenAirport system.
// This includes flight retrieval, creation, updates, and management functions.
package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"mindenairport/models"
	"strconv"
	"time"

	"github.com/godror/godror"
)

// GetFlightByID retrieves a specific flight from the database by its unique identifier.
// This function calls the Oracle stored procedure GetFlightByID to fetch flight details.
//
// Parameters:
//   - id: The unique flight identifier
//
// Returns:
//   - models.Flight: The flight record with all details
//   - error: Any database error that occurred during retrieval
func (db Database) GetFlightByID(id string) (models.Flight, error) {
	stmt, err := db.Prepare(`
	BEGIN 
	GetFlightByID(:1, :2); END;
	`)
	if err != nil {
		return models.Flight{}, fmt.Errorf("error preparing statement: %w", err)
	}

	var cursor driver.Rows
	_, err = stmt.Exec(id, sql.Out{Dest: &cursor})
	if err != nil {
		return models.Flight{}, fmt.Errorf("error executing statement: %w", err)
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var flight models.Flight

	err = cursor.Next(r)
	if err == nil {
		flight.ID = r[0].(string)
		flight.From = r[1].(string)
		flight.To = r[2].(string)
		flight.PilotID = r[3].(string)
		flight.PlaneID = r[4].(string)
		if r[5] != nil {
			flight.TerminalID = r[5].(string)
		}
		if r[6] != nil {
			flight.StatusID, _ = strconv.Atoi(r[6].(godror.Number).String())
		}
		if r[7] != nil {
			flight.ScheduledDeparture = r[7].(time.Time)
		}
		if r[8] != nil {
			t := r[8].(time.Time)
			flight.ActualDeparture = &t
		}
		if r[9] != nil {
			flight.ScheduledArrival = r[9].(time.Time)
		}
		if r[10] != nil {
			t := r[10].(time.Time)
			flight.ActualArrival = &t
		}
		if r[11] != nil {
			flight.Gate = r[11].(string)
		}
		if r[12] != nil {
			flight.BaggageClaim = r[12].(string)
		}
	}

	return flight, nil
}

// GetFlights retrieves all flights from the database.
//
// DEPRECATED: This function lacks proper error handling and context management.
// Use GetAllFlights instead for production code.
//
// This function executes a direct SQL query to fetch all flight records
// and returns them as a slice. It logs fatal errors instead of returning them.
//
// Returns:
//   - []models.Flight: Slice of all flight records in the database
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
	countStmt, err := db.Prepare(`BEGIN GetFlightCount(:1); END;`)
	if err != nil {
		return nil, 0, err
	}
	_, err = countStmt.Exec(sql.Out{Dest: &total})
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get flights with pagination using stored procedure
	stmt, err := db.Prepare(`
	BEGIN 
	GetAllFlights(:1, :2, :3); END;
	`)
	if err != nil {
		return nil, 0, err
	}

	var cursor driver.Rows
	_, err = stmt.Exec(offset, limit, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, 0, err
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var flight models.Flight
		flight.ID = r[0].(string)
		flight.From = r[1].(string)
		flight.To = r[2].(string)
		flight.PilotID = r[3].(string)
		flight.PlaneID = r[4].(string)
		if r[5] != nil {
			flight.TerminalID = r[5].(string)
		}
		if r[6] != nil {
			flight.StatusID, _ = strconv.Atoi(r[6].(godror.Number).String())
		}
		if r[7] != nil {
			flight.ScheduledDeparture = r[7].(time.Time)
		}
		if r[8] != nil {
			t := r[8].(time.Time)
			flight.ActualDeparture = &t
		}
		if r[9] != nil {
			flight.ScheduledArrival = r[9].(time.Time)
		}
		if r[10] != nil {
			t := r[10].(time.Time)
			flight.ActualArrival = &t
		}
		if r[11] != nil {
			flight.Gate = r[11].(string)
		}
		if r[12] != nil {
			flight.BaggageClaim = r[12].(string)
		}
		flightList = append(flightList, flight)
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
