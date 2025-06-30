package database

import (
	"database/sql"
	"log"
	"mindenairport/models"
	"mindenairport/utils"
	"time"

	"github.com/google/uuid"
)

// GetUserByEmail retrieves a user by email
func (db Database) GetUserByEmail(email string) (*models.AirportUser, error) {
	var user models.AirportUser

	query := `SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE, ROLE 
			  FROM AIRPORTUSER WHERE EMAIL = :1`

	err := db.QueryRow(query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Birthdate,
		&user.Password,
		&user.Active,
		&user.Email,
		&user.Phone,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByID retrieves a user by ID
func (db Database) GetUserByID(id string) (*models.AirportUser, error) {
	var user models.AirportUser

	query := `SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE, ROLE 
			  FROM AIRPORTUSER WHERE ID = :1`

	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Birthdate,
		&user.Password,
		&user.Active,
		&user.Email,
		&user.Phone,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user
func (db Database) CreateUser(req models.RegisterRequest) (*models.AirportUser, error) {
	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Generate new UUID for user
	userID := uuid.New().String()

	// Insert user into database
	query := `INSERT INTO AIRPORTUSER (ID, FIRSTNAME, LASTNAME, BIRTHDATE, PASSWORD, ACTIVE, EMAIL, PHONE) 
			  VALUES (:1, :2, :3, :4, :5, :6, :7, :8)`

	_, err = db.Exec(query,
		userID,
		req.FirstName,
		req.LastName,
		req.Birthdate,
		hashedPassword,
		1, // Active by default
		req.Email,
		req.Phone,
	)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	// Return the created user
	user := &models.AirportUser{
		ID:        userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Birthdate: req.Birthdate,
		Password:  hashedPassword,
		Active:    true,
		Email:     req.Email,
		Phone:     req.Phone,
	}

	return user, nil
}

// UpdateUserLastLogin updates the user's last login time
func (db Database) UpdateUserLastLogin(userID string) error {
	query := `UPDATE AIRPORTUSER SET LAST_LOGIN = :1 WHERE ID = :2`
	_, err := db.Exec(query, time.Now(), userID)
	return err
}

// DeactivateUser deactivates a user account
func (db Database) DeactivateUser(userID string) error {
	query := `UPDATE AIRPORTUSER SET ACTIVE = 0 WHERE ID = :1`
	_, err := db.Exec(query, userID)
	return err
}

// CheckEmailExists checks if an email already exists in the database
func (db Database) CheckEmailExists(email string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM AIRPORTUSER WHERE EMAIL = :1`
	err := db.QueryRow(query, email).Scan(&count)
	return count > 0, err
}
