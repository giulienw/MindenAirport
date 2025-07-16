package database

import (
	"database/sql"
	"log"
	"mindenairport/models"

	"github.com/google/uuid"
)

// GetBaggageByID retrieves a specific baggage by ID
func (db Database) GetBaggageByID(id string) (*models.Baggage, error) {
	query := `BEGIN GetBaggageByID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, id, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var baggage models.Baggage

	if cursor.Next() {
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
			return nil, err
		}
		return &baggage, nil
	}

	return nil, nil
}

// GetBaggageByUserID retrieves all baggage for a specific user
func (db Database) GetBaggageByUserID(userID string) ([]models.Baggage, error) {
	query := `BEGIN GetBaggageByUserID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, userID, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var baggageList []models.Baggage

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
			return nil, err
		}
		baggageList = append(baggageList, baggage)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return baggageList, nil
}

// GetBaggageByFlightID retrieves all baggage for a specific flight
func (db Database) GetBaggageByFlightID(flightID string) ([]models.Baggage, error) {
	query := `BEGIN GetBaggageByFlightID(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, flightID, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var baggageList []models.Baggage

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
			return nil, err
		}
		baggageList = append(baggageList, baggage)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
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
	query := `BEGIN CreateBaggage(:1, :2, :3, :4, :5, :6, :7, :8); END;`
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
	query := `BEGIN UpdateBaggage(:1, :2, :3, :4, :5, :6, :7, :8); END;`
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
	query := `BEGIN DeleteBaggage(:1); END;`
	_, err := db.Exec(query, id)
	return err
}

// GetBaggageByTrackingNumber retrieves baggage by tracking number
func (db Database) GetBaggageByTrackingNumber(trackingNumber string) (*models.Baggage, error) {
	query := `BEGIN GetBaggageByTrackingNumber(:1, :2); END;`

	var cursor *sql.Rows
	_, err := db.Exec(query, trackingNumber, sql.Out{Dest: &cursor})
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	var baggage models.Baggage

	if cursor.Next() {
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
			return nil, err
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
	countQuery := `BEGIN GetBaggageCount(:1); END;`
	_, err := db.Exec(countQuery, sql.Out{Dest: &total})
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get baggage with pagination using stored procedure
	query := `BEGIN GetAllBaggage(:1, :2, :3); END;`
	var cursor *sql.Rows
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
