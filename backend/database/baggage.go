package database

import (
	"database/sql"
	"log"
	"mindenairport/models"

	"github.com/google/uuid"
)

// GetBaggageByID retrieves a specific baggage by ID
func (db Database) GetBaggageByID(id string) (*models.Baggage, error) {
	var baggage models.Baggage

	query := `SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
			  FROM BAGGAGE WHERE ID = :1`

	err := db.QueryRow(query, id).Scan(
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &baggage, nil
}

// GetBaggageByUserID retrieves all baggage for a specific user
func (db Database) GetBaggageByUserID(userID string) ([]models.Baggage, error) {
	var baggageList []models.Baggage

	query := `SELECT * FROM BAGGAGE WHERE AIRPORTUSER = :1 ORDER BY ID DESC`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var baggage models.Baggage
		err := rows.Scan(
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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return baggageList, nil
}

// GetBaggageByFlightID retrieves all baggage for a specific flight
func (db Database) GetBaggageByFlightID(flightID string) ([]models.Baggage, error) {
	var baggageList []models.Baggage

	query := `SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
			  FROM BAGGAGE WHERE FLIGHT = :1 ORDER BY ID DESC`

	rows, err := db.Query(query, flightID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var baggage models.Baggage
		err := rows.Scan(
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

	if err = rows.Err(); err != nil {
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

	query := `INSERT INTO BAGGAGE (ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING)
			  VALUES (:1, :2, :3, :4, :5, :6, :7, :8)`

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
	query := `UPDATE BAGGAGE SET 
			  AIRPORTUSER = :1, FLIGHT = :2, SIZE = :3, WEIGHT = :4, 
			  TRACKING_NUMBER = :5, STATUS = :6, SPECIAL_HANDLING = :7
			  WHERE ID = :8`

	result, err := db.Exec(query,
		baggage.AirportUserID,
		baggage.FlightID,
		baggage.Size,
		baggage.Weight,
		baggage.TrackingNumber,
		baggage.Status,
		baggage.SpecialHandling,
		id,
	)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	// Set the ID and return the updated baggage
	baggage.ID = id
	return &baggage, nil
}

// DeleteBaggage deletes a baggage entry
func (db Database) DeleteBaggage(id string) error {
	query := `DELETE FROM BAGGAGE WHERE ID = :1`

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// GetBaggageByTrackingNumber retrieves baggage by tracking number
func (db Database) GetBaggageByTrackingNumber(trackingNumber string) (*models.Baggage, error) {
	var baggage models.Baggage

	query := `SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
			  FROM BAGGAGE WHERE TRACKING_NUMBER = :1`

	err := db.QueryRow(query, trackingNumber).Scan(
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &baggage, nil
}

// GetAllBaggage retrieves all baggage with pagination for admin
func (db Database) GetAllBaggage(page, limit int) ([]models.Baggage, int, error) {
	var baggageList []models.Baggage
	var total int

	// First get the total count
	countQuery := `SELECT COUNT(*) FROM BAGGAGE`
	err := db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get baggage with pagination
	query := `SELECT ID, AIRPORTUSER, FLIGHT, SIZE, WEIGHT, TRACKING_NUMBER, STATUS, SPECIAL_HANDLING 
			  FROM BAGGAGE 
			  ORDER BY ID DESC
			  OFFSET :1 ROWS FETCH NEXT :2 ROWS ONLY`

	rows, err := db.Query(query, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var baggage models.Baggage
		err := rows.Scan(
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

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return baggageList, total, nil
}
