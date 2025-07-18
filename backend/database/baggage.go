package database

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"mindenairport/models"
	"strconv"

	"github.com/godror/godror"
	"github.com/google/uuid"
)

// GetBaggageByID retrieves a specific baggage by ID
func (db Database) GetBaggageByID(id string) (*models.Baggage, error) {
	stmt, err := db.Prepare(`BEGIN MindenAirport.GetBaggageByID(:1, :2); END;`)
	if err != nil {
		return nil, err
	}

	var cursor driver.Rows
	_, err = stmt.Exec(id, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var baggage models.Baggage

	err = cursor.Next(r)
	if err == nil {
		baggage.ID = r[0].(string)
		baggage.AirportUserID = r[1].(string)
		baggage.FlightID = r[2].(string)
		baggage.Size, _ = strconv.Atoi(r[3].(godror.Number).String())
		baggage.Weight, _ = strconv.ParseFloat(r[4].(godror.Number).String(), 64)
		baggage.TrackingNumber = r[5].(string)
		baggage.Status = r[6].(string)
		if r[7] != nil {
			baggage.SpecialHandling = r[7].(string)
		}
		return &baggage, nil
	}

	return nil, nil
}

// GetBaggageByUserID retrieves all baggage for a specific user
func (db Database) GetBaggageByUserID(userID string) ([]models.Baggage, error) {
	stmt, err := db.Prepare(`BEGIN MindenAirport.GetBaggageByUserID(:1, :2); END;`)
	if err != nil {
		return nil, err
	}

	var cursor driver.Rows
	_, err = stmt.Exec(userID, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var baggageList []models.Baggage

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var baggage models.Baggage
		baggage.ID = r[0].(string)
		baggage.AirportUserID = r[1].(string)
		baggage.FlightID = r[2].(string)
		baggage.Size, _ = strconv.Atoi(r[3].(godror.Number).String())
		baggage.Weight, _ = strconv.ParseFloat(r[4].(godror.Number).String(), 64)
		baggage.TrackingNumber = r[5].(string)
		baggage.Status = r[6].(string)
		if r[7] != nil {
			baggage.SpecialHandling = r[7].(string)
		}
		baggageList = append(baggageList, baggage)
	}

	return baggageList, nil
}

// GetBaggageByFlightID retrieves all baggage for a specific flight
func (db Database) GetBaggageByFlightID(flightID string) ([]models.Baggage, error) {
	stmt, err := db.Prepare(`BEGIN MindenAirport.GetBaggageByFlightID(:1, :2); END;`)
	if err != nil {
		return nil, err
	}

	var cursor driver.Rows
	_, err = stmt.Exec(flightID, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var baggageList []models.Baggage

	for {
		err := cursor.Next(r)
		if err != nil {
			break
		}
		var baggage models.Baggage
		baggage.ID = r[0].(string)
		baggage.AirportUserID = r[1].(string)
		baggage.FlightID = r[2].(string)
		baggage.Size, _ = strconv.Atoi(r[3].(godror.Number).String())
		baggage.Weight, _ = strconv.ParseFloat(r[4].(godror.Number).String(), 64)
		baggage.TrackingNumber = r[5].(string)
		baggage.Status = r[6].(string)
		if r[7] != nil {
			baggage.SpecialHandling = r[7].(string)
		}
		baggageList = append(baggageList, baggage)
	}

	return baggageList, nil
}

// CreateBaggage creates a new baggage entry
func (db Database) CreateBaggage(baggage models.Baggage) (*models.Baggage, error) {
	// Generate new UUID if not provided
	if baggage.ID == "" {
		baggage.ID = uuid.New().String()
	}

	// Generate tracking number if not provided
	if baggage.TrackingNumber == "" {
		baggage.TrackingNumber = "BAG" + uuid.New().String()[:8]
	}

	// Call stored procedure
	query := `BEGIN MindenAirport.CreateBaggage(:1, :2, :3, :4, :5, :6, :7, :8); END;`
	_, err := db.Exec(query,
		baggage.ID,
		baggage.AirportUserID,
		baggage.FlightID,
		baggage.Size,
		baggage.Weight,
		baggage.TrackingNumber,
		baggage.Status,
		baggage.SpecialHandling,
	)

	if err != nil {
		log.Printf("Error creating baggage: %v", err)
		return nil, err
	}

	return &baggage, nil
}

// UpdateBaggage updates an existing baggage entry
func (db Database) UpdateBaggage(id string, baggage models.Baggage) (*models.Baggage, error) {
	// Call stored procedure
	query := `BEGIN MindenAirport.UpdateBaggage(:1, :2, :3, :4, :5, :6, :7, :8); END;`
	_, err := db.Exec(query,
		id,
		baggage.AirportUserID,
		baggage.FlightID,
		baggage.Size,
		baggage.Weight,
		baggage.TrackingNumber,
		baggage.Status,
		baggage.SpecialHandling,
	)

	if err != nil {
		return nil, err
	}

	// Set the ID and return the updated baggage
	baggage.ID = id
	return &baggage, nil
}

// DeleteBaggage deletes a baggage entry
func (db Database) DeleteBaggage(id string) error {
	// Call stored procedure
	query := `BEGIN MindenAirport.DeleteBaggage(:1); END;`
	_, err := db.Exec(query, id)
	return err
}

// GetBaggageByTrackingNumber retrieves baggage by tracking number
func (db Database) GetBaggageByTrackingNumber(trackingNumber string) (*models.Baggage, error) {
	stmt, err := db.Prepare(`BEGIN MindenAirport.GetBaggageByTrackingNumber(:1, :2); END;`)
	if err != nil {
		return nil, err
	}

	var cursor driver.Rows
	_, err = stmt.Exec(trackingNumber, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}

	r := make([]driver.Value, len(cursor.Columns()))
	defer cursor.Close()

	var baggage models.Baggage

	err = cursor.Next(r)
	if err == nil {
		baggage.ID = r[0].(string)
		baggage.AirportUserID = r[1].(string)
		baggage.FlightID = r[2].(string)
		baggage.Size, _ = strconv.Atoi(r[3].(godror.Number).String())
		baggage.Weight, _ = strconv.ParseFloat(r[4].(godror.Number).String(), 64)
		baggage.TrackingNumber = r[5].(string)
		baggage.Status = r[6].(string)
		if r[7] != nil {
			baggage.SpecialHandling = r[7].(string)
		}
		return &baggage, nil
	}

	return nil, nil
}

// GetAllBaggage retrieves all baggage with pagination for admin
func (db Database) GetAllBaggage(page, limit int) ([]models.Baggage, int, error) {
	var baggageList []models.Baggage
	var total int

	// First get the total count using stored procedure
	countQuery := `BEGIN MindenAirport.GetBaggageCount(:1); END;`
	_, err := db.Exec(countQuery, sql.Out{Dest: &total})
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get baggage with pagination using stored procedure
	query := `BEGIN MindenAirport.GetAllBaggage(:1, :2, :3); END;`
	var cursor sql.Rows
	_, err = db.Exec(query, offset, limit, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close()

	for cursor.Next() {
		var baggage models.Baggage
		err := cursor.Scan(
			&baggage.ID,
			&baggage.AirportUserID,
			&baggage.FlightID,
			&baggage.Size,
			&baggage.Weight,
			&baggage.TrackingNumber,
			&baggage.Status,
			&baggage.SpecialHandling,
		)
		if err != nil {
			return nil, 0, err
		}
		baggageList = append(baggageList, baggage)
	}

	if err = cursor.Err(); err != nil {
		return nil, 0, err
	}

	return baggageList, total, nil
}
